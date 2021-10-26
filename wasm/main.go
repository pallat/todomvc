package main

import (
	"github.com/pallat/todomvc/wasm/todo"
)

var signal = make(chan int)

func keepAlive() {
	for {
		<-signal
	}
}

func main() {
	todo.Start()
	keepAlive()
}
