// 基准测试
// 一种测试代码性能的方法
//  想要测试解决同一问题的不同方案的性能，以及查看哪种解决方案的性能更好时，基准测试就会很有用
//  识别某段代码的 CPU或者内存效率问题，保证能最大化系统的吞吐量
package benchmark

import (
	"fmt"
	"strconv"
	"testing"
)

///////////////// 演示代码的功能： 找出将整数值转为字符串的最快方法 //////////////

// 基准测试框架默认会在持续 1 秒的时间内，反复调用需要测试的函数。测试框架每次调用测
// 试函数时，都会增加 b.N 的值。第一次调用时，b.N 的值为 1。需要注意，一定要将所有要进
// 行基准测试的代码都放到循环里，并且循环要使用 b.N 的值。否则，测试的结果是不可靠的。

// 命令：   go test -v -run="none" -bench="."
// output:
/*
// 函数-并发数           运行次数                每次操作的时间
BenchmarkSprintf-12     20000000                 74.1 ns/op
BenchmarkFormate-12     1000000000               2.36 ns/op
BenchmarkItoa-12        500000000                3.51 ns/op
PASS
ok      mytest/benchmark 6.611s （总耗时6s）
*/

// 命令： go test -v -run="none" -bench="." -benchmem
/*
BenchmarkSprintf-12     20000000                78.9 ns/op            16 B/op          2 allocs/op
BenchmarkFormate-12     1000000000               2.34 ns/op            0 B/op          0 allocs/op
BenchmarkItoa-12        500000000                3.51 ns/op            0 B/op          0 allocs/op
PASS
ok      mytest/benchmark 6.682s
*/
// BenchmarkSprintf-12     20000000                 74.1

// 单位为 ns/op的值表示每次操作op花费的时间为纳秒
// 单位为 allocs/op 的值表示每次操作从堆上分配内存的次数
// 单位为 B/op 的值表示每次操作分配的字节数

/*
# 生成文件
go test -cpuprofile cpu.prof -memprofile mem.prof -bench .

# 查看火焰图、性能等数据
go tool pprof -http=:8080 cpu.prof
*/

func BenchmarkSprintf(b *testing.B) {
	number := 10
	//  代码开始执行循环之前需要进行初始化时，这个方法用来重置计时器
	//  保证测试代码执行前的初始化代码，不会干扰计时器的结果。
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}

func BenchmarkFormate(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	number := 10
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
