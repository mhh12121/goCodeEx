package main

import "runtime"

func main() {
	var buf []byte
	runtime.Stack(buf, true)
}
