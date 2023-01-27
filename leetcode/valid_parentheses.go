package main

import "fmt"

func isValid(s string) bool {
	stack := make([]string, 0)

	for _, v := range s {
		n := len(stack)
		switch v {
			case '(', '{', '[': stack = append(stack, string(v))
			case ')': if n != 0 && stack[n - 1] == "(" {stack = stack[:n - 1]
			} else {return false}
			case '}': if n != 0 && stack[n - 1] == "{" {stack = stack[:n - 1]
			} else {return false}
			case ']': if n != 0 && stack[n - 1] == "[" {stack = stack[:n - 1]
			} else {return false}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("((){(()())}[[]]]){}"))
}