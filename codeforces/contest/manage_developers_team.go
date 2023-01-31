package main

import (
    "fmt"
    "bufio"
    "os"
)

func Abs(x int) int {
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
		fmt.Println(devList)
		
		for j := 0; j <= len(devList) - 1; j++ {
			fmt.Println(len(devList))
			for k := 1; k <= len(devList) - 1; k++ {
				if k == 4 {
				 	fmt.Fprintln(out, j + 1, " ", k + 1)
				}
				if Abs(devList[j] - devList[k]) == 1 {
					fmt.Fprintln(out, j + 1, " ", k + 1)
					devList = append(devList[j:k], devList[k:]...)
					break
				}
			}
		}
		fmt.Println(devList)
	}
}


// func main() {
//     in := bufio.NewReader(os.Stdin)
//     out := bufio.NewWriter(os.Stdout)
//     defer out.Flush()

//     var teams int
//     fmt.Fscan(in, &teams)

//     for i := 0; i < teams; i++ {
// 		var developersCount int
// 		devList := make([]int, 0)
// 		fmt.Fscan(in, &developersCount)
//         for l := 0; l < developersCount; l++ {
// 			var devLevel int
// 			fmt.Fscan(in, &devLevel)
// 			devList = append(devList, devLevel)
// 		}
// 		fmt.Println(devList)
		
// 		for j := 0; j <= len(devList) - 1; j++ {
// 			fmt.Println(len(devList))
// 			for k := 1; k <= len(devList) - 1; k++ {
// 				if k == 4 {
// 				 	fmt.Fprintln(out, j + 1, " ", k + 1)
// 				}
// 				if Abs(devList[j] - devList[k]) == 1 {
// 					fmt.Fprintln(out, j + 1, " ", k + 1)
// 					devList = append(devList[j:k], devList[k:]...)
// 					break
// 				}
// 			}
// 		}
// 		fmt.Println(devList)
// 	}
// }

	

	

	
