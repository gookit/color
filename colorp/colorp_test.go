package colorp_test

import (
	"testing"

	"github.com/gookit/color/colorp"
)

func TestColorPrint(t *testing.T) {
	// code gen by: kite gen parse colorp/_demo/gen-code.tpl
	colorp.Redp("p:red color message, ")
	colorp.Redf("f:%s color message, ", "red")
	colorp.Redln("ln:red color message print in cli.")
	colorp.Bluep("p:blue color message, ")
	colorp.Bluef("f:%s color message, ", "blue")
	colorp.Blueln("ln:blue color message print in cli.")
	colorp.Cyanp("p:cyan color message, ")
	colorp.Cyanf("f:%s color message, ", "cyan")
	colorp.Cyanln("ln:cyan color message print in cli.")
	colorp.Grayp("p:gray color message, ")
	colorp.Grayf("f:%s color message, ", "gray")
	colorp.Grayln("ln:gray color message print in cli.")
	colorp.Greenp("p:green color message, ")
	colorp.Greenf("f:%s color message, ", "green")
	colorp.Greenln("ln:green color message print in cli.")
	colorp.Yellowp("p:yellow color message, ")
	colorp.Yellowf("f:%s color message, ", "yellow")
	colorp.Yellowln("ln:yellow color message print in cli.")
	colorp.Magentap("p:magenta color message, ")
	colorp.Magentaf("f:%s color message, ", "magenta")
	colorp.Magentaln("ln:magenta color message print in cli.")

	colorp.Infop("p:info color message, ")
	colorp.Infof("f:%s color message, ", "info")
	colorp.Infoln("ln:info color message print in cli.")
	colorp.Successp("p:success color message, ")
	colorp.Successf("f:%s color message, ", "success")
	colorp.Successln("ln:success color message print in cli.")
	colorp.Warnp("p:warn color message, ")
	colorp.Warnf("f:%s color message, ", "warn")
	colorp.Warnln("ln:warn color message print in cli.")
	colorp.Errorp("p:error color message, ")
	colorp.Errorf("f:%s color message, ", "error")
	colorp.Errorln("ln:error color message print in cli.")
}
