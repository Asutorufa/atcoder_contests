package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
				><<>><<<<>

				压缩成类似上面的字符串，然后直接找>东西
				前面的个数乘以后面的个数就是这个子串的总是

				如上面的
				> 1
				< 2
				> 2
				< 4
				> 1

		Ai-1<Ai>Ai+1 是 <>
		Ai-1>Ai<Ai+1 是 ><

		由上可知，找到 > , 就可以和前面的组合成<>, 和后面的组合成 ><
		而和前面的组合的是前面字串的最后一个字符，和后面的组合的是后面字串的第一个字符
		**那么剩余的字串就是可以自由组合的字符**
		又由于必须 A1<A2，加上前面的是<,所以符合题目的要求
	*/
	br := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(br, &N)

	p := []int{}

	for i := 0; i < N; i++ {
		var A int
		fmt.Fscan(br, &A)

		p = append(p, A)
	}

	queue := []*S{}
	for i := 0; i < N-1; i++ {
		if p[i] < p[i+1] {
			if len(queue) == 0 || queue[len(queue)-1].char == '>' {
				queue = append(queue, &S{'<', 1})
			} else {
				queue[len(queue)-1].count++
			}
		} else {
			if len(queue) == 0 || queue[len(queue)-1].char == '<' {
				queue = append(queue, &S{'>', 1})
			} else {
				queue[len(queue)-1].count++
			}
		}
	}

	ans := 0
	for i := 1; i < len(queue)-1; i++ {
		if queue[i].char == '>' {
			ans += queue[i-1].count * queue[i+1].count
		}
	}

	fmt.Println(ans)
}

type S struct {
	char  byte
	count int
}
