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

