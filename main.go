package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//注）<-chanは，型名の一部！　×チャネル名
//channelは，受信専用で使う
func input(r io.Reader) <-chan string {
	ch := make(chan string)

	//このgoroutineでは，channelは送信専用で使う
	go func(chan<- string) {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		ch = nil
	}(ch)

	//×<-ch. chから受け取った値が戻り値になってしまうため．
	return ch
}

func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print(">")
		fmt.Println(<-ch)
	}
}
