package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gookit/color"
	// "github.com/gookit/goutil/dump"
)

// go run ./_examples/envcheck.go
//
// bash:
// 	COLOR_DEBUG_MODE=on go run ./_examples/envcheck.go
// cmd.exe:
//   set COLOR_DEBUG_MODE=on
//   go run ./_examples/envcheck.go
func main() {
	fmt.Println("current OS:", runtime.GOOS)

	fmt.Println("Terminal Color Level:", color.TermColorLevel())
	fmt.Println("Support Basic Color:", color.SupportColor())
	fmt.Println("Support 256 Color:", color.Support256Color())
	fmt.Println("Support True Color:", color.SupportTrueColor())

	fmt.Println("------- Re-Detect by Enable Debug Mode -------")
	color.EnableDebug()
	fmt.Println("Detected Color Level:", color.DetectColorLevel())

	if es := color.InnerErrs(); len(es) > 0 {
		fmt.Println("inner errors:", es)
	}

	termVal := os.Getenv("TERM")
	fmt.Println("ENV TERM:", termVal)
	// fmt.Println(
	// 	termVal[0:1],
	// 	strconv.FormatUint(uint64(termVal[0]), 16),
	// )

	fmt.Println("------- Test Color Output -------")
	fmt.Println("\x1b[34mHello \x1b[35mWorld\x1b[0m!")
	_, _ = fmt.Fprintf(os.Stdout, "\x1b[34mHello \x1b[35mWorld\x1b[0m!\n")
	// fmt.Println("os.Environ:")
	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }
}
