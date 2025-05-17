package main

import "fmt"

const NMAX int = 99
type tabInt[NMAX]int

func bacaData(data tabInt) {
	
}

func tampilkanData(data tabInt) {
	var i int
	for i = 0; i < len(data); i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()
}

func insertionsortKecilBesar(data *tabInt) {
	var i, j, key int
	for i = 1; i < len(data); i++ {
		key = data[i]
		j = i - 1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
	}
}

func insertionsortBesarKecil(data *tabInt) {
	var i, j, key int
	for i = 1; i < len(data); i++ {
		key = data[i]
		j = i - 1
		for j >= 0 && data[j] < key {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
	}
}

func main() {
	var data tabInt
	var nData, i, t int

	fmt.Print("Masukkan banyak data: ")
	fmt.Scan(&nData)
	fmt.Print("Masukkan data: ")
	for i = 0; i < nData; i++ {
		fmt.Scan(&data[i])
	}
	fmt.Println()

	fmt.Println("TAMPILKAN DATA")
	fmt.Println("1. Urutkan data dari terbesar ke terkecil")
	fmt.Println("2. Urutkan data dari terkecil ke terbesar")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&t)

	if t == 1 {
		insertionsortBesarKecil(&data, nData)
		tampilkanData(data, nData)
	} else if t == 2 {
		insertionsortKecilBesar(&data, nData)
		tampilkanData(data, nData)
	}
}