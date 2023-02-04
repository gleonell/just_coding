package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type timeInterval struct{
	start time.Time
	end time.Time
}

func (t *timeInterval) compare() int {
	if t.end.After(t.start) {return 1}
	if t.end.Before(t.start) {return -1}
	return 0 // if Equal
}

func main() {
	in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()
	
	var requests int
	fmt.Fscan(in, &requests)
	for r := 0; r < requests; r++ {
		var sets int
		fmt.Fscan(in, &sets)

		intervalsList := make([]timeInterval, 0)
		wrongSet := false
		for s := 0; s < sets; s++ {
			
			var timePair string
			fmt.Fscan(in, &timePair)
			var interval timeInterval
			
			timePairSlice := strings.Split(timePair, "-")
			firstTime, err := time.Parse("15:04:05", timePairSlice[0])
			if err != nil {
				wrongSet = true
			}
			secondTime, err := time.Parse("15:04:05", timePairSlice[1])
			if err != nil {
				wrongSet = true
			}

			interval.start, interval.end = firstTime, secondTime
			if interval.compare() == -1 {
				wrongSet = true
			}

			intervalsList = append(intervalsList, interval)

			if s == sets - 1 {
				if wrongSet == true {
					fmt.Fprintln(out, "NO")
					break
				}
				sort.Slice(intervalsList, func(i, j int) bool {
					return intervalsList[i].start.Before(intervalsList[j].start)
				})
				for i := 0; i < len(intervalsList) - 1; i++ {
					if (intervalsList[i+1].start.Before(intervalsList[i].end) || intervalsList[i+1].start.Equal(intervalsList[i].end)) && 
					intervalsList[i+1].start.After(intervalsList[i].start) || intervalsList[i+1].start.Equal(intervalsList[i].start) {
						fmt.Fprintln(out, "NO")
						wrongSet = true
						break
					}
				}
				if !wrongSet {
					fmt.Fprintln(out, "YES")
				}
			}
		}	
	}
}
