/*
    The package is used to show gofmt & godoc.
    gofmt:
        -l -w xx.go
    godoc:
        -o
*/
package main

import (
	"fmt"
)

type T struct {
	name  string // name of obj
	value int    // it's value
}

func main() {
	t := Y{name: "gofmt", value: 2014}
	fmt.Printf("T: %v\n", t)
}
