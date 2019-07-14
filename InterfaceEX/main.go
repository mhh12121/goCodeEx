package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (bc *ByteCounter) wordsCount(p []byte) (int, error) {
	flag := false //有单词
	temp := 0
	for _, v := range p {
		if v == ' ' && flag {
			temp++
			flag = false
		} else if v != ' ' {
			flag = true
		}

	}
	if flag {
		temp++
	}
	*bc += ByteCounter(temp)
	return temp, nil

}

type LineCounter int

// func (lc *LineCounter) lineCount(p []byte) (int, error) {

// }
func main() {
	f, err := os.Open("../InterfaceEX/inter.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		adv, token, _ := bufio.ScanWords(line, false)
		fmt.Printf("adv:%v,token:%v", adv, string(token))
		var c ByteCounter
		c.wordsCount(line)
		fmt.Println(c)
		if err != nil {
			if err == io.EOF {
				fmt.Println("end")
				break
			}
		}
	}

	// y := []byte("dsafdfaf fdfdfd fqwwerqw  ffdfd  ")

}
