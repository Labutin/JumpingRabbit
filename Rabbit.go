package main

import (
	"flag"
	"fmt"
	"strconv"
)

func solve(path string, N int, K int) {
	if N > 0 {
		for i := 1; i <= K && i <= N; i++ {
			newPath := path
			if path != "" {
				newPath += " + "
			}
			newPath += strconv.Itoa(i)
			solve(newPath, N - i, K)
		}
	} else {
		fmt.Printf("%s\n", path)
	}
}

func main() {
	N := flag.Int("N", 1, "N")
	K := flag.Int("K", 1, "K")
	flag.Parse()

	solve("", *N, *K);
}
