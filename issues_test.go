package color

import (
	"fmt"
	"strings"
	"testing"
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

	fmt.Printf("topBar:\n%q\n%q\n", topBar, ClearCode(topBar))
	fmt.Printf("title:\n%q\n%q\n", title, ClearCode(title))
	fmt.Printf("Split:\n%#v\n", strings.Split(ClearCode(topBar), ClearCode(title)))
}
