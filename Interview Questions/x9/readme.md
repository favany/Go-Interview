# 9

### 下面这段代码能否通过编译？不能的话，原因是什么？如果通过，输出什么？

```go
func main() {
    sn1 := struct {
        age  int
        name string
	}{age: 11, name: "qq"}
	sn2 := struct {
        age  int
        name string
	}{age: 11, name: "11"}

    if sn1 == sn2 {
        fmt.Println("sn1 == sn2")
    }

    sm1 := struct {
        age int
        m   map[string]string
    }{age: 11, m: map[string]string{"a": "1"}}
    sm2 := struct {
        age int
        m   map[string]string
    }{age: 11, m: map[string]string{"a": "1"}}

    if sm1 == sm2 {
        fmt.Println("sm1 == sm2")
    }
}
```

```shell
# command-line-arguments
./main.go:29:5: invalid operation: sm1 == sm2 (struct containing map[string]string cannot be compared)
```

考点是结构体的比较，有几个需要注意的地方：

结构体只能比较是否相等，但是不能比较大小；
1. 相同类型的结构体才能进行比较，结构体是否相同不但与属性类型有关，还与属性顺序相关；
2. 如果struct的所有成员都可以比较，则该struct就可以通过==或!=进行比较是否相同，比较时逐个项进行比较，如果每一项都相等，则两个结构体才相等，否则不相等；

那有什么是可以比较的呢？ 常见的有：bool、数值型、字符、指针、数组等
不能比较的有：slice、map、函数