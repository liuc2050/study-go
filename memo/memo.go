//Package memo provides a memoization of a function of type Func
package memo

import (
	"sync"
)

// Memo caches the results of calling a Func
type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

//Func is the type of function of memoize
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} //closed when res is ready
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (m *Memo) Get(key string) (interface{}, error) {
	m.mu.Lock()
	e := m.cache[key]
	if e != nil {
		m.mu.Unlock()
		<-e.ready
	} else {
		e = &entry{ready: make(chan struct{})}
		m.cache[key] = e
		m.mu.Unlock()
		e.res.value, e.res.err = m.f(key)
		close(e.ready)
	}
	return e.res.value, e.res.err
}
