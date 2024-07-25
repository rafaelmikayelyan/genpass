package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const lowercase = "abcdefghijklmnopqrstuvwxyz"
const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const alphanumericReadibility = "ABCDEFGHIJKLMNPQRSTUVWXYZ0123456789" 
const alphanumericReatibilityFull = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ0123456789" 
const symbols = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
const reduced = "!@#$%^&*"
const defaultPswLength = 16
const defaultCharSet = lowercase + uppercase + numbers + symbols
const version = "1.0"
const helpMSG = "\nusage: genpass [-lnrsu] [int]\n\noptional flags:\n-l : use lowercase\n-u : use uppercase\n-n : use numbers\n-s : use symbols\n-r : use redused set of symbols : !@#$%^&*\n"
const errorMSG = "\nError\n" + helpMSG

func main() {
	charSet := defaultCharSet
	pswLength := defaultPswLength
	argLength := len(os.Args)

	if argLength > 3 {
		fmt.Println(errorMSG)
		os.Exit(0)
	} else if argLength > 1 {
		for i, v := range os.Args {
			if i == 0 {
				continue
			}

			num, err := strconv.Atoi(v)
			if err == nil {
				pswLength = num
			} else {
				charSet = evaluateFlags(v)
			}
		}

	}

	generatePassword(pswLength, charSet)
}

func generatePassword(pswLength int, charSet string) {
	var password string
	for i := 0; i < pswLength; i++ {
		randomNum := getRandInt(len(charSet))
		password += charSet[randomNum:randomNum + 1]
	}
	fmt.Println(password)
}

func getRandInt(max int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}
	return int(nBig.Int64())
}

func evaluateFlags(flags string) string {
	charSet := ""

	flagSet := map[string]string {
		"u": uppercase,
		"l": lowercase,
		"n": numbers,
		"s": symbols,
		"r": reduced,
	}

	if len(flags) < 2 || flags[0:1] != "-" {
		fmt.Println(errorMSG)
		os.Exit(0)
	}
	for i := 1; i < len(flags); i++ {
		flag, exists := flagSet[flags[i:i+1]]
		if !exists {
			fmt.Println(errorMSG)
			os.Exit(0)
		}
		charSet += flag
	}
	return charSet
}
