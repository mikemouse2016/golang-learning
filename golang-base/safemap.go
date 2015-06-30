// safemap
package main

import (
	"fmt"
	"sync"
)

type BeeMap struct {
	lock *sync.RWMutex
	bm   map[interface{}]interface{}
}

func NewBeeMap() *BeeMap {
	return &BeeMap{
		lock: new(sync.RWMutex),
		bm:   make(map[interface{}]interface{}),
	}
}

func (m *BeeMap) Get(k interface{}) interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if val, ok := m.bm[k]; ok {
		return val
	}
	return nil
}

func (m *BeeMap) Set(k interface{}, v interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	if val, ok := m.bm[k]; !ok {
		m.bm[k] = v
	} else if val != v {
		m.bm[k] = v
	} else {
		return false
	}
	return true
}

func (m *BeeMap) Check(k interface{}) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if _, ok := m.bm[k]; !ok {
		return false
	}
	return true
}

func (m *BeeMap) Delete(k interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.bm, k)
}

func main() {
	bee := NewBeeMap()
	bee.Set("123", 345)
	bee.Set("345", "abc")
	fmt.Println(bee.Get("123"))
	fmt.Println(bee.Check("123"))
	bee.Delete("123")
	fmt.Println(bee.Check("123"))
}
