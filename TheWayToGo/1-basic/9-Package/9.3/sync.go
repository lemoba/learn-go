package main

import "sync"

type Info struct {
	mu sync.Mutex
}

func main() {

}

func update(info *Info) {
	info.mu.Lock()
	// info.Str = "ch1"
	info.mu.Unlock()
}
