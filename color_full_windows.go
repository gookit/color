// +build windows

// support print color text on windows cmd.exe
package color

import (
	"syscall"
	"unsafe"
)

// related docs
// https://docs.microsoft.com/zh-cn/windows/console/console-virtual-terminal-sequences
// https://docs.microsoft.com/zh-cn/windows/console/console-virtual-terminal-sequences#samples
var (
	kernel32Dll    = syscall.NewLazyDLL("Kernel32.dll")
	setConsoleMode = kernel32Dll.NewProc("SetConsoleMode")
)

// docs https://docs.microsoft.com/zh-cn/windows/console/getconsolemode#parameters
const (
	// equals to docs page's ENABLE_VIRTUAL_TERMINAL_PROCESSING 0x0004
	EnableVirtualTerminalProcessingMode uint32 = 0x4
)

// EnableVirtualTerminalProcessing Enable virtual terminal processing
// ref from github.com/konsorten/go-windows-terminal-sequences
// Usage:
// 	err := EnableVirtualTerminalProcessing(syscall.Stdout, true)
// 	// support print color text
// 	err = EnableVirtualTerminalProcessing(syscall.Stdout, false)
func EnableVirtualTerminalProcessing(stream syscall.Handle, enable bool) error {
	var mode uint32
	// Check if it is currently in the terminal
	err := syscall.GetConsoleMode(syscall.Stdout, &mode)
	if err != nil {
		return err
	}

	if enable {
		mode |= EnableVirtualTerminalProcessingMode
	} else {
		mode &^= EnableVirtualTerminalProcessingMode
	}

	ret, _, err := setConsoleMode.Call(uintptr(unsafe.Pointer(stream)), uintptr(mode))
	if ret == 0 {
		return err
	}

	return nil
}

// EnableCmdColorRender enable cmd color render.
func EnableCmdColorRender(fn func()) {
	err := EnableVirtualTerminalProcessing(syscall.Stdout, true)
	// if is not in terminal, will clear color tag.
	if err != nil {
		// panic(err)
		fn()
		return
	}

	// force open color render
	old := ForceOpenColor()
	fn()
	// revert color setting
	isSupportColor = old

	err = EnableVirtualTerminalProcessing(syscall.Stdout, false)
	if err != nil {
		panic(err)
	}
}
