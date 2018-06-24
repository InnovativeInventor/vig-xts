package main

import (
  "testing"
  "crypto/rand"
  "encoding/base64"
)

func TestReversible(t *testing.T) {
    // Checks if decrypt will reverse encrypt
    
    plaintext,err1 := generateRandomString(214)
    key,err2 := generateRandomString(13)
    
    if err1 != nil || err2 != nil {
        t.Errorf("There was an error with generating random strings for the TestReversible function.")
    }
    
    ciphertext := Encrypt(plaintext,key)
    answeredPlaintext := Decrypt(ciphertext,key)
    if plaintext != answeredPlaintext {
       t.Errorf("Response was incorrect, got: %v, want: %v.", answeredPlaintext, plaintext)
    }
}

func generateRandomBytes(n int) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if err != nil {
        return nil, err
    }

    return b, nil
}

func generateRandomString(s int) (string, error) {
    b, err := generateRandomBytes(s)
    return base64.URLEncoding.EncodeToString(b), err
}


func TestAdd(t *testing.T) {
    // Should be one + two + prev (mod 95)
    
    arrayOne := []int{31,41,59,26,53}
    arrayTwo := []int{62,8,31}
    arrayExpected := []int{93,47,42,35,1}
    arrayAnswered := addMod(arrayOne, arrayTwo)
    equal := testEq(arrayAnswered, arrayExpected)
    if !equal {
       t.Errorf("Response was incorrect, got: %d, want: %d.", arrayAnswered, arrayExpected)
    }
}

func TestSub(t *testing.T) {
    // Should be one - two - prevOne (mod 95)
    
    arrayOne := []int{31,41,59,26,53}
    arrayTwo := []int{62,8,31}
    arrayExpected := []int{64,2,82,0,19}
    arrayAnswered := subMod(arrayOne, arrayTwo)
    equal := testEq(arrayAnswered, arrayExpected)
    if !equal {
       t.Errorf("Response was incorrect, got: %d, want: %d.", arrayAnswered, arrayExpected)
    }
}

func testEq(a, b []int) bool {

    if a == nil && b == nil { 
        return true; 
    }

    if a == nil || b == nil { 
        return false; 
    }

    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}

func TestTurnASCII(t *testing.T) {
    stringOne := "Test!"
    arrayExpected := []int{52,69,83,84,1}
    arrayAnswered := turnASCII(stringOne)
    equal := testEq(arrayAnswered, arrayExpected)
    if !equal {
       t.Errorf("Response was incorrect, got: %d, want: %d.", arrayAnswered, arrayExpected)
    }
}

func TestASCII(t *testing.T) {
    arrayOne := []rune{'t','E','$','T'}
    arrayExpected := []int{84,37,4,52}
    arrayAnswered := make([]int,0)
    
    for _, char := range arrayOne {
        output := ASCII(char)
        arrayAnswered = append(arrayAnswered,output)
    }
    
    equal := testEq(arrayAnswered, arrayExpected)
    if !equal {
       t.Errorf("Response was incorrect, got: %d, want: %d.", arrayAnswered, arrayExpected)
    }
}

func TestTurnString(t *testing.T) {
    stringExpected := "Test!"
    arrayOne := []int{52,69,83,84,1}
    stringAnswered := turnString(arrayOne)
    if stringExpected != stringAnswered {
       t.Errorf("Response was incorrect, got: %v, want: %v.", stringExpected, stringAnswered)
    }
}
