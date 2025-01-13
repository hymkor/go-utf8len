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
            t.Fatalf("expect %#v,but %#v for %#v", expect, result, s)
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
```

**go test -bench . -benchmem**

```make bench |
go test -bench . -benchmem
goos: windows
goarch: amd64
pkg: github.com/hymkor/go-utf8len
cpu: Intel(R) Core(TM) i5-6500T CPU @ 2.50GHz
BenchmarkFromFirstByte-4        	378031695	         3.256 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecodeRuneInString-4   	138281330	         8.378 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/hymkor/go-utf8len	4.058s
```
