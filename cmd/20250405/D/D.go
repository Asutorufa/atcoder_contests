package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
)

type Queue[T any] struct {
	queue *list.List
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{list.New()}
}

func (q *Queue[T]) PushBack(v T) {
	q.queue.PushBack(v)
}

func (q *Queue[T]) PushFront(v T) {
	q.queue.PushFront(v)
}

func (q *Queue[T]) PopFront() T {
	v := q.queue.Front()
	q.queue.Remove(v)
	return v.Value.(T)
}

func (q *Queue[T]) PopBack() T {
	v := q.queue.Back()
	q.queue.Remove(v)
	return v.Value.(T)
}

func (q *Queue[T]) Empty() bool {
	return q.queue.Len() == 0
}

/*
#include <bits/stdc++.h>
using namespace std;

	int main() {
		int h,w;
		int a,b,c,d;
		int x,y,z,nx,ny;
		bool wall;
		deque<int>dq;
		const int dx[4]={-1,1,0,0};
		const int dy[4]={0,0,-1,1};
		cin>>h>>w;
		vector<string>s(h);
		vector<vector<int> >ans(h,vector<int>(w,h*w));
		for(int i=0;i<h;i++)cin>>s[i];
		cin>>a>>b>>c>>d;
		a--,b--,c--,d--;

		ans[a][b]=0;
		dq.push_front(a*w+b);
		while(!dq.empty()){
			z=dq.front();
			dq.pop_front();
			if(z==(c*w+d)){
				cout<<ans[c][d]<<endl;
				return 0;
			}
			x=z/w,y=z%w;
			for(int i=0;i<4;i++){
				wall=false;
				for(int j=1;j<=2;j++){
					nx=x+dx[i]*j;
					ny=y+dy[i]*j;
					if((nx<0)||(nx>=h)||(ny<0)||(ny>=w))break;
					if(s[nx][ny]=='#')wall=true;
					if(!wall){
						if(j==1){
							if(ans[nx][ny]>ans[x][y]){
								ans[nx][ny]=ans[x][y];
								dq.push_front(nx*w+ny);
							}
						}
					}
					else{
						if(ans[nx][ny]>ans[x][y]+1){
							ans[nx][ny]=ans[x][y]+1;
							dq.push_back(nx*w+ny);
						}
					}
				}
			}
		}
		return 0;
	}
*/

func solve() {
	fmt.Scan(&H, &W)

	mas = make([][]byte, H)

	for i := range mas {
		var s string
		fmt.Scan(&s)

		mas[i] = []byte(s)
	}

	fmt.Scan(&A, &B, &C, &D)

	A--
	B--
	C--
	D--

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	ans := make([][]int, H)
	for i := range ans {
		ans[i] = make([]int, W)

		for j := range ans[i] {
			ans[i][j] = H * W
		}
	}

	q := NewQueue[int]()

	/*
			二维转一维
			一维索引公式

			如果你有一个二维数组（比如 matrix[h][w]，高为 h，宽为 w），想把它变成一个一维数组 array[h * w]，那么二维坐标 (x, y) 转为一维索引的公式是：

				index = x * w + y

			其中：
		    	x 是行号（row，纵坐标）；
		    	y 是列号（column，横坐标）；
		    	w 是每一行的列数，也就是数组的宽度。



			一维转二维

			如果你知道了一维数组中的索引 index，想转回二维坐标 (x, y)，那公式就是：

				x = index // w   # 整除，算出行
				y = index % w    # 取余，算出列
	*/
	ans[A][B] = 0
	q.PushFront(A*W + B)

	for !q.Empty() {
		z := q.PopFront()

		if z == C*W+D {
			fmt.Println(ans[C][D])
			return
		}

		/*
			一维转二维

			如果你知道了一维数组中的索引 index，想转回二维坐标 (x, y)，那公式就是：

				x = index // w   # 整除，算出行
				y = index % w    # 取余，算出列
		*/
		x := z / W
		y := z % W

		for i := 0; i < 4; i++ {
			wall := false
			for j := 1; j <= 2; j++ {
				nx := x + dx[i]*j
				ny := y + dy[i]*j
				if (nx < 0) || (nx >= H) || (ny < 0) || (ny >= W) {
					break
				}
				if mas[nx][ny] == '#' {
					wall = true
				}

				if !wall {
					if j == 1 {
						if ans[nx][ny] > ans[x][y] {
							ans[nx][ny] = ans[x][y]
							q.PushFront(nx*W + ny)
						}
					}
				} else {
					if ans[nx][ny] > ans[x][y]+1 {
						ans[nx][ny] = ans[x][y] + 1
						q.PushBack(nx*W + ny)
					}
				}
			}
		}
		/*
			z=dq.front();
			dq.pop_front();
			if(z==(c*w+d)){
				cout<<ans[c][d]<<endl;
				return 0;
			}
			x=z/w,y=z%w;
			for(int i=0;i<4;i++){
				wall=false;
				for(int j=1;j<=2;j++){
					nx=x+dx[i]*j;
					ny=y+dy[i]*j;
					if((nx<0)||(nx>=h)||(ny<0)||(ny>=w))break;
					if(s[nx][ny]=='#')wall=true;
					if(!wall){
						if(j==1){
							if(ans[nx][ny]>ans[x][y]){
								ans[nx][ny]=ans[x][y];
								dq.push_front(nx*w+ny);
							}
						}
					}
					else{
						if(ans[nx][ny]>ans[x][y]+1){
							ans[nx][ny]=ans[x][y]+1;
							dq.push_back(nx*w+ny);
						}
					}
				}
			}
		*/
	}
}

// 使用 01-bfs 题解
// dijkstra也可解决
func solve2() {
	br := bufio.NewReader(os.Stdin)

	fmt.Fscan(br, &H, &W)

	mas := make([]string, H)

	for i := range mas {
		fmt.Fscan(br, &mas[i])

	}

	fmt.Fscan(br, &A, &B, &C, &D)

	A--
	B--
	C--
	D--

	d := []Point{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	ans := make([][]int, H)
	for i := range ans {
		ans[i] = make([]int, W)

		for j := range ans[i] {
			ans[i][j] = H * W
		}
	}

	target := Point{C, D}
	q := NewQueue[Point]()

	ans[A][B] = 0
	// 起点
	q.PushFront(Point{A, B})

	for !q.Empty() {
		z := q.PopFront()

		if z == target {
			fmt.Println(ans[C][D])
			return
		}

		for _, i := range d {
			wall := false

			// 为什么需要j=2
			// 下面题目中 1 つ前の区画および 2 つ前の区画について、その区画が壁ならば道に変えることができる
			// 指的是 高橋くん　踢一次可以把前面两个区块（假设两个区块都为墙）变成道路，那么这里往前走两步总代价都是1
			// 		如果只有 j=1 会导致往前走两步总代价变为2
			//
			// 高橋君は次の
			//
			// 2 種類の行動を好きな順番で繰り返し行うことができます。
			//
			// 上下左右に隣接する、町の中の区画であって、道であるようなものに移動する。
			// 上下左右の方向を一つ決め、その方向に前蹴りを行う。
			// 高橋君が前蹴りを行うと、現在いる区画からその方向に
			//
			// 1 つ前の区画および 2 つ前の区画について、その区画が壁ならば道に変えることができる。
			// ここで、1 つ前の区画または 2 つ前の区画が町の外である場合でも前蹴りを行うことができるが、町の外が変化することはない。
			for _, j := range []int{1, 2} {
				nx := z.X + i.X*j
				ny := z.Y + i.Y*j

				if (nx < 0) || (nx >= H) || (ny < 0) || (ny >= W) {
					break
				}

				if mas[nx][ny] == '#' {
					wall = true
				}

				if !wall {
					// 无墙时 不跳跃
					if j == 1 {
						// 无墙优先级升高
						if ans[nx][ny] > ans[z.X][z.Y] {
							ans[nx][ny] = ans[z.X][z.Y]
							q.PushFront(Point{nx, ny})
						}
					}
				} else {
					// 有墙优先级降低
					//
					// 代价+1，允许以墙的位置作为之后的路径

					if ans[nx][ny] > ans[z.X][z.Y]+1 {
						ans[nx][ny] = ans[z.X][z.Y] + 1
						q.PushBack(Point{nx, ny})
					}
				}
			}
		}
	}
}

func main() {
	solve2()
}

type Point struct {
	X, Y int
}

var H, W int
var A, B, C, D int
var mas [][]byte
var History map[Point]bool = map[Point]bool{}

// -----------------------------------------------------------
// 无法用dfs解决
func dfs(A, B int) int {
	if A == C && B == D {
		return 0
	}

	if A < 0 || A >= H || B < 0 || B >= W {
		return -2
	}

	if mas[A][B] == '#' {
		return -3
	}

	if History[Point{A, B}] {
		return -1
	}

	History[Point{A, B}] = true
	defer delete(History, Point{A, B})

	A1 := dfs(A-1, B)
	A2 := dfs(A+1, B)
	B1 := dfs(A, B-1)
	B2 := dfs(A, B+1)

	z := Min(A1, A2, B1, B2)
	if z >= 0 {
		return z
	}

	if A1 == -2 && A2 == -2 && B1 == -2 && B2 == -2 {
		return -2
	}

	if A1 == -1 && A2 == -1 && B1 == -1 && B2 == -1 {
		return -1
	}

	if A1 == -3 && A-1 >= 0 {
		mas[A-1][B] = '.'
		A1 = dfs(A-1, B)
		mas[A-1][B] = '#'
	}

	if A2 == -3 && A+1 < H {
		mas[A+1][B] = '.'
		A2 = dfs(A+1, B)
		mas[A+1][B] = '#'
	}

	if B1 == -3 && B-1 >= 0 {
		mas[A][B-1] = '.'
		B1 = dfs(A, B-1)
		mas[A][B-1] = '#'
	}

	if B2 == -3 && B+1 < W {
		mas[A][B+1] = '.'
		B2 = dfs(A, B+1)
		mas[A][B+1] = '#'
	}

	if A1 == -2 && A2 == -2 && B1 == -2 && B2 == -2 {
		return -2
	}

	if A1 == -1 && A2 == -1 && B1 == -1 && B2 == -1 {
		return -1
	}

	x := Min(A1, A2, B1, B2)
	if x >= 0 {
		fmt.Println(x, A1, A2, B1, B2)
		return x + 1
	}

	// fmt.Println(A1, A2, B1, B2, x, A, B)
	return -1
}

func Min(A, B, C, D int) int {
	min := math.MaxInt

	if A >= 0 {
		min = A
	}

	if B >= 0 && min > B {
		min = B
	}

	if C >= 0 && min > C {
		min = C
	}

	if D >= 0 && min > D {
		min = D
	}

	if min == math.MaxInt {
		return -3
	}

	return min
}

/*
入力例 1

10 10
..........
#########.
#.......#.
#..####.#.
##....#.#.
#####.#.#.
.##.#.#.#.
###.#.#.#.
###.#.#.#.
#.....#...
1 1 7 1

出力例 1

1

*/
