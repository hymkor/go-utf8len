utf8len.FromFirstByte
=====================

```go
func FromFirstByte(c byte) int
```

A function that calculates the byte length of the first UTF-8 character using only the first byte.

This approach is slightly faster than Go's standard utf8.DecodeRuneInString, as it does not compute the Unicode code point of the character.

If an invalid value is provided as the first byte of a UTF-8 sequence, the function returns zero.

**main_test.go**

```main_test.go
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
```

**go test -bench . -benchmem**

```make bench |
go test -bench . -benchmem
goos: windows
goarch: amd64
pkg: github.com/hymkor/go-utf8len
cpu: Intel(R) Core(TM) i5-6500T CPU @ 2.50GHz
BenchmarkFromFirstByte-4        	11031307	       102.1 ns/op	      16 B/op	       1 allocs/op
BenchmarkDecodeRuneInString-4   	10946996	       103.8 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	github.com/hymkor/go-utf8len	2.808s
```
