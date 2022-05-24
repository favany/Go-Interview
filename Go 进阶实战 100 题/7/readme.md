# 7. 下⾯代码能运⾏吗？为什么

```go
type Param map[string]interface{}
type Show struct {
 Param
}
func main1() {
 s := new(Show)
 s.Param["RMB"] = 10000
}
```

<details>
  <summary>答案</summary>
  共发现两个问题：
  1. main 函数不能加数字。
  2. new 关键字⽆法初始化 Show 结构体中的 Param 属性，所以直接对 s.Param 操作会出错。
</details>