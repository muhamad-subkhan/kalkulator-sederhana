package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)


	fmt.Println("\n\nSelamat Datang di Aplikasi Kalkulator Sederhana")
	for {
		fmt.Println("\nMasukan Pilihan:")
		fmt.Println("1. Penjumlahan")
		fmt.Println("2. Pengurahan")
		fmt.Println("3. Perkalian")
		fmt.Println("4. Pembagian")

		scanner.Scan()
		scanner.Text()
		menu, _ := strconv.Atoi(scanner.Text())

		switch menu {
		case 1:
			fmt.Print("\nMasukan Angka pertama: ")
			scanner.Scan()
			scanner.Text()
			scanner.Text()
			Angka1, _ := strconv.Atoi(scanner.Text())

			fmt.Print("\nMasukan Angka Kedua: ")
			scanner.Scan()
			scanner.Text()
			scanner.Text()
			Angka2, _ := strconv.Atoi(scanner.Text())

			hasil := Angka1 + Angka2
			fmt.Println("\nHasil: ", hasil)
		
		case 2:
			fmt.Print("\nMasukan Angka pertama: ")
			scanner.Scan()
			scanner.Text()
			Angka1, _ := strconv.Atoi(scanner.Text())

			fmt.Print("\nMasukan Angka Kedua: ")
			scanner.Scan()
			scanner.Text()
			Angka2, _ := strconv.Atoi(scanner.Text())

			hasil := Angka1 - Angka2
			fmt.Println("\nHasil: ", hasil)
		case 3:
			fmt.Print("\nMasukan Angka pertama: ")
			scanner.Scan()
			scanner.Text()
			Angka1, _ := strconv.Atoi(scanner.Text())

			fmt.Print("\nMasukan Angka Kedua: ")
			scanner.Scan()
			scanner.Text()
			Angka2, _ := strconv.Atoi(scanner.Text())

			hasil := Angka1 * Angka2
			fmt.Println("\nHasil: ", hasil)
		case 4:
			fmt.Print("\nMasukan Angka pertama: ")
			scanner.Scan()
			scanner.Text()
			Angka1, _ := strconv.Atoi(scanner.Text())

			fmt.Print("\nMasukan Angka Kedua: ")
			scanner.Scan()
			scanner.Text()
			Angka2, _ := strconv.Atoi(scanner.Text())

			hasil := Angka1 / Angka2
			fmt.Println("\nHasil: ", hasil)
		}
	}
}
