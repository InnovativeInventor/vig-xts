package main

import (
  "fmt"
  "log"
  "os"
  "bufio"
  "github.com/urfave/cli"
  "bytes"
)

func main() {
  app := cli.NewApp()
  app.Name = "Vig-xts"
  app.Usage = "Use xts with the Vigenere cipher"
  app.Version = "0.1.6"
  app.Commands = []cli.Command{
    {
      Name:    "encrypt",
      Aliases: []string{"e"},
      Usage:   "Encrypt text",
      Action:  func(c *cli.Context) error {
        text, key := getInput()
        ciphertext := Encrypt(text,key)
        fmt.Println(ciphertext)
        return nil
      },
    },
    {
      Name:    "decrypt",
      Aliases: []string{"d"},
      Usage:   "Decrypt text",
      Action:  func(c *cli.Context) error {
        text, key := getInput()
        plaintext := Decrypt(text,key)
        fmt.Println(plaintext)
        return nil
      },
    },
   }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}

func Encrypt(plaintext string ,key string) (ciphertext string) {
    plaintextArray := turnASCII(plaintext)
    keyArray := turnASCII(key)
    
    encryptedArray := addMod(plaintextArray, keyArray) 
    ciphertext = turnString(encryptedArray)
    return ciphertext
}

func Decrypt(ciphertext string, key string) (plaintext string) {
    ciphertextArray:= turnASCII(ciphertext)
    keyArray := turnASCII(key)
    
    decryptedArray := subMod(ciphertextArray, keyArray) 
    plaintext = turnString(decryptedArray)
    return plaintext
}

func addMod(text []int, key []int,) (output []int) {
    // Adds the numbers in the array, mod 95
    output = make([]int,0)
    for count, letter := range text {
        pos := count%len(key)
        
        if count == 0 {
            result := letter+key[pos]
            output = append(output,result%95)
        } else {
            result := letter+output[count-1]+key[pos]
            output = append(output,result%95)
        }

    }
    return output
}

func subMod(text []int, key []int,) (output []int) {
    // Subtracts the numbers in the array, mod 95
    output = make([]int,0)
    for count, letter := range text {
        pos := count%len(key)
        
        if count == 0 {
            result := letter-key[pos]+95
            resultMod := ((result % 95) + 95) % 95
            output = append(output,resultMod)
        } else {
            result := letter-key[pos]-text[count-1]+95+95
            resultMod := ((result % 95) + 95) % 95
            output = append(output,resultMod)
        }

    }
    return output
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
    return text, key
}