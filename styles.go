package color

// some defined style tags <tag>some text</>
const (
	// basic
	Red     = "red"
	Blue    = "blue"
	Cyan    = "cyan"
	Black   = "black"
	Green   = "green"
	Brown   = "brown"
	White   = "white"
	Normal  = "normal" // no color
	Yellow  = "yellow"
	Magenta = "magenta"

	// alert
	Suc     = "suc" // same "green" and "bold"
	Success = "success"
	Info    = "info"    // same "green"
	Comment = "comment" // same "brown"
	Note    = "note"
	Notice  = "notice"
	Warn    = "warn"
	Warning = "warning"
	Danger  = "danger" // same "red"
	Err     = "err"
	Error   = "error"

	// option
	Bold       = "bold"
	Underscore = "underscore"
	Reverse    = "reverse"
)

/**
 * Some internal defined styles
 * custom style: fg;bg;opt
 */
var Styles = map[string]string{
	// basic
	"red":     "0;31",
	"blue":    "0;34",
	"cyan":    "0;36",
	"black":   "0;30",
	"green":   "0;32",
	"brown":   "0;33",
	"white":   "1;37",
	"default": "39", // no color
	"normal":  "39", // no color
	"yellow":  "1;33",
	"magenta": "1;35",

	// alert
	"suc":     "1;32", // same "green" and "bold"
	"success": "1;32",
	"info":    "0;32", // same "green",
	"comment": "0;33", // same "brown"
	"note":    "36;1",
	"notice":  "36;4",
	"warn":    "0;30;43",
	"warning": "0;30;43",
	"danger":  "0;31", // same "red"
	"err":     "30;41",
	"error":   "30;41",

	// more
	"lightRed":      "1;31",
	"light_red":     "1;31",
	"lightGreen":    "1;32",
	"light_green":   "1;32",
	"lightBlue":     "1;34",
	"light_blue":    "1;34",
	"lightCyan":     "1;36",
	"light_cyan":    "1;36",
	"lightDray":     "37",
	"light_gray":    "37",
	"darkDray":      "90",
	"dark_gray":     "90",
	"lightYellow":   "93",
	"light_yellow":  "93",
	"lightMagenta":  "95",
	"light_magenta": "95",

	// extra
	"lightRedEx":     "91",
	"light_red_ex":   "91",
	"lightGreenEx":   "92",
	"light_green_ex": "92",
	"lightBlueEx":    "94",
	"light_blue_ex":  "94",
	"lightCyanEx":    "96",
	"light_cyan_ex":  "96",
	"whiteEx":        "97",
	"white_ex":       "97",

	// option
	"bold":       "1",
	"underscore": "4",
	"reverse":    "7",
}
