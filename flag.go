package main

import (
	"flag"
)

var flagLength uint

func init() {
	flag.UintVar(&flagLength, "length", 32, "length of generated password")

	flag.Parse()
}
