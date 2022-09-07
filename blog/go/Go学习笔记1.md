## 入门

### 官方安装方法

https://go.dev/doc/install

网站会根据系统自动识别应该下载的安装包，下载对应的安装包，点击安装即可。在终端输入`go version` 输出 go 版本号代表安装成功，不生效重启终端试试。

![](https://raw.githubusercontent.com/chen-huicheng/Blogs/master/img/image-20220907220203682.png)

### 源码安装方式

1.   下载 Go

     打开Golang官网下载地址：https://go.dev/dl/  根据系统选择对应的包，复制下载链接。

     ```shell
     # 下载tar文件
     wget https://go.dev/dl/go1.19.1.linux-amd64.tar.gz
     ```

     ![](https://raw.githubusercontent.com/chen-huicheng/Blogs/master/img/image-20220907221237783.png)

2.   安装

     ```shell
     # 解压文件
     tar -zxvf go1.19.1.linux-amd64.tar.gz
     # 将解压出来的文件 移动端 /usr/local
     sudo mv -v go /usr/local
     ```

     编辑  ~/.bash_profile 将 /usr/local/go/bin 添加到环境变量，并设置 GOPATH

     ```shell
     export GOPATH=$HOME/go
     export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
     ```

     使 ~/.bash_profile 生效

     ```shell
     source ~/.bash_profile
     ```

     验证是否安装成功

     ```shell
     go version
     # 输出版本号表示安装成功
     # go version go1.19.1 linux/amd64
     ```

3.   第一个 Go 程序

     ```shell
     # 进入Go工作空间
     cd $GOPATH
     # 没有的话先创建
     mkdir $GOPATH && cd $GOPATH
     # 通常在 $GOPATH 下创建 bin src pkg 三个文件夹
     mkdir bin src pkg
     cd src
     vim hello.go
     ```
     在 hello.go 中写入以下内容
     ```go
     package main
     import "fmt"
     func main(){
         fmt.Println("Hello go")
     }
     ```

     运行 go 文件

     ```shell
     go run hello.go
     # output
     # Hello go
     ```

### Go学习文档

+   [Go语言之旅](https://tour.go-zh.org/welcome/1) 推荐

+   [Go语言圣经](https://books.studygolang.com/gopl-zh/)  推荐

+   [Go by Example](https://gobyexample.com/) 推荐

+   [Go语言中文网](https://books.studygolang.com/)

+   [Go package](https://pkg.go.dev/)

    



## 接口

### fmt.Stringer

定制自定义类型的打印格式/内容。

fmt 包内有有一个 Stringer 接口，接口定义如下

```go
type Stringer interface {
    String() string
}
```

fmt.Print 时会有如下调用 

`go/src/fmt/print.go:611`

```go
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

当打印变量实现了 `String()` 方法时，即实现了 `Stringer` 接口，调用 `String()` 方法打印其内容。对于使用`struct` 自定义类型，可以通过定义 `String()` 方法来实现，`fmt.Println/Printf` 的格式。

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



## 方法



### 结构体嵌入

### 封装

无论是函数、方法、自定义类型等，Go语言只有一种控制可见性的手段：大写首字母的标识符会从定义它们的包中被导出，小写字母的则不会。这种基于名字的手段使得在语言中最小的封装单元是package。

```go
func add(a,b int){} //所在包可见，其他包不可见
func Add(a,b int){} //首字母大写，所有包可见

type Student struct { //自定义类型 Student 所有包可见 
	id     int			// 包内可见
	Name   string   	// 首字母大写，所有包可见
	gender byte			// 包内可见
	age    int			// 包内可见
	Addr   string		// 首字母大写，所有包可见
}
```

