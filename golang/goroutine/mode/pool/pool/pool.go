package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

/* 资源池基本管理
创建: 初始化资源池（设置大小，存放什么资源）
申请: 判段资源池是否关闭，是否大小满，满了创建新
释放: 判断资源池是否关闭，判断是否大小满，满了直接释放资源
关闭: 关闭资源池，在读完通道内数据
*/

// Pool 管理一组可以安全地在多个 goroutine 间共享的资源
// 被管理的资源必须实现 io.Closer 接口
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrPoolClosed 表示请求已关闭的池
var ErrPoolClosed = errors.New("Pool has been closed")

// New 创建管理资源的池
//   池需要分配资源的函数,并规定池的大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value too small")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire 从池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire: shared resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire: new resource")
		return p.factory()
	}
}

// Release 将一个使用后的资源放回池里
func (p *Pool) Release(r io.Closer) {
	// 保证操作和close操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	// 如果池被关闭，销毁这个资源
	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("release: in Queue")

	default:
		log.Println("release: closing")
		r.Close()
	}
}

// Close 会让资源池停止工作，并关闭所有现有的资源
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	// 在清空通道里的资源之前，将通道关闭
	// 如果不这样做，会发生死锁
	close(p.resources)

	// 不关闭通道，通道为空时for将阻塞，发生死锁
	for r := range p.resources {
		r.Close()
	}
}
