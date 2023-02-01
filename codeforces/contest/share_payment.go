package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    var ordersNum int
    fmt.Fscan(in, &ordersNum)

    for i := 0; i < ordersNum; i++ {
		var result, itemsNum int
		setItems := make(map[int]int)
		fmt.Fscan(in, &itemsNum)
        for l := 0; l < itemsNum; l++ {
			var itemId int
			fmt.Fscan(in, &itemId)
			setItems[itemId]++
		}
		for k, v := range setItems {
			result += k * (v - v/3)
		}
		fmt.Fprintln(out, result)
	}
}

	

	

	
