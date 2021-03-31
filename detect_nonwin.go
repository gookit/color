// +build !windows

// The method in the file has no effect
// Only for compatibility with non-Windows systems

package color

import "syscall"

// func renderColorCodeOnCmd(_ func()) {}

// IsTerminal returns true if the given file descriptor is a terminal.
//
// Usage:
// 	IsTerminal(os.Stdout.Fd())
func IsTerminal(fd uintptr) bool {
	return fd == uintptr(syscall.Stdout) || fd == uintptr(syscall.Stdin) || fd == uintptr(syscall.Stderr)
}
