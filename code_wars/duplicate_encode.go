// The goal of this exercise is to convert a string to a new string where each character in the new string
// is "(" if that character appears only once in the original string, or ")" if that character appears more than once in the original string.
// Ignore capitalization when determining if a character is a duplicate.

// Examples
// "din"      =>  "((("
// "recede"   =>  "()()()"
// "Success"  =>  ")())())"
// "(( @"     =>  "))(("
// Notes
// Assertion messages may be unclear about what they display in some languages. If you read "...It Should encode XXX", the "XXX" is the expected result, not the input!

package main

import (
	"fmt"
	"strings"
)

func DuplicateEncode(word string) string {
  var resStr string
  mapChars := make(map[string]int)
  
  word = strings.ToLower(word)

  for i, v := range word {
    mapChars[string(v)] = strings.Count(word, string(word[i]))
  }

  for _, v := range word {
	  if mapChars[string(v)] > 1 {
		  resStr += ")"
	  } else {
		  resStr += "("
	  }
  }
  return resStr
}

func main() {
	fmt.Println(DuplicateEncode("Aaabcc"))
}

// Time: 1632ms Passed: 4Failed: 0
// Test Results:
// Test Example
// Basic tests:
// should return the correct value
// Test Passed
// Completed in 0.0271ms
// should ignore case
// Test Passed
// Completed in 0.0158ms
// Tests with '(' and ')'
// should return the correct value
// Test Passed
// Completed in 0.0062ms
// Some random tests: 
// should return correct values
// Test Passed
// Completed in 0.0479ms
// You have passed all of the tests! :)




/*
package main

import (
	"fmt"
	"unicode"
)

func DuplicateEncode(word string) string {
  var resStr string
  mapChars := make(map[string]int)

  for _, v := range word {
    if unicode.IsUpper(v) {  
      v = unicode.ToLower(v)
    }
    mapChars[string(v)]++
  }

  for _, v := range word {
    if mapChars[string(v)] != 1 {
      resStr += ")"
    } else {
      resStr += "("
    }
  }
  return resStr
}
*/