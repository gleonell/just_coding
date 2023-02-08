package main

import "fmt"

func most_common_symbol(str string) byte {
	smblsCount := make(map[byte]int)
	var ret byte 
	for i := 0; i < len(str); i++ {
		smblsCount[str[i]]++
	}
	tmp := 0
	for k, count := range smblsCount {
		if count > tmp {ret, tmp = k, count}
	}
	return ret
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(string(most_common_symbol(str)))
}