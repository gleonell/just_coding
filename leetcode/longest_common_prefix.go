// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

 

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	for i, val1:=  range strs[0] {
		for _ , substr := range strs[1:]{
			if i == len(substr) || val1 != rune(substr[i]) {return strs[0][:i]}
		}
	}
	return strs[0]
}

func main() {
	//strs := []string{"flower","flowedlf","flowght"}
	strs := []string{"aaaa","aaa","aa", "aaa"}
	fmt.Println(longestCommonPrefix(strs))
}

// Runtime
// 0 ms
// Beats
// 100%
// Memory
// 2.2 MB
// Beats
// 53.35%