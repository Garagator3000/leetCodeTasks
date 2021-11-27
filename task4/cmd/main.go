//Сделать конвейер чисел

package main

import (
	"fmt"
)

func main() {
	channelIn := make(chan int)
	channelOut := make(chan int)

	go func(chin, chout chan int) {
		for i := 0; i < 5; i++ {
			chin <- i
		}
		close(chin)
	}(channelIn, channelOut)

	go func(chin, chout chan int) {
		for i := range chin {
			chout <- i*i*i
		}
		close(chout)
	}(channelIn, channelOut)

	for i := range channelOut {
		fmt.Println(i)
	}
}
