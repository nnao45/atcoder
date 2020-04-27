package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	A := nextInt() // takahashi physical
	B := nextInt() // talahashi attack
	C := nextInt() // aoki physical
	D := nextInt() // aoki attack

	for {
		C = C - B
		if C <= 0 {
			fmt.Print("Yes")
			break
		}
		A = A - D
		if A <= 0 {
			fmt.Print("No")
			break
		}
	}
}
