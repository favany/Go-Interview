# interface02

> 26

下面这段代码输出什么？

```go
func main() {
    var i interface{}
    if i == nil {
        fmt.Println("nil")
        return
    }
    fmt.Println("not nil")
}
```

- A. nil
- B. not nil
- C. compilation error

<details>
  <summary>答案</summary>
  **答：A**

  **解析：**
  当且仅当接口的动态值和动态类型都为 nil 时，接口类型值才为 nil
</details>
