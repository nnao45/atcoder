package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	S := nextStr()
	T := nextStr()
	if S == T[:(len(T)-1)] {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
