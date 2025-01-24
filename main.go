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
	fmt.Println("Length of your password (default: 32): ")
	// Ask user
	var i int
        _, err := fmt.Scanf("%d", &i) //Getting the user input
	if(err != nil){
		genpwd(alpha + other, 32) //Default to 32 if user input is nil
	}else{
		genpwd(alpha + other, i) //Using the user input
	}

}

// letter
func genpwd(charList string, sizePswd int) string {
	pssw := ""
	if (sizePswd == 0){
		return pssw
	}
	check := 0
	for true {
		var nb int = rand.Intn(len(charList)) + 0
		pssw += string(charList[nb])
		check++
		if check == sizePswd {
			fmt.Println(pssw)
			break
		}
	}
	return pssw
}
