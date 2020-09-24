package memo

import (
	"fmt"
)

// Func is the type of the function to memoize.
type Func func(key string, done <-chan struct{}) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result   // the client wants a single result
	done     <-chan struct{} // 获取取消通知
}

type Memo struct {
	requests chan request
	cancels  chan string
}

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{make(chan request), make(chan string)}
	go memo.server(f)
	return memo
}

// 通过传入的 done 通道，可以在 Get 调用中的某个时刻，通知是正常请求并缓存，还是取消请求并清空缓存
func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {

	response := make(chan result)
	req := request{key, response, done}
	memo.requests <- req // 发出请求

	res := <-response // 阻塞等待响应

	select {
	case <-done:
		memo.cancels <- key // 宣告取消，并清空缓存
	default:
		// nothing to do
	}

	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry) // 在 server 内部由单一 goroutine 访问，没有并发冲突问题
MainLoop:
	for {
	CancelChecker:
		// 请求之前，判断是否已取消缓存
		// 保证被标记为失效的缓存，不会返回给任何 goroutine
		for {
			select {
			case key := <-memo.cancels:
				fmt.Printf("key: [%s] cancelled before request\n", key)
				delete(cache, key)
			default:
				break CancelChecker
			}
		}

		// 如果不执行前面的 for + select 组合，只执行以下的 select，则有可能会出现如下情况：
		// 对于同时有效的 case，select 会随机选择
		// 若同名的 key 在 memo.requests 和 memo.cancels 同时存在，且该 key 对应的请求已经缓存过
		// 而 select 随机选择了 memo.requests，就会返回一个被标记为失效但尚未来得及删除的缓存结果
		select {
		case key := <-memo.cancels: // 在所有 goroutine 间共享
			fmt.Printf("key: [%s] cancelled\n", key)
			delete(cache, key)
			continue MainLoop
		case req := <-memo.requests: // 在所有 goroutine 间共享
			fmt.Printf("request key: %s\n", req.key)
			e := cache[req.key]
			if e == nil {
				// This is the first request for this key.
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key, req.done) // call f(key, done)
			}
			go e.deliver(req.response)
		}
	}
}

func (e *entry) call(f Func, key string, done <-chan struct{}) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key, done) // 要求 Func 的实现 f 对于 done 给予支持，比如取消 http 请求等
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}
