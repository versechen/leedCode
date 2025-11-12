package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40, 50}

	for i, j := 0, 0; i < len(nums)-1 && j < len(nums)-1; i++, j++ {
		fmt.Printf("i: %d, j: %d, nums[i]: %d, nums[j]: %d\n", i, j, nums[i], nums[j])
	}
}
