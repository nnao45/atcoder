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

func stringsToInts(strings []string) []int {
	r := []int{}
	for _, s := range strings {
		i, e := strconv.Atoi(s)
		if e != nil {
			panic(e)
		}
		r = append(r, i)
	}
	return r
}

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
	one := stringsFiledsToInts()
	M := one[1]
	two := stringsFiledsToInts()
	result := map[int]int{}
	for _, t := range two {
		result[t] = 0
	}
	for i := 1; i <= M; i++ {
		sc.Scan()
		r := strings.Fields(sc.Text())
		A, _ := strconv.Atoi(r[0])
		B, _ := strconv.Atoi(r[1])
		higher := 0
		if two[A-1] > two[B-1] {
			higher = A - 1
		} else {
			higher = B - 1
		}
		result[higher]++
	}
	ret := 0
	for _, v := range result {
		if v != 0 {
			ret++
		}
	}
	fmt.Print(ret)
}
