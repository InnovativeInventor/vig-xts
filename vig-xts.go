package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	lowerBound := 33 // These values are both inclusive
	upperBound := 126
	// Getting input and turning it into a slice
	plaintext, key := getInput()

	plaintextSlice := turnSlice(plaintext, lowerBound, upperBound)
	keySlice := turnSlice(key, lowerBound, upperBound)

	encryptSlice := encrypt(plaintextSlice, keySlice, lowerBound, upperBound)
	ciphertext := turnString(encryptSlice, lowerBound, upperBound)
	fmt.Println(ciphertext)
}

func getInput() (plaintext string, key string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Type in the text you want to encrypt")
	plaintext, _ = reader.ReadString('\n')

	fmt.Println("Type in your key")
	key, _ = reader.ReadString('\n')
	return
}

func turnSlice(input string, lowerBound int, upperBound int) []int {
	var plaintextSlice []int
	var num int

	for pos, char := range input {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			continue
		}
		num = int([]rune(input)[pos])
		if lowerBound <= num && upperBound >= num {
			plaintextSlice = append(plaintextSlice, num-lowerBound) // This makes it impossible to enter anything besides basic latin
		}
	}

	return plaintextSlice
}

func turnString(input []int, lowerBound int, upperBound int) string {
	output := make([]rune, 0, len(input))
	for _, i := range input {
		output = append(output, rune(i+lowerBound))
	}
	return string(output)
}

func encrypt(plaintextSlice []int, keySlice []int, lowerBound int, upperBound int) (encryptSlice []int) {
	var encryptedInt int
	length := len(keySlice)
	mod := upperBound - lowerBound

	for pos := range plaintextSlice {
		key := keySlice[pos%length]
		encryptedInt = (key + pos + encryptedInt) % mod

		encryptSlice = append(encryptSlice, encryptedInt)
	}
	return
}
