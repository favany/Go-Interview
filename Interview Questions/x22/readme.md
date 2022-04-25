# 22

下面这段代码输出什么？

```go
func main() {  
    a := 5
    b := 8.1
    fmt.Println(a + b)
}
```

A. 13.1
B. 13
C. compilation error

```shell

```

答：C
解析： a 的类型是int ，b 的类型是float ，两个不同类型的数值不能相加，编译报错。