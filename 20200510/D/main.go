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
	NK := stringsFiledsToInts()
	// N := NK[0]
	K := NK[1]

	state := 1

	inputs := stringsFiledsToInts()
	mapper := map[int]int{}

	for i := 1; i <= K; i++ {
		next := inputs[state-1]
		state = next
		mapper[next]++

		if i == 1000 {
			mmapper := mapper
			last := 0
			zeroCount := 0
			for {
				for j, m := range mmapper {
					zeroCount = 0
					if m != 0 {
						mmapper[j]--
					} else {
						zeroCount++
					}
				}
				if zeroCount == last {
					break
				}
				last = zeroCount
			}
		}
	}
	fmt.Print(state)
}
