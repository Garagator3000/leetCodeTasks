//На вход подаются два неупорядоченных слайса любой длины.
//Надо написать функцию, которая возвращает их пересечение

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sliceA, sliceB []int

	ReadSlice(&sliceA, os.Stdin)
	ReadSlice(&sliceB, os.Stdin)

	fmt.Println(sliceA)
	fmt.Println(sliceB)

	fmt.Println(Intersection(sliceA, sliceB))
}

func ReadSlice(slice *[]int, reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	for _, elem := range strings.Fields(scanner.Text()) {
		integer, err := strconv.Atoi(elem)
		if err == nil {
			*slice = append(*slice, integer)
		}
	}

}

func Intersection(sliceA, sliceB []int) []int {
	var intersection []int
	var counter = make(map[int]int)

	for _, elem := range sliceA {
		if _, ok := counter[elem]; !ok {
			counter[elem] = 1
		} else {
			counter[elem]++
		}
	}
	for _, elem := range sliceB {
		if count, ok := counter[elem]; ok && count > 0 {
			counter[elem] = 0
			intersection = append(intersection, elem)
		}
	}

	return intersection
}
