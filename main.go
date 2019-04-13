package main

import (
	"fmt"
)

func scanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}

func main() {
	var s string
	fmt.Scan(&s)

	rune_s := []rune(s)
	c := 0
	for i, v := range rune_s {
		if i == 0 {
			continue
		}
		if int32(rune_s[i-1]) == v {
			if v == 48 {
				rune_s[i] = 49
			} else {
				rune_s[i] = 48
			}
			c++
		}
	}

	fmt.Println(c)
}
