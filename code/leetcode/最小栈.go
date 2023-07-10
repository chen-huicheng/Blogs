package leetcode

// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

// 实现 MinStack 类:

// MinStack() 初始化堆栈对象。
// void push(int val) 将元素val推入堆栈。
// void pop() 删除堆栈顶部的元素。
// int top() 获取堆栈顶部的元素。
// int getMin() 获取堆栈中的最小元素。

type MinStack struct {
	sta    []int
	minSta []int
}

func Constructor() MinStack {
	return MinStack{sta: make([]int, 0), minSta: make([]int, 0)}
}

func (s *MinStack) Push(val int) {
	s.sta = append(s.sta, val)
	mn := len(s.minSta)
	if mn == 0 {
		s.minSta = append(s.minSta, val)
	} else {
		s.minSta = append(s.minSta, min(s.minSta[mn-1], val))
	}
}

func (s *MinStack) Pop() {
	n := len(s.sta)
	if n == 0 {
		return
	}
	s.sta = s.sta[:n-1]
	s.minSta = s.minSta[:n-1]
}

func (s *MinStack) Top() int {
	n := len(s.sta)
	if n == 0 {
		return 0
	}
	v := s.sta[n-1]
	s.Pop()
	return v

}

func (s *MinStack) GetMin() int {
	n := len(s.minSta)
	if n == 0 {
		return 0
	}
	return s.minSta[n-1]
}

// func (s *MinStack) down(i int) {
// 	n := len(s.sta)
// 	for i*2+1 < n {
// 		l := i*2 + 1
// 		r := l + 1
// 		if r < n && s.sta[r] < s.sta[l] {
// 			l = r
// 		}
// 		if s.sta[l] >= s.sta[i] {
// 			break
// 		}
// 		s.sta[i], s.sta[l] = s.sta[l], s.sta[i]
// 	}
// }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
