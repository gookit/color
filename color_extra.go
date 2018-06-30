package color

import "fmt"

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
const Set256Fg = "\x1b[38;5;%dm"
const Set256Bg = "\x1b[48;5;%dm"

// 24 byte RGB color
// RGB:
// 	R 0-255 G 0-255 B 0-255
// 	R 00-FF G 00-FF B 00-FF (16进制)
// format:
// 	ESC[ … 38;2;<r>;<g>;<b> … m // 选择RGB前景色
// 	ESC[ … 48;2;<r>;<g>;<b> … m // 选择RGB背景色
// links:
// 	https://zh.wikipedia.org/wiki/ANSI转义序列#24位
const SetRgbFg = "\x1b[38;2;%dm"
const SetRgbBg = "\x1b[48;2;%dm"

// uint8 at 0 - 255
// 10进制和16进制都可 0x98 = 152
type Bt8Color uint8
type RgbColor [3]uint8

// Theme 主题 Color wheel 色盘

// Byte8Color use 256 color
func Byte8Color(str string, val Bt8Color) {
	fmt.Printf("\x1b[38;5;242;48;5;208m%s\x1b[0m\n", str)
	fmt.Printf("\x1b[38;5;%dm%s\x1b[0m\n", val, str)
}

// Byte24Color use RGB color
func Byte24Color(str string, rgb RgbColor)  {

}

// Byte24Color use RGB color
func TrueColor(str string, rgb RgbColor)  {

}
