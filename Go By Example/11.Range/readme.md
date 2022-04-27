# 11. for ... Range 循环遍历

range 能迭代各种数据结构。让我们来看看如何在已经学过的数据结构上使用 `for ... range` 做循环遍历。

这里我们使用 `range` 来对 slice 中的元素求和。对于数组也可以采用这种方法。
```go
nums := []int{2, 3, 4}
sum := 0
for _, num := range nums {
    sum += num
}
fmt.Println("sum:", sum)
```

`range` 在数组和 slice 中提供对其中每一项的索引 index 和值 value 的访问。
上面我们不需要索引，所以我们使用空白标识符 `_` 来忽略它。有时候我们实际上是需要这个索引的。

```go
for i, num := range nums {
    if num == 3 {
		fmt.Println("index:", i)
    }
}
```

`range` 在 map 中迭代键值对。
```go
kvs := map[string]string{"a": "apple", "b": "banana"}
for k, v:=range kvs {
    fmt.Printf("%s -> %s\n", k, v)
}
```

`range` 也可以只遍历 map 的键。

```go
for k := range kvs {
    fmt.Println("key:", k)
}
```

`range` 在字符串中迭代 unicode 码点(code point)。
第一个返回值是字符的起始字节位置，然后第二个是字符本身。
```go
for i, c := range "go" {
    fmt.Println(i, c, string(c))
}
```


