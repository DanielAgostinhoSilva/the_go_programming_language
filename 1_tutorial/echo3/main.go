package main

import (
	"fmt"
	"os"
	"strings"
)

// $ go run main.go a b c d e f
func main() {
	fmt.Println(strings.Join(os.Args[1:], ""))
}
