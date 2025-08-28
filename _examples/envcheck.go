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

	if es := color.InnerErrs(); len(es) > 0 {
		fmt.Println("inner errors:", es)
	}

	termVal := os.Getenv("TERM")
	fmt.Println("ENV TERM:", termVal)
	// fmt.Println(
	// 	termVal[0:1],
	// 	strconv.FormatUint(uint64(termVal[0]), 16),
	// )

	// fmt.Println("os.Environ:")
	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }
}
