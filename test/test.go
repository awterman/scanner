package main

import (
	"fmt"

	"github.com/awterman/scanner"
)

func main() {
	s := scanner.OpenFile("test.go")
	for s.Scan() {
		fmt.Println(s.Text())
	}
	fmt.Println(s.Err())
}
