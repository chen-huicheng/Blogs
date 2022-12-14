package main

import "fmt"

type Printer interface {
	Print()
}

type Node struct {
	key string
	val int
}

func (n Node) Print() {
	fmt.Printf("key:%s-val:%d\n", n.key, n.val)
}

type Student struct {
	id     int
	name   string
	gender byte
	age    int
	addr   string
}

func (s *Student) String() string {
	return fmt.Sprintf("å§å:\t%s\nå°å:\t%s", s.name, s.addr)
}

func main() {
	n := Node{"a", 3}
	var p Printer = n
	p.Print()

	var stu *Student
	var ster fmt.Stringer = stu
	fmt.Println(ster.String())
}
