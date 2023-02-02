package main

import (
	"bufio"
	"fmt"
	"os"
)

type Empty struct{}

func main() {
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    var reportsNum int
    fmt.Fscan(in, &reportsNum)

    for i := 0; i < reportsNum; i++ {
		var tasksCount int
		fmt.Fscan(in, &tasksCount)
		reportList := make([]int, 0)
		for j := 0; j < tasksCount; j++ {
			var task int
			fmt.Fscan(in, &task)
			reportList = append(reportList, task)
		}

		setTasks := make(map[int]Empty, len(reportList))

		for k := 0; k <= len(reportList) - 1; k++ {
			if _, ok := setTasks[reportList[k]]; ok && reportList[k] != reportList[k-1] {
				fmt.Fprintln(out, "NO")
				break
			}
			setTasks[reportList[k]] = Empty{}
			if k == len(reportList) - 1 {
				fmt.Fprintln(out, "YES")
			}
		}
	}
}
