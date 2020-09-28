package main

import "fmt"

func main() {
	var e float32 = 0
	var n float32 = 0
	var lim float32
	fmt.Scanln(&lim)
	for n <= lim {
		var sum float32 = 1
		var nf float32 = 1
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
