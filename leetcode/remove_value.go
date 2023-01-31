package main

import "fmt"

func removeElement(nums []int, val int) int {
    i := 0
    for _, v := range nums {
        if v != val {
            nums[i] = v
            i++
        }
    }
	fmt.Println(nums)
    return i
}

func main() {
	nums := []int{1,2,2,3,3,3,3,4,5,5, 6}
	fmt.Println(removeElement(nums, 3))
}
