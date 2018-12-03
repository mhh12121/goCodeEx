package main

import (
	"errors"
	"fmt"
)

func main() {
	s := ""
	boo := validPara(s)
	fmt.Println(boo)
}

//1. thread-unsafe
//
type stack struct {
	val  []byte
	size int
}

func (s *stack) Pop() (byte, error) {
	if s.size > 0 {
		temp := s.val[s.size-1]
		s.val = s.val[:s.size-1]
		s.size--
		return temp, nil
	} else {
		return ' ', errors.New("shit no data")
	}

}
func (s *stack) Push(b byte) error {
	s.val = append(s.val, b)
	s.size++
	return nil
}
func validPara(s string) bool {
	//two stack
	s1 := &stack{size: 0} //save left open
	s2 := &stack{size: 0} //right closed
	mbyte := []byte(s)
	if len(mbyte)%2 != 0 {
		return false
	}
	for _, v := range mbyte {
		switch v {
		case '{':
			{
				s1.Push(v)
			}
		case '}':
			{
				s2.Push(v)
				temp1, _ := s1.Pop()
				if temp1 == '{' {
					s2.Pop()
				} else {
					s1.Push(temp1)
				}
			}
		case '[':
			{
				s1.Push(v)
			}
		case ']':
			{
				s2.Push(v)
				temp1, _ := s1.Pop()
				if temp1 == '[' {
					s2.Pop()
				} else {
					s1.Push(temp1)
				}
			}
		case '(':
			{
				s1.Push(v)
			}
		case ')':
			{
				s2.Push(v)
				temp1, _ := s1.Pop()
				if temp1 == '(' {
					s2.Pop()
				} else {
					s1.Push(temp1)
				}
			}
		default:
			{
				fmt.Println("wtf the byte")
			}
		}

	}
	if s1.size == 0 && s2.size == 0 {
		return true
	}

	return false
}

//2. use 游标卡尺法 if met left brackets just add it to array and move 游标 to next index
//if met right brackets, can see if current value is corresponding left bracket, yes==> index--, no, return false
