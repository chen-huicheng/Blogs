package leetcode

type MyQueue struct {
	sta1 []int
	sta2 []int
}

func Constructor() MyQueue {
	return MyQueue{sta1: make([]int, 0), sta2: make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.sta1 = append(this.sta1, x)
}

func (this *MyQueue) Pop() int {
	n := len(this.sta2)
	if n > 0 {
		v := this.sta2[n-1]
		this.sta2 = this.sta2[:n-1]
		return v
	}
	this.stuck()
	n = len(this.sta2)
	v := this.sta2[n-1]
	this.sta2 = this.sta2[:n-1]
	return v

}

func (this *MyQueue) Peek() int {

}

func (this *MyQueue) Empty() bool {

}

func (this *MyQueue) stuck() {
	for i := len(this.sta1) - 1; i >= 0; i-- {
		this.sta2 = append(this.sta2, this.sta1[i])
	}
	this.sta1 = this.sta1[:0]
}
