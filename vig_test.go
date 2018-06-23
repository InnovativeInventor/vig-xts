package main

import "testing"

func TestAdd(t *testing.T) {
    arrayOne := []int{31,41,59,26,53}
    arrayTwo := []int{62,8,31}
    arrayExpected := []int{93,47,42,35,1}
    arrayAnswered := addMod(arrayOne, arrayTwo)
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