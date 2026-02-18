package main

import (
	"fmt"
	"sync"
)

var msg string

var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func race() {
	msg = "Hello, world"
	var mux sync.Mutex
	wg.Add(2)
	go updateMessage("Hello, universe", &mux)
	go updateMessage("Hello, cosmos", &mux)
	wg.Wait()
	fmt.Println(msg)
}
