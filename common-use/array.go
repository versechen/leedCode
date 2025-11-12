package main

import "fmt"

func commonUse1(s []int, i int) {
	// 删除第i个元素，其他保持顺序,注意，s和之前公共一份底层数据，底层数据也变了
	s = append(s[:i], s[i+1:]...)
	// 假如s := []int{1, 2, 3, 4, 5}, i = 2
	// 覆盖后的底层数组：[1, 2, 4, 5, 5]
	// s = [1, 2, 4, 5]

	// copy(s[i:], s[i+1:])
	// var zero T
	// s[len(s)-1] = zero // 如果 T 是指针/包含指针，清零有助于 GC
	// s = s[:len(s)-1]
	// type Big struct {
	//     buf []byte  // 很大的 slice
	// }

	// arr := make([]*Big, 10_000) // 很多 *Big
	// 删除某个元素时，如果不把最后那个指针清掉，GC 会认为还在被引用
	// 所以要显式清零
	// s[len(s)-1] = zero

}

func commonUse2(s []int, i int) {
	// 删除，但不保持顺序
	s[i] = s[len(s)-1]
	s = s[:len(s)-1]
	// 手动回收最后一个元素
	// var zero T
	// s[len(s)-1] = zero // 如果 T 是指针/包含指针，清零有助于 GC
}

func commonUse3(s []int, i int, target int) {
	// 在i处插入元素
	s = append(s, 0)     // 预填充一个元素，扩展长度
	copy(s[i+1:], s[i:]) // 向后移动一位, copy复制min(len(s)-i-1, len(s)-i)
	s[i] = target
}

func main() {
	// 数组转切片
	var a = [5]int{1, 2, 3, 4, 5}
	s := a[1:4] // s == []int{1,2,3}  len=3 cap=4(=5-1)
	fmt.Printf("len is %d, cap is %d\n", len(s), cap(s))

	// 创建切片
	var s1 []int           //nil 切片；len=0 cap=0；可直接 append
	s1 = []int{}           // 空切片（非nil）：len=0, cap=0
	s1 = make([]int, 3)    // len=3, cap=3 ==> [0,0,0]
	s1 = make([]int, 3, 8) //len=3 cap=8(预存容量，减少扩容)
	fmt.Println("预占3个0值", s1)

	// 基本切片语法
	// b :=a[low:high] [low,hight) 半开区间
	// b :=a[:n] 从头开始
	// b:=a[m:] 到末尾
	// b:=a[:] 整个数组、切片

	// 三下标
	// c:=s[low:high:max] len = high-low; cap = max -low
	// d:= s[:i:i]  让 cap==len, 后续 append强制新分配，避免“污染”原切片
	a1 := []int{1, 2, 3, 4, 5}
	sub := a1[1:3] // [2,3], len=2 cap=4(=5-1)
	fmt.Println("1) 初始：     a =", a1, " sub =", sub, " len,cap =", len(sub), cap(sub))
	sub = append(sub, 99) // 仍在同一个底层数组里扩展，覆盖了 a[3]
	fmt.Println("   append 后：a1 =", a1, " sub =", sub, "  (a1[3] 被改成 99)")

	// 2) 用 full slice 限制 cap：b[1:3:3] => len=2, cap=2
	b := []int{1, 2, 3, 4, 5}
	subB := b[1:3:3] // 第三个下标 max=3 => cap = 3-1 = 2
	fmt.Println("\n2) 初始：     b =", b, " subB =", subB, " len,cap =", len(subB), cap(subB))
	subB = append(subB, 99) // cap 已满，必须新分配，不会改动 b
	fmt.Println("   append 后：b =", b, " subB =", subB, "  (b 保持不变)")

	// 3) 常见写法：d := s[:i:i]，强制 cap==len
	s2 := []int{10, 20, 30, 40, 50}
	i := 3
	d := s[:i:i] // len=3 cap=3
	fmt.Println("\n3) 初始：     s2 =", s2, " d =", d, " len,cap =", len(d), cap(d))
	d = append(d, 100) // 必新分配，不会污染 s2
	fmt.Println("   append 后：s2 =", s2, " d =", d, "  (s 保持不变)")

	// 4) 验证公式：c := s[low:high:max] => len=high-low, cap=max-low
	low, high, max := 1, 4, 5
	c := s2[low:high:max]
	fmt.Printf("\n4) c = %v, len=%d, cap=%d (期望 len=%d, cap=%d)\n",
		c, len(c), cap(c), high-low, max-low)

	// 追加与拷贝
	// s = append(s, x)       // 可能原地，也可能新分配并复制（cap 不够时）
	// s = append(s, a, b, c) // 追加多个
	// s = append(s, t...)    // 追加另一个切片

	// dst := make([]int, 3)
	// n := copy(dst, s) // 返回实际拷贝数量 = min(len(dst), len(s))

}
