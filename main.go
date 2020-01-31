package main

import "fmt"

import "strings"

// check1 simply makes sure that the string length is even, if it isn't then we're missing a character and thus
// we are not balaneced.
func check1(input string) bool {
	if len(input)%2 == 0 {
		return true
	}
	return false
}

// check2 simply checks that the string starts with either a (, [, or a { and if not, we're not balanced.
func check2(input string) bool {
	if strings.HasPrefix(input, "(") || strings.HasPrefix(input, "[") || strings.HasPrefix(input, "{") {
		return true
	}
	return false
}

// check3 appends the expected closing brackets (based on the rune number) to the expected array while
// the actual closing order is added to the reality array
// from there we loop through the reality array backwards and compare it to the expected array
// if a rune number does not match the expected, then we are not balanced
func check3(input string) bool {
	expected := []int{}
	reality := []int{}
	length := len(input) / 2
	for _, letternum := range input {
		if int(letternum) != 123 && int(letternum) != 91 && int(letternum) != 40 {
			reality = append(reality, int(letternum))
		} else {
			if letternum == 123 || letternum == 91 {
				expected = append(expected, int(letternum+2))
			} else {
				expected = append(expected, int(letternum+1))
			}
		}
	}
	i := 0
	for ii := length; ii > 0; ii-- {
		if reality[ii-1] != expected[i] {
			return false
		}
		i++
	}
	return true
}

func main() {
	// modify test to be whatever you want, if it's balanced it will pass all three checks, otherwise it will fail a check.
	test := "(([[{}]]))"
	pass := true
	if check1(test) {
		fmt.Println("Passed check 1!")
	} else {
		fmt.Println("Failed check 1!")
		pass = false
	}
	if check2(test) {
		fmt.Println("Passed check 2!")
	} else {
		fmt.Println("Failed check 2!")
		pass = false
	}
	if check3(test) {
		fmt.Println("Passed check 3!")
	} else {
		fmt.Println("Failed check 3!")
		pass = false
	}
	if pass {
		fmt.Printf("%s is balanced!", test)
	} else {
		fmt.Printf("%s is not balanced!", test)
	}
}
