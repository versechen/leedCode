package main

import (
	"fmt"
	"strings"
)

func main() {
	// 输入字符串
	var str string
	fmt.Println("请输入一个字符串：")
	fmt.Scanln(&str)

	// 统计 ":-)" 和 ":-(" 的数量
	countHappy := strings.Count(str, ":-)")
	countSad := strings.Count(str, ":-(")

	// 输出结果
	fmt.Printf("开心表情 ':-)' 出现次数: %d\n", countHappy)
	fmt.Printf("伤心表情 ':-(' 出现次数: %d\n", countSad)
	if countHappy > countSad {
		fmt.Printf("开心")
	} else if countHappy < countSad {
		fmt.Printf("伤心")
	} else {
		fmt.Printf("just so so")
	}

}
