# usage: kite gen parse testdata/gen-printf.tpl
names =[Red, Blue, Cyan, Gray, Green, Yellow, Magenta]

###

{{ foreach ($names as $var): }}

// {{ $var }}f print message with {{ $var }} color
func {{ $var }}f(format string, a ...interface{}) {
	{{ $var }}.Printf(format, a...)
}

{{ endforeach}}
