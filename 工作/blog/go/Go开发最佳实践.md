# Go开发最佳实践

## 编程规范

+ 避免多余的else语句

```Go
//example 差
func absoluteNumber(x int)int{
    if x>=0{
        return x
    }else{
        return -x
    }
}
//好
func absoluteNumber(x int)int{
    if x>=0{
        return x
    }
    return -x
}
```

- 保持正常代码路径为最小缩进

> 正常代码路径是指在没有异常情况下代码所执行的路径。
>
> 保持正常代码路径不缩进，能提高代码的可读性。如果正常代码路径路径很长，那么理解正常代码路径上的代码需要记住很多上下文，增加理解的难度。

```Go
//example 代码对比 错误
func fun()error{
    err:=doSomething()// 正常代码
    if err == nil{
        err:= doAnotherthing()// 正常代码
        if err == nil{
            return nil// 正常代码
        }
        return err
    }
    return err
}
//正确
func fun()error{
    err:=doSomething()// 正常代码
    if err != nil{
        return err
    }
    err:= doAnotherthing()// 正常代码
    if err != nil{
        return err
    }
    return nil// 正常代码
}
```

- 代码注释
  - 表达具体功能（What）而不是实现（How）
  - 包的注释放在与包同名的文件下，或者单独的doc.go中。
  - 避免不必要的注释。
  - 注释应该使用完整的句子，开头使用函数或者变量的名字。

- Panic
  - Panic 是运行时异常。
  - 一旦执行到panic后：
    - 当前执行的函数会停止。
    - 代码之前调用的defer均会执行。
    - 如果不使用recover处理panic，整个程序也会终止。
  - 什么时候使用panic
    - 没法恢复的error，导致程序没法继续执行。
    - 认为产生的错误。
  - 如何处理panic
    - 在defer调用的函数中使用recover。
    - defer 函数位于目标函数之前。

- Error
  - error是一个包含Error() string 函数的接口。
  - 自定义错误，实现Error() string函数即可，Unwrap() string 
  - 如何匹配error
    - error.Is()
    - error.As()

- 指针
  - 什么时候使用指针
    - 方法/函数需要修改对象/参数时
    - 传递大的数据时

## 匿名函数

匿名函数捕获外部变量时使用的引用类型，而不是值类型。**警告：捕获迭代变量**，在捕获迭代变量时，捕获的时该迭代变量，而不是某个时刻该迭代变量的值。

样例如下所示：

```go
//正确
var rmdirs []func()
for _, d := range tempDirs() {
    dir := d // NOTE: necessary!
    os.MkdirAll(dir, 0755) // creates parent directories too
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir)
    })
}
// ...do some work…
for _, rmdir := range rmdirs {
    rmdir() // clean up
}
//错误
var rmdirs []func()
for _, dir := range tempDirs() {
    os.MkdirAll(dir, 0755)
    rmdirs = append(rmdirs, func() {
      //dir是for作用域变量，rmdirs中所用 func() 类型的函数均捕获该变量，在for执行结束时，dir为tempDirs最后一个元素，执行rmdir()时，仅删除最有一个文件夹。
        os.RemoveAll(dir) // NOTE: incorrect!
    })
}
```

## deferred函数

```go

package main
import (
	"log"
	"time"
)
//!+main
func bigSlowOperation() {
    // defer 仅将最后一级函数作为函数结束时执行的函数，其余函数作为表达式直接执行。
    // 如下所示：多级函数（函数的返回结果是一个返回函数的函数）前两级直接执行
	defer trace("bigSlowOperation")()() // don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}
func trace(msg string) func() func() {
  
	start := time.Now()
	log.Printf("enter %s (%s)", msg, start)
	process := func() func() {
		log.Println("test defer")
		return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
	}
	return process
}
func main() {
	bigSlowOperation()
}


/*
defer trace("bigSlowOperation")()()
!+output
2022/09/02 11:36:46 enter bigSlowOperation (2022-09-02 11:36:46.959925 +0800 CST m=+0.000141764)
2022/09/02 11:36:46 test defer
2022/09/02 11:36:56 exit bigSlowOperation (10.000647126s)
!-output
*/
// 修改 defer trace("bigSlowOperation")()() 为 defer trace("bigSlowOperation")() 执行结果如下：
/* 
!+output
2022/09/02 11:52:36 enter bigSlowOperation (2022-09-02 11:52:36.534406 +0800 CST m=+0.000128712)
2022/09/02 11:52:46 test defer
!-output
最后一级函数不执行
*/

```


