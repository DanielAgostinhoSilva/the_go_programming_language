package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() string {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func echo2() string {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3() string {
	return strings.Join(os.Args[1:], " ")
}

func main() {
	start := time.Now()
	result1 := echo1()
	duration1 := time.Since(start)
	fmt.Printf("Echo1 took %v to run and produced: %v\n", duration1, result1)

	start = time.Now()
	result2 := echo2()
	duration2 := time.Since(start)
	fmt.Printf("Echo2 took %v to run and produced: %v\n", duration2, result2)

	start = time.Now()
	result3 := echo3()
	duration3 := time.Since(start)
	fmt.Printf("Echo3 took %v to run and produced: %v\n", duration3, result3)
}
