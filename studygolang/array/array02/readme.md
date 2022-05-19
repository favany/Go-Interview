# array02

下面这段代码输出什么？

```go
func main() {
    a := [2]int{5, 6}
    b := [3]int{5, 6}
    if a == b {
        fmt.Println("equal")
    } else {
        fmt.Println("not equal")
    }
}
```

- A. compilation error
- B. equal
- C. not equal

<details>
  <summary>答案</summary>
  **答：A**

  ```shell
  invalid operation: cannot compare a == b (mismatched types [2]int and [3]int)
  ```

  **解析：**
  Go中的数组是值类型，可比较，另外一方面，数组的长度也是数组类型的组成部分，所以 `a` 和 `b` 是不同的类型，是不能比较的，所以编译错误。
</details>