// UVa 583 - Prime Factors

package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"strconv"
)

func factorize(n int) []string {
	p := make([]string, 0)

	if n < 0 {
		p = append(p, "-1")
		n *= -1
	}

	t := 2
	done := false
	for !done {
		for n % t == 0 {
			p = append(p, strconv.Itoa(t))
			if n /= t; n == 1 {
				done = true
				break
			}
		}

		if !done {
			t ++
			if t > int(math.Sqrt(float64(n)) + .5) {
				p = append(p, strconv.Itoa(n))
				break
			}
		}
	}
	return p
}

func main() {
	in, _ := os.Open("583.in")
	out, _ := os.Create("583.out")
	var n int
	for {
		fmt.Fscanf(in, "%d", &n)
		if n == 0 {
			break
		}
		p := factorize(n)
		fmt.Fprintf(out, "%d = ", n)
		fmt.Fprintln(out, strings.Join(p, " x "))
	}
	in.Close()
	out.Close()
}