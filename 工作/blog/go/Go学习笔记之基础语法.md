## 入门

### 官方安装方法

https://go.dev/doc/install

网站会根据系统自动识别应该下载的安装包，下载对应的安装包，点击安装即可。在终端输入`go version` 输出 go 版本号代表安装成功，不生效重启终端试试。![image-20220907220203682](https://raw.githubusercontent.com/chen-huicheng/ImageHub/main/typora_img/202307101250560.png)

### 源码安装方式

1.   下载 Go

     打开Golang官网下载地址：https://go.dev/dl/  根据系统选择对应的包，复制下载链接。

     ```shell
     # 下载tar文件
     wget https://go.dev/dl/go1.19.1.linux-amd64.tar.gz
     ```

     ![image-20220907221237783](https://raw.githubusercontent.com/chen-huicheng/ImageHub/main/typora_img/202307101251426.png)

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

## Go基础语法

### 变量

```go
//变量声明  
var i int  	// var [varname] [vartype]  
var i,j int // 多变量声明 默认初始化为0值
i:=1    	// [varname]:=val 自动推导类型 :=  推导类型并赋值
/*go中变量类型有 
bool
uint8,uint16,uint32,uint64
int8,int16,int32,int64
float32,float64
byte(uint8别名)
rune(int32 别名)
string
指针
*/
type char byte // 定义一个新的类型 类型别名
var cc char

// 匿名变量 _
i,j := 1,2
temp,_ := i,j // 忽略第二个变量

// 常量定义
const i int = 1i
const j = i

// 字符串
var str1 string    			// 默认值为 ""
str1 = "hello go"
str2 := "hello go"
str := str1 + " " + str2	// 字符串连接

// 指针
var pi *int // 定义指向int型的指针，默认值为空：nil
i := 10
p = &a      // 取地址
fmt.Println(*pi) // 取指针指向的值

// 切片
var s []int  // 声明一个int类型切片
s = append(s, 6) // 添加一个元素到切片
fmt.Printf("len = %d, cap = %d\n", len(s), cap(s)) //打印切片长度和容量
/*
len(s):切片长度，切片包含元素数量
cap(s):切片容量，切片底层数组的长度
当 append 时 len(s)==cap(s) 即没有多余的空间，先扩容，再添加
*/
s2 := make([]int, 5, 10) //使用make创建一个切片 可以指定 len 和 cap
s3 := s2[i:j] //截取切片
s4 := s2[:j]  //可忽略首
s5 := s3[i:]  //可忽略尾
/*
s2 s3 s4 s5 底层是同一个数组，因此修改元素会导致其他切片同时修改
*/

// map 
// var m1 map[keyType]valType  keyType必须是可比较类型
var m1 map[int]string
m3 := make(map[int]string, 5)
dict := map[int]string{}
dict[1] = "hello"   //添加一个key:value
val,ok := dict[key]
if ok{ //判断 key 是否存在，ok:true 表示存在，ok:false 表示不存在 
    fmt.Println(val)
}
delete(dict, 1) //删除key

//channel
var ch = make(chan int) //创建方式  make(chan Type)

ch <- v    // 发送数据v到管道，管道满的时候,发送数据会阻塞
v := <-ch  // 从管道接收数据赋值该v，管道空的时候，接收数据会阻塞

ch := make(chan int, 100) // 带缓冲的管道

// 自定义类型
//大写首字母的标识符会从定义它们的包中被导出，小写字母的则不会
type Student struct {
	Name   string    // 首字母大写,可导出;其他包可使用
	age    int       // 首字母小写,不可导出;仅所在包可使用
}
var s Student
s.Name = "zhang"
s.age = 15  //同一包下正常，其他包取不到 age
s1 := Student{"zhang", 25}

```

### 控制语句

#### for

```go
// for 初始化语句;判断条件;循环条件 {}
sum := 0
for i := 0; i < 10; i++ {
	sum += i
}
// 初始化语句,循环条件 可省略
sum := 0
for sum < 10{  // 相当于 “while”
    sum+=sum
}
// 死循环
for {
}
// 结合 range 遍历 slice map 等
for i,key := range s{
}
for i := range s{ //key 可以省略
}
for _,key := range s{ //不使用 i 时，使用 _ 表示
}
```

#### If else

```go
// if 表达式 {}
if i<10{    
}
// if 初始化语句;表达式{}
if val,ok := dict[key];ok{
    fmt.Println(val)
}
//等价于
val,ok := dict[key]
if ok{ //判断 key 是否存在，ok:true 表示存在，ok:false 表示不存在 
    fmt.Println(val)
}

// if else
if v < 0 {
    return -v
} else {
    return v
}
```

#### switch

```go
switch os := runtime.GOOS; os {
    case "darwin":
    fmt.Println("OS X.")
    break                  // 默认 break 加不加效果一致
    case "linux":
    fmt.Println("Linux.")
    fallthrough            //使用fallthrough 不跳出switch
    default:
    fmt.Printf("%s.\n", os)
}
// case 可以是一个表达式
t := time.Now()
switch {
    case t.Hour() < 12:
    fmt.Println("Good morning!")
    case t.Hour() < 17:
    fmt.Println("Good afternoon.")
    default:
    fmt.Println("Good evening.")
}
```

#### defer

```go
func main() {
	defer fmt.Println("world")	// 后执行
	fmt.Println("hello")       	// 先执行
}
/*+output
hello
world
*/
func main() {
	defer fmt.Println("hello")	// 多个 defer 前面的后执行
    defer fmt.Println("world")	
    defer fmt.Println("go")		
	fmt.Println("hello")       	// 先执行
}
```



