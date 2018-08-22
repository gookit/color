package color

import (
	"fmt"
	"strings"
)

// Tag value is a defined style name
type Tag string

// Print messages
func (tg Tag) Print(args ...interface{}) {
	name := string(tg)
	if stl := GetStyle(name); !stl.IsEmpty() {
		stl.Print(args...)
		return
	}

	str := RenderCode(GetTagCode(name), args...)
	fmt.Print(str)
}

// Printf format and print messages
func (tg Tag) Printf(format string, args ...interface{}) {
	name := string(tg)
	if stl := GetStyle(name); !stl.IsEmpty() {
		stl.Printf(format, args...)
		return
	}

	str := RenderCode(GetTagCode(name), fmt.Sprintf(format, args...))
	fmt.Print(str)
}

// Println messages line
func (tg Tag) Println(args ...interface{}) {
	name := string(tg)
	if stl := GetStyle(name); !stl.IsEmpty() {
		stl.Println(args...)
		return
	}

	str := RenderCode(GetTagCode(name), args...)
	fmt.Println(str)
}

// Sprint render messages
func (tg Tag) Sprint(args ...interface{}) string {
	name := string(tg)
	if stl := GetStyle(name); !stl.IsEmpty() {
		return stl.Render(args...)
	}

	return RenderCode(GetTagCode(name), args...)
}

// Tips will add color for all text
// value is a defined style name
type Tips string

// Print messages
func (t Tips) Print(args ...interface{}) (int, error) {
	name := string(t)
	upName := strings.ToUpper(name)

	if isLikeInCmd {
		return GetStyle(name).Println(upName, ": ", fmt.Sprint(args...))
	}

	str := RenderCode(GetTagCode(name), upName, ": ", fmt.Sprint(args...))
	return fmt.Println(str)
}

// Println messages line
func (t Tips) Println(args ...interface{}) (int, error) {
	return t.Print(args...)
}

// Printf format and print messages
func (t Tips) Printf(format string, args ...interface{}) (int, error) {
	name := string(t)
	upName := strings.ToUpper(name)

	if isLikeInCmd {
		return GetStyle(name).Println(upName, ": ", fmt.Sprintf(format, args...))
	}

	str := RenderCode(GetTagCode(name), upName, ": ", fmt.Sprintf(format, args...))
	return fmt.Println(str)
}

// LiteTips will only add color for tag name
// value is a defined style name
type LiteTips string

// Print messages
func (t LiteTips) Print(args ...interface{}) (int, error) {
	tag := string(t)

	if isLikeInCmd {
		GetStyle(tag).Print(strings.ToUpper(tag), ": ")
		return fmt.Println(args...)
	}

	str := RenderCode(GetTagCode(tag), strings.ToUpper(tag), ":")
	return fmt.Println(str, fmt.Sprint(args...))
}

// Println messages line
func (t LiteTips) Println(args ...interface{}) (int, error) {
	return t.Print(args...)
}

// Printf format and print messages
func (t LiteTips) Printf(format string, args ...interface{}) (int, error) {
	tag := string(t)
	if isLikeInCmd {
		GetStyle(tag).Print(strings.ToUpper(tag), ": ")
		return fmt.Printf(format+"\n", args...)
	}

	str := RenderCode(GetTagCode(tag), strings.ToUpper(tag), ":")
	return fmt.Println(str, fmt.Sprintf(format, args...))
}
