// display color on windows
// ref:
//  golang.org/x/sys/windows
// 	golang.org/x/crypto/ssh/terminal
// 	https://docs.microsoft.com/en-us/windows/console

// +build windows

package color

import (
	"syscall"
	"fmt"
	"unsafe"
	"os"
)

type WColor uint16
type WStyle []WColor

// color on windows
// you can see on windows by command: COLOR /?
// windows color build by: Bg + Fg
// Consists of any two of the following:
// the first is the background color, and the second is the foreground color
// 颜色属性由两个十六进制数字指定
//  - 第一个对应于背景，第二个对应于前景。
// 	- 当只传入一个值时，则认为是前景色
// 每个数字可以为以下任何值:
// more see: https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/cmd
const (
	// Foreground colors.
	WinFgBlack  WColor = 0x00 // 0 黑色
	WinFgBlue   WColor = 0x01 // 1 蓝色
	WinFgGreen  WColor = 0x02 // 2 绿色
	WinFgAqua   WColor = 0x03 // 3 浅绿 skyblue
	WinFgRed    WColor = 0x04 // 4 红色
	WinFgPurple WColor = 0x05 // 5 紫色
	WinFgYellow WColor = 0x06 // 6 黄色
	WinFgWhite  WColor = 0x07 // 7 白色
	WinFgGray   WColor = 0x08 // 8 灰色

	WinFgLightBlue   = 0x09 // 9 淡蓝色
	WinFgLightGreen  = 0x0a // 10 淡绿色
	WinFgLightAqua   = 0x0b // 11 淡浅绿色
	WinFgLightRed    = 0x0c // 12 淡红色
	WinFgLightPurple = 0x0d // 13 淡紫色
	WinFgLightYellow = 0x0e // 14 淡黄色
	WinFgLightWhite  = 0x0f // 15 亮白色

	// Background colors.
	WinBgBlack  = 0x00 // 黑色
	WinBgBlue   = 0x10 // 蓝色
	WinBgGreen  = 0x20 // 绿色
	WinBgAqua   = 0x30 // 浅绿 skyblue
	WinBgRed    = 0x40 // 红色
	WinBgPink   = 0x50 // 紫色
	WinBgYellow = 0x60 // 黄色
	WinBgWhite  = 0x70 // 白色
	WinBgGray   = 0x80 // 128 灰色

	WinBgLightBlue   = 0x90 // 淡蓝色
	WinBgLightGreen  = 0xa0 // 淡绿色
	WinBgLightAqua   = 0xb0 // 淡浅绿色
	WinBgLightRed    = 0xc0 // 淡红色
	WinBgLightPink   = 0xd0 // 淡紫色
	WinBgLightYellow = 0xe0 // 淡黄色
	WinBgLightWhite  = 0xf0 // 240 亮白色

	// bg black, fg white
	defSetting = WinBgBlack | WinFgWhite

	// see https://docs.microsoft.com/en-us/windows/console/char-info-str
	WinFgIntensity uint16 = 0x0008 // 8 前景强度
	WinBgIntensity uint16 = 0x0080 // 128 背景强度

	WinOpLeading    WColor = 0x0100 // 前导字节
	WinOpTrailing   WColor = 0x0200 // 尾随字节
	WinOpHorizontal WColor = 0x0400 // 顶部水平
	WinOpReverse    WColor = 0x4000 // 反转前景和背景
	WinOpUnderscore WColor = 0x8000 // 32768 下划线
)

var colorsMap = map[Color]WColor{}

var (
	// for cmd.exe
	escChar = ""
	// isMSys bool
	kernel32 *syscall.LazyDLL

	// procGetConsoleMode *syscall.LazyProc
	// procSetConsoleMode *syscall.LazyProc

	procSetTextAttribute           *syscall.LazyProc
	procGetConsoleScreenBufferInfo *syscall.LazyProc

	// console screen buffer info
	// eg {size:{x:215 y:3000} cursorPosition:{x:0 y:893} attributes:7 window:{left:0 top:882 right:214 bottom:893} maximumWindowSize:{x:215 y:170}}
	defScreenInfo consoleScreenBufferInfo
)

func init() {
	// Byte8Color("test 8 byte color", 208)
	// Byte24Color("test 24 byte color")
	// os.Exit(0)

	// if at linux, mac, or windows's ConEmu, Cmder, putty
	if isSupportColor {
		return
	}

	// isMSys = utils.IsMSys()
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	// https://docs.microsoft.com/en-us/windows/console/setconsolemode
	// procGetConsoleMode = kernel32.NewProc("GetConsoleMode")
	// procSetConsoleMode = kernel32.NewProc("SetConsoleMode")

	procSetTextAttribute = kernel32.NewProc("SetConsoleTextAttribute")
	// https://docs.microsoft.com/en-us/windows/console/getconsolescreenbufferinfo
	procGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")

	// fetch console screen buffer info
	getConsoleScreenBufferInfo(uintptr(syscall.Stdout), &defScreenInfo)

	fmt.Printf("%+v\n", WinOpUnderscore)

	// 2|8 = 2+8 = 10, 'A' = 65
	// 8|4|2 = 14
	// fmt.Println(9|8|2, '\x10', 0x0a, 0xa)
	WinPrint("test [OK];\n", WinFgRed)
	// revertDefault()
	os.Exit(0)
}

// win 设置终端字体颜色
// 使用方法，直接调用即可输出带颜色的文本
// WPrint("[OK];", 2|8) //亮绿色
func WinPrint(s string, val WColor) {
	// kernel32 := syscall.NewLazyDLL("kernel32.dll")
	// proc := kernel32.NewProc("SetConsoleTextAttribute")
	fmt.Print("val: ", val, " ")

	handle, _, _ := procSetTextAttribute.Call(uintptr(syscall.Stdout), uintptr(val))

	fmt.Print(s)

	// handle, _, _ = procSetTextAttribute.Call(uintptr(syscall.Stdout), uintptr(7))

	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}

// revertDefault
func revertDefault() bool {
	return setConsoleTextAttr(uintptr(syscall.Stdout), uint16(defSetting))
}

// setConsoleTextAttr
func setConsoleTextAttr(consoleOutput uintptr, winAttr uint16) bool {
	ret, _, _ := procSetTextAttribute.Call(consoleOutput, uintptr(winAttr))

	return ret != 0
}

// IsTty returns true if the given file descriptor is a terminal.
// func IsTty(fd uintptr) bool {
// 	var st uint32
// 	r, _, e := syscall.Syscall(procGetConsoleMode.Addr(), 2, fd, uintptr(unsafe.Pointer(&st)), 0)
// 	return r != 0 && e == 0
// }

// IsTerminal returns true if the given file descriptor is a terminal.
// fd := os.Stdout.Fd()
// fd := uintptr(syscall.Stdout) for windows
// func IsTerminal(fd int) bool {
// 	var st uint32
// 	r, _, e := syscall.Syscall(procGetConsoleMode.Addr(), 2, uintptr(fd), uintptr(unsafe.Pointer(&st)), 0)
// 	return r != 0 && e == 0
// }

// from package: golang.org/x/sys/windows
type (
	short int16
	word uint16

	// coord cursor position coordinates
	coord struct {
		x short
		y short
	}

	smallRect struct {
		left   short
		top    short
		right  short
		bottom short
	}

	// Used with GetConsoleScreenBuffer to retreive information about a console
	// screen buffer. See
	// https://docs.microsoft.com/en-us/windows/console/console-screen-buffer-info-str
	// for details.
	consoleScreenBufferInfo struct {
		size              coord
		cursorPosition    coord
		attributes        word // is windows color setting
		window            smallRect
		maximumWindowSize coord
	}
)

// GetSize returns the dimensions of the given terminal.
func getSize(fd int) (width, height int, err error) {
	var info consoleScreenBufferInfo

	if err := getConsoleScreenBufferInfo(uintptr(fd), &info); err != nil {
		return 0, 0, err
	}

	return int(info.size.x), int(info.size.y), nil
}

// from package: golang.org/x/sys/windows
func getConsoleScreenBufferInfo(consoleOutput uintptr, info *consoleScreenBufferInfo) (err error) {
	r1, _, e1 := syscall.Syscall(procGetConsoleScreenBufferInfo.Addr(), 2, consoleOutput, uintptr(unsafe.Pointer(info)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = e1
		} else {
			err = syscall.EINVAL
		}
	}

	return
}

/**
	The follow codes from package: golang.org/x/crypto/ssh/terminal
 */
const (
	enableLineInput       = 2
	enableEchoInput       = 4
	enableProcessedInput  = 1
	enableWindowInput     = 8
	enableMouseInput      = 16
	enableInsertMode      = 32
	enableQuickEditMode   = 64
	enableExtendedFlags   = 128
	enableAutoPosition    = 256
	enableProcessedOutput = 1
	enableWrapAtEolOutput = 2
)

const (
	keyCtrlD       = 4
	keyCtrlU       = 21
	keyEnter       = '\r'
	keyEscape      = 27
	keyBackspace   = 127
	keyUnknown     = 0xd800 /* UTF-16 surrogate area */ + iota
	keyUp
	keyDown
	keyLeft
	keyRight
	keyAltLeft
	keyAltRight
	keyHome
	keyEnd
	keyDeleteWord
	keyDeleteLine
	keyClearScreen
	keyPasteStart
	keyPasteEnd
)

var (
	crlf       = []byte{'\r', '\n'}
	pasteStart = []byte{keyEscape, '[', '2', '0', '0', '~'}
	pasteEnd   = []byte{keyEscape, '[', '2', '0', '1', '~'}
)
