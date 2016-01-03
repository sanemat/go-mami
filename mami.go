package main

import (
	"fmt"
)

func mami() string {
	return "foo/123.txt"
}

func main() {
	output := mami()
	fmt.Println(output)
}
