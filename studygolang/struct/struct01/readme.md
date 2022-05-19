# struct01

> 2022-05-19

下面代码是否可以编译通过？为什么？

```go
package main

import "fmt"

func main() {

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

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


<details>
  <summary>答案</summary>
  结构体比较⚠️
  规则1：只有相同类型的结构体才可以比较。结构体是否相同不但与属性类型个数有关，还与属性顺序相关。
  比如：

  sn1 := struct {
    age  int
    name string
  }{age: 11, name: "qq"}

  sn3:= struct {
      name string
      age  int
  }{age:11, name:"qq"}
  sn3与sn1就不是相同的结构体了，不能比较。

  规则2：结构体是相同的，但是结构体属性中有不可以比较的类型，如map,slice，则结构体不能用==比较。
  可以使用reflect.DeepEqual进行比较

  ```go
  if reflect.DeepEqual(sm1, sm2) {
      fmt.Println("sm1 == sm2")
  } else {
      fmt.Println("sm1 != sm2")
  }
  ```
</details>

