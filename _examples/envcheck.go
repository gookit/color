package main

import (
	"fmt"
	"runtime"

	"github.com/gookit/color"
	// "github.com/gookit/goutil/dump"
)

// go run ./_examples/envcheck.go
// COLOR_DEBUG_MODE=on go run ./_examples/envcheck.go
func main() {
	fmt.Println("current OS:", runtime.GOOS)

	fmt.Println("Terminal Color Level:", color.TermColorLevel())
	fmt.Println("Support Basic Color:", color.SupportColor())
	fmt.Println("Support 256 Color:", color.Support256Color())
	fmt.Println("Support True Color:", color.SupportTrueColor())

	if es := color.InnerErrs(); len(es) > 0 {
		fmt.Println(es)
	}

	// dump.P(os.Environ())
	// fmt.Println(os.Environ())
}
