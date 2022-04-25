package main

import "fmt"

func main() {

	// 这里我们创建了一个数组 `a` 来存放刚好 5 个 `int`
	// 元素的类型和长度都是数组类型的一部分
	// 数组默认是零值的，对于 `int` 数组来说也就是 `0`
	var a [5]int
	fmt.Println("声明一个五个元素的整数数组，结果是:", a)

	// 我们可以使用 `array[index] = value` 语法来设置数组
	// 指定位置的值，或者用 `array[index]` 得到值
	a[4] = 100
	fmt.Println("设置a[4]=100，得到a:", a)
	fmt.Println("取得a[4]:", a[4])

	// 使用内置函数 `len` 返回数组的长度
	fmt.Println("获取数组长度len(a)，得到结果：", len(a))

	// 在一行内同时声明和初始化数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("同时声明和初始化数组的结果 b是:", b)

	// 数组类型是一维的，但是你可以组合构造多维的数据结构。
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i*10 + j
		}
	}
	fmt.Println("二维数组示例: ", twoD)
}
