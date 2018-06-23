package main

import (
  "fmt"
  "log"
  "os"
  "bufio"
  "github.com/urfave/cli"
  "bytes"
//  "strings"
)
/*
 * This program accepts anything from ASCII 31 to 127 (95 characters)
 */

func main() {
  app := cli.NewApp()
  app.Name = "Vig-xts"
  app.Usage = "Use xts with the Vigenere cipher"
  app.Action = func(c *cli.Context) error {
    encrypt()
    return nil
  }
  
  app.Commands = []cli.Command{
    {
      Name:    "encrypt",
      Aliases: []string{"e"},
      Usage:   "Encrypt text",
      Action:  func(c *cli.Context) error {
        encrypt()
        return nil
      },
    },
    {
      Name:    "decrypt",
      Aliases: []string{"d"},
      Usage:   "Decrypt text",
      Action:  func(c *cli.Context) error {
        decrypt()
        return nil
      },
    },
   }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}

func encrypt() {
    fmt.Println("Encrypting")
    plaintext, key := getInput()
    
    plaintextArray := turnASCII(plaintext)
    fmt.Println(plaintextArray)
    keyArray := turnASCII(key)
    
    encryptedArray := addMod(plaintextArray, keyArray) 
    ciphertext := turnString(encryptedArray)
    fmt.Println(encryptedArray) //debug
    
    fmt.Println(ciphertext)
    
}

func decrypt() {
    fmt.Println("Decrypting")
    ciphertext, key := getInput()

    ciphertextArray:= turnASCII(ciphertext)
    keyArray := turnASCII(key)
    
    decryptedArray := subMod(ciphertextArray, keyArray) 
    fmt.Println(decryptedArray) //debug
    
    plaintext := turnString(decryptedArray)
    fmt.Println(plaintext)
}

func addMod(text []int, key []int,) (output []int) {
    output = make([]int,0)
    fmt.Println("Adding your numbers")
    for count, letter := range text {
        pos := count%len(key)
        result := letter+key[pos]
        output = append(output,result%95)
    }
    return
}

func subMod(text []int, key []int,) (output []int) {
    output = make([]int,0)
    fmt.Println("Subtracting your numbers")
    for count, letter := range text {
        pos := count%len(key)
        result := letter-key[pos]+95
        
        // This makes sure mod is always positive, but I think the +95 in the line above fixes that
        resultMod := ((result % 95) + 95) % 95
        output = append(output,resultMod)
    }
    return
}

func ASCII(r rune) int {
    num := int(r)
    if num > 31 {   
        return num-32   
    } else {
        return 0
    }
}

func turnASCII(text string) (slice []int) {
    slice = make([]int,0)
    for _, char := range text {
        num := ASCII(char)
        slice = append(slice, num)
    }
    return slice
}

func turnString(slice []int) (text string) {
    var buffer bytes.Buffer
    for _, char := range slice {
        character := string(char+32)
        buffer.WriteString(character)
    }
    return buffer.String()
}

func getInput() (text string, key string) {
    reader := bufio.NewReader(os.Stdin)
    
    fmt.Println("Type in the text you want to encrypt/decrypt")
    text, _ = reader.ReadString('\n')
    text = text[:len(text)-1]
    
    fmt.Println("Type in your key")
    key, _ = reader.ReadString('\n')
    key = key[:len(key)-1]
    return
}