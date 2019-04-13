package main

import "fmt"

func scanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}

func main() {
	var n int
	fmt.Scan(&n)

	h := scanNums(n)

	max := 0
	c_sea := 0
	for _, v := range h {
		if v >= max {
			max = v
			c_sea++
		}
	}
	fmt.Println(c_sea)

}
