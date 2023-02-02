package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func longestCommonSuffix(setWords []string, currentRequest string) string {
	lastEnterIndex := 0

	for i := len(currentRequest) - 1; i >= 0; i-- {
		for lei , v := range setWords {
			if strings.HasSuffix(v, currentRequest[i:]) && currentRequest != v {
				lastEnterIndex = lei
			}
		}
	}

	return setWords[lastEnterIndex]
}

func main() {
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    var setSize int
    fmt.Fscan(in, &setSize)
	setWords := make([]string, setSize)

    for i := 0; i < setSize; i++ {
        fmt.Fscan(in, &setWords[i])
    }

	var requestCount int 
	fmt.Fscan(in, &requestCount)

	for i := 0; i < requestCount; i ++ {
		var currentRequest string
		fmt.Fscan(in, &currentRequest)
		fmt.Fprintln(out, longestCommonSuffix(setWords, currentRequest))
	}
}