package color_test

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
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

// https://github.com/gookit/color/issues/95
func TestIssues_95(t *testing.T) {
	// Test for race condition in Theme.Tips when called concurrently
	// We use a thread-safe buffer to avoid low-level I/O races
	// and focus on testing the application-level race condition
	buf := &safeBuffer{Buffer: &bytes.Buffer{}}
	color.SetOutput(buf)
	defer color.ResetOutput()

	const numGoroutines = 20
	var wg sync.WaitGroup
	wg.Add(numGoroutines * 2)

	// Start multiple goroutines calling Tips concurrently
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			color.Warn.Tips("warning message %d", id)
		}(i)

		go func(id int) {
			defer wg.Done()
			color.Error.Tips("error message %d", id)
		}(i)
	}

	wg.Wait()
	output := buf.String()

	// Clean the output of ANSI codes to check logical structure
	cleanOutput := color.ClearCode(output)
	lines := strings.Split(strings.TrimSpace(cleanOutput), "\n")
	
	// Verify that each line is properly formatted (without ANSI codes)
	for i, line := range lines {
		if line == "" {
			continue
		}
		// Each line should start with either "WARNING:" or "ERROR:"
		if !strings.HasPrefix(line, "WARNING:") && !strings.HasPrefix(line, "ERROR:") {
			t.Errorf("Line %d is malformed due to race condition: %q", i, line)
		}
		// Should not contain both prefixes (indicating interleaved output)
		if strings.Contains(line, "WARNING:") && strings.Contains(line, "ERROR:") {
			t.Errorf("Line %d contains interleaved output: %q", i, line)
		}
	}
}

// Thread-safe buffer for testing
type safeBuffer struct {
	*bytes.Buffer
	mu sync.Mutex
}

func (sb *safeBuffer) Write(p []byte) (n int, err error) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	return sb.Buffer.Write(p)
}
