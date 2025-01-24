package main

/*

small program to learn Golang

*/

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var alpha string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
var other string = "$!/()-£!?={}*#%+&"
var number string = "0123456789"

func main() {
	fmt.Println("--- PSSWD GENERATOR ---")
	// Ask user

	wierdSymbolBool, numberBool := askUser()

	if wierdSymbolBool {
		alpha += other
	}
	if numberBool {
		alpha += number
	}
	genpwd(alpha, 32)

}

// letter
func genpwd(charList string, sizearg int) string {
	pssw := ""
	//var sizePswd int = 32
	check := 0
	for true {
		var nb int = rand.Intn(len(charList)) + 0
		pssw += string(charList[nb])
		check++
		if check == sizearg {
			fmt.Print(pssw)
			break
		}
	}
	return pssw
}

func askUser() (bool, bool) {
	var wierdBool bool
	var numberBool bool
	scan := bufio.NewScanner(os.Stdin)
	//ask $£!
	for {
		fmt.Print("Wierd symbol ? -> [Y/N] : ")
		scan.Scan()
		if strings.EqualFold(strings.ToLower(scan.Text()), "y") {
			wierdBool = true
			break
		} else if strings.EqualFold(strings.ToLower(scan.Text()), "n") {
			break
		} else {
			fmt.Println(" --- Wrong --- ")
			continue
		}
	}
	//ask number
	for {
		fmt.Print("Number? -> [Y/N] : ")
		scan.Scan()
		if strings.EqualFold(strings.ToLower(scan.Text()), "y") {
			numberBool = true
			break
		} else if strings.EqualFold(strings.ToLower(scan.Text()), "n") {
			break
		} else {
			fmt.Println(" --- Wrong --- ")
			continue
		}
	}
	return wierdBool, numberBool
}
