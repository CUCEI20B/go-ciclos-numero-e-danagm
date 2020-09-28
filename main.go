package main

import "fmt"

func main() {
	var e float64 = 0
	var n float64 = 0
	var lim float64
	fmt.Scanln(&lim)
	for n <= lim {
		var sum float64 = 1
		var nf float64 = 1
		for nf <= n {
			sum *= nf
			nf++
		}
		e += (1 / sum)
		n++
	}

	fmt.Println(e)
}
