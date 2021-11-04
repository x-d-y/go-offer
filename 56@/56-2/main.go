package main

import "fmt"

// 使用自动机机，求出状态转移方程

// 设 当前位为 one， 高位为 two

// 总共状态有三种， 分别是余数为 00 -> 01 -> 10 -> 00

// 如果输入的当前位为 0， 则 one 不用动， two 也不用动

// 如果输入的当前位为 1， 则 one 进位， two 看情况变化

// 最终的状态状态转移方程为： one = one ^ n & ~two ； two = two ^ n & ~one

// 只需要计算当前位的个数即可， 因为如果是 3 的倍数， 对应为为 0

func main() {
	fmt.Println(singleNumber([]int{1, 1, 1, 5,5,5,7}))

}

func singleNumber(nums []int) int {
	one, two := 0, 0
	for _, v := range nums {
		one = one ^ v & ^two
		two = two ^ v & ^one
	}
	return one
}
