package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(br, &N, &M)

	var strs = make(map[int]*Node, M)

	for i := 0; i < M; i++ {
		var K int
		fmt.Fscan(br, &K)

		str := make(map[int]struct{}, K)

		for j := 0; j < K; j++ {
			var A int
			fmt.Fscan(br, &A)
			str[A] = struct{}{}

			node, ok := strs[A]
			if !ok {
				node = NewNode()
				strs[A] = node
			}

			node.Add(str)

		}

	}

	count := 0

	remove := []*list.Element{}

	for i := 0; i < N; i++ {
		var B int
		fmt.Fscan(br, &B)

		BB := strs[B]

		if BB == nil {
			fmt.Println(count)
			continue
		}

		BB.Range(func(z *list.Element) bool {
			str, ok := z.Value.(map[int]struct{})
			if !ok {
				return true
			}

			if len(str) == 0 {
				remove = append(remove, z)
				return true
			}

			delete(str, B)

			if len(str) == 0 {
				count++
				remove = append(remove, z)
			}

			return true
		})

		for _, v := range remove {
			BB.Remove(v)
		}

		remove = remove[:0]

		fmt.Println(count)
	}
}

type Node struct {
	list *list.List
}

func NewNode() *Node {
	return &Node{
		list: list.New(),
	}
}

func (n *Node) Add(v map[int]struct{}) {
	n.list.PushBack(v)
}

func (n *Node) Remove(v *list.Element) {
	n.list.Remove(v)
}

func (n *Node) Range(f func(z *list.Element) bool) {
	for e := n.list.Front(); e != nil; e = e.Next() {
		if !f(e) {
			break
		}
	}
}

/*
10 8
1 4
5 6 9 7 4 3
4 2 4 1 3
1 1
5 7 9 8 1 5
2 9 8
1 2
1 1
6 5 2 7 8 4 1 9 3 11
*/
