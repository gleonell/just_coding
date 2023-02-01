package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func main() {
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    var teams int
    fmt.Fscan(in, &teams)

    for i := 0; i < teams; i++ {
		var developersCount int
		devList := make([]int, 0)
		fmt.Fscan(in, &developersCount)
        for l := 0; l < developersCount; l++ {
			var devLevel int
			fmt.Fscan(in, &devLevel)
			devList = append(devList, devLevel)
		}
		
		teamsCount := 0
		for i := 0; i <= len(devList) - 1; i++ {
			if devList[i] == -1 { continue }
			tmpAbsMin := 9999999
			lastIndex := 0
			for j := 1; j <= len(devList) - 1; j++ {
				if devList[j] != -1 && j != i && abs(devList[i] - devList[j]) < tmpAbsMin {
					tmpAbsMin = abs(devList[i] - devList[j])
					lastIndex = j
				}
				if j == len(devList) - 1 {
					fmt.Fprintln(out, i+1, " ", lastIndex+1)
					devList[i] = -1
					devList[lastIndex] = -1
					teamsCount++
				}
			}
		}
		fmt.Fprintln(out)
	}
}
