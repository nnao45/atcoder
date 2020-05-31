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
	sc.Split(bufio.ScanWords)
	A := float64(nextInt())
	B := float64(nextInt())
	N := float64(nextInt())
	maxLeft := math.Floor((A * N) / B)
	maxRight := math.Floor((N) / B)
	// answer := int(math.Floor((A*x)/B)) - int(A*math.Floor(x/B))
	fmt.Print(maxLeft, maxRight)
}
