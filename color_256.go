package color

import (
	"fmt"
	"strings"
)

/*
from wikipedia:
   ESC[ … 38;5;<n> … m选择前景色
   ESC[ … 48;5;<n> … m选择背景色
     0-  7：标准颜色（同 ESC[30–37m）
     8- 15：高强度颜色（同 ESC[90–97m）
    16-231：6 × 6 × 6 立方（216色）: 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
   232-255：从黑到白的24阶灰度色
*/

// TplFg256 8 bit 256 color(`2^8`)
//
// format:
// 	ESC[ … 38;5;<n> … m // 选择前景色
//  ESC[ … 48;5;<n> … m // 选择背景色
//
// example:
//  fg "\x1b[38;5;242m"
//  bg "\x1b[48;5;208m"
//  both "\x1b[38;5;242;48;5;208m"
//
// links:
// 	https://zh.wikipedia.org/wiki/ANSI%E8%BD%AC%E4%B9%89%E5%BA%8F%E5%88%97#8位
const TplFg256 = "38;5;%d"
const TplBg256 = "48;5;%d"

// Bit8Color 256 (8 bit) color, uint8 range at 0 - 255
//
// 颜色值使用10进制和16进制都可 0x98 = 152
//
// 颜色有两位uint8组成, 0: color value 1: color type, Fg(0) or Bg(^0)
// 	fg color: [152, 0]
//  bg color: [152, 1]
type Bit8Color [2]uint8

// Bit8 create a color256
func Bit8(val uint8, isBg ...bool) Bit8Color {
	return C256(val, isBg...)
}

// C256 create a color256
func C256(val uint8, isBg ...bool) Bit8Color {
	bc := Bit8Color{val}

	// mark is bg color
	if len(isBg) > 0 && isBg[0] {
		bc[1] = 1
	}

	return bc
}

func (c Bit8Color) Print(args interface{}) {

}

// String convert to string
func (c Bit8Color) String() string {
	if c[1] == 0 { // 0 is Fg
		return fmt.Sprintf(TplFg256, c[0])
	}

	// ^0 is Bg
	return fmt.Sprintf(TplBg256, c[0])
}

type Attribute uint16

// Style256 definition
//
// 前/背景色
// 都是由两位uint8组成, 第一位是色彩值；
// 第二位与Bit8Color不一样的是，在这里表示是否设置了值
type Style256 struct {
	Name   string
	fg, bg Bit8Color
}

// S256 Color256
// Usage:
// 	s := color.S256()
// 	s := color.S256(132)
// 	s := color.S256(132, 203)
func S256(values ...uint8) *Style256 {
	s := &Style256{}
	vl := len(values)

	// with fg
	if vl > 0 {
		s.fg = Bit8Color{values[0], 1}

		// and with bg
		if vl > 1 {
			s.bg = Bit8Color{values[1], 1}
		}
	}

	return s
}

// SetBg set bg color value
func (s *Style256) SetBg(bgVal uint8) {
	s.bg = Bit8Color{bgVal, 1}
}

// SetFg set fg color value
func (s *Style256) SetFg(fgVal uint8) {
	s.fg = Bit8Color{fgVal, 1}
}

// Print print message
func (s *Style256) Print(a ...interface{}) (n int, err error) {
	return fmt.Printf(FullColorTpl, s.String(), fmt.Sprint(a...))
}

func (s *Style256) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(FullColorTpl, s.String(), fmt.Sprintf(format, a...))
}

func (s *Style256) Println( a ...interface{}) (n int, err error) {
	return fmt.Printf(FullColorNlTpl, s.String(), fmt.Sprint(a...))
}

// String convert to string
func (s *Style256) String() string {
	var ss []string
	if s.fg[1] > 0 {
		ss = append(ss, fmt.Sprintf(TplFg256, s.fg[0]))
	}

	if s.bg[1] > 0 {
		ss = append(ss, fmt.Sprintf(TplBg256, s.bg[0]))
	}

	return strings.Join(ss, ";")
}

func Color256Table() {

}

// Byte8Color use 8 byte, 0 - 255 color
func Byte8Color(str string, val Bit8Color) {
	fmt.Printf("\x1b[38;5;242;48;5;208m%s\x1b[0m\n", str)
	fmt.Printf("\x1b[38;5;%dm%s\x1b[0m\n", val, str)
}

// 16-231：6 × 6 × 6 立方（216色）: 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
func RGBto216(n int) int {
	if n < 0 {
		return 0
	}

	if n > 5 {
		return 5
	}

	return n
}
