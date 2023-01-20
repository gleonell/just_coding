// Take 2 strings s1 and s2 including only letters from a to z. 
// Return a new sorted string, the longest possible, containing distinct letters - each taken only once - coming from s1 or s2.

// Examples:
// a = "xyaabbbccccdefww"
// b = "xxxxyyyyabklmopq"
// longest(a, b) -> "abcdefklmopqwxy"

// a = "abcdefghijklmnopqrstuvwxyz"
// longest(a, a) -> "abcdefghijklmnopqrstuvwxyz"

package main

import (
	"fmt"
	"sort"
	"strings"
)
//решение в лоб
func sortString(str string) string{
	res := strings.Split(str, "")
	sort.Strings(res)
	return strings.Join(res, "")
}

func TwoToOne(s1 string, s2 string) string {
	mapUnique := make(map[rune]string)
	resStr := s1 + s2
	for _, v := range resStr {
		mapUnique[v] = string(v) 
	}
	resStr = ""
	for _, v := range mapUnique {
		resStr += v
	}

	return sortString(resStr)
}

func main() {
	fmt.Print(TwoToOne("xyaabbbccccdefww", "xxxxyyyyabklmopq"))
}

// Time: 1924ms Passed: 3Failed: 0
// Test Results:
// Extensive Test
// additional fixed tests
// Test Passed
// Completed in 0.2634ms
// 200 random tests
// Test Passed
// Completed in 1.6565ms
// Test Example
// fixed tests
// Test Passed
// Completed in 0.0217ms
// You have passed all of the tests! :)


// best practices *******************
/*
func TwoToOne(s1 string, s2 string) string {
	s := s1 + s2
	result := ""
	for _, char := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
		if strings.Contains(s, char) {
			result += char
		}
	}
	return result
}
*/