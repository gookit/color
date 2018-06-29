// display color on windows

// +build windows

package color

import (
	"syscall"
	"fmt"
	"os"
)

// color on windows
// you can see on windows by command: COLOR /?
// windows color build by: Bg + Fg
// Consists of any two of the following,
// the first is the background color,
// and the second is the foreground color
const (
	// 0 - 9 and A - F
	winBlack  = iota
	winBlue
	winGreen
	winLightGreen  // 浅绿
	winRed
	winPurple      // 紫色
	winYellow
	winWhite
	winGray
	winLightBlue   // 淡蓝色

	winPaleGreen      = "A" // 淡绿色
	winPaleLightGreen = "B" // 淡浅绿色
	winLightRed       = "C" // 淡红色
	winLightPurple    = "D" // 淡紫色
	winLightYellow    = "E" // 淡黄色
	winBrightWhite    = "F" // 亮白色

	// bg black, fg white
	defSetting = 07
)

// ref from package `issue9/term`
// windows 预定义的颜色值
const (
	winFgBlue      = 1
	winFgGreen     = 2
	winFgRed       = 4
	winFgIntensity = 8

	winBgBlue      = 16
	winBgGreen     = 32
	winBgRed       = 64
	winBgIntensity = 123

	// 增强前景色
	winFgYellow  = winFgRed + winFgGreen
	winFgCyan    = winFgGreen + winFgBlue
	winFgMagenta = winFgBlue + winFgRed
	winFgWhite   = winFgRed + winFgBlue + winFgGreen
	winFgDefault = winFgWhite

	// 增强背景色
	winBgYellow  = winBgRed + winBgGreen
	winBgCyan    = winBgGreen + winBgBlue
	winBgMagenta = winBgBlue + winBgRed
	winBgWhite   = winBgRed + winBgBlue + winBgGreen
	winBgDefault = 0 // 默认背景为黑

	defColor = winFgDefault + winBgDefault
)

var (
	isMSys bool

	kernel32                   *syscall.LazyDLL
	setConsoleTextAttribute    *syscall.LazyProc
	getConsoleScreenBufferInfo *syscall.LazyProc
)

func init() {
	if len(os.Getenv("MSYSTEM")) > 0 { // msys 环境
		isMSys = true
		return
	}

	kernel32 = syscall.NewLazyDLL("kernel32.dll")
	setConsoleTextAttribute = kernel32.NewProc("SetConsoleTextAttribute")
	getConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
}

// win 设置终端字体颜色
// 使用方法，直接调用即可输出带颜色的文本
// WPrint("[OK];", 2|8) //亮绿色
func WPrint(s string, i int) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")

	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))

	fmt.Print(s)

	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}
