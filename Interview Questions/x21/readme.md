# 21

下面这段代码输出什么？

```go
func hello(num ...int) {
    num[0] = 18
}

func main() {
    i := []int{5, 6, 7}
    hello(i...)
    fmt.Println(i[0])
}
```

A. 18
B. 5
C. Compilation error

解析： 可变参数传递过去，改变了第一个值。
