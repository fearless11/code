package memochannel

type Func func(key string) (interface{}, error)

type entry struct {
	res   result
	ready chan struct{}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}

type result struct {
	value interface{}
	err   error
}

type request struct {
	key      string
	response chan<- result
}

type Memo struct {
	requests chan request
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)} // channel使用未关闭
	go memo.server(f)
	return memo
}

func (m *Memo) server(f Func) {
	cache := make(map[string]*entry)

	for req := range m.requests { // 通过channel获取请求实体
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response) // 向channel塞入响应结果
	}
}

func (m *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	m.requests <- request{key, response} // 向channel中塞一个请求实体
	res := <-response                    // 通过channel获取响应结果
	return res.value, res.err
}

func (m *Memo) Close() {
	close(m.requests) // 手动关闭channel
}
