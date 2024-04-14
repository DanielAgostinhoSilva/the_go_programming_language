package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[:], ""))
	for index, value := range os.Args {
		fmt.Printf("index: %d value: %s\n", index, value)
	}
}
