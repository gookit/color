package main

import "github.com/gookit/color"

// run: go run ./_examples/demo_tag.go
func main() {
	text := `
  <mga1>gookit/color:</>
     A <green>command-line</> 
     <cyan>color library</> with <fg=167;bg=232>256-color</>
     and <fg=11aa23;op=bold>True-color</> support,
     <fg=mga;op=i>universal API</> methods
     and <cyan>Windows</> support.
`
	color.Print(text)
}
