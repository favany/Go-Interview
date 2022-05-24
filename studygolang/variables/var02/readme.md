# var02

下面这段代码能否通过编译，如果可以，输出什么？

```go
var(
	size := 1024
	max_size = size*2
)

func main() {
	fmt.Println(size,max_size)
}
```

不能 var 里面只能声明变量 不能有表达式的计算

<details>
  <summary>答案</summary>
参考答案：不能通过编译。

参考解析：这道题的主要知识点是变量声明的简短模式，形如：x := 100.
但这种声明方式有限制：

1. 必须使用显示初始化；
2. 不能提供数据类型，编译器会自动推导；
3. 只能在函数内部使用简短模式；

</details>
