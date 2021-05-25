package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	helloStr := "Hello, OTUS!"
	fmt.Printf("%s", stringutil.Reverse(helloStr))
}
