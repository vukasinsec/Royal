#!/bin/bash

# Pokreće ngrok u background modu i čeka 2 sekunde
ngrok http 8080 > /dev/null &
sleep 2

# Uzimanje novog ngrok URL-a
NGROK_URL=$(curl -s http://localhost:4040/api/tunnels | jq -r '.tunnels[0].public_url')

# Proveri da li je BOT_TOKEN postavljen
if [[ -z "$BOT_TOKEN" ]]; then
    echo "❌ Greška: BOT_TOKEN nije postavljen! Postavi ga sa 'export BOT_TOKEN=...' "
    exit 1
fi

# Proveri da li je URL dobijen
if [[ $NGROK_URL == http* ]]; then
    echo "🔗 Novi NGROK URL: $NGROK_URL"
    
    # Postavljanje Telegram Webhook-a
    curl -F "url=$NGROK_URL/webhook" \
    "https://api.telegram.org/bot$BOT_TOKEN/setWebhook"

    echo "✅ Webhook podešen na: $NGROK_URL/webhook"
else
    echo "❌ Greška: NGROK URL nije dobijen!"
fi
