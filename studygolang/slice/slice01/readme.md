# slice01

> 91

下面代码输出什么？
```go
func main() {
    x := []string{"a", "b", "c"}
    for v := range x {
        fmt.Print(v)
    }
}

```

<details>
  <summary>答案</summary>
  答案： 0 1 2
  解析：
  注意区分以下代码段
  func main() {
    x := []string{"a", "b", "c"}
    for _, v := range x {
        fmt.Print(v)     //输出 a b c
    }
  }

注意区别下面代码段：
</details>



