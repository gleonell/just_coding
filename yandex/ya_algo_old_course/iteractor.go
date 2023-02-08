package main

import "fmt"

func iteractor_verdict(retCode, itr, checker int) int {
	if !(itr >= 0 && itr <= 7) && !(retCode >= -127 && retCode <= 127) {
		return 0
	}

	switch  {
		case (itr == 0 || itr == 4) && retCode != 0 : return 3
		case ((itr == 0 || itr == 1) && retCode == 0) || itr == 1 : return checker
		case itr == 6: return 0
		case itr == 7: return 1
	default:
		return itr
	}
}

func main() {
	var retCode, iteractor, checker int
	fmt.Scan(&retCode, &iteractor, &checker)
	fmt.Println(iteractor_verdict(retCode, iteractor, checker))
}