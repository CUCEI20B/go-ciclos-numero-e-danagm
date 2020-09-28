package main

import "fmt"

func main() {
	var e uint64 = 0
	var n uint64 = 0
	var lim uint64
	fmt.Scanln(&lim)
	for n <= lim {
		var sum uint64 = 1
		var nf uint64 = 1
		for nf < n {
			sum *= nf
			nf++
		}
		if sum != 0 {
			e += (1 / sum)
		}
		n++
	}

	fmt.Println(e)
}
