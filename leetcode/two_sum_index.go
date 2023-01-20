// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

 

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].


package main 

import "fmt"


//*************     O(n'2)    *************//

// func twoSum(nums []int, target int) []int {
//     res := make([]int, 0)

//     for i, v1 := range nums {
//         for j, v2 := range nums {
// 			if j == i {continue}
//             if v1 + v2 == target {
//                 res = append(res, i, j)
// 				return res
//             }
//         }
//     }
//     return res
// }

//*************     O(n)    *************//

func twoSum(nums []int, target int) []int {
    resmap := make(map[int]int)

    for i, curNum := range nums {
        if _, ok := resmap[target - curNum]; ok {
            return []int{resmap[target - curNum], i}
        } else {
            resmap[curNum] = i
        }
    }
    return nil
}

func main() {
	fmt.Println(twoSum([]int{2,4,2}, 4))
}

//16'