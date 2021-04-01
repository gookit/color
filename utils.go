package color

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"syscall"

	"github.com/xo/terminfo"
)

/*************************************************************
 * helper methods for detect color supports
 *************************************************************/

// DetectColorLevel for current env
//
// NOTICE: The method will detect terminal info each times,
// 	if only want get current color level, please direct call SupportColor() or TermColorLevel()
func DetectColorLevel() terminfo.ColorLevel {
	level, _ := detectTermColorLevel()
	return level
}

// IsWindows OS env
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// func TermColorLevel(refresh bool) uint8 {
// 	syncOnce.Do(func() {
// 		IsSupportColor()
// 	})
//
// 	return colorLevel
// }

// refer https://github.com/Delta456/box-cli-maker
func detectTermColorLevel() (level terminfo.ColorLevel, needVTP bool) {
	var err error
	level, err = terminfo.ColorLevelFromEnv()
	if err != nil {
		// if on windows OS
		if runtime.GOOS == "windows" {
			level, needVTP = detectSpecialTermColor()
		} else {
			saveInternalError(err)
		}
		return
	}

	// Detect WSL as it has True Color support
	if level == terminfo.ColorLevelNone && runtime.GOOS == "windows" {
		level, needVTP = detectSpecialTermColor()
	}
	return
}

var detectedWSL bool
var wslContents string

// https://github.com/Microsoft/WSL/issues/423#issuecomment-221627364
func detectWSL() bool {
	if !detectedWSL {
		b := make([]byte, 1024)
		// `cat /proc/version`
		// on mac:
		// 	!not the file!
		// on linux(debian,ubuntu,alpine):
		//	Linux version 4.19.121-linuxkit (root@18b3f92ade35) (gcc version 9.2.0 (Alpine 9.2.0)) #1 SMP Thu Jan 21 15:36:34 UTC 2021
		// on win git bash, conEmu:
		// 	MINGW64_NT-10.0-19042 version 3.1.7-340.x86_64 (@WIN-N0G619FD3UK) (gcc version 9.3.0 (GCC) ) 2020-10-23 13:08 UTC
		// on WSL:
		//  Linux version 4.4.0-19041-Microsoft (Microsoft@Microsoft.com) (gcc version 5.4.0 (GCC) ) #488-Microsoft Mon Sep 01 13:43:00 PST 2020
		f, err := os.Open("/proc/version")
		if err == nil {
			_, _ = f.Read(b) // ignore error
			f.Close()
			wslContents = string(b)
		}
		detectedWSL = true
	}
	return strings.Contains(wslContents, "Microsoft")
}

// IsSupportColor check current console is support color.
//
// NOTICE: The method will detect terminal info each times,
// 	if only want get current color level, please direct call SupportColor() or TermColorLevel()
func IsSupportColor() bool {
	return IsSupport16Color()
}

// IsSupportColor check current console is support color.
//
// NOTICE: The method will detect terminal info each times,
// 	if only want get current color level, please direct call SupportColor() or TermColorLevel()
func IsSupport16Color() bool {
	level, _ := detectTermColorLevel()
	return level > terminfo.ColorLevelNone
}

// IsSupport256Color render check
//
// NOTICE: The method will detect terminal info each times,
// 	if only want get current color level, please direct call SupportColor() or TermColorLevel()
func IsSupport256Color() bool {
	level, _ := detectTermColorLevel()
	return level > terminfo.ColorLevelBasic
}

// IsSupportRGBColor check. alias of the IsSupportTrueColor()
//
// NOTICE: The method will detect terminal info each times,
// 	if only want get current color level, please direct call SupportColor() or TermColorLevel()
func IsSupportRGBColor() bool {
	return IsSupportTrueColor()
}

// IsSupportTrueColor render check.
//
// NOTICE: The method will detect terminal info each times,
// 	if only want get current color level, please direct call SupportColor() or TermColorLevel()
//
// ENV:
// "COLORTERM=truecolor"
// "COLORTERM=24bit"
func IsSupportTrueColor() bool {
	level, _ := detectTermColorLevel()
	return level > terminfo.ColorLevelHundreds
}

/*************************************************************
 * helper methods for check env
 *************************************************************/

// IsConsole Determine whether w is one of stderr, stdout, stdin
func IsConsole(w io.Writer) bool {
	o, ok := w.(*os.File)
	if !ok {
		return false
	}

	fd := o.Fd()

	// fix: cannot use 'o == os.Stdout' to compare
	return fd == uintptr(syscall.Stdout) || fd == uintptr(syscall.Stdin) || fd == uintptr(syscall.Stderr)
}

// IsMSys msys(MINGW64) environment, does not necessarily support color
func IsMSys() bool {
	// like "MSYSTEM=MINGW64"
	if len(os.Getenv("MSYSTEM")) > 0 {
		return true
	}

	return false
}

/*************************************************************
 * print methods(will auto parse color tags)
 *************************************************************/

// Print render color tag and print messages
func Print(a ...interface{}) {
	Fprint(output, a...)
}

// Printf format and print messages
func Printf(format string, a ...interface{}) {
	Fprintf(output, format, a...)
}

// Println messages with new line
func Println(a ...interface{}) {
	Fprintln(output, a...)
}

// Fprint print rendered messages to writer
// Notice: will ignore print error
func Fprint(w io.Writer, a ...interface{}) {
	_, err := fmt.Fprint(w, Render(a...))
	saveInternalError(err)

	// if isLikeInCmd {
	// 	renderColorCodeOnCmd(func() {
	// 		_, _ = fmt.Fprint(w, Render(a...))
	// 	})
	// } else {
	// 	_, _ = fmt.Fprint(w, Render(a...))
	// }
}

// Fprintf print format and rendered messages to writer.
// Notice: will ignore print error
func Fprintf(w io.Writer, format string, a ...interface{}) {
	str := fmt.Sprintf(format, a...)
	_, err := fmt.Fprint(w, ReplaceTag(str))
	saveInternalError(err)
}

// Fprintln print rendered messages line to writer
// Notice: will ignore print error
func Fprintln(w io.Writer, a ...interface{}) {
	str := formatArgsForPrintln(a)
	_, err := fmt.Fprintln(w, ReplaceTag(str))
	saveInternalError(err)
}

// Lprint passes colored messages to a log.Logger for printing.
// Notice: should be goroutine safe
func Lprint(l *log.Logger, a ...interface{}) {
	l.Print(Render(a...))
}

// Render parse color tags, return rendered string.
// Usage:
//	text := Render("<info>hello</> <cyan>world</>!")
//	fmt.Println(text)
func Render(a ...interface{}) string {
	if len(a) == 0 {
		return ""
	}

	return ReplaceTag(fmt.Sprint(a...))
}

// Sprint parse color tags, return rendered string
func Sprint(a ...interface{}) string {
	if len(a) == 0 {
		return ""
	}

	return ReplaceTag(fmt.Sprint(a...))
}

// Sprintf format and return rendered string
func Sprintf(format string, a ...interface{}) string {
	return ReplaceTag(fmt.Sprintf(format, a...))
}

// String alias of the ReplaceTag
func String(s string) string {
	return ReplaceTag(s)
}

// Text alias of the ReplaceTag
func Text(s string) string {
	return ReplaceTag(s)
}

/*************************************************************
 * helper methods for print
 *************************************************************/

// new implementation, support render full color code on pwsh.exe, cmd.exe
func doPrintV2(code, str string) {
	_, err := fmt.Fprint(output, RenderString(code, str))
	saveInternalError(err)

	// if isLikeInCmd {
	// 	renderColorCodeOnCmd(func() {
	// 		_, _ = fmt.Fprint(output, RenderString(code, str))
	// 	})
	// } else {
	// 	_, _ = fmt.Fprint(output, RenderString(code, str))
	// }
}

// new implementation, support render full color code on pwsh.exe, cmd.exe
func doPrintlnV2(code string, args []interface{}) {
	str := formatArgsForPrintln(args)
	_, err := fmt.Fprintln(output, RenderString(code, str))
	saveInternalError(err)
}

// if use Println, will add spaces for each arg
func formatArgsForPrintln(args []interface{}) (message string) {
	if ln := len(args); ln == 0 {
		message = ""
	} else if ln == 1 {
		message = fmt.Sprint(args[0])
	} else {
		message = fmt.Sprintln(args...)
		// clear last "\n"
		message = message[:len(message)-1]
	}
	return
}

/*************************************************************
 * helper methods
 *************************************************************/

// its Win system. linux windows darwin
// func isWindows() bool {
// 	return runtime.GOOS == "windows"
// }

// equals: return ok ? val1 : val2
func compareVal(ok bool, val1, val2 uint8) uint8 {
	if ok {
		return val1
	}
	return val2
}

func saveInternalError(err error) {
	if err != nil {
		errors = append(errors, err)
	}
}

func stringToArr(str, sep string) (arr []string) {
	str = strings.TrimSpace(str)
	if str == "" {
		return
	}

	ss := strings.Split(str, sep)
	for _, val := range ss {
		if val = strings.TrimSpace(val); val != "" {
			arr = append(arr, val)
		}
	}
	return
}
