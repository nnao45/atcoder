package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	S := nextStr()
	intS, e := strconv.Atoi(S)
	if e != nil {
		panic(e)
	}
	i := 1
	result := 0
	for {
		num := 2019 * i
		if intS >= num {
			i++
			result += strings.Count(S, fmt.Sprint(num))
		} else {
			break
		}
	}

	fmt.Print(result)
}
