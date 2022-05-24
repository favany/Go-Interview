# 9. 写出打印的结果

```go
type People struct {
  name string `json:"name"`
}

func main() {
  js := `{
    "name":"11"
  }`
  var p People
  err := json.Unmarshal([]byte(js), &p)
  if err != nil {
    fmt.Println("err: ", err)
    return
  }
  fmt.Println("people: ", p)
}
```

按照 go 的语法，⼩写开头的⽅法、属性或 struct 是私有的，同样，在 json 解码或转码的时候也⽆法实现私有属性的转换。
题⽬中是⽆法正常得到 People 的 name 值的。⽽且，私有属性 name 也不应该加 json 的标签。
