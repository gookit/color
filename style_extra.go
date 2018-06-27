package color

import (
	"fmt"
	"strings"
)

// value is a defined style name
type tagName string

// Tag more please Styles
func Tag(name string) *tagName {
	if !IsStyle(name) {
		panic("unknown style name: " + name)
	}

	tg := tagName(name)
	return &tg
}

// Print
func (tg tagName) Print(args ...interface{})  {
	str := buildColoredText(
		GetStyleCode(string(tg)),
		fmt.Sprint(args...),
	)

	fmt.Print(str)
}

// Println
func (tg tagName) Println(args ...interface{})  {
	str := buildColoredText(
		GetStyleCode(string(tg)),
		fmt.Sprint(args...),
	)

	fmt.Println(str)
}

// value is a defined style name
type Tips string

// Print
func (t Tips) Print(args ...interface{})  {
	tag := string(t)
	str := buildColoredText(
		GetStyleCode(tag),
		strings.ToUpper(tag) + ": ",
	)

	fmt.Print(str, fmt.Sprint(args...))
}

// Printf
func (t Tips) Printf(format string, args ...interface{})  {
	tag := string(t)
	str := buildColoredText(
		GetStyleCode(tag),
		strings.ToUpper(tag) + ": ",
	)

	fmt.Print(str, fmt.Sprintf(format, args...))
}

// value is a defined style name
type BlockTips string

// Print
func (t BlockTips) Print(args ...interface{})  {
	tag := string(t)
	str := buildColoredText(
		GetStyleCode(tag),
		strings.ToUpper(tag) + ": " + fmt.Sprint(args...),
	)

	fmt.Print(str)
}

// Printf
func (t BlockTips) Printf(format string, args ...interface{})  {
	tag := string(t)
	str := buildColoredText(
		GetStyleCode(tag),
		strings.ToUpper(tag) + ": " + fmt.Sprintf(format, args...),
	)

	fmt.Print(str)
}

// Logger console logger
type Logger struct {
	style string
	fields map[string]string
}

// log level to cli color style
var LogLevel2tag = map[string]string{
	"info":    "info",
	"warn":    "warning",
	"warning": "warning",
	"debug":   "cyan",
	"notice":  "notice",
	"error":   "error",
}

func NewLog(fields map[string]string) *Logger {
	return &Logger{Info, fields}
}

func (l *Logger) Info(args ...interface{}) {

}

func (l *Logger) Log(args ...interface{}) {

}
