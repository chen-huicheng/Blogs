## 函数

### 定义

go 语言中使用 `func` 关键字定义函数，go中的函数支持多返回值。具体内容看下面的代码

```go
/* 函数定义模版
func funcName(v1 valType1, v2 valType2, v3, v4 valType3)(r1 valType4, r2 valtype5){
}
*/
```
无返回值函数
```go
// 无参数无返回值函数
func MyFunc1() { // 在main前后都可执行
	// do
}
// 单参数函数
func MyFunc2(v int) {
    // do(v)
}
// 多参数函数
func MyFunc3(v1 int, v2 int) { // 等价 func MyFunc3(v1, v2 int) {
    // do(v1, v2)
}
//不定参数
func MyFunc4(args ...int) { //不定长参数只能在最后一个
	fmt.Println("len ", len(args))
	for i, item := range args {
		fmt.Printf("idx : %d, char : %d\n", i, item)
	}
}
```
有返回值函数
```go
// 有返回值 返回值类型写在 函数参数列表后 花括号前
func MyFunc5() int {
	return 666
}

//给返回值 起一个名字  常用写法
func MyFunc6() (ret int) {
	ret = 666
	return
}

// 返回多个返回值
// 等价于 func MyFunc7() (r1 int, r2 int, r3 int)
// func MyFunc7() (r1, r2, r3 int)
func MyFunc7() (int, int, int) { 
	return 1, 2, 3
}
func MyFunc8() (v1 int, v2 int, v3 int) {
	v1, v2, v3 = 1, 2, 3
	return
}
```

## 方法

方法是附着在结构体上的函数。

无论是函数、方法、自定义类型等，Go语言只有一种控制可见性的手段：大写首字母的标识符会从定义它们的包中被导出，小写字母的则不会。这种基于名字的手段使得在语言中最小的封装单元是package，也就是说，以小写字母定义的类型、方法等，仅在该 package 中可见，无论是否在同一个文件中，以大写字母开头的类型方法对所有包可见。

### 定义

```go
// 结构体定义
type Student struct { //自定义类型 Student 所有包可见 
	id     int			// 包内可见
	Name   string   	// 首字母大写，所有包可见
	gender byte			// 包内可见
	age    int			// 包内可见
	Addr   string		// 首字母大写，所有包可见
}
// 方法定义
// 在函数声明的基础上，在函数名之前加一个变量，即是一个方法
func (s Student)SetAge(age int){
    s.age=age
}
s1 := Student{}
s1.SetAge(18)
fmt.Println(s1.age)
// output 0
```

定义的方法是在Student类型上，当调用时，s1 是以值传递的方式传递给SetAge中的s，因此方法中的 s 是调用 s1 的拷贝，修改 s 的值并不能修改 s1。想要修改 s1 的值只需要以指针的形式传递便可以实现，本质上此处也是值传递，只不过传递的值是 s1 的地址，可以通过地址来修改 s1。

```go
func (s *Student)SetAge(age int){
    s.age=age
}
s1 := Student{}
s1.SetAge(18)
fmt.Println(s1.age)
// output 18
```

go 语言中指针类型变量的方法调用和成员使用

```go
var s *Student  //定义一个指针
s1 := Student{}
s = &s1   // 使用 & 符号取地址，并赋值给 s
s.SetAge(18) // 调用 func (s *Student)SetAge(age int)
s1.SetAge(18) // 等价于 (&s1).SetAge(18)
```

### 结构体嵌入

Go 语言中没有继承的概念，可以通过结构体嵌入实现类似的效果。

```go
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
s := Student{People{Name: "zhangsan", Age: 18}, "01101", "高三"}// 嵌入类型的初始化
// Student 类型变量 可以直接使用 People 函数和成员变量
s.Name = "lisi"
s.Display() // 等价于 s.People.Display()，如果 Student 也定义了 Display 只能通过这种显示形式调用
// output# My name is zhangsan. I'm 18 years old.
```

## 接口

go中接口是一个方法集，凡事实现了接口中的所有方法的类型，均可以把该类型赋值给接口类型。一个类型如果拥有一个接口需要的所有方法，那么该类型就实现了这个接口。

```go
package main
import "fmt"
type Printer interface {
	Print()
}
type Node struct {
	key string
	val int
}
func (n Node) Print() { //Node 定义了 Print 方法，即实现了 Printer 接口
	fmt.Printf("key:%s-val:%d\n", n.key, n.val)
}

func main() {
	n := Node{"a", 3}
	var p Printer = n
	p.Print()
}
```

接口类型之间也可以赋值，如果 A 接口的方法集是 B 接口方法集的子集，则可将 B 接口赋值给 A 接口。

### fmt.Stringer

定制自定义类型的打印格式/内容。

fmt 包内有有一个 Stringer 接口，接口定义如下：

```go
type Stringer interface {
    String() string
}
```

fmt 中打印有如下调用，会判断打印变量的类型是否是`error`或`Stringer`接口类型，如果是的话调用`Error()`或`String()`。

```go
// go/src/fmt/print.go:616
switch v := p.arg.(type) {
    case error:
    	handled = true
    	defer p.catchPanic(p.arg, verb, "Error")
    	p.fmtString(v.Error(), verb)
    	return
    case Stringer:
    	handled = true
    	defer p.catchPanic(p.arg, verb, "String")
    	p.fmtString(v.String(), verb)
    	return
}
```

当打印变量实现了 `String()` 方法时，即实现了 `Stringer` 接口，调用 `String()` 方法打印其内容。对于使用`struct` 自定义类型，可以通过定义 `String()` 方法来改变`fmt.Println/Printf` 的打印样式。

```go
type Student struct {
	id     int
	name   string
	gender byte
	age    int
	addr   string
}
func (s Student) String() string {
	return fmt.Sprintf("姓名:\t%s\n地址:\t%s", s.name, s.addr)
}
/* 
func (s *Student) String() string {
	return fmt.Sprintf("姓名:\t%s\n地址:\t%s", s.name, s.addr)
}*/
func main() {
    s := Student{1, "zhang", 0, 25, "beijing"}
	fmt.Println(s)
}
/* output
姓名:   zhang
地址:   beijing
*/
```

使用 `fmt.Print` 什么类型，给什么类型定义 `String()` 方法，`Student` 与 `*Student` 是不同类型。

接口的另一种理解：利用一个接口值可以持有各种具体类型值的能力，将这个接口认为是这些类型的联合。

例如：interface{} ，可以接受任意类型，但是对于一个仅可以处理基本类型(如:int uint float double 等)的函数而言，interface{} 是这些类型的集合。

```go
func PrintType(x interface{}){
    switch x := x.(type){
        case int:
        fmt.Println("int",x)
        case float,double:
        fmt.Println("float",x)
        case bool:
        fmt.Println("bool",x)
        case string:
        fmt.Println("string",x)
        //...  实现支持的类型
        default:
        panic(fmt.Sprintf("unexpected type %T: %v", x, x))// 不支持的类型 panic
    }
}
```

### 接口的实现

接口由两个部分组成，一个具体的类型和那个类型的值，被称为接口的动态类型和动态值。根据接口两部分存储的值是否为空可以分为三类，如图所示：![Slice.drawio](https://raw.githubusercontent.com/chen-huicheng/Blogs/master/img/202209232111871.png)

```go
//第一种
var ster fmt.Stringer = nil

//第二种
var stu *Student = nil
var ster fmt.Stringer = stu
fmt.Println(ster.String())
/* output
panic: runtime error: invalid memory address or nil pointer dereference
*/
//第三种
var stu *Student = &Student{1, "zhang", 0, 25, "beijing"}
var ster fmt.Stringer = stu
fmt.Println(ster.String())

// 陷阱
// 仅当 Type 和 Value 均为 nil 时，才成立，所以当 ster 处于第二种状态时下边的语句会报 panic
if ster != nil{ 
    fmt.Println(ster.String())
}
```

### 类型断言

通过在接口上使用类型断言可以推断接口的动态类型。

```go
// 简化 fmt.print 中的代码
// x 是 interface{} 类型
switch v := x.(type) {
    case error: // 如果 x 是 error 接口类型，则调用 Error() 方法
    	res = v.Error()
    case Stringer:// 如果 x 是 Stringer 接口类型，则调用 String() 方法
    	res = v.String()
}
```

可使用类型断言 + switch 来推断接口的动态类型，基于接口类型断言查询接口动态类型是否实现了某个行为，即如果通过断言 `x`是 Stringer 类型，则该接口指向的动态类型必然实现了 String() 方法，可通过类似方式判断一个接口是否实现了一个具体的方法，通过定义一个仅包含该方法的接口，然后使用类型断言来判断实现。 

