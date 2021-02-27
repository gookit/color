package color

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
)

// mark/flag
var supColorMark string

// Support color:
// 	"TERM=xterm"
// 	"TERM=xterm-vt220"
// 	"TERM=xterm-256color"
// 	"TERM=screen-256color"
// 	"TERM=tmux-256color"
// 	"TERM=rxvt-unicode-256color"
// Don't support color:
// 	"TERM=cygwin"
var specialColorTerms = map[string]bool{
	"alacritty": true,
}

// SupColorMark get
func SupColorMark() string {
	return supColorMark
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

// IsSupportColor check current console is support color.
//
// Supported:
// 	linux, mac, or windows's ConEmu, Cmder, putty, git-bash.exe
// Not support:
// 	windows cmd.exe, powerShell.exe
func IsSupportColor() bool {
	envTerm := os.Getenv("TERM")
	if strings.Contains(envTerm, "term") {
		supColorMark = "TERM=" + envTerm
		return true
	}

	// it's special color term
	if _, ok := specialColorTerms[envTerm]; ok {
		supColorMark = "TERM=" + envTerm
		return true
	}

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

	// up: if support 256-color, can also support basic color.
	return isSupport256Color(envTerm)
}

// IsSupport256Color render
func IsSupport256Color() bool {
	return isSupport256Color(os.Getenv("TERM"))
}

func isSupport256Color(termVal string) bool {
	// "TERM=xterm-256color"
	// "TERM=screen-256color"
	// "TERM=tmux-256color"
	// "TERM=rxvt-unicode-256color"
	supported := strings.Contains(termVal, "256color")
	if !supported {
		// up: if support true-color, can also support 256-color.
		return IsSupportTrueColor()
	}

	supColorMark = "TERM=" + termVal
	return supported
}

// IsSupportRGBColor render. alias of the IsSupportTrueColor()
func IsSupportRGBColor() bool {
	return IsSupportTrueColor()
}

// IsSupportTrueColor render.
func IsSupportTrueColor() bool {
	v := os.Getenv("COLORTERM")
	// "COLORTERM=truecolor"
	// "COLORTERM=24bit"
	ok := strings.Contains(v, "truecolor") || strings.Contains(v, "24bit")

	if ok {
		supColorMark = "COLORTERM=" + v
	}
	return ok
}

/*************************************************************
 * helper methods for converts
 *************************************************************/

// Colors2code convert colors to code. return like "32;45;3"
func Colors2code(colors ...Color) string {
	if len(colors) == 0 {
		return ""
	}

	var codes []string
	for _, color := range colors {
		codes = append(codes, color.String())
	}

	return strings.Join(codes, ";")
}

// Hex2rgb alias of the HexToRgb()
func Hex2rgb(hex string) []int { return HexToRgb(hex) }

// HexToRGB alias of the HexToRgb()
func HexToRGB(hex string) []int { return HexToRgb(hex) }

// HexToRgb convert hex color string to RGB numbers
// Usage:
// 	rgb := HexToRgb("ccc") // rgb: [204 204 204]
// 	rgb := HexToRgb("aabbcc") // rgb: [170 187 204]
// 	rgb := HexToRgb("#aabbcc") // rgb: [170 187 204]
// 	rgb := HexToRgb("0xad99c0") // rgb: [170 187 204]
func HexToRgb(hex string) (rgb []int) {
	hex = strings.TrimSpace(hex)
	if hex == "" {
		return
	}

	// like from css. eg "#ccc" "#ad99c0"
	if hex[0] == '#' {
		hex = hex[1:]
	}

	hex = strings.ToLower(hex)
	switch len(hex) {
	case 3: // "ccc"
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	case 8: // "0xad99c0"
		hex = strings.TrimPrefix(hex, "0x")
	}

	// recheck
	if len(hex) != 6 {
		return
	}

	// convert string to int64
	if i64, err := strconv.ParseInt(hex, 16, 32); err == nil {
		color := int(i64)
		// parse int
		rgb = make([]int, 3)
		rgb[0] = color >> 16
		rgb[1] = (color & 0x00FF00) >> 8
		rgb[2] = color & 0x0000FF
	}

	return
}

// Rgb2hex alias of the RgbToHex()
func Rgb2hex(rgb []int) string { return RgbToHex(rgb) }

// RgbToHex convert RGB-code to hex-code
// Usage:
//	hex := RgbToHex([]int{170, 187, 204}) // hex: "aabbcc"
func RgbToHex(rgb []int) string {
	hexNodes := make([]string, len(rgb))
	for _, v := range rgb {
		hexNodes = append(hexNodes, strconv.FormatInt(int64(v), 16))
	}

	return strings.Join(hexNodes, "")
}

// Rgb2ansi alias of the RgbToAnsi()
func Rgb2ansi(r, g, b uint8, isBg bool) uint8 {
	return RgbToAnsi(r, g, b, isBg)
}

// RgbToAnsi convert RGB-code to 16-code
// refer https://github.com/radareorg/radare2/blob/master/libr/cons/rgb.c#L249-L271
func RgbToAnsi(r, g, b uint8, isBg bool) uint8 {
	var bright, c, k uint8

	base := compareVal(isBg, BgBase, FgBase)

	// eco bright-specific
	if r == 0x80 && g == 0x80 && b == 0x80 { // 0x80=128
		bright = 53
	} else if r == 0xff || g == 0xff || b == 0xff { // 0xff=255
		bright = 60
	} // else bright = 0

	if r == g && g == b {
		// 0x7f=127
		// r = (r > 0x7f) ? 1 : 0;
		r = compareVal(r > 0x7f, 1, 0)
		g = compareVal(g > 0x7f, 1, 0)
		b = compareVal(b > 0x7f, 1, 0)
	} else {
		k = (r + g + b) / 3

		// r = (r >= k) ? 1 : 0;
		r = compareVal(r >= k, 1, 0)
		g = compareVal(g >= k, 1, 0)
		b = compareVal(b >= k, 1, 0)
	}

	// c = (r ? 1 : 0) + (g ? (b ? 6 : 2) : (b ? 4 : 0))
	c = compareVal(r > 0, 1, 0)

	if g > 0 {
		c += compareVal(b > 0, 6, 2)
	} else {
		c += compareVal(b > 0, 4, 0)
	}

	return base + bright + c
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
