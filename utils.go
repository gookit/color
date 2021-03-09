package color

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"syscall"
)

var (
	// support color:
	// 	"TERM=xterm"
	// 	"TERM=xterm-vt220"
	// 	"TERM=xterm-256color"
	// 	"TERM=screen-256color"
	// 	"TERM=tmux-256color"
	// 	"TERM=rxvt-unicode-256color"
	// Don't support color:
	// 	"TERM=cygwin"
	specialColorTerms = map[string]bool{
		"alacritty": true,
	}

	// mark/flag string.
	supColorMark string

	// the color support level for current terminal
	colorLevel = LevelNo
	// syncOnce = sync.Once{}
)

// ColorLevel value
func ColorLevel() uint8 {
	return colorLevel
}

// SupColorMark get
func SupColorMark() string {
	return supColorMark
}

// func ColorLevel(refresh bool) uint8 {
// 	syncOnce.Do(func() {
// 		IsSupportColor()
// 	})
//
// 	return colorLevel
// }

// func findColorLevel() uint8 {
// }

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

// IsSupportColor check current console is support color.
//
// Supported:
// 	linux, mac, or windows's ConEmu, Cmder, putty, git-bash.exe
// Not support:
// 	windows cmd.exe, powerShell.exe
func IsSupportColor() bool {
	// check true color.
	envCTerm := os.Getenv("COLORTERM")
	if isSupportTrueColor(envCTerm) {
		colorLevel = LevelRgb
		return true
	}

	envTerm := os.Getenv("TERM")

	// check 256 color
	ok := isSupport256Color(envTerm)
	if ok {
		colorLevel = Level256
	} else if ok = isSupport16Color(envTerm); ok { // check 16 color
		colorLevel = Level16
	}

	return ok
}

// IsSupportColor check current console is support color.
//
// Supported:
// 	linux, mac, or windows's ConEmu, Cmder, putty, git-bash.exe
// Not support:
// 	windows cmd.exe, powerShell.exe
func IsSupport16Color() bool {
	envTerm := os.Getenv("TERM")
	return isSupport16Color(envTerm)
}

func isSupport16Color(envTerm string) bool {
	if strings.Contains(envTerm, "term") {
		supColorMark = "TERM=" + envTerm
		return true
	}

	return false
}

// IsSupport256Color render
func IsSupport256Color() bool {
	return isSupport256Color(os.Getenv("TERM"))
}

func isSupport256Color(envTerm string) bool {
	// like on ConEmu software, e.g "ConEmuANSI=ON"
	if os.Getenv("ConEmuANSI") == "ON" {
		supColorMark = "ConEmuANSI=ON"
		return true
	}

	// like on ConEmu software, e.g "ANSICON=189x2000 (189x43)"
	if val := os.Getenv("ANSICON"); val != "" {
		supColorMark = "ANSICON=" + val
		return true
	}

	// it's special color term
	if _, ok := specialColorTerms[envTerm]; ok {
		supColorMark = "TERM=" + envTerm
		return true
	}

	// "TERM=xterm-256color"
	// "TERM=screen-256color"
	// "TERM=tmux-256color"
	// "TERM=rxvt-unicode-256color"
	ok := strings.Contains(envTerm, "256color")
	if ok {
		supColorMark = "TERM=" + envTerm
	}
	return ok
}

// IsSupportRGBColor render. alias of the IsSupportTrueColor()
func IsSupportRGBColor() bool {
	return IsSupportTrueColor()
}

// IsSupportTrueColor render.
func IsSupportTrueColor() bool {
	envCTerm := os.Getenv("COLORTERM")
	return isSupportTrueColor(envCTerm)
}

func isSupportTrueColor(envCTerm string) bool {
	// "COLORTERM=truecolor"
	// "COLORTERM=24bit"
	ok := strings.Contains(envCTerm, "truecolor") || strings.Contains(envCTerm, "24bit")
	if ok {
		supColorMark = "COLORTERM=" + envCTerm
	}
	return ok
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

// refer https://github.com/Delta456/box-cli-maker
// func detectTerminalColor() terminfo.ColorLevel {
// 	level, err := terminfo.ColorLevelFromEnv()
// 	if err != nil {
// 		saveInternalError(err)
// 		return terminfo.ColorLevelNone
// 	}
//
// 	// Detect WSL as it has True Color support
// 	if level == terminfo.ColorLevelNone && runtime.GOOS == "windows" {
// 		wsl, err := ioutil.ReadFile("/proc/sys/kernel/osrelease")
// 		if err != nil {
// 			saveInternalError(err)
// 			return level
// 		}
//
// 		// Microsoft for WSL and microsoft for WSL 2
// 		content := strings.ToLower(string(wsl))
// 		if strings.Contains(content, "microsoft") {
// 			return terminfo.ColorLevelMillions
// 		}
// 	}
//
// 	return level
// }
