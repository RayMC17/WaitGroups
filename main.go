package main

import (
	"flag"
	"fmt"
	"strings"
)

//domnstrate flags

func main(){
	msg := flag.String("msg", "Howdy, stranger!", "the message")
	num:= flag.Int("num", 1 ,"How many times to print the message")
	caps:= flag.Int("num", 1 ,"Should print in all caps")
	
	flag.Parse()

	for i := 0; i < *num; i++{
		fmt.Println(*msg)
	}


	}
