# other04

下面这段代码输出什么？

```go
func main() {
    i := -5
    j := +5
    fmt.Printf("%+d %+d", i, j)
}
```

- A. -5 +5
- B. +5 +5
- C. 0  0

<details>
  <summary>答案</summary>
  **答：A**
  解析：`%d`表示输出十进制数字，`+`表示输出数值的符号。这里不表示取反。
</details>