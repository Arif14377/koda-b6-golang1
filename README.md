# Belajar Golang Dasar (Day 1)

Membuat program hitung luas dan keliling lingkaran

## Update:
Membuat menjadi interaktif CLI.

## Pembelajaran:

- mendeklarasikan variable dengan := (wajib inisiasi value jika pakai :=)

  ```go
  pi := 3.14
  r := 18.0
  ```

- membuat function dengan awalan func dan mendeklarasikan tipe data parameter dan return function

  ```go
  func luasLingkaran(pi float64, r float64) float64 {
      return pi * r * r
  }

  func kelilingLingkaran(pi float64, r float64) float64 {
      return 2 * pi *r
  }
  ```

- memakai fmt.Printf() dari package fmt untuk print literal string di console (tidak bisa print dengan tipe data berbeda jika println)

  ```go
  fmt.Printf("hasil luas lingkaran adalah: ,%f\n", luasLingkaran)
  fmt.Printf("hasil keliling lingkaran adalah: ,%f\n", kelilingLingkaran)
  ```
