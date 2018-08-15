# color

golang下的命令行色彩使用库

**[EN Readme](README.md)**

## 功能特色

- 使用简单方便
- 支持丰富的颜色输出
- 同时支持html标签式的颜色渲染
- 兼容Windows

## 获取安装

- 使用 dep 包管理

```bash
dep ensure -add gopkg.in/gookit/color.v1 // 推荐
// OR
dep ensure -add github.com/gookit/color
```

- 使用 go get

```bash
go get gopkg.in/gookit/color.v1 // 推荐
// OR
go get -u github.com/gookit/color
```

- git 克隆

```bash
git clone https://github.com/gookit/color
```

## Godoc

- [godoc for gopkg](https://godoc.org/gopkg.in/gookit/color.v1)
- [godoc for github](https://godoc.org/github.com/gookit/color)

## 快速开始

如下，引入当前包就可以快速的使用

```bash
import "gopkg.in/gookit/color.v1" // 推荐
// or
import "github.com/gookit/color"
```

### 如何使用

```go
package main

import (
	"fmt"
    "github.com/gookit/color"
)

func main() {
	// simple usage
	color.FgCyan.Printf("Simple to use %s\n", "color")

    // use like func
    red := color.FgRed.Render
    green := color.FgGreen.Render
    fmt.Printf("%s line %s library\n", red("Command"), green("color"))

	// custom color
	color.New(color.FgWhite, color.BgBlack).Println("custom color style")

	// can also:
	color.Style{color.FgCyan, color.OpBold}.Println("custom color style")
	
	// use style tag
	color.Print("<suc>he</><comment>llo</>, <cyan>wel</><red>come</>\n")

	// set a style tag
	color.Tag("info").Println("info style text")

	// use info style tips
	color.Tips("info").Print("tips style text")

	// use info style blocked tips
	color.LiteTips("info").Print("blocked tips style text")
}
```

运行 demo: `go run ./demo/app.go`

### 颜色输出展示

![colored-out](demo/images/colored-out.jpg)

### 构建风格

```go
// 仅设置前景色
color.FgCyan.Printf("Simple to use %s\n", "color")
// 仅设置背景色
color.BgRed.Printf("Simple to use %s\n", "color")

// 完全自定义 前景色 背景色 选项
style := color.New(color.FgWhite, color.BgBlack, color.OpBold)
style.Println("custom color style")

// can also:
color.Style{color.FgCyan, color.OpBold}.Println("custom color style")
```

```go
// 设置console颜色
color.Set(color.FgCyan)

// 输出信息
fmt.Print("message")

// 重置console颜色
color.Reset()
```

### 使用内置风格

#### 基础颜色方法

> 支持在windows `cmd.exe` 使用

- `color.Bold(args ...interface{})`
- `color.Black(args ...interface{})`
- `color.White(args ...interface{})`
- `color.Gray(args ...interface{})`
- `color.Red(args ...interface{})`
- `color.Green(args ...interface{})`
- `color.Yellow(args ...interface{})`
- `color.Blue(args ...interface{})`
- `color.Magenta(args ...interface{})`
- `color.Cyan(args ...interface{})`

```go
color.Bold("bold message")
color.Yellow("yellow message")
```

#### 扩展风格方法 

> 支持在windows `cmd.exe` 使用

- `color.Info(args ...interface{})`
- `color.Note(args ...interface{})`
- `color.Light(args ...interface{})`
- `color.Error(args ...interface{})`
- `color.Danger(args ...interface{})`
- `color.Notice(args ...interface{})`
- `color.Success(args ...interface{})`
- `color.Comment(args ...interface{})`
- `color.Primary(args ...interface{})`
- `color.Warning(args ...interface{})`
- `color.Question(args ...interface{})`
- `color.Secondary(args ...interface{})`

```go
color.Info("Info message")
color.Success("Success message")
```

#### 使用颜色html标签

> **不** 支持在windows `cmd.exe` 使用，但不影响使用，会自动去除颜色标签

使用颜色标签可以非常方便简单的构建自己需要的任何格式

```go
// 使用内置的 color tag
color.Print("<suc>he</><comment>llo</>, <cyan>wel</><red>come</>")
color.Println("<suc>hello</>")
color.Println("<error>hello</>")
color.Println("<warning>hello</>")

// 自定义颜色属性
color.Print("<fg=yellow;bg=black;op=underscore;>hello, welcome</>\n")
```

#### 使用 `color.Tag`

给后面输出的文本信息加上给定的颜色风格标签

```go
// set a style tag
color.Tag("info").Print("info style text")
color.Tag("info").Printf("%s style text", "info")
color.Tag("info").Println("info style text")
```

### 内置的标签

这里列出了内置的标签，基本上涵盖了各种风格和颜色搭配。它们都可用作颜色html标签，或者作为 `color.Tag` `color.Tips` 等的参数

```text
// Some internal defined style tags
// usage: <tag>content text</>

// basic tags
- red
- blue
- cyan
- black
- green
- brown
- white
- default  // no color
- normal// no color
- yellow  
- magenta 

// alert tags like bootstrap's alert
- suc // same "green" and "bold"
- success 
- info // same "green"
- comment  // same "brown"
- note 
- notice  
- warn
- warning 
- primary 
- danger // same "red"
- err 
- error

// more tags
- lightRed
- light_red
- lightGreen
- light_green
- lightBlue 
- light_blue
- lightCyan
- light_cyan
- lightDray
- light_gray
- gray
- darkGray
- dark_gray
- lightYellow
- light_yellow  
- lightMagenta  
- light_magenta 

// extra
- lightRedEx
- light_red_ex
- lightGreenEx
- light_green_ex 
- lightBlueEx
- light_blue_ex  
- lightCyanEx
- light_cyan_ex  
- whiteEx
- white_ex

// option
- bold
- underscore 
- reverse
```

## 参考项目

- `issue9/term` https://github.com/issue9/term
- `beego/bee` https://github.com/beego/bee
- `inhere/console` https://github/inhere/php-console
- [ANSI转义序列](https://zh.wikipedia.org/wiki/ANSI转义序列)
- [Standard ANSI color map](https://conemu.github.io/en/AnsiEscapeCodes.html#Standard_ANSI_color_map)

## License

MIT
