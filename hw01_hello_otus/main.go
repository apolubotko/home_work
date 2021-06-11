package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	helloStr := "Hello, OTUS!"
	fmt.Println(stringutil.Reverse(helloStr))
}
