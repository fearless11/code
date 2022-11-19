package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	s "strings"
	"sync"
	"sync/atomic"
	"text/template"
	"time"
	"unicode/utf8"
)

// https://gobyexample.com/

func main() {
	// helloworld()
	// values()
	// variables()
	// constants()
	// onlyfor()
	// ifelse()
	// switchCase()
	// arrays()
	// slices()
	// maps()
	// ranges()
	// functions()
	// pointers()
	// stringsAndRunes()
	// structs()
	// interfaces()
	// structEmbedding()
	// generics()
	// errorss()
	// goroutines()
	// channels()
	// selects()
	// times()
	// workerPools()
	// waitGroups()
	// rateLimiting()
	// atomicCounters()
	// mutexes()
	// statefulGoroutines()
	// sorting()
	// panics()
	// defers()
	// recovers()
	// stringsFunc()
	// fmtFunc()
	// textTemplates()
	// regularExpressions()
	// jsons()
	// xmls()
	// timeFormat()
	// randomNum()
	// numberParse()
	// urlParse()
	// sha256Hashes()
	// base64Encoding()
	// readFiles()
	// writeFiles()
	// lineFilters()
	// filePaths()
	directories()

}

func directories() {
	fmt.Println("HELLO WORLD")
}

func filePaths() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))
	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)
	fmt.Println(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

}

// echo "hello\ngo\n" > /tmp/data1
// cat /tmp/data1 | go run example.go
func lineFilters() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}
}

// overwrite write
func writeFiles() {
	b1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/data1", b1, 0664)
	check(err)

	f, err := os.Create("/tmp/data2")
	check(err)
	defer f.Close()

	b2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(b2)
	check(err)
	fmt.Printf("write %d bytes\n", n2)

	n3, err := f.WriteString("happy\n")
	check(err)
	fmt.Printf("write %d bytes\n", n3)
	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("bufio\n")
	check(err)
	fmt.Printf("write %d bytes\n", n4)
	w.Flush()
}

func readFiles() {
	data, err := os.ReadFile("/tmp/data")
	check(err)
	fmt.Print(string(data))

	f, err := os.Open("/tmp/data")
	check(err)
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes@ %d: %v\n", n2, o2, string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes@ %d: %v\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	f.Close()

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func base64Encoding() {
	data := "abc123!?$*&()'-=@~"
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}

// For example, TLS/SSL certificates use SHA256 to compute a certificate’s signature.
func sha256Hashes() {
	s := "sha256 this string"
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%d %x\n", len(bs), bs)

	h1 := sha512.New()
	h1.Write([]byte(s))
	bs1 := h1.Sum(nil)
	fmt.Printf("%d %x\n", len(bs1), bs1)

}

func urlParse() {
	s := "postgres://user:pass@host.com:5432/path?k=v&k1=v1#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User.String())
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}

func numberParse() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)

}

func randomNum() {
	p := fmt.Println
	// 0 <= n < 100
	p(rand.Intn(100), rand.Intn(100))
	// 0.0 <= f < 1.0
	p(rand.Float64())
	// 5.0 <= f < 10.0
	p((rand.Float64() * 5) + 5)

	// nice
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", r1.Intn(100))
	}
	p()

	// If you seed a source with the same number,
	// it produces the same sequence of random numbers.
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", r2.Intn(100))
	}
	p()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", r3.Intn(100))
	}
}

// time comment
// 1: month (January, Jan, 01, etc)
// 2: day (02)
// 3: hour (15 is 3pm on a 24 hour clock)
// 4: minute (04)
// 5: second (05)
// 6: year (2006)
// 7: timezone (GMT-7 is MST)
func timeFormat() {
	p := fmt.Println
	now := time.Now()
	p(now)
	then := time.Date(2022, 11, 19, 8, 20, 30, 10, time.UTC)
	p(then)
	p(then.Year(), then.Month(), then.Day())
	p(then.Hour(), then.Minute(), then.Minute(), then.Nanosecond())
	p(then.Location())
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())
	fmt.Println(time.Unix(0, 0))
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))

	p(now.Format(time.RFC3339))
	t1, _ := time.Parse(time.RFC3339, "2020-11-20T01:22:40Z")
	p(t1)
	p(now.Format("3:04PM"))
	p(now.Format("Mon Jan _2 15:04:05 2006"))
	p(now.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	asinc := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(asinc, "8:41PM")
	p(e)

	// 2006 01 02 03 04 05
	// year month day hour minute second
	cus := "2006-01-2 15:04:05"
	t3, e := time.Parse(cus, "2022-11-20 23:19:59")
	p(t3, e)
}

func xmls() {
	coffee := &plant{ID: 27, Name: "coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	out, _ := xml.MarshalIndent(coffee, " ", " ")
	fmt.Println(string(out))
	fmt.Println(xml.Header + string(out))

	var p plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &plant{ID: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*plant `xml:"parent>child>plant"`
	}
	nesting := &Nesting{}
	nesting.Plants = []*plant{coffee, tomato}
	out, _ = xml.MarshalIndent(nesting, " ", " ")
	fmt.Println(string(out))
}

type plant struct {
	XMLName xml.Name `xml:"plant"`
	ID      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p plant) String() string {
	return fmt.Sprintf("plant id=%d,name=%v,origin=%v", p.ID, p.Name, p.Origin)
}

func jsons() {
	bolB, _ := json.Marshal(true)
	fmt.Println(bolB, string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(intB, string(intB))
	fltB, _ := json.Marshal(3.14)
	fmt.Println(fltB, string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(strB, string(strB))
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	slcB, _ = json.MarshalIndent(slcD, " ", " ")
	fmt.Println(slcB, string(slcB))
	mapD := map[string]int{"apple": 5, "pear": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(mapB, string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	// pretty format
	res1B, _ = json.MarshalIndent(res1D, " ", " ")
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":3.14,"strs":["a","b"]}`)
	// nice
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(float64)
	fmt.Println(num)
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page":1,"fruits":["apple","pear"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res, res.Page)
	fmt.Printf("%+v\n", res)

	// stream JSON encodings
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "pear": 7}
	enc.Encode(d)
}

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func regularExpressions() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))
	fmt.Println("idx:", r.FindStringIndex("peach punch"))
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindAllString("peach punch", -1))
	fmt.Println("idx:", r.FindAllStringIndex("peach punch", -1))
	fmt.Println("idx:", r.FindAllStringSubmatchIndex("peach punch", -1))
	fmt.Println(r.Match([]byte("peach")))
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)
	fmt.Println(r.ReplaceAllString("a peach", "fruit"))
	in := []byte("peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}

func textTemplates() {
	t1 := template.New("t1")
	t1, err := t1.Parse("value is {{ . }}\n")
	if err != nil {
		panic(err)
	}
	// the template.Must function to panic in case Parse returns an error.
	t1 = template.Must(t1.Parse("value is {{ . }}\n"))
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{"Go", "Python", "Javascript"})

	create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := create("t1", "Name: {{.Name}}\n")
	t2.Execute(os.Stdout, struct {
		Name string
	}{"Mary"})
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Tom",
	})

	t3 := create("t3",
		"{{ if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, " not empty")
	t3.Execute(os.Stdout, "")

	t4 := create("t4",
		"Range: {{range .}} {{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{"Go", "C", "Python"})
}

func fmtFunc() {
	p := fpoint{1, 2}
	fmt.Printf("struct: %v\n", p)
	fmt.Printf("struct: %+v\n", p)
	fmt.Printf("struct: %#v\n", p)
	fmt.Printf("type: %T\n", p)
	fmt.Printf("pointer: %p\n", &p)

	fmt.Printf("bool: %t\n", true)
	fmt.Printf("int: %d\n", 123)
	fmt.Printf("bin: %b\n", 14)
	fmt.Printf("char: %c\n", 33)
	fmt.Printf("hex: %x\n", 10)
	fmt.Printf("float: %f\n", 78.9)
	fmt.Printf("float: %e\n", 123400000.0)
	fmt.Printf("float: %E\n", 123400000.0)
	fmt.Printf("str: %s\n", "\"string\"")
	fmt.Printf("str: %q\n", "\"string\"")
	fmt.Printf("str: %x\n", "a b")

	fmt.Printf("wdith: |%6d|%6d|\n", 12, 345)
	fmt.Printf("wdith: |%-6d|%-6d|\n", 12, 345)
	fmt.Printf("wdith: |%6.3f|%6.3f|\n", 1.2, 3.45)
	fmt.Printf("wdith: |%-6.3f|%-6.3f|\n", 1.2, 3.45)
	fmt.Printf("wdith: |%6s|%6s|\n", "foo", "b")
	fmt.Printf("wdith: |%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}

type fpoint struct {
	x, y int
}

type point struct {
	x, y int
}

func stringsFunc() {
	var p = fmt.Println
	p("contains: ", s.Contains("test", "es"))
	p("count:", s.Count("test", "t"))
	p("hasSuffix:", s.HasPrefix("test", "te"))
	p("hasSuffix:", s.HasSuffix("test", "st"))
	p("index:", s.Index("test", "e"))
	p("join:", s.Join([]string{"a", "b"}, "-"))
	p("repeat:", s.Repeat("a", 5))
	p("replace:", s.Replace("foo", "o", "a", -1))
	p("replace:", s.Replace("foo", "o", "a", 1))
	p("split:", s.Split("a-b-c-d", "-"))
	p("ToLower:", s.ToLower("TEST"))
	p("ToUpper:", s.ToUpper("test"))
}

func recovers() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered. Error: ", r)
		}
	}()

	panic("a problem")
	fmt.Println("after panic")
}

func defers() {
	// LIFO
	defer func() {
		fmt.Println("one")
	}()
	defer func() {
		fmt.Println("two")
	}()

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func panics() {
	panic("a problem")
	fmt.Println("panic")
}

func sorting() {
	strs := []string{"c", "a", "b", "d"}
	sort.Strings(strs)
	fmt.Println("strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted:", s)

	sortingByFunc()
}

func sortingByFunc() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func statefulGoroutines() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println(readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println(writeOpsFinal)
}

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func mutexes() {
	c := counter{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.Inc(name)
		}
		wg.Done()
	}
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	wg.Wait()
	fmt.Println(c.counters)

}

type counter struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *counter) Inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func atomicCounters() {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				// because the goroutines would interfere with each other. Moreover, we’d get data race failures when running with the -race flag.
				// ops++
				// atomics safely
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("ops:", ops)
}

func rateLimiting() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("bursty")

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

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

	fmt.Println("interface word String,", new(word))

	r := recta{name: "recta", width: 3, height: 5}
	c := circle{name: "circle", radius: 5}
	measure(r)
	measure(c)
}

type word struct{}

func (w *word) String() string {
	return "hello world"
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
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
