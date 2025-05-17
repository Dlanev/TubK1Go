package main

import "fmt"
import "time"

/*func bacaData(data tabSpend) {
	var t string = time.Now()
	
}

func tampilkanData(data tabSpend) {
	var i int
	for i = 0; i < len(data); i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()
}

func insertionsortKecilBesar(data *tabSpend, n int) {
	var i, j, key int
	for i = 1; i < n; i++ {
		key = data[i]
		j = i - 1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
	}
}

func insertionsortBesarKecil(data *tabSpend) {
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

