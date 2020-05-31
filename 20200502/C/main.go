package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	sc.Scan()
	init := strings.Fields(sc.Text())
	N, _ := strconv.Atoi(init[0])
	M, _ := strconv.Atoi(init[1])
	// Q, _ := strconv.Atoi(init[2])
	A := [][]int{}
	for i := 1; i <= N; i++ {
		r := []int{}
		for i := 1; i <= M; i++ {
			r = append(r, i)
		}
		A = append(A, r)
	}
	fmt.Println(A)

	nums := [][]int{}
	result := 0

	for sc.Scan() {
		raw := strings.Fields(sc.Text())
		r0, _ := strconv.Atoi(raw[0])
		r1, _ := strconv.Atoi(raw[1])
		r2, _ := strconv.Atoi(raw[2])
		r3, _ := strconv.Atoi(raw[3])

		nums = append(nums, []int{r0, r1, r2, r3})
	}

	for i := 1; i <= M; i++ {
		maybeResult := 0
		for k := 1; k <= M; k++ {
			for _, num := range nums {
				fmt.Println(A[num[0]-1][i-1], A[num[1]-1][k-1])
				if A[num[1]-1][i-1]-A[num[0]-1][k-1] == num[2] {
					maybeResult += num[3]
				}
			}
		}
		if result <= maybeResult {
			result = maybeResult
		}
	}

	fmt.Print(result)
}
