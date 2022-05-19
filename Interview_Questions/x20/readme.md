# 20

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


