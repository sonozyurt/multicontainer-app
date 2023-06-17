package main

func calcFib(n int) int {
	if n <= 1 {
		return n
	}
	return calcFib(n-1) + calcFib(n-2)

}
