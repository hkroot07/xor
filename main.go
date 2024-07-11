package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"xor/cipherer"
)

var mode = flag.String("mode", "cipher", "Set to 'cipher' or 'decipher'. Default is 'cipher'.")
var secretKey = flag.String("secret", "", "Your secret key. Must contain at least 1 character")

func main() {
	flag.Parse()

	// This is how you can get the pointer length:
	//fmt.Printf("The size of the pointer is: %d bytes\n", unsafe.Sizeof(secretKey))

	if len(*secretKey) == 0 {
		fmt.Fprintln(os.Stderr, "No sicret is provided! Exiting now...")
		os.Exit(1)
	}

	switch *mode {
	case "cipher":
		plaintext := getUserInput("Enter your text cipher: ")

		cipheredText, err := cipherer.Cipher(plaintext, *secretKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error encrypting text: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(cipheredText)
	case "decipher":
		cipheredText := getUserInput("Enter your ciphered data to decipher: ")

		fmt.Println(cipherer.Decipher(cipheredText, *secretKey))
	default:
		fmt.Println("invalid mode. Use 'cipher' or decipher.")
		os.Exit(1)
	}

}

func getUserInput(msg string) string {
	fmt.Println(msg)

	reader := bufio.NewReader(os.Stdin)
	for {
		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error! Please try again...")
			continue
		}
		return strings.TrimRight(result, "\r\n")
	}

}
