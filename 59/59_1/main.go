package main

func main() {

}

type MaxQueue struct {
	queue   []int
	dequeue []int
}

func Constructor() MaxQueue {
	m := MaxQueue{
		queue:   make([]int, 0),
		dequeue: make([]int, 0),
	}
	return m
}

func (this *MaxQueue) Max_value() int {
	if len(this.dequeue) == 0 {
		return -1
	}
	return this.dequeue[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)
	for i, v := range this.dequeue {
		if v < value {
			this.dequeue = this.dequeue[:i]
			this.dequeue = append(this.dequeue, value)
			return
		}
	}
	this.dequeue = append(this.dequeue, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	res := this.queue[0]
	this.queue = this.queue[1:]
	if res == this.dequeue[0] {
		this.dequeue = this.dequeue[1:]
	}
	return res
}
