// +build !windows

// The method in the file has no effect
// Only for compatibility with non-Windows systems

package color

import (
	"io/ioutil"
	"runtime"
	"strings"
	"syscall"

	"github.com/xo/terminfo"
)

// refer
//  https://github.com/Delta456/box-cli-maker/blob/7b5a1ad8a016ce181e7d8b05e24b54ff60b4b38a/detect_unix.go#L27-L45
// detect WSL as it has True Color support
func detectSpecialTermColor() (level terminfo.ColorLevel, needVTP bool) {
	// Detect WSL as it has True Color support
	if runtime.GOOS == "windows" {
		// `cat /proc/sys/kernel/osrelease`
		// on mac:
		//	!not the file!
		// on linux:
		// 	4.19.121-linuxkit
		// on WSL Output:
		//  4.4.0-19041-Microsoft
		wsl, err := ioutil.ReadFile("/proc/sys/kernel/osrelease")
		if err != nil {
			saveInternalError(err)
			return
		}

		// it gives "Microsoft" for WSL and "microsoft" for WSL 2
		content := strings.ToLower(string(wsl))
		if strings.Contains(content, "microsoft") {
			level = terminfo.ColorLevelMillions
		}
	}

	return
}

// IsTerminal returns true if the given file descriptor is a terminal.
//
// Usage:
// 	IsTerminal(os.Stdout.Fd())
func IsTerminal(fd uintptr) bool {
	return fd == uintptr(syscall.Stdout) || fd == uintptr(syscall.Stdin) || fd == uintptr(syscall.Stderr)
}
