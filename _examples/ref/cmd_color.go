//+build windows

package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"
)

// https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences#samples
// go run _examples/ref/cmd_color.go
func main() {
	err := EnableVirtualTerminalProcessing(syscall.Stdout, true)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	//noinspection GoUnhandledErrorResult
	defer EnableVirtualTerminalProcessing(syscall.Stdout, false)

	fmt.Println("\x1b[34mHello \x1b[35mWorld\x1b[0m!")
	_, _ = fmt.Fprintf(os.Stdout, "\x1b[34mHello \x1b[35mWorld\x1b[0m!\n")
}

var (
	kernel32Dll    *syscall.LazyDLL  = syscall.NewLazyDLL("Kernel32.dll")
	setConsoleMode *syscall.LazyProc = kernel32Dll.NewProc("SetConsoleMode")
)

// https://docs.microsoft.com/zh-cn/windows/console/console-virtual-terminal-sequences#samples
// from github.com/konsorten/go-windows-terminal-sequences
// Usage:
// 	EnableVirtualTerminalProcessing()
func EnableVirtualTerminalProcessing(stream syscall.Handle, enable bool) error {
	const EnableVirtualTerminalProcessing uint32 = 0x4

	var mode uint32
	err := syscall.GetConsoleMode(syscall.Stdout, &mode)
	if err != nil {
		return err
	}

	if enable {
		mode |= EnableVirtualTerminalProcessing
	} else {
		mode &^= EnableVirtualTerminalProcessing
	}

	ret, _, err := setConsoleMode.Call(uintptr(unsafe.Pointer(stream)), uintptr(mode))
	if ret == 0 {
		return err
	}

	return nil
}
