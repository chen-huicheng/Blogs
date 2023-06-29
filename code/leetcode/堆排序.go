package main

import "fmt"

type cmpFunc func(a, b int) bool

func heapSort(arr []int) {
	// 第一个非叶子节点 构建小顶堆
	i := (len(arr) + 1) / 2
	var cmp cmpFunc
	cmp = func(a, b int) bool {
		if a > b {
			return true
		}
		return false
	}
	for ; i >= 0; i-- {
		down(arr, i, cmp)
	}
	for i = len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		down(arr[:i], 0, cmp)
	}
}

func down(arr []int, i int, cmp cmpFunc) {
	n := len(arr)
	for 2*i+1 < n {
		j := 2*i + 1
		if j+1 < n && cmp(arr[j+1], arr[j]) {
			j++
		}
		if !cmp(arr[j], arr[i]) {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i = j
	}
}

func main() {
	arr := []int{10, 5, 4, 3, 2, 6, 7, 9, 8}
	heapSort(arr)
	fmt.Println(arr)
}
