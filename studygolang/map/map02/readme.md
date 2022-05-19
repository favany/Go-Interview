# map02

下面这段代码输出什么？

```go
type person struct {
    name string
}

func main() {
    var m map[person]int
    p := person{"make"}
    fmt.Println(m[p])
}
```

A. 0
B. 1
C. Compilation error

<details>
    <summary>答案</summary>
    **答：A**
    **解析：**
    打印一个map中不存在的值时，返回元素类型的零值。这个例子中，m的类型是map[person]int，因为m中 不存在p，所以打印int类型的零值，即0。
</details>
A.
