// The method in the file has no effect
// Only for compatibility with non-Windows systems

// +build !windows

package color

// winSet
func winSet(colors ...Color) (n int, err error) {
	return
}

// winReset
func winReset() (n int, err error) {
	return
}

// winPrint
func winPrint(str string, colors ...Color) (n int, err error) {
	return
}

// winPrintln
func winPrintln(str string, colors ...Color) (n int, err error) {
	return
}
