package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	expression := flag.Arg(0)

	fmt.Println(expression)
}
