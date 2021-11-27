//Написать генератор случайных чисел

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	max, _ := strconv.Atoi(os.Args[2])


	for num := range RandNumsGenerator(max, n) {
		fmt.Println(num)
	}
}

func RandNumsGenerator(max, n int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var out = make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			out <- r.Intn(max)
		}
		close(out)
	}()
	return out
}
