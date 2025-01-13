package utf8len_test

import (
	"fmt"
	"testing"
	"unicode/utf8"

	"github.com/hymkor/go-utf8len"
)

var source = [0x110000]string{}

func init() {
	//println("setup test table")
	for c := range source {
		source[c] = fmt.Sprintf("%c", c)
	}
	//println("done")
}

func TestFromFirstByte(t *testing.T) {
	for _, s := range source {
		_, expect := utf8.DecodeRuneInString(s)
		result := utf8len.FromFirstByte(s[0])
		if expect != result {
			t.Fatalf("expect %v,but %v for %v", expect, result, source)
		}
		// println(source,expect,result)
	}
}

func BenchmarkFromFirstByte(b *testing.B) {
	for c := 0; c < b.N; c++ {
		utf8len.FromFirstByte(source[c%len(source)][0])
	}
}

func BenchmarkDecodeRuneInString(b *testing.B) {
	for c := 0; c < b.N; c++ {
		_, _ = utf8.DecodeRuneInString(source[c%len(source)])
	}
}
