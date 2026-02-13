package main

import "fmt"

func main() {
	// pi := 3.14
	// r := 18.0
	var appHitung int

	PI := 3.14
	var r float64

	fmt.Println("# Hitung Luas dan Keliling Lingkaran.")
	fmt.Println("Pilih aplikasi hitung:")
	fmt.Println("1. Luas lingkaran")
	fmt.Println("2. Keliling lingkaran")

	for true {
		fmt.Printf("%s", "Masukkan pilihan: ")
		fmt.Scanf("%d", &appHitung)

		if appHitung > 2 || appHitung < 1 {
			fmt.Println("pilihan tidak tersedia.")
			continue
		}

		switch appHitung {
		case 1:
			fmt.Printf("%s", "Masukkan jari-jari lingkaran: ")
			fmt.Scanf("%f", &r)
			luasLingkaran := luasLingkaran(PI, r)
			fmt.Printf("hasil luas lingkaran adalah: %f\n", luasLingkaran)
		case 2:
			fmt.Printf("%s", "Masukkan jari-jari lingkaran: ")
			fmt.Scanf("%f", &r)
			kelilingLingkaran := kelilingLingkaran(PI, r)
			fmt.Printf("hasil keliling lingkaran adalah: %f\n", kelilingLingkaran)
		}
		break
	}

}

func luasLingkaran(pi float64, r float64) float64 {
	return pi * r * r
}

func kelilingLingkaran(pi float64, r float64) float64 {
	return 2 * pi * r
}
