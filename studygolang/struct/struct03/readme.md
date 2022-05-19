# struct03

> 63

下面这段代码输出什么？

```go
type Direction int

const (
    North Direction = iota
    East
    South
    West
)

func (d Direction) String() string {
    return [...]string{"North", "East", "South", "West"}[d]
}

func main() {
    fmt.Println(South)
}

```

<details>
  <summary>答案</summary>
  **答：South**

  知识点：iota 的用法、类型的 String() 方法。

  当定义了一个有很多方法的类型时，一般你都会使用 String() 方法来定制本类型的可阅读性和打印性的、字符串形式的输出。
  如果类型定义了 String() 方法，它会被用在 fmt.Printf() 中生成默认的输出：等同于使用格式化描述符 %v 产生的输出。
  还有 fmt.Print() 和 fmt.Println() 也会自动使用 String() 方法。

  参考文献：https://learnku.com/docs/the-way-to-go/106-method/3644

</details>

