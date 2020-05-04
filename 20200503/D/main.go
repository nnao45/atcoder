package main

import (
	"bufio"
	"fmt"
	"math"
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
	X := float64(nextInt())
	for i := -120; i <= 120; i++ {
		A := float64(i)
		for k := -120; k <= 120; k++ {
			B := float64(k)
			if math.Pow(A, 5)-math.Pow(B, 5) == X {
				fmt.Print(A, B)
				goto FIN
			}
		}
	}
FIN:
}
