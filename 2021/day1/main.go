package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// getNums gets slice of nums from raw string
func getNums(input string) (nums []int) {

	s := strings.Split(input, "\n")
	for _, v := range s {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		nums = append(nums, n)
	}
	return nums
}

// depthCompare comapres a depth value to the next in the slice
// and sums the total
func depthCompare(nums []int) (c int) {
	for i := range nums {
		if i == len(nums)-1 {
			break
		}
		if nums[i] < nums[i+1] {
			c++
		}
	}
	return c
}

// depthCompareWindow compares the next three depth values with the
// immediate next three depth values and returns the sum.
func depthCompareWindow(nums []int) (c int) {
	for i := range nums {
		if i == len(nums)-3 {
			break
		}
		if nums[i]+nums[i+1]+nums[i+2] < nums[i+1]+nums[i+2]+nums[i+3] {
			c++
		}
	}
	return c
}

func main() {
	nums := getNums(rawInput)

	// single depth count
	count := depthCompare(nums)

	// sliding window depth count
	countWindow := depthCompareWindow(nums)

	//  print results
	fmt.Printf("Num times depth increases: '%d'\n", count)
	fmt.Printf("Num times depth increases in window: '%d'\n", countWindow)

}
