package main

import (
	"fmt"
	"math/rand"
)

var alpha string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
var other string = "$!/()-_éà£!?={}*#%&0123456789"

func main() {
	fmt.Println("--- PSSWD GENERATOR ---")
	genpwd()
}

// letter
func genpwd() string {
	pssw := ""
	var size_pswd int = 16
	check := 0
	for true {
		var nb int = rand.Intn(len(alpha)) + 0
		pssw += string(alpha[nb])
		check++
		if check == size_pswd {
			fmt.Print(pssw)
			break
		}
	}
	return pssw
}

func add() {

}
