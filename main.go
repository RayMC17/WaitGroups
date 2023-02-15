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
	reverse:= flag.Bool("r", false, "reverse the string (true /false)")

	flag.Parse()
	// check if it is should be uppercase before printing it.
	if *caps {
		*msg = strings.ToUpper(*msg)
	}
	//check if we should reverse the string
			
		//reverse string
		//step:
		//1. create an empty string var result string
		//2. Iterate over the string that you want to reverse for _, value := range *msg
	if *reverse{
		var result string
		for _, value := range *msg{
			result = string(value) + result
		}
		//write the reverse string to *msg
		*msg = result
	}

	//print the s string
	for i := 0; i < *num; i++ {
		fmt.Println(*msg)
	}

}
