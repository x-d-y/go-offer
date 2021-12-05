package main

import (
	"fmt"
	"math"
)

func main() {
	m := Constructor()
	input := []int{155, 66, 114, 0, 60, 73, 109, 26, 154, 0, 107, 75, 9, 57, 53, 6, 85, 151, 12, 110, 64, 103, 42, 103, 126, 3, 88, 142, 79, 88, 147, 47, 134, 27, 82, 95, 26, 124, 71, 79, 130, 91, 131, 67, 64, 16, 60, 156, 9, 65, 21, 66, 49, 108, 80, 17, 159, 24, 90, 79, 31, 79, 113, 39, 54, 156, 139, 8, 90, 19, 10, 50, 89, 77, 83, 13, 3, 71, 52, 21, 50, 120, 159, 45, 22, 69, 144, 158, 19, 109, 52, 50, 51, 62, 20, 22, 71, 95, 47, 12, 21, 32, 17, 130, 109, 8, 61, 13, 48, 107, 14, 122, 62, 54, 70, 96, 11, 141, 129, 157, 136, 41, 40, 78, 141, 16, 137, 127, 19, 70, 15, 16, 65, 96, 157, 111, 87, 95, 52, 42, 12, 60, 17, 20, 63, 56, 37, 129, 67, 129, 106, 107, 133, 80, 8, 56, 72, 81, 143, 90, 0}
	check := []float64{155, 110.5, 114, 90, 66, 69.5, 73, 69.5, 73, 69.5, 73, 74, 73, 69.5, 66, 63, 66, 69.5, 66, 69.5, 66, 69.5, 66, 69.5, 73, 69.5, 73, 74, 75, 77, 79, 77, 79, 77, 79, 80.5, 79, 80.5, 79, 79, 79, 80.5, 82, 80.5, 79, 79, 79, 79, 79, 77, 75, 74, 73, 74, 75, 74, 75, 74, 75, 77, 75, 77, 79, 77, 75, 77, 79, 77, 79, 77, 75, 74, 75, 76, 77, 76, 75, 74, 73, 72, 71, 72, 73, 72, 71, 71, 71, 72, 71, 72, 71, 71, 71, 70, 69, 68, 69, 70, 69, 68, 67, 66.5, 66, 66.5, 67, 66.5, 66, 66, 66, 66, 66, 66, 66, 65.5, 66, 66, 66, 66, 66, 66.5, 67, 66.5, 66, 66.5, 67, 66.5, 67, 68, 67, 68, 67, 66.5, 66, 66.5, 67, 68, 69, 69.5, 69, 68, 67, 66.5, 66, 66, 66, 65.5, 65, 65.5, 66, 66, 66, 66.5, 67, 67, 67, 66.5, 67, 67, 67, 68, 67}
	for i, v := range input {
		m.AddNum(v)
		fmt.Println(m.FindMedian())
		if check[i] != m.FindMedian() {
			fmt.Println("")
		}
	}
}

type MedianFinder struct {
	maxheap []int
	minheap []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *MedianFinder) minheapPush(num int) {
	this.minheap = append(this.minheap, math.MaxInt64)
	for i := 0; i < len(this.minheap); i++ {
		if num < this.minheap[i] {
			num, this.minheap[i] = this.minheap[i], num
		}
	}
}

func (this *MedianFinder) maxheapPush(num int) {
	this.maxheap = append(this.maxheap, math.MaxInt64)
	for i := 0; i < len(this.maxheap); i++ {
		if num < this.maxheap[i] {
			num, this.maxheap[i] = this.maxheap[i], num
		}
	}
}

func (this *MedianFinder) AddNum(num int) {
	if num == 107 {
		fmt.Println("")
	}
	if len(this.minheap) <= len(this.maxheap) {
		// todo 插入到小数堆中

		if len(this.minheap) == 0 || num <= this.maxheap[0] {
			this.minheapPush(num)
		} else {
			this.minheapPush(this.maxheap[0])
			this.maxheap = this.maxheap[1:]
			this.maxheapPush(num)
		}
	} else {
		// todo 插入到大数堆中
		if len(this.maxheap) == 0 {
			if num < this.minheap[len(this.minheap)-1] {
				num, this.minheap[len(this.minheap)-1] = this.minheap[len(this.minheap)-1], num
			}
		}
		if len(this.maxheap) == 0 || num >= this.minheap[len(this.minheap)-1] {
			this.maxheapPush(num)
		} else {
			this.maxheapPush(this.minheap[len(this.minheap)-1])
			this.minheap = this.minheap[:len(this.minheap)-1]
			this.minheapPush(num)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.maxheap) == len(this.minheap) {
		return float64(this.maxheap[0]+this.minheap[len(this.minheap)-1]) / 2.0
	}
	return float64(this.minheap[len(this.minheap)-1])
}
