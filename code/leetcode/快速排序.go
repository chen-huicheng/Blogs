package main

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
	v := arr[l]
	for l < r {
		for l < r && arr[r] > v {
			r--
		}
		arr[l] = arr[r]
		for l < r && arr[l] <= v {
			l++
		}
		arr[r] = arr[l]
	}
	arr[l] = v
	return l
}

func main() {

}
