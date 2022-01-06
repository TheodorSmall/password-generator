package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func generate(length uint) {
	if length == 0 {
		fmt.Println("cannot generate empty password")
		return
	}

	valid := "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	vlen := big.NewInt(int64(len(valid)))
	passwd := make([]byte, length)

	for i := range passwd {
		vi, err := rand.Int(rand.Reader, vlen)
		if err != nil {
			fmt.Printf("rand.Int(...) failed: %s", err)
		}

		passwd[i] = valid[vi.Int64()]
	}

	fmt.Println(string(passwd))
}
