package main

import "fmt"

func main() {
	c := Constructor()
	fmt.Println(c.DeleteHead())
	c.AppendTail(0)
	fmt.Println(c.DeleteHead())
	c.AppendTail(1)

	c.AppendTail(2)
	c.AppendTail(3)
	c.AppendTail(4)

	fmt.Println(c.DeleteHead())
	fmt.Println(c.DeleteHead())

	fmt.Println(c.DeleteHead())

	fmt.Println(c.DeleteHead())

}

type CQueue struct {
	list1 []int
	list2 []int
}

func Constructor() CQueue {
	return CQueue{
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.list1 = append(this.list1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.list2) > 0 {
		last := this.list2[len(this.list2)-1]
		this.list2 = this.list2[:len(this.list2)-1]
		return last
	}
	if len(this.list1) == 1 {
		last := this.list1[0]
		this.list1 = []int{}
		return last
	}
	lenList1 := len(this.list1)
	for i := lenList1 - 1; i >= 0; i-- {
		if i == 0 {
			this.list1 = []int{}
			return this.list1[0]
		}
		this.list2 = append(this.list2, this.list1[i])
	}
	return -1
}
