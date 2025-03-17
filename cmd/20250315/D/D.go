package main

import (
	"fmt"
	"math"
)

// a = 3
// b = 3d
// c = d^2 - m
func sol(a, b, c int) int {
	l, r := 0, 600000001

	for r-l > 1 {
		mid := (l + r) / 2

		if a*mid*mid+b*mid+c <= 0 {
			l = mid
		} else {
			r = mid
		}
	}

	if a*l*l+b*l+c == 0 {
		return l
	} else {
		return -1
	}
}

/*
#include <bits/stdc++.h>
using namespace std;
using ll = long long;
ll sol(ll a, ll b, ll c) {
    // ax^2 + bx + c = 0の解
    ll l = 0, r = 600000001;
    while (r - l > 1) {
        ll mid = (l + r) / 2;
        if (a * mid * mid + b * mid + c <= 0)
            l = mid;
        else
            r = mid;
    }
    if (a * l * l + b * l + c == 0)
        return l;
    return -1;
}
int main() {
    ll n;
    cin >> n;
    for (ll d = 1; d * d * d <= n; ++d) {
        // (k+d)^3 - k^3 = d^3 + 3*d^2k + 3*d*k^2 = n
        if (n % d != 0)
            continue;
        ll m = n / d; // =3*k^2 + 3*dk + d^2
        ll k = sol(3, 3 * d, d * d - m);
        if (k > 0) {
            cout << k + d << " " << k << endl;
            return 0;
        }
    }
    cout << -1 << endl;
    return 0;
}
*/

/*
x - y = d
x^3 - y^3 = N
x^3 - y^3 = (x-y)(x^2+xy+y^2) = d(x^2+xy+y^2)

x^2 + xy + y^2 > x^2 - xy + y^2 = d^2

	==> x^2 + xy + y^2 > d^2
			|
			|
			v

d^3 <= d(x^2+xy+y^2) = x^3 - y^3 = N

x - y = d よって, x = d + y

d^3 <= (d+y)^3 - y^3 = N

	==> d^3 <= d^3 + 3d^2y + 3dy^2 = N

且つ　d^3 <= N より、 d <= sqrt3(N)

(k+d)^3 - k^3 = N 　（kは？？？）

(k+d)^3 = k^3 + 3k^2d + 3kd^2 + d^3
*/
func main() {
	var N int
	fmt.Scan(&N)

	for d := 1; math.Pow(float64(d), 3) <= float64(N); d++ {
		if N%d != 0 {
			// d^3 + 3d^2y + 3dy^2 = N よって d^2 + 3dy + 3y^2 = N/d
			//  ==> N % d == 0
			continue
		}

		// d^2 + 3dy + 3y^2 = N/d
		m := N / d

		y := sol(3, 3*d, d*d-m)

		if y > 0 {
			fmt.Println(y+d, y)
			return
		}
		// d^3 + 3d^2y + 3dy^2 = N よって d^3 + 3d^2y + 3dy^2 - N = 0

	}

	fmt.Println(-1)
}
