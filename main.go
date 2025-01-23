package main

/*

small program to learn Golang

*/

import (
	"fmt"
	"math/rand"
)

var alpha string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
var other string = "$!/()-£!?={}*#%&0123456789"

func main() {
	fmt.Println("--- PSSWD GENERATOR ---")
	// List of char

	// Ask user
	var charByDefault bool = true
	if charByDefault {
		genpwd(alpha + other)
	}

}

// letter
func genpwd(charList string) string {
	pssw := ""
	var sizePswd int = 32
	check := 0
	for true {
		var nb int = rand.Intn(len(charList)) + 0
		pssw += string(charList[nb])
		check++
		if check == sizePswd {
			fmt.Print(pssw)
			break
		}
	}
	return pssw
}
