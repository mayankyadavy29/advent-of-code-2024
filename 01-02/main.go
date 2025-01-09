package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	l1, l2 := []int{}, map[int]int{}
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, _ := reader.ReadLine()
		l := strings.Split(string(line), "   ")
		if l[0] == "end" {
			break
		}
		x, _ := strconv.Atoi(l[0])
		l1 = append(l1, x)
		y, _ := strconv.Atoi(l[1])
		if _, ok := l2[y]; !ok {
			l2[y] = 0
		}
		l2[y] += 1
	}
	ans := 0
	for i := 0; i < len(l1); i++ {
		if v, ok := l2[l1[i]]; ok {
			fmt.Println(l1[i], v)
			ans += l1[i] * v
		}
	}
	fmt.Println(ans, len(l1))
}
