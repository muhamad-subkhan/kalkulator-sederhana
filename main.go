package main

import "fmt"

func main() {
	fmt.Println("Selamat Datang di Aplikasi Kalkulator Sederhana")
	
	var menu int
	fmt.Println("\nMasukan Pilihan:")
	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurahan")
	fmt.Println("3. Perkalian")
	fmt.Println("4. Pembagian")
	fmt.Scanln(&menu)

	switch menu{
	case 1:
		var Angka1 int
		fmt.Println("\nMasukan Angka pertama:")
		fmt.Scanln(&Angka1)

		var Angka2 int
		fmt.Println("\nMasukan Angka Kedua:")
		fmt.Scanln(&Angka2)

		hasil := Angka1 + Angka2
		fmt.Println("\nHasil: ", hasil)
	case 2:
		var Angka1 int
		fmt.Println("\nMasukan Angka pertama:")
		fmt.Scanln(&Angka1)

		var Angka2 int
		fmt.Println("\nMasukan Angka Kedua:")
		fmt.Scanln(&Angka2)

		Hasil := Angka1 - Angka2
		fmt.Println("\nHasil: ", Hasil)

	case 3:
		var Angka1 int
		fmt.Println("\nMasukan Angka pertama:")
		fmt.Scanln(&Angka1)

		var Angka2 int
		fmt.Println("\nMasukan Angka Kedua:")
		fmt.Scanln(&Angka2)

		hasil := Angka1 * Angka2
		fmt.Println("\nHasil: ", hasil)


	case 4:
		var Angka1 int
		fmt.Println("\nMasukan Angka pertama:")
		fmt.Scanln(&Angka1)

		var Angka2 int
		fmt.Println("\nMasukan Angka Kedua:")
		fmt.Scanln(&Angka2)

		hasil := Angka1 / Angka2
		fmt.Println("\nHasil: ", hasil)

	}

}