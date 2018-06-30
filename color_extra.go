package color

import (
	"fmt"
	"strings"
)

// 8 byte 256 color(`2^8`)
// format:
// 	ESC[ … 38;5;<n> … m // 选择前景色
//  ESC[ … 48;5;<n> … m // 选择背景色
// example:
//  fg "\x1b[38;5;242m"
//  bg "\x1b[48;5;208m"
//  both "\x1b[38;5;242;48;5;208m"
// links:
// 	https://zh.wikipedia.org/wiki/ANSI转义序列#8位
const Tpl256Fg = "38;5;%d"
const Tpl256Bg = "48;5;%d"

// uint8 at 0 - 255
// 10进制和16进制都可 0x98 = 152
// use 8 byte, 0 - 255 color
type Bt8Color uint8

type Bt8Style struct {
	Fg, Bg Bt8Color
}

// Byte8Color use 8 byte, 0 - 255 color
func Byte8Color(str string, val Bt8Color) {
	fmt.Printf("\x1b[38;5;242;48;5;208m%s\x1b[0m\n", str)
	fmt.Printf("\x1b[38;5;%dm%s\x1b[0m\n", val, str)
}

func (s *Bt8Style) Print(args ...interface{}) (n int, err error) {
	return fmt.Printf(FullColorTpl, s.String(), fmt.Sprint(args...))
}

func (s *Bt8Style) String() string {
	var ss []string
	if s.Fg > 0 {
		ss = append(ss, fmt.Sprintf(Tpl256Fg, s.Fg))
	}

	if s.Bg > 0 {
		ss = append(ss, fmt.Sprintf(Tpl256Bg, s.Bg))
	}

	return strings.Join(ss, ";")
}

// 24 byte RGB color
// RGB:
// 	R 0-255 G 0-255 B 0-255
// 	R 00-FF G 00-FF B 00-FF (16进制)
// format:
// 	ESC[ … 38;2;<r>;<g>;<b> … m // 选择RGB前景色
// 	ESC[ … 48;2;<r>;<g>;<b> … m // 选择RGB背景色
// links:
// 	https://zh.wikipedia.org/wiki/ANSI转义序列#24位
const SetRgbFg = "\x1b[38;2;%d;%d;%dm"
const SetRgbBg = "\x1b[48;2;%d;%d;%dm"

type RgbColor [3]uint8

type RgbStyle struct {
	Fg, Bg RgbColor
}

// Byte24Color use RGB color
func Byte24Color(str string) {
	fmt.Printf("\x1b[38;2;30;144;255m%s\x1b[0m\n", str)
}

// Byte24Color use RGB color
func TrueColor(str string, rgb RgbColor)  {
	// RgbStyle{RgbColor{'dd', 'cc', 'dd'}, RgbColor{'dd', 'cc', 'dd'}}
}
