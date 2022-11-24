# Wire 使用指南

翻译 https://github.com/google/wire/blob/main/docs/guide.md

## 前言

#### IoC : Inversion of Control 控制反转

控制反转是一种设计原则，用来降低代码模块之间的耦合度。

因为大多数应用程序都是由两个或是更多的类合作来实现业务逻辑，这使得每个对象都需要获取与其合作的对象（也就是它所依赖的对象）的引用。如果这个获取过程要靠自身实现，那么这将导致代码高度耦合并且难以维护和调试。

控制反转就是”依赖对象的获得“被反转了。

#### DI：Dependency Injection 依赖注入

**依赖注入**是实现**控制反转**的一种**实现方法**。

例如：类 A 中定义一个的B对象，A 依赖B。

```go
// 不使用 控制反转
type A struct{}
type B struct{}

func (a *A) Controller() {
	b := new(B) //A 直接 new 来获得依赖对象 B 的引用
	b.Server()
}
func (b *B) Server() {
	fmt.Println("no IoC")
}
func main() {
	a := new(A)
	a.Controller()
}

// 使用 控制反转
type A struct {
	b *B
}
type B struct{}

func (a *A) Controller() {
	a.b.Server()
}
func (b *B) Server() {
	fmt.Println("use IoC")
}
func initA() *A { // 依赖注入控制程序
	b := new(B)
	return &A{b}
}
func main() {
	a := initA()
	a.Controller()
}
```

使用依赖注入的方式是通过相关的控制程序来将 B 对象在外部 new 出来并注入到 A 类里的引用中。

**依赖反转**由 A 直接 new 来获得依赖对象 B 的引用**反转**为通过相关的控制程序来将依赖对象 B 在外部 new 出来并注入到对象 A 里的引用中。

采用依赖注入技术之后，A 的代码只需要定义一个 private 的B对象，不需要直接 new 来获得这个对象，而是通过相关的容器控制程序来将B对象在外部new出来并注入到A类里的引用中。

Wire 是一个代码生成工具，它使用依赖注入自动连接组件。也就是说 wire 帮我们自动生成上述的 initA 函数。

## 安装和使用

运行一下命令安装wire

```shell
go install github.com/google/wire/cmd/wire@latest
```

确保 `$GOPATH/bin` 已经添加到环境变量 `$PATH`中。不然没法运行。

编写一下代码 `wire.go`

```go
/* Service.go */
type A struct {
	b *B
}
type B struct{}

func (a *A) Controller() {
	a.b.Server()
}
func (b *B) Server() {
	fmt.Println("use IoC")
}
func NewA(b *B)*A{
    return &A{b}
}
func NewB()*B{
    return &B{}
}

/* wire.go */
//+build wireinject

func InitA() (A, error) {
	wire.Build(NewA, NewB)  // 函数注入
    // wire.Build(wire.Struct(new(A), "*"), wire.Struct(new(B), "*")) // 结构注入
	return A{}, nil
}

/* wire_gen.go */ //wire 生成代码
//go:build !wireinject
// +build !wireinject

// 函数注入 生成代码
func InitA() (*A, error) {
	b := NewB()
	a := NewA(b)
	return a, nil
}
// 结构注入 生成代码
func InitA() (*A, error) {
	b := &B{}
	a := &A{
		b: b,
	}
	return a, nil
}
```

`service.go` 业务代码

`wire.go` 按照 wire 的标准编写。`//+build wireinject` 标注该文件是 wire 的输入文件

`wire_gen.go` 是 wire 生成的实现依赖注入的代码。

## 高级用法

wire 由两个核心概念：**providers**和**injectors**。

providers：可以是一个返回特定值的函数，或者一个 struct。上述中的 NewA 和 NewB函数。

injectors：一个按依赖顺序调用 providers 的函数。上述中 wire.go 中的 InitA 函数。

### providers

可以将常用的 providers 放到一个 Set 中。

```go
var SuperSet = wire.NewSet(NewA, NewB)
wire.Build(SuperSet)
```

**绑定接口** 和 **结构提供者**

```go
/* Service.go */
type A struct {
	is IServer  // A 依赖接口 IServer
}
type IServer interface{
    Server()
}
type B struct{}

func (a *A) Controller() {
	a.is.Server()
}
func (b *B) Server() {  // B 实现了 IServer
	fmt.Println("use IoC")
}
/* wire.go */
//+build wireinject

package app

var SuperSet = wire.NewSet(
    wire.Struct(new(B),"*"), // 填充 B 中所有字段
    wire.Struct(new(A),"*"), // 填充 A 中所有字段
    wire.Bind(new(IServer),new(*B)), // 使用 *B 填充 IServer 字段
)
func InitA() (A, error) {
    wire.Build(SuperSet) // 结构注入
	return A{}, nil
}

/* wire_gen.go */ //wire 生成代码
//go:build !wireinject
// +build !wireinject

// Injectors from wire.go:
func InitA() (A, error) {
	b := &B{}
	a := A{
		is: b,
	}
	return a, nil
}
```



