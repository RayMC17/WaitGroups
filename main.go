package main

import (
	"flag"
	"fmt"
	"strings"
)

//domnstrate flags

func main() {
	msg := flag.String("msg", "Howdy, stranger!", "the message")
	num := flag.Int("num", 1, "How many times to print the message")
	caps := flag.Bool("U", false, "Specify whether to uppercase the string(true/false)")

	flag.Parse()
	// check if it is should be uppercase before printing it.
	if *caps {
		*msg = strings.ToUpper(*msg)
	}
	//print the s string
	for i := 0; i < *num; i++ {
		fmt.Println(*msg)
	}

}
