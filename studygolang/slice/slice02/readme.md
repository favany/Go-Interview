# slice02

下面的代码有什么问题？

```go
func main() {
    fmt.Println([...]int{1} == [2]int{1})
    fmt.Println([]int{1} == []int{1})
}
```

<details>
  <summary>答案</summary>
  **答：有两处错误**
  **解析：**
  - go 中不同类型是不能比较的，而数组长度是数组类型的一部分，所以 `[…]int{1}` 和 `[2]int{1}` 是两种不同的类型，不能比较；
  - 切片是不能比较的；
</details>
