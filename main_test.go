package utf8len_test

import (
	"fmt"
	"testing"
	"unicode/utf8"

	"github.com/hymkor/go-utf8len"
)

func TestFromFirstByte(t *testing.T) {
	for c := '\u0020'; c < '\U0010FFFF'; c++ {
		source := fmt.Sprintf("%c", c)
		_, expect := utf8.DecodeRuneInString(source)
		result := utf8len.FromFirstByte(source[0])
		if expect != result {
			t.Fatalf("expect %v,but %v for %v", expect, result, source)
		}
		// println(source,expect,result)
	}
}

func BenchmarkFromFirstByte(b *testing.B) {
	for c := 0; c < b.N; c++ {
		source := fmt.Sprintf("%c", c)
		utf8len.FromFirstByte(source[0])
	}
}

func BenchmarkDecodeRuneInString(b *testing.B) {
	for c := 0; c < b.N; c++ {
		source := fmt.Sprintf("%c", c)
		_, _ = utf8.DecodeRuneInString(source)
	}
}
