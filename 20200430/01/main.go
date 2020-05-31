package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// 0以上n以下の整数の個数
// n=12345なら、00000〜12345の12346個を数えています。
func main() {
	s := nextStr()
	digit := len(s)
	dp := [][]int{}
	for i := 0; i < digit+1; i++ {
		dp = append(dp, []int{0, 0})
	}
	dp[0][0] = 1
	fmt.Printf("%v", dp)
	for _, i := range dp {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = dp[i-1][1]*10 + dp[i-1][0]*int(s[i-1])
	}
}
