package commands

import (
	"fmt"
	"os/exec"
	"runtime"
)

// ExecuteCommand izvršava shell komande i vraća rezultat
func ExecuteCommand(cmd string) string {
	var shell, flag string

	if runtime.GOOS == "windows" {
		shell, flag = "cmd", "/C"
	} else {
		shell, flag = "sh", "-c"
	}

	out, err := exec.Command(shell, flag, cmd).Output()
	if err != nil {
		return fmt.Sprintf("Error: %s", err)
	}
	return string(out)
}
