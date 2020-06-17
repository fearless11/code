// package memo concurrency safe
//  通过共享变量 cache map[string]*entry实现各个goroutine间的通信
package memoshare

import (
	"sync"
)

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (m *Memo) Get(key string) (interface{}, error) {

	//  不重复
	//  这次加锁,同时确保,同一个url只会get一次
	// 不存在的url进来,先向map插入一条初始化的值,后续查询同类url的将直接跳转到else块
	// 因为else块的执行,依赖channel。只有当channel关闭,说明该数据有值时才会返回。
	m.mu.Lock()
	e := m.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		m.cache[key] = e
		m.mu.Unlock()

		//  无阻塞
		e.res.value, e.res.err = m.f(key)
		close(e.ready)
	} else {
		m.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}
