package main

import (
	"fmt"

	stack "github.com/golang-collections/collections/stack"
)

const (
	a = 1 << iota //1
	b             //2
	c             //4
	x = iota      //3

)

// var i int

func dosomething() func(int) int {
	i := 1
	return func(delta int) int {
		i += delta
		return i
	}
}
func FibonacciClosure() func() int {
	i, j := 0, 1
	return func() int {
		i, j = j, i+j
		fmt.Println("number now", j)
		return j
	}
}

type douRoute struct {
	val  string
	next []*douRoute
}

func main() {

	// target := 5
	// result := twoSum(intputnums, target)
	// fmt.Println(result)

	//----two sum-----
	// l1 := &ListNode{}
	// l1.Val = 0
	// l2 := &ListNode{}
	// l2.Val = 1
	// l2.Next = &ListNode{Val: 8, Next: nil}
	// addTwoNumbers(l1, l2)
	//------------------
	// fmt.Println(a, b, c, d, e, x)

	// fmt.Println(a, b, c, x)

	// s := 'a'
	// fmt.Println(s)
	// // x := dosomething()
	// y := FibonacciClosure()
	// for i := 0; i < 5; i++ {
	// 	y()
	// }

	// url := "/main/v{version}/getSomething"
	// urlArr := strings.Split(url, "/")
	// route := make([]string, len(urlArr))

	// // route[urlArr[i]]
	// reg, _ := regexp.Compile("$/[*]/main/controller/\\w*^")
	// fmt.Println(reg.MatchString("/\\w*/main/controller/\\w*"))

	// fmt.Println(x(1), x(2), x(3))
	// m := &minStack{stack1: stack.New(), stack2: stack.New()}
	// m.push(4)
	// m.push(6)
	// m.push(8)
	// m.push(2)
	// fmt.Println(m.min())
	// m.pop()
	// fmt.Println(m.min())

	// c := cron.New()
	// context.WithValue()
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// c.AddFunc("@every 1s", func() {
	// 	cmdx := exec.CommandContext(ctx, "date", "-R")
	// 	cmdx.Stdout = os.Stdout
	// 	err := cmdx.Run()
	// 	if err != nil {

	// 		fmt.Println("cmd run error:", err)
	// 	}
	// 	// fmt.Println("yoyo")
	// })
	// c.Start()
	// select {}

	// y := make([]int, 0)
	// y = append(y, 1)
	// x := y[1:len(y)]
	// fmt.Println(x)

	// a := make([]int, 3, 4)
	// a[0] = 1
	// a[1] = 2
	// a[2] = 3

	// b := append(a, 4)
	// c := append(a, 100)
	// fmt.Println(&a[0], &b[0], &c[0]) //Found that they use the same memory address
	// fmt.Println(a, b, c)             //
	// c[0] = 101
	// fmt.Println(a, b, c)
	s := []int{1, 1, 1}
	do(s)
	fmt.Println(s)
}
func do(s []int) {
	for x, _ := range s {
		s[x]++
	}
	// for i := range s {
	// 	s[i]++
	// }
}

type minStack struct {
	stack1 *stack.Stack
	stack2 *stack.Stack
}

func (ms *minStack) pop() {
	if ms.stack1.Peek() == ms.stack2.Peek() {
		ms.stack2.Pop()
	}
	ms.stack1.Pop()
}
func (ms *minStack) push(node int) {
	ms.stack1.Push(node)
	if ms.stack2.Len() == 0 {
		ms.stack2.Push(node)
	} else if node <= ms.stack2.Peek().(int) {
		ms.stack2.Push(node)
	}
}

func (ms *minStack) top() int {
	return ms.stack1.Pop().(int)
}

func (ms *minStack) min() int {
	return ms.stack2.Peek().(int)
}
