package main

import (
	"fmt"
	"os"
)

type pairs struct {
	a int
	b int
}

func main() {
	var (
		word string
		q    int
	)
	fmt.Fscan(os.Stdin, &word, &q)
	dataSlice := make([]pair, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(os.Stdin, &dataSlice[i].a, &dataSlice[i].b)
	}

}
