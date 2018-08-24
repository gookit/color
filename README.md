# cli color

[![GoDoc](https://godoc.org/github.com/gookit/color?status.svg)](https://godoc.org/github.com/gookit/color)
[![Build Status](https://travis-ci.org/gookit/color.svg?branch=master)](https://travis-ci.org/gookit/color)
[![Coverage Status](https://coveralls.io/repos/github/gookit/color/badge.svg?branch=master)](https://coveralls.io/github/gookit/color?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/color)](https://goreportcard.com/report/github.com/gookit/color)

Command line color library. rich color rendering output, universal API method, compatible with Windows system

**[中文说明](README_cn.md)**

## Features

- Simple to use
- Supports rich color output 16色(4bit)，256色(8bit)，RGB色彩(24bit)
  - 16 color (4bit) is the most commonly used and most widely supported, supporting Windows `cmd.exe`
  - 256 and RGB color support `linux` `mac` and Windows `ConEmu` `git-bash` `mintty` part terminal
- Generic API method: `Print` `Printf` `Println` `Sprint` `Sprintf`
- Supports html tab-style color rendering. like: `<green>message</>`
- Basic color: `Bold` `Black` `White` `Gray` `Red` `Green` `Yellow` `Blue` `Magenta` `Cyan`
- Extra style: `Info` `Note` `Light` `Error` `Danger` `Notice` `Success` `Comment` `Primary` `Warning` `Question` `Secondary`
- Compatible with Windows system environment

## Install

- use dep

```bash
dep ensure -add gopkg.in/gookit/color.v1 // is recommended
// OR
dep ensure -add github.com/gookit/color
```

- go get

```bash
go get gopkg.in/gookit/color.v1 // is recommended
// OR
go get -u github.com/gookit/color
```

## Godoc

- [godoc for gopkg](https://godoc.org/gopkg.in/gookit/color.v1)
- [godoc for github](https://godoc.org/github.com/gookit/color)

## Quick start

```bash
import "gopkg.in/gookit/color.v1" // is recommended
// or
import "github.com/gookit/color"
```

```go
package main

import (
	"fmt"
	"github.com/gookit/color"
)

func main() {
	// simple usage
	color.Cyan.Printf("Simple to use %s\n", "color")

	// use like func
	red := color.FgRed.Render
	green := color.FgGreen.Render
	fmt.Printf("%s line %s library\n", red("Command"), green("color"))

	// custom color
	color.New(color.FgWhite, color.BgBlack).Println("custom color style")

	// can also:
	color.Style{color.FgCyan, color.OpBold}.Println("custom color style")
	
	// internal theme/style:
	color.Info.Tips("message")
	color.Info.Prompt("message")
	color.Info.Println("message")
	color.Warn.Println("message")
	color.Error.Println("message")
	
	// use style tag
	color.Print("<suc>he</><comment>llo</>, <cyan>wel</><red>come</>\n")

	// apply a style tag
	color.Tag("info").Println("info style text")

	// prompt message
	color.Info.Prompt("prompt style message")
	color.Warn.Prompt("prompt style message")

	// tips message
	color.Info.Tips("tips style message")
	color.Warn.Tips("tips style message")
}
```

> run demo: `go run ./_examples/app.go`

![colored-out](_examples/images/color-demo.jpg)

## More usage

### Basic color

> support on windows `cmd.exe`

- `color.Bold`
- `color.Black`
- `color.White`
- `color.Gray`
- `color.Red`
- `color.Green`
- `color.Yellow`
- `color.Blue`
- `color.Magenta`
- `color.Cyan`

```go
color.Bold.Println("bold message")
color.Yellow.Println("yellow message")
```

> run demo: `go run ./_examples/basiccolor.go`

![basic-color](_examples/images/basic-color.png)

### Extra styles 

> support on windows `cmd.exe`

- `color.Info`
- `color.Note`
- `color.Light`
- `color.Error`
- `color.Danger`
- `color.Notice`
- `color.Success`
- `color.Comment`
- `color.Primary`
- `color.Warning`
- `color.Question`
- `color.Secondary`

```go
color.Info.Print("Info message")
color.Success.Print("Success message")
```

> run demo: `go run ./_examples/theme_style.go`

![theme-style](_examples/images/theme-style.jpg)

### Use like html tag

> **not** support on windows `cmd.exe`

```go
// use style tag
color.Print("<suc>he</><comment>llo</>, <cyan>wel</><red>come</>")
color.Println("<suc>hello</>")
color.Println("<error>hello</>")
color.Println("<warning>hello</>")

// custom color attributes
color.Print("<fg=yellow;bg=black;op=underscore;>hello, welcome</>\n")
```

- `color.Tag`

```go
// set a style tag
color.Tag("info").Print("info style text")
color.Tag("info").Printf("%s style text", "info")
color.Tag("info").Println("info style text")
```

> run demo: `go run ./_examples/colortag.go`

![color-tags](_examples/images/color-tags.jpg)

## Use 256 color

### Use foreground or background color
 
- `color.C256(val uint8, isBg ...bool) Color256`

```go
c := color.C256(132) // fg color
c.Println("message")
c.Printf("format %s", "message")

c := color.C256(132, true) // bg color
c.Println("message")
c.Printf("format %s", "message")
```

### Use 256 color style

> Can set foreground and background colors at the same time
 
- `color.S256(fgAndBg ...uint8) *Style256`

```go
s := color.S256(32, 203)
s.Println("message")
s.Printf("format %s", "message")
```

> run demo: `go run ./_examples/color256.go`

![color-tags](_examples/images/256-color.jpg)

## Use RGB color

### Use foreground or background color 

- `color.RGB(r, g, b uint8, isBg ...bool) RGBColor`

```go
c := color.RGB(30,144,255) // fg color
c.Println("message")
c.Printf("format %s", "message")

c := color.RGB(30,144,255, true) // bg color
c.Println("message")
c.Printf("format %s", "message")
```

- `color.HEX(hex string, isBg ...bool) RGBColor` Create from hexadecimal color string

```go
c := HEX("ccc") // 也可以写为: "cccccc" "#cccccc"
c.Println("message")
c.Printf("format %s", "message")

c = HEX("aabbcc", true) // as bg color
c.Println("message")
c.Printf("format %s", "message")
```

### Use RGB color style

> Can set foreground and background colors at the same time

- `color.NewRGBStyle(fg RGBColor, bg ...RGBColor) *RGBStyle`

```go
s := NewRGBStyle(RGB(20, 144, 234), RGB(234, 78, 23))
s.Println("message")
s.Printf("format %s", "message")
```

- `color.HEXStyle(fg string, bg ...string) *RGBStyle` Create from hexadecimal color string

```go
s := HEXStyle("11aa23", "eee")
s.Println("message")
s.Printf("format %s", "message")
```

## Refer

- `issue9/term` https://github.com/issue9/term
- `beego/bee` https://github.com/beego/bee
- `inhere/console` https://github/inhere/php-console
- [ANSI escape code](https://en.wikipedia.org/wiki/ANSI_escape_code)

## License

**MIT**
