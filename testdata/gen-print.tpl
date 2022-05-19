# usage: kite gen parse tmp/gen-print.tpl
names=[Info, Error, Warn]

###

{{ foreach ($names as $var): }}

// {{ $var }}p print message with {{ $var }} color
func {{ $var }}p(a ...interface{}) {
	{{ $var }}.Print(a...)
}

{{ endforeach}}
