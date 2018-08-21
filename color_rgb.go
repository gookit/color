package color

import (
	"fmt"
	"strconv"
	"strings"
)

// TplFgRGB 24 bit RGB color
// RGB:
// 	R 0-255 G 0-255 B 0-255
// 	R 00-FF G 00-FF B 00-FF (16进制)
//
// format:
// 	ESC[ … 38;2;<r>;<g>;<b> … m // 选择RGB前景色
// 	ESC[ … 48;2;<r>;<g>;<b> … m // 选择RGB背景色
//
// links:
// 	https://zh.wikipedia.org/wiki/ANSI%E8%BD%AC%E4%B9%89%E5%BA%8F%E5%88%97#24位
//
// example:
// 	fg: \x1b[38;2;30;144;255mMESSAGE\x1b[0m
// 	bg: \x1b[48;2;30;144;255mMESSAGE\x1b[0m
// 	both: \x1b[38;2;233;90;203;48;2;30;144;255mMESSAGE\x1b[0m
const TplFgRGB = "38;2;%d;%d;%d"
const TplBgRGB = "48;2;%d;%d;%d"

// RGBColor definition.
//
// The first to third digits represent the color value,
// The last digit represents the foreground(0) or background(^0)
//
// Usage:
// 	// 0, 1, 2 is R,G,B. 3 is Fg(0) or Bg(^0).
// 	RGBColor{30,144,255, 0}
// 	RGBColor{30,144,255, 1}
type RGBColor [4]uint8

// RGB color create
func RGB(r, g, b uint8, isBg ...bool) RGBColor {
	rgb := RGBColor{r, g, b}

	// mark is bg color
	if len(isBg) > 0 && isBg[0] {
		rgb[3] = 1
	}

	return rgb
}

// Print print message
func (c RGBColor) Print(a ...interface{}) {
	fmt.Printf(FullColorTpl, c.String(), fmt.Sprint(a...))
}

// Printf format and print message
func (c RGBColor) Printf(format string, a ...interface{}) {
	fmt.Printf(FullColorTpl, c.String(), fmt.Sprintf(format, a...))
}

// Println print message with newline
func (c RGBColor) Println(a ...interface{}) {
	fmt.Printf(FullColorNlTpl, c.String(), fmt.Sprint(a...))
}

// Sprint returns rendered message
func (c RGBColor) Sprint(a ...interface{}) string {
	return fmt.Sprintf(FullColorTpl, c.String(), fmt.Sprint(a...))
}

// Sprint returns format and rendered message
func (c RGBColor) Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(FullColorTpl, c.String(), fmt.Sprintf(format, a...))
}

// String to string
func (c RGBColor) String() string {
	if c[3] == 0 { // 0 is Fg
		return fmt.Sprintf(TplFgRGB, c[0], c[1], c[2])
	}

	// ^0 is Bg
	return fmt.Sprintf(TplBgRGB, c[0], c[1], c[2])
}

// RGBStyle definition.
//
// 前/背景色
// 都是由4位uint8组成, 前三位是色彩值；
// 最后一位与RGBColor不一样的是，在这里表示是否设置了值
type RGBStyle struct {
	Name   string
	fg, bg RGBColor
}

// String convert to string
func (s *RGBStyle) String() string {
	var ss []string
	if s.fg[3] > 0 {
		ss = append(ss, fmt.Sprintf(TplFgRGB, s.fg[0], s.fg[1], s.fg[2]))
	}

	if s.bg[3] > 0 {
		ss = append(ss, fmt.Sprintf(TplBgRGB, s.bg[0], s.bg[1], s.bg[2]))
	}

	return strings.Join(ss, ";")
}

// Bit24Color use RGB color
func Bit24Color(str string) {
	fmt.Printf("\x1b[38;2;30;144;255m%s\x1b[0m\n", str)
}

// TrueColor use RGB color
func TrueColor(str string, rgb RGBColor) {
	// RGBStyle{RGBColor{'dd', 'cc', 'dd'}, RGBColor{'dd', 'cc', 'dd'}}

}

func RGBto256(r, g, b uint8) {

}

// HexToRGB hex color string to RGB numbers
// Usage:
// 	rgb := HexToRGB("ccc") // rgb: [204 204 204]
// 	rgb := HexToRGB("aabbcc") // rgb: [170 187 204]
// 	rgb := HexToRGB("0xad99c0") // rgb: [170 187 204]
func HexToRGB(hex string) (rgb []int) {
	hex = strings.TrimSpace(hex)
	if hex == "" {
		return
	}

	hex = strings.ToLower(hex)
	switch len(hex) {
	case 3: // "ccc"
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	case 6: // "ad99c0"
	case 8: // "0xad99c0"
		hex = strings.TrimPrefix(hex, "0x")
	default: // invalid
		return
	}

	// 字串到数据整型
	i64, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		// panic("invalid color string, error: " + err.Error())
		return
	}

	color := int(i64)
	// parse int
	rgb = make([]int, 3)
	rgb[0] = color >> 16
	rgb[1] = (color & 0x00FF00) >> 8
	rgb[2] = color & 0x0000FF

	return
}
