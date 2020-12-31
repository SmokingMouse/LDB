package main

import (
	"LDB/buffer_manager"
)

func main() {
	lru := buffer_manager.NewLru(1)
	lru.Set(20, 20)
	lru.Set(30, 30)
}
