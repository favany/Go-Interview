# slice04

```go
package main

import (
  "fmt"
)

func main(){
  s := make([]int, 10)

  s = append(s, 1, 2, 3)

  fmt.Println(s)
}
```

<details>
  <summary>答案</summary>

[0 0 0 0 0 0 0 0 0 0 1 2 3]
考点：切片追加, make 初始化均为 0

</details>
