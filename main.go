package main

import "fmt"

func main() {
	pi := 3.14
	r := 18.0

	luasLingkaran := luasLingkaran(pi, r)
	kelilingLingkaran := kelilingLingkaran(pi, r)

	fmt.Printf("hasil luas lingkaran adalah: ,%f\n", luasLingkaran)
	fmt.Printf("hasil keliling lingkaran adalah: ,%f\n", kelilingLingkaran)
}

func luasLingkaran(pi float64, r float64) float64 {
	return pi * r * r
}

func kelilingLingkaran(pi float64, r float64) float64 {
	return 2 * pi *r
}