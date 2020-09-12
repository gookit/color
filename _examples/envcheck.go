package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gookit/color"
	"github.com/gookit/goutil/dump"
)

func main() {
	fmt.Println("OS", runtime.GOOS)

	fmt.Println("IsSupport256Color", color.IsSupport256Color())
	fmt.Println("IsSupportColor", color.IsSupportColor())

	dump.P(os.Environ())
}
