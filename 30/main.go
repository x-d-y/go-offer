type MinStack struct {
	normal []int
	min    []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		normal: make([]int, 0),
		min:    make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.normal = append(this.normal, x)
	if len(this.min) == 0 || this.min[len(this.min)-1] >= x {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	last := this.normal[len(this.normal)-1]
	this.normal = this.normal[:len(this.normal)-1]
	if last == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	return this.normal[len(this.normal)-1]
}

func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}