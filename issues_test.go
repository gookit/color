package color_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gookit/color"
)

// https://github.com/gookit/color/issues/51
func TestIssues_51(t *testing.T) {
	topBarRs := []rune{
		9484, 32, 66, 111, 120, 32, 32, 32, 32, 32, 67, 76, 73, 32, 32, 32, 32, 32, 77, 97, 107, 101, 114, 32, 32, 32, 128230, 32, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472,
		9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9488,
	}

	topBar := string(topBarRs)

	titleRs := []rune{
		66, 111, 120, 32, 32, 32, 67, 76, 73, 32, 32, 32, 32, 32, 77, 97, 107, 101, 114, 32, 32, 32, 128230,
	}
	title := string(titleRs)

	fmt.Printf("topBar:\n%q\n%q\n", topBar, color.ClearCode(topBar))
	fmt.Printf("title:\n%q\n%q\n", title, color.ClearCode(title))
	fmt.Printf("Split:\n%#v\n", strings.Split(color.ClearCode(topBar), color.ClearCode(title)))
}

// https://github.com/gookit/color/issues/52
func TestIssues_52(t *testing.T) {
	test1 := `FAILS:
one <bg=lightGreen;fg=black>two</> <three>
foo <bg=lightGreen;fg=black>two</> <four>
`

	test2 := `WORKS:
one <bg=lightGreen;fg=black>two <three></>
foo <bg=lightGreen;fg=black>two <four></>
`

	test3 := `WORKS:
one <bg=lightGreen;fg=black>two</> three
foo <bg=lightGreen;fg=black>two</> four
`

	// colorp.Info
	color.Print(test1)
	color.Print(test2)
	color.Print(test3)
}
