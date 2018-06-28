package color

import (
	"io"
	"fmt"
	"os"
)

// Print
func Print(args ...interface{}) (int, error) {
	return fmt.Fprint(os.Stdout, Render(args...))
}

// Printf
func Printf(format string, args ...interface{}) (int, error) {
	return fmt.Fprint(os.Stdout, Render(fmt.Sprintf(format, args...)))
}

// Println
func Println(args ...interface{}) (int, error) {
	return fmt.Fprintln(os.Stdout, Render(args...))
}

// Fprint
func Fprint(w io.Writer, args ...interface{}) (int, error) {
	return fmt.Fprint(w, Render(args...))
}

// Fprintf
func Fprintf(w io.Writer, format string, args ...interface{}) (int, error) {
	return fmt.Fprint(w, Render(fmt.Sprintf(format, args...)))
}

// Fprintln
func Fprintln(w io.Writer, args ...interface{}) (int, error) {
	return fmt.Fprintln(w, Render(args...))
}
