package main

import (
	"fmt"
	"math/rand"
	"time"
)

//切片 变长数组
func main() {
	arr := []int{10, 20, 30, 0, 0}
	fmt.Println("arr = ", arr)

	s := arr[0:3]
	fmt.Println("s = ", s)
	// s[3] = 2
	// fmt.Println("arr = ", arr)
	fmt.Println("len(s) = ", len(s))
	fmt.Println("cap(s) = ", cap(s))

	//创建方式
	s1 := []int{1, 2, 3}
	fmt.Println(s1)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
	s2 := make([]int, 5, 10) // make(切片类型，长度，容量)
	fmt.Println(s2)
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))
	s3 := make([]int, 5)
	fmt.Println(s3)
	fmt.Printf("len = %d, cap = %d\n", len(s3), cap(s3))

	//只能在 len(arr)范围内操作 添加元素使用append
	// s2[5] = 6//runtime error: index out of range [5] with length 5"
	s3 = append(s3, 6)
	fmt.Println(s3)
	fmt.Printf("len = %d, cap = %d\n", len(s3), cap(s3))

	//切片的截取
	array := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(array[1:3])
	fmt.Println(array[:3])
	fmt.Println(array[3:])
	fmt.Println(array[4])
	fmt.Println(array[:])
	fmt.Println(len(array))
	fmt.Println(cap(array))
	// [2 3]
	// [1 2 3]
	// [4 5 6]
	// 5
	// [1 2 3 4 5 6]
	// 6
	// 6

	// 切片底层原理
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[2:5]
	b[1] = 666
	fmt.Println(a)
	fmt.Println(b)
	// [1 2 3 666 5 6 7]
	// [3 666 5]

	c := b[2:7] //可以取到不在自己范围内的    a,b,c 操作同一个底层数组
	c[3] = 777
	fmt.Println(a)
	fmt.Println(c)
	// [1 2 3 666 5 6 7 777 9 10]
	// [5 6 7 777 9]

	// append
	aa := make([]int, 0, 0)
	// aa := []int{}
	oldCap := cap(aa)
	for i := 0; i < 15; i++ {
		aa = append(aa, i)
		if newCap := cap(aa); oldCap < newCap {
			fmt.Printf("cap : %d ===> %d\n", oldCap, newCap)
			oldCap = newCap
		}
	}
	fmt.Println(aa)

	//copy
	srcSlice := []int{11, 12}
	dstSlice := []int{3, 3, 3, 3, 3, 3}

	dstSlice = append(dstSlice, 8)
	copy(dstSlice, srcSlice) // 把 src 的数据按照下标位置 拷贝到 dst 中
	fmt.Println(dstSlice)    // [11 12 3 3 3 3 8]

	// fmt.Printf("%T  %T", srcSlice, dstSlice)

	//切片做函数参数

	sortS := make([]int, 20)
	InitData(sortS)
	sort(sortS)
	fmt.Println(sortS)
	modifySlice(sortS)
	fmt.Println(sortS)

}
func InitData(s []int) { //切片作为参数传递的是引用
	rand.Seed(time.Now().UnixNano())
	for i := range s {
		s[i] = rand.Intn(100)
	}
}

func sort(s []int) {
	n := len(s)
	flag := false
	for i := n - 1; i > 0; i-- {
		flag = true
		for j := 0; j < i; j++ {
			if s[j] > s[j+1] {
				flag = false
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
		if flag {
			break
		}
	}
	// fmt.Println("arr = ", arr)
}

func modifySlice(s []int) {
	fmt.Println(s)
	s = append(s, 2048)
	s[0] = 1024
}
