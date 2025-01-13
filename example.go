//go:build run

package main

import (
	"fmt"
	"github.com/hymkor/go-utf8len"
)

func main() {
	fmt.Printf("A=%d\n", utf8len.FromFirstByte('A'))
}
