

### go test


[godoc-testing](https://godoc.org/testing)

```bash
-run|-bench    # 接受任意正则表达式
-benchmem      # 查看每次操作分配内存的次数，以及总共分配内存的字节数

```

```bash
go test -run ''      # Run all tests.
go test -run Foo     # Run top-level tests matching "Foo", such as "TestFooBar".
go test -run Foo/A=  # For top-level tests matching "Foo", run subtests matching "A=".
go test -run /A=1    # For all top-level tests, run subtests matching "A=1".

# 因单元测试名中有none，所以排除所有单元测试，运行基准测试
go test -v  -run="none" -bench="BenchmarkSprintf" 

go test -run Benchmark -bench Benchmark    
go test -race       # running parallel subtests

# 运行所有基准测试，同时查看内存情况
go test -v -run="none" -bench="." -benchmem  
```

  
### 单元测试
- 格式： `func TestXxx(*testing.T)`

- 基本测试
- 表组测试
- mock测试
- 服务端测试

```go
func TestAbs(t *testing.T) {
	got := abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d ; want 1", got)
	}
}

usage: go test
output: 
PASS
ok      mytest/test        0.348s
```


### 示例测试

- 格式: `func Example() {}`

- 示例代码的函数必须是已经存在的公开的函数或方法，首字母大写

```
示例函数命名 a function F, a type T and method M on type T 
func Example() { ... }
func ExampleF() { ... }
func ExampleT() { ... }
func ExampleT_M() { ... }
```

```go
func Example() {
	// test ok
	fmt.Println("hello")
	// Output: hello

	// test fail
	fmt.Println("hi")
	// Output: hello
}

usage: go test
output: 
--- FAIL: Example (0.00s)
got:
hello
hi
want:
hello
FAIL
exit status 1
FAIL    mytest/test        0.367s
```



### 性能测试
格式： `func BenchmarkXxx(*testing.B)`

```go
func BenchmarkHello(b *testing.B) {
  b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintln("hello")
	}
}

usage: go test -run Benchmark -bench=Benchmark
output:
goos: windows
goarch: amd64
pkg: mytest/test
// 函数-并发数                            操作次数                每次操作的时间
BenchmarkHello-12                       30000000                56.8 ns/op
PASS
ok      mytest/test        5.915s // 总耗时
```

