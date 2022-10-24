package main

import "fmt"

type People struct {
	Name string
	Age  int
}

func (p People) Display() {
	fmt.Printf("My name is %s. I'm %d years old.\n", p.Name, p.Age)
}

type Student struct {
	People        // 嵌入 people 类型
	StuNo  string //自定义成员
	Grade  string
}

// // 结构体定义
// type Student struct { //自定义类型 Student 所有包可见
// 	id     int    // 包内可见
// 	Name   string // 首字母大写，所有包可见
// 	gender byte   // 包内可见
// 	age    int    // 包内可见
// 	Addr   string // 首字母大写，所有包可见
// }

// 方法定义
// 在函数声明的基础上，在函数名之前加一个变量，即是一个方法
// func (s *Student) SetAge(age int) {
// 	s.age = age
// }

func main() {
	// s1 := Student{}
	// s1.SetAge(18)
	// fmt.Println(s1.age)

	// var s *Student
	// s1 = Student{}
	// s = &s1
	// fmt.Println(s)
	// s1.SetAge(18)
	// fmt.Println(s)

	s := Student{People{Name: "zhangsan", Age: 18}, "01101", "高三"}
	s.Display()
	s.People.Display()
}
