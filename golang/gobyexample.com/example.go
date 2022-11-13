package main

import (
	"errors"
	"fmt"
	"math"
	"sync"
	"time"
	"unicode/utf8"
)

// https://gobyexample.com/

func main() {
	// rateLimit()
	waitGroups()
	// workerPools()
	// times()
	// selects()
	// channels()
	// goroutines()
	// errorss()
	// generics()
	// structEmbedding()
	// interfaces()
	// structs()
	// stringsAndRunes()
	// pointers()
	// functions()
	// ranges()
	// maps()
	// slices()
	// arrays()
	// switchCase()
	// ifelse()
	// onlyfor()
	// constants()
	// variables()
	// values()
	// helloworld()
}

func rateLimit() {
	fmt.Println("hello world")
}

func waitGroups() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// Avoid re-use of the same i value in each goroutine closure.
		i := i
		go func() {
			defer wg.Done()
			worker3(i)
		}()
	}
	wg.Wait()
}

func worker3(id int) {
	fmt.Printf("worked %d staring\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worked %d done\n", id)
}

func workerPools() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 1; i <= 3; i++ {
		go worker2(i, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// 后面的执行与关闭代码位置无关
	close(jobs)

	for K := 1; K <= numJobs; K++ {
		res := <-results
		fmt.Println(res)
	}
}

func worker2(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("woker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("woker", id, "finished job", j)
		results <- j * 2
	}
}

func times() {
	timeouts()
	timeWait()
	timeTicker()
}

func timeTicker() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick at ", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	done <- true
	ticker.Stop()
	fmt.Println("ticker stopped")

}

func timeWait() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}

// Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time.
func timeouts() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

// Go’s select lets you wait on multiple channel operations.
func selects() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received:", msg1)
		case msg2 := <-c2:
			fmt.Println("received:", msg2)
		}
	}
}

// Channels are the pipes that connect concurrent goroutines.
func channels() {
	channel()
	channelBuffer()
	channelSync()
	channelDirections()
	channelNonBlock()
	channelClose()
	channelRange()
}

func channelRange() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

func channelClose() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			i, more := <-jobs
			if more {
				fmt.Println("received job", i)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("send job", i)
	}
	// deadlock if not closed
	close(jobs)
	fmt.Println("send all jobs")

	<-done
}

// Basic sends and receives on channels are blocking.
// Must be prepared at the same time
func channelNonBlock() {
	messages := make(chan string)
	signals := make(chan bool)

	// go func() {
	// 	<-messages
	// 	signals <- true
	// }()
	// time.Sleep(1 * time.Second)

	select {
	case msg := <-messages:
		fmt.Println("recevied message", msg)
	default:
		// run result
		fmt.Println("no message received")
	}

	// because the channel has no buffer and there is no receiver
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		// run result
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:

		fmt.Println("no activity")
	}

	time.Sleep(2 * time.Second)
}

// This specificity increases the type-safety of the program.
func channelDirections() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func channelSync() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func channelBuffer() {
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func channel() {
	message := make(chan string)
	go func() {
		message <- "ping"
	}()

	msg := <-message
	fmt.Println(msg)
}

func goroutines() {
	f("direct")
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func errorss() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg, ae.prob)
	}

}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

// 泛型 类似C++中的类模板
func generics() {
	var m = map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println("keys:", gmapKeys(m))

	m1 := gmapKeys[int, string](m)
	fmt.Println("m1:", m1)

	lst := glist[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}

func gmapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type glist[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *glist[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *glist[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func structEmbedding() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	fmt.Printf("co={num:%v,str:%v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describe:", d.describe())

}

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

// Interfaces are named collections of method signatures.
func interfaces() {
	r := recta{name: "recta", width: 3, height: 5}
	c := circle{name: "circle", radius: 5}
	measure(r)
	measure(c)
}

type geometry interface {
	area() float64
	perim() float64
}

type recta struct {
	name          string
	width, height float64
}

func (r recta) area() float64 {
	return r.width * r.height
}

func (r recta) perim() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	name   string
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// Go’s structs are typed collections of fields.
func structs() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.name, sp.age)
	sp.age = 51
	fmt.Println(sp, s)

	// methods
	r := rect{width: 10, height: 5}
	fmt.Printf("area:%v perim:%v\n", r.area(), r.perim())
	rp := &r
	fmt.Printf("area:%v perim:%v\n", rp.area(), rp.perim())
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

type rect struct {
	width, height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func stringsAndRunes() {
	// A Go string is a read-only slice of bytes.
	// byte: uint8
	// rune: int32
	const s = "嗨咯，世界"
	fmt.Printf("value:%v,type:%T\n", s, s)
	byt := []byte(s)
	fmt.Printf("len(s):%v,type:%T,byte:%v\n", len(s), byt, byt)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x,type:%T\t", s[i], s[i])
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	for idx, runeValue := range s {
		fmt.Printf("%#U stars at %d\n", runeValue, idx)
	}
	fmt.Println("\nUsing DecodeRuneInstring")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == '嗨' {
		fmt.Printf("found 嗨,type:%T\n", r)
	}
}

func pointers() {
	i := 1
	fmt.Println("init:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func functions() {
	a := function1(1, 2)
	fmt.Println(a)
	b := function2(1, 2, 3)
	fmt.Println(b)

	c, d := multipleReturnValues()
	fmt.Println(c, d)
	_, e := multipleReturnValues()
	fmt.Println(e)

	variadicFunctions(1, 2)
	variadicFunctions(3, 4, 5)
	nums := []int{6, 7, 8}
	variadicFunctions(nums...)

	//anonymous functions
	nextInt := closures()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := closures()
	fmt.Println(newInts())

	fmt.Println(recursion(7))
	// Closures can also be recursive, but this requires the closure to be declared with a typed var explicitly before it’s defined.
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}

func recursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursion(n-1)
}

func closures() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func variadicFunctions(nums ...int) {
	fmt.Printf("nums:%v, type:%T ", nums, nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total:", total)
}

func multipleReturnValues() (int, int) {
	return 3, 7
}

func function2(a, b, c int) int {
	return a + b + c
}

func function1(a int, b int) int {
	return a + b
}

func ranges() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key: ", k)
	}

	for i, c := range "go" {
		fmt.Printf("key:%v type:%T,\t value:%v string:%v type:%T\n", i, i, c, string(c), c)
	}
}

func maps() {
	var a map[string]int
	if a == nil {
		fmt.Println("a is nul ", a)
	}
	fmt.Printf("emp: %v, type: %T\n", a, a)

	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	m["k3"] = 12
	fmt.Println("map: ", m)
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	v, ok := m["k3"]
	if ok {
		fmt.Printf("exist v: %v, is true: %v\n", v, ok)
	} else {
		fmt.Printf("not exist v: %v, is false: %v\n", v, ok)
	}

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)
}

func slices() {
	var a []string
	if a == nil {
		fmt.Println("a is nul ", a)
	}
	fmt.Printf("emp: %v, type: %T, len: %v\n", a, a, len(a))
	s := make([]string, 3)
	fmt.Printf("emp: %v, type:%T,len: %v\n", s, s, len(s))
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s))
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)
	l := s[2:5]
	fmt.Println("sl1: ", l)
	l = s[:5]
	fmt.Println("sl2: ", l)
	l = s[2:]
	fmt.Println("sl3: ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func arrays() {
	var a [5]int
	fmt.Printf("emp: %v, type: %T\n", a, a)
	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	fmt.Println("len: ", len(a))
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func switchCase() {
	i := 2
	fmt.Print("zrite ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch { // if/else logic
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	case t.Hour() > 12:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func ifelse() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divsible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func onlyfor() {
	i := 1
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}
	for {
		fmt.Println("loop")
		break
	}
	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	a := []int{1, 2, 3}
	for k, v := range a {
		fmt.Println(k, v)
	}
}

func constants() {
	const s string = "constant"
	fmt.Println(s)
	// a numeric constant has no type until it’s given one
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}

func variables() {
	var a = "init"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e bool
	fmt.Println("bool:", e)
	var f int
	fmt.Println("int:", f)
	var g string
	fmt.Printf("string:%v type:%T \n", g, g)
	var h byte // uint8 代表一个ASCII码
	fmt.Printf("byte:%v type:%T \n", h, h)
	var i rune // int32 代表一个UTF-8字符
	fmt.Printf("rune:%v type:%T \n", i, i)
	j := "golang"
	fmt.Println(j)
}

func values() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func helloworld() {
	fmt.Println("hello world")
}
