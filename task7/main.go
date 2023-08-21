package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

type ConcurrentMap struct {
	storage map[string]string
	sync.RWMutex
}

func (cm *ConcurrentMap) Read(key string) string {
	cm.RLock()
	defer cm.RUnlock()

	value, _ := cm.storage[key]
	return value
}

func (cm *ConcurrentMap) Write(key string, value string) {
	cm.Lock()
	defer cm.Unlock()

	cm.storage[key] = value
}

func NewConcurrentMap() *ConcurrentMap {
	m := make(map[string]string)
	return &ConcurrentMap{storage: m}
}

// if a lot of cores and high load we use sync.Map{} to avoid cache contention problem

func main() {
	cm := NewConcurrentMap()
	cm.Write("test_key1", "test_value1")
	cm.Write("test_key2", "test_value2")

	fmt.Println(cm.Read("test_key1"))
	fmt.Println(cm.Read("test_key2"))
}
