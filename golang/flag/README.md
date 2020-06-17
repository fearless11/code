[toc]

#### 简介
[godoc-flag](https://godoc.org/flag#example-package) 

[标准库-命令行参数解析flag](http://blog.studygolang.com/2013/02/%E6%A0%87%E5%87%86%E5%BA%93-%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%8F%82%E6%95%B0%E8%A7%A3%E6%9E%90flag/)

[cobra-解析命令](https://github.com/spf13/cobra)

功能：flag实现命令行解析
```shell
# 用法
go run main.go -flag
go run main.go -flag=x
go run main.go -flag x  #bool类型的flag不能用,当x为false代表文件名将引起歧义。
```

#### 使用
##### 普通类型
```go

// 两种用法
// 参数： flag名, 默认值, 说明
var city = flag.String("city", "xian", "your name")

// 参数： 变量, flag名, 默认值, 说明
var name string

func init() {
	flag.StringVar(&name, "name", "vera", "your name")
}

func main() {
	flag.Parse()
	fmt.Println(*city, name)
}

//  usage: go run main.go -city shanghai -name lina
// output: shanghai lina
```
##### 自定义类型
```go
// 自定义类型要实现接口Set()、String()方法
type interval []time.Duration

func (i *interval) String() string {
	return fmt.Sprint(*i)
}

func (i *interval) Set(value string) error {
    if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

// 在init中调用flag.Var()实现
var intervalFlag interval
func init() {
    flag.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events")
}

func main() { 
    flag.Parse()
    fmt.Println(intervalFlag)
}

//  usage: go run main.go -deltaT 10s,9s,8s
// output: [10s 9s 8s]
```
##### usage

```go
func usage() {
	fmt.Fprintf(os.Stderr, `test: test/1.1.0
	Usage: test [-city]
Options:
`)
	flag.PrintDefaults()
}

var city = flag.String("city", "xian", "the city name")

func init() {
    // 重置默认的 Usage
    flag.Usage = usage
}

func main() { 
    flag.Parse()
    fmt.Println(*city)
}

//  usage: go run main.go -h 
/* output: 
test: test/1.1.0
        Usage: test [-city]
Options:
  -city string
        the city name (default "xian")
*/
```

#### 源码
#####  理解
- 标准库先定义一个通用的类型与方法`FlagSet`
  ```go
  type FlagSet struct {
		Usage func()
		name          string
		parsed        bool
		actual        map[string]*Flag   // 存放符合条件的命令flag
		formal        map[string]*Flag   // 存放程序预定义的flag
		args          []string     // arguments after flags
		errorHandling ErrorHandling
		output        io.Writer   // nil means stderr; use out() accessor
	}
  ```
- 再实例化一个`FlagSet`类型的实例`CommandLine`
   ```go
   func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet {
		f := &FlagSet{
			name:          name,
			errorHandling: errorHandling,
		}
		f.Usage = f.defaultUsage
		return f
	}

   var CommandLine = NewFlagSet(os.Args[0], ExitOnError)
   ```

- 所有可直接调用的方法是通过实例`CommandLine`实现
	```go
	func Parse() {
		// Ignore errors; CommandLine is set for ExitOnError.
		CommandLine.Parse(os.Args[1:])
	}
	```

##### 分析

一个`flag`的数据结构
```go
type Flag struct {
	Name     string // name as it appears on command line
	Usage    string // help message
	Value    Value  // value as set
	DefValue string // default value (as text); for usage message
}
```
接口灵活设定和打印不同类型数据
```go
// set()    : 实现从命令行接收到对应flag时重新设置flag的值
// string() : 两个功能，一是程序启动时读默认值赋值，二是fmt.Println打印
//  Value接口类型
type Value interface {
	String() string
	Set(string) error
}
```

	
###### 以`flag.String`为例

- 初始化关键步骤
  
  ```go
  // main包中设定
  var species = flag.String("species", "gopher", "the species we are studying")


  // flag包
  // 调用CommandLine实例
  func String(name string, value string, usage string) *string {
  	return CommandLine.String(name, value, usage)
  }

  // 和flag.StringVar()本质一样调用flagset.stringVar()
  func (f *FlagSet) String(name string, value string, usage string) *string {
		p := new(string)
		f.StringVar(p, name, value, usage)
		return p
   }

   // 本质调用f.Var()第一个参数Value是接口类型
   func (f *FlagSet) StringVar(p *string, name string, value string, usage string) {
		f.Var(newStringValue(value, p), name, usage)
   }

   // 自定义潜在类型stringValue, 实现了Set()和String()
   type stringValue string

	func newStringValue(val string, p *string) *stringValue {
		*p = val
		return (*stringValue)(p)
	}


   // 初始化的核心
   func (f *FlagSet) Var(value Value, name string, usage string) {
		// 读flag的默认值
		// value.String()调用的为stringValue.String()
		flag := &Flag{name, usage, value, value.String()} 
		// 判断是否flag重复设定
		_, alreadythere := f.formal[name]
		...
		// 将flag存放在formal里
		f.formal[name] = flag
	}

- 解析关键步骤
  ```go
  // main包
  flag.Parse()

  // flag包
  func Parse() {
	CommandLine.Parse(os.Args[1:])
  }

  func (f *FlagSet) Parse(arguments []string) error {
	seen, err := f.parseOne()
	...
   }

  // 解析核心
  func (f *FlagSet) parseOne() (bool, error) {
	  m := f.formal
	  // 读取存在的flag
	  flag, alreadythere := m[name] // BUG
	  // 设置新值
	  if err := flag.Value.Set(value); err != nil {				
	  	return false, f.failf("invalid value %q for flag -%s: %v", value, name, err)
	  }
	  ...
	  // 存放命令行flag结构中
	  f.actual[name] = flag
	  return true, nil
  }
