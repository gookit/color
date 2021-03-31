// +build windows

// Display color on windows
// refer:
//  golang.org/x/sys/windows
// 	golang.org/x/crypto/ssh/terminal
// 	https://docs.microsoft.com/en-us/windows/console
package color

import (
	"os"
	"syscall"
	"unsafe"

	"github.com/xo/terminfo"
	"golang.org/x/sys/windows"
)

// related docs
// https://docs.microsoft.com/zh-cn/windows/console/console-virtual-terminal-sequences
// https://docs.microsoft.com/zh-cn/windows/console/console-virtual-terminal-sequences#samples
var (
	// isMSys bool
	kernel32 *syscall.LazyDLL

	procGetConsoleMode *syscall.LazyProc
	procSetConsoleMode *syscall.LazyProc
)

func init() {
	// if at windows's ConEmu, Cmder, putty ... terminals
	if supportColor {
		return
	}

	isLikeInCmd = true

	// if disabled.
	if !Enable {
		return
	}

	// -------- try force enable colors on windows terminal -------

	// init simple color code info
	// initWinColorsMap()

	// load related windows dll
	// isMSys = utils.IsMSys()
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	// https://docs.microsoft.com/en-us/windows/console/setconsolemode
	procGetConsoleMode = kernel32.NewProc("GetConsoleMode")
	procSetConsoleMode = kernel32.NewProc("SetConsoleMode")

	// enable colors on windows terminal
	if tryApplyOnCONOUT() {
		// NOTICE: update var `supportColor` to TRUE.
		supportColor = true
		colorLevel, colorMark = LevelRgb, "VirtualTerminal"
		return
	}

	if tryApplyOnStdout() {
		// NOTICE: update var `supportColor` to TRUE.
		supportColor = true
		colorLevel, colorMark = LevelRgb, "VirtualTerminal"
	}

	// fetch console screen buffer info
	// err := getConsoleScreenBufferInfo(uintptr(syscall.Stdout), &defScreenInfo)
}

func tryApplyOnCONOUT() bool {
	outHandle, err := syscall.Open("CONOUT$", syscall.O_RDWR, 0)
	if err != nil {
		saveInternalError(err)
		return false
	}

	err = EnableVirtualTerminalProcessing(outHandle, true)
	if err != nil {
		saveInternalError(err)
		return false
	}

	return true
}

func tryApplyOnStdout() bool {
	// try direct open syscall.Stdout
	err := EnableVirtualTerminalProcessing(syscall.Stdout, true)
	if err != nil {
		saveInternalError(err)
		return false
	}

	return true
}

// Get the Windows Version and Build Number
var (
	winVersion, _, buildNumber = windows.RtlGetNtVersionNumbers()
)

// refer from https://github.com/Delta456/box-cli-maker/blob/master/detect_windows.go
// detectTerminalColor detects the Color Level Supported
func detectTerminalColor() terminfo.ColorLevel {
	if os.Getenv("ConEmuANSI") == "ON" {
		// ConEmuANSI is "ON" for generic ANSI support
		// but True Color option is enabled by default
		// I am just assuming that people wouldn't have disabled it
		// Even if it is not enabled then ConEmu will auto round off
		// accordingly
		return terminfo.ColorLevelMillions
	}

	// Before Windows 10 Build Number 10586, console never supported ANSI Colors
	if buildNumber < 10586 || winVersion < 10 {
		// Detect if using ANSICON on older systems
		if os.Getenv("ANSICON") != "" {
			conVersion := os.Getenv("ANSICON_VER")
			// 8 bit Colors were only supported after v1.81 release
			if conVersion >= "181" {
				return terminfo.ColorLevelHundreds
			}
			return terminfo.ColorLevelBasic
		}
		return terminfo.ColorLevelNone
	}

	// True Color is not available before build 14931 so fallback to 8 bit color.
	if buildNumber < 14931 {
		return terminfo.ColorLevelHundreds
	}
	return terminfo.ColorLevelMillions
}

/*************************************************************
 * render full color code on windows(8,16,24bit color)
 *************************************************************/

// docs https://docs.microsoft.com/zh-cn/windows/console/getconsolemode#parameters
const (
	// equals to docs page's ENABLE_VIRTUAL_TERMINAL_PROCESSING 0x0004
	EnableVirtualTerminalProcessingMode uint32 = 0x4
)

// EnableVirtualTerminalProcessing Enable virtual terminal processing
//
// ref from github.com/konsorten/go-windows-terminal-sequences
// doc https://docs.microsoft.com/zh-cn/windows/console/console-virtual-terminal-sequences#samples
//
// Usage:
// 	err := EnableVirtualTerminalProcessing(syscall.Stdout, true)
// 	// support print color text
// 	err = EnableVirtualTerminalProcessing(syscall.Stdout, false)
func EnableVirtualTerminalProcessing(stream syscall.Handle, enable bool) error {
	var mode uint32
	// Check if it is currently in the terminal
	// err := syscall.GetConsoleMode(syscall.Stdout, &mode)
	err := syscall.GetConsoleMode(stream, &mode)
	if err != nil {
		// fmt.Println("EnableVirtualTerminalProcessing", err)
		return err
	}

	if enable {
		mode |= EnableVirtualTerminalProcessingMode
	} else {
		mode &^= EnableVirtualTerminalProcessingMode
	}

	ret, _, err := procSetConsoleMode.Call(uintptr(stream), uintptr(mode))
	if ret == 0 {
		return err
	}

	return nil
}

// renderColorCodeOnCmd enable cmd color render.
// func renderColorCodeOnCmd(fn func()) {
// 	err := EnableVirtualTerminalProcessing(syscall.Stdout, true)
// 	// if is not in terminal, will clear color tag.
// 	if err != nil {
// 		// panic(err)
// 		fn()
// 		return
// 	}
//
// 	// force open color render
// 	old := ForceOpenColor()
// 	fn()
// 	// revert color setting
// 	supportColor = old
//
// 	err = EnableVirtualTerminalProcessing(syscall.Stdout, false)
// 	if err != nil {
// 		panic(err)
// 	}
// }

/*************************************************************
 * render simple color code on windows
 *************************************************************/

// IsTty returns true if the given file descriptor is a terminal.
func IsTty(fd uintptr) bool {
	var st uint32
	r, _, e := syscall.Syscall(procGetConsoleMode.Addr(), 2, fd, uintptr(unsafe.Pointer(&st)), 0)
	return r != 0 && e == 0
}

// IsTerminal returns true if the given file descriptor is a terminal.
//
// Usage:
// 	fd := os.Stdout.Fd()
// 	fd := uintptr(syscall.Stdout) // for windows
// 	IsTerminal(fd)
func IsTerminal(fd uintptr) bool {
	var st uint32
	r, _, e := syscall.Syscall(procGetConsoleMode.Addr(), 2, fd, uintptr(unsafe.Pointer(&st)), 0)
	return r != 0 && e == 0
}
