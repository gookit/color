// The method in the file has no effect
// Only for compatibility with non-Windows systems

// +build !windows

package color

// winSet
func winSet(_ ...Color) (n int, err error) {
	return
}

// winReset
func winReset() (n int, err error) {
	return
}

// winPrint
func winPrint(_ string, _ ...Color) (n int, err error) {
	return
}

// winPrintln
func winPrintln(_ string, _ ...Color) (n int, err error) {
	return
}
