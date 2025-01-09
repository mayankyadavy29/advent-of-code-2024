package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	l1, l2 := []int{}, []int{}
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
		l2 = append(l2, y)
	}
	sort.Ints(l1)
	sort.Ints(l2)
	ans := 0
	for i := 0; i < len(l1); i++ {
		dis := l1[i] - l2[i]
		if dis < 0 {
			ans += -1 * dis
		} else {
			ans += dis
		}
	}
	fmt.Println(ans)
}
