# 2022-5-17

Go 1.15 中 var i interface{} = a 会有额外堆内存分配吗？

具体代码是：

```go
var a  int = 3
// 以下有额外内存分配吗？
var i interface{} = a
```

<details open>
<summary>答案</summary>
在 Go 中，接口被实现为一对指针（请参阅 Russ Cox 的 Go 数据结构：接口）：指向有关类型信息的指针和指向值的指针。可以简单的表示为：

type iface struct {
    tab  *itab
    data unsafe.Pointer
}

然而，**Go 1.15 发行说明**在 runtime 部分中提到了一个有趣的改进：

Converting a small integer value into an interface value no longer causes allocation.
意思是说，将小整数转换为接口值不再需要进行内存分配。小整数是指 0 到 255 之间的数。

我们实际简单测试一下。

创建一个包 smallint，在包中创建文件 smallint.go，加上如下代码：

package smallint

func Convert(val int) []interface{} {
    var slice = make([]interface{}, 100)
    for i := 0; i < 100; i++ {
        slice[i] = val
    }

    return slice
}
为了更好的看到效果，函数中进行了 100 次 int 到 interface 的转换。写个基准测试 smallint_test.go：

package smallint_test

import (
    "testing"
    "test/smallint"
)

func BenchmarkConvert(b *testing.B) {
    for i := 0; i < b.N; i++ {
        result := smallint.Convert(12)
        _ = result
    }
}

分别使用 Go1.14 和 Go1.15 版本进行测试：

$ go version
go version go1.14.7 darwin/amd64
$ go test -bench . -benchmem ./...
goos: darwin
goarch: amd64
pkg: test/smallint
BenchmarkConvert-8      569830       1966 ns/op     2592 B/op      101 allocs/op
PASS
ok   test/smallint 1.647s
$ go version
go version go1.15 darwin/amd64
$ go test -bench . -benchmem ./...
goos: darwin
goarch: amd64
pkg: test/smallint
BenchmarkConvert-8     1859451        655 ns/op     1792 B/op        1 allocs/op
PASS
ok   test/smallint 2.178s
接着讲 smallint_test.go 中调用 Convert 的参数由 12 改为 256，再次使用 Go1.15 运行，结果如下：

$ go test -bench . -benchmem ./...
goos: darwin
goarch: amd64
pkg: test/smallint
BenchmarkConvert-8      551546       2049 ns/op     2592 B/op      101 allocs/op
PASS
ok   test/smallint 1.502s

证明了上面提到的优化点。

那么，你想过它大概怎么实现的吗？因为上文提到，Go 中接口的实现，使用一个指针字段指向接口值。现在竟然不再额外进行内存分配，说明做了什么特殊的事情。

其实答案非常简单。如果你对 Python、Java 等语言熟悉，应该知道大概如何实现的。Go 中如何做的，可以在 Go CL 216401 中（合并到**此提交**中了，GitHub 上更易于阅读）找到。
具体来说就是 Go 中定义了一个特殊的静态数组，该数组由 256 个整数组成（0 到 255）。当必须分配内存以将整数存储在堆上，并将其转换为接口的一部分时，它首先检查是否它可以只返回指向数组中适当元素的指针。这种经常使用的值的静态分配，是一种很常见的优化手段。例如，Python 对小整数执行类似的操作，Java 也有常量池，进行类似的优化处理。

实际上，Go 以前有一个优化，如果你将 0 转换为接口值，它将返回一个指向特殊静态零值的指针。这次新的 0-255 优化替代了该值。

对具体实现细节感兴趣的，可以阅读下上文提到的提交。

答案解析来自：https://mp.weixin.qq.com/s/1r0nt8nA3foDRRrbRp4omg
</details>