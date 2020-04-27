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

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	input := []string{}
	n := nextInt()
	for i := 0; i < n; i++ {
		input = append(input, nextStr())
	}
	resultMap := map[string]int{}

	for _, s := range input {
		resultMap[s]++
	}

	fmt.Print(len(resultMap))
}
