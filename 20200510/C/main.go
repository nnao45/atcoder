package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func stringsFiledsToInts() []int {
	sc.Scan()
	r := []int{}
	ss := strings.Fields(sc.Text())
	for _, s := range ss {
		i, e := strconv.Atoi(s)
		if e != nil {
			panic(e)
		}
		r = append(r, i)
	}
	return r
}

func main() {
	NMX := stringsFiledsToInts()
	N := NMX[0]
	M := NMX[1]
	X := NMX[2]
	A := [][]int{}

	allCost := make([]int, M+1)

	for i := 1; i <= N; i++ {
		book := stringsFiledsToInts()
		for j, b := range book {
			allCost[j] += b
		}
		A = append(A, book)
	}

	for _, cost := range allCost {
		if cost < X {
			fmt.Print("-1")
			return
		}
	}

	for {
		counter := 0
		for _, a := range A {
			for k, aa := range a {
				if k == 0 {
					continue
				}

				if allCost[k]-aa < X {
					goto INIT
				}
			}
			for k, aa := range a {
				fmt.Println(allCost[k])
				allCost[k] -= aa
				counter++
			}
		INIT:
		}
		if counter == 0 {
			break
		}
	}
	fmt.Print(allCost[0])
}
