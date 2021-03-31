package color

import "fmt"

/*************************************************************
 * colored message Printer
 *************************************************************/

// PrinterFace interface
type PrinterFace interface {
	fmt.Stringer
	Sprint(a ...interface{}) string
	Sprintf(format string, a ...interface{}) string
	Print(a ...interface{})
	Printf(format string, a ...interface{})
	Println(a ...interface{})
}

// Printer a generic color message printer.
// Usage:
// 	p := &Printer{"32;45;3"}
// 	p.Print("message")
type Printer struct {
	// TODO NoColor
	// NoColor bool
	// ColorCode color code string. eg "32;45;3"
	ColorCode string
}

// NewPrinter instance
func NewPrinter(colorCode string) *Printer {
	return &Printer{colorCode}
}

// String returns color code string. eg: "32;45;3"
func (p *Printer) String() string {
	// panic("implement me")
	return p.ColorCode
}

// Sprint returns rendering colored messages
func (p *Printer) Sprint(a ...interface{}) string {
	return RenderCode(p.String(), a...)
}

// Sprintf returns format and rendering colored messages
func (p *Printer) Sprintf(format string, a ...interface{}) string {
	return RenderString(p.String(), fmt.Sprintf(format, a...))
}

// Print rendering colored messages
func (p *Printer) Print(a ...interface{}) {
	doPrintV2(p.String(), fmt.Sprint(a...))
}

// Printf format and rendering colored messages
func (p *Printer) Printf(format string, a ...interface{}) {
	doPrintV2(p.String(), fmt.Sprintf(format, a...))
}

// Println rendering colored messages with newline
func (p *Printer) Println(a ...interface{}) {
	doPrintlnV2(p.ColorCode, a)
}

// IsEmpty color code
func (p *Printer) IsEmpty() bool {
	return p.ColorCode == ""
}
