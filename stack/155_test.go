package stack

import (
	"log"
	"testing"
)

type MinStack struct {
	data   []int
	min    []int
	p1, p2 int
}

func Constructor() MinStack {
	return MinStack{data: make([]int, 0, 2000), min: make([]int, 0, 2000), p1: 0, p2: 0}
}

func (this *MinStack) Push(val int) {
	// log.Printf("when push the %v %v", this.p1, this.p2)
	if this.p1 == len(this.data) {
		this.data = append(this.data, val)
	} else {
		this.data[this.p1] = val
	}
	this.p1++
	if this.p1 == 1 || this.min[this.p2-1] >= val {
		if this.p2 == len(this.min) {
			this.min = append(this.min, val)
		} else {
			this.min[this.p2] = val
		}
		this.p2++
	}
}

func (this *MinStack) Pop() {
	if this.data[this.p1-1] == this.min[this.p2-1] {
		this.p2--
	}
	this.p1--
}

func (this *MinStack) Top() int {
	return this.data[this.p1-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.p2-1]
}

func TestStack(t *testing.T) {

	arr := make([]int, 0, 10)
	temp := -1
	log.Println("the len of arr is", len(arr))
	if true || arr[temp] == 1 {
		log.Println("HI")
	}

	s := Constructor()
	s.Push(1)
	s.Push(1)
	s.Push(1)
	log.Println(s.p1)
	s.Pop()
	s.Pop()
	s.Pop()
	log.Println(s.p1)
	s.Push(2)

}
