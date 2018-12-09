package main

import (
	"fmt"
	"sync/atomic"
	"testing"
)

var value int32

func Test_atmoicChange(t *testing.T) {
	delta := 2
	SetValue(int32(delta))
	fmt.Println(value)
}
func SetValue(delta int32) {

	v := value
	if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {

	}

}
