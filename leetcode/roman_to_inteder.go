// Given a roman numeral, convert it to an integer.

// Example 1:

// Input: s = "III"
// Output: 3
// Explanation: III = 3.

package main 

import "fmt"

func romanToInt(s string) int {
    res := 0
	mapRoman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := len(s) - 1; i >= 0; i-- {
		switch string(s[i]) {
			case "V", "X":
				if i != 0 && string(s[i-1]) == "I" {
					res += mapRoman[string(s[i])] - 1
					i--
				} else {
					res += mapRoman[string(s[i])]
				}
			case "L", "C":
				if i != 0 && string(s[i-1]) == "X" {
					res += mapRoman[string(s[i])] - 10
					i--
					 
				} else {
					res += mapRoman[string(s[i])]
				}
			case "D", "M":
				if i != 0 && string(s[i-1]) == "C" {
					res += mapRoman[string(s[i])] - 100
					i--
				} else {
					res += mapRoman[string(s[i])]
				}
			default:
				res += mapRoman[string(s[i])]
		}
	}
	return res
}

func main() {
	fmt.Println(romanToInt("III"), 3)
	fmt.Println(romanToInt("LVIII"), 58)
	fmt.Println(romanToInt("MCMXCIV"), 1994)
}

// Runtime
// 11 ms
// Beats
// 63.1%
// Memory
// 2.8 MB
// Beats
// 100%