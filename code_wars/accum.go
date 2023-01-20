// This time no story, no theory. The examples below show you how to write function accum:

// Examples:
// accum("abcd") -> "A-Bb-Ccc-Dddd"
// accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
// accum("cwAt") -> "C-Ww-Aaa-Tttt"
// The parameter of accum is a string which includes only letters from a..z and A..Z.

package main

import (
	"fmt"
	"strings"
)

func Accum(s string) string {
	s = strings.ToUpper(s)
	resStr := strings.Split(s, "")
	for i, v := range resStr{
		if i >= 1 {
			v += strings.Repeat(strings.ToLower(v), i) 
			resStr[i] = v
		}
	}
	
	return strings.Join(resStr, "-")
}

func main() {
	fmt.Printf("%v\n%v\n", Accum("abc"), "A-Bb-Ccc")
	fmt.Printf("%v\n%v\n", Accum("RqaEzty"), "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy")
	fmt.Printf("%v\n%v\n", Accum("cwAt"),"C-Ww-Aaa-Tttt")
}

// Time: 1565ms Passed: 2Failed: 0
// Test Results:
// Test Example
// should handle basic cases
// Test Passed
// Completed in 0.1057ms
// should handle random cases
// Test Passed
// Completed in 0.3638ms
// You have passed all of the tests! :)