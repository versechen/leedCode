package main

import (
	"fmt"
	"math"
	"strconv"
)

var (
	minAbs = math.MaxInt
	c      = 0
)

func dfs(nums string, num int, sum int) {
	if len(nums) == num {
		minAbs = int(math.Min(float64(minAbs), float64(math.Abs(float64(sum-c)))))
		println(minAbs)
		return
	}

	nowNum, _ := strconv.Atoi(string(nums[num]))
	dfs(nums, num+1, sum-nowNum)
	dfs(nums, num+1, sum+nowNum)
}

func main() {
	strs := ""
	fmt.Scan(&strs)
	fmt.Scan(&c)
	first, _ := strconv.Atoi(string(strs[0]))
	dfs(strs, 1, first)
	println(minAbs)
}
