package main

import (
	"math/rand"
)

func quickSort(arr []int) {
	_quickSort(arr, 0, len(arr)-1)
}
func _quickSort(arr []int, l, r int) {
	if l < r {
		m := partition(arr, l, r)
		_quickSort(arr, l, m-1)
		_quickSort(arr, m+1, r)
	}
}
func partition(arr []int, l, r int) int {
	p := rand.Intn(r-l+1) + l // 随机位置
	arr[l], arr[p] = arr[p], arr[l]
	for l < r {
		for l < r && arr[l] < arr[r] { // 交换支点位置 arr所有值都一样的情况下 l r 趋近于中间
			r--
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
		}
		for l < r && arr[l] < arr[r] {
			l++
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
			r--
		}
	}
	return l
}

func main() {

}
