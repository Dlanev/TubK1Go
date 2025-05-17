package main

import "fmt"
import "time"

type tSpend struct {
	jumlah int
	tipe   string
	time.Time
}

type bdgt struct {
	tots int
	trip int
}

type waktu struct {
	jam, menit         int
	hari, bulan, tahun int
}

const NMAX int = 50

type tabSpend [NMAX]tSpend

func main() {
	var data tabSpend
	var Budget bdgt
	var option, count, idx int
	count = 1
	for option != 6 {
		fmt.Print("-------------------------------------------------")
		fmt.Println()
		fmt.Printf("Total Balance: %d", Budget.tots)
		fmt.Println()
		fmt.Printf("Trip Budget: %d\n", Budget.trip)
		fmt.Print("-------------------------------------------------\n")
		fmt.Println("1.Set Budget")
		if Budget.trip != 0 {
			fmt.Println("2.End Trip")
		} else {
			fmt.Println("2.Start Trip")
		}
		fmt.Println("3.Spend")
		fmt.Println("4.History")
		fmt.Println("5.Spending")
		fmt.Println("6.Exit")
		fmt.Scan(&option)
		switch option {
		case 1:
			fmt.Print("Input Balance:")
			fmt.Scan(&Budget.tots)
		case 2:
			if count%2 != 0 {
				if Budget.tots == 0 {
					fmt.Println("No Balance")
				} else {
					fmt.Print("Set a budget:")
					fmt.Scan(&Budget.trip)
					if Budget.trip > Budget.tots {
						fmt.Println("Insufficient Balance")
						Budget.trip = 0
					} else {
						Budget.tots -= Budget.trip
						count++
					}
				}
			} else {
				Budget.tots += Budget.trip
				Budget.trip = 0
			}
		case 3:
			sPend(&data, idx, &Budget.trip)
			idx++
		case 4:
			hIst(data, idx)
		case 6:
			fmt.Println("Terima Kasih")
		case 5:

		}
	}
}

func sPend(A *tabSpend, n int, B *int){
	fmt.Print("-------------------------------------------------\n")
	fmt.Print("Input Your Spending And Spending Type: ")
	bacaData(A, n)
	fmt.Print(A[n].jumlah)
	*B = *B - A[n].jumlah
}

func hIst(A tabSpend, n int){
	var opsi int
	fmt.Print("-------------------------------------------------\n")
	fmt.Print("Spending History:\n")
	cetakData(A, n)
	fmt.Print("-------------------------------------------------\n")
	fmt.Print("Filter:\n")
	fmt.Println("1.Sort")
	fmt.Println("2.Search")
	fmt.Println("3.Exit")
	fmt.Scan(&opsi)
	switch opsi {
	case 1:
		choice1(A, n)
	}
}

func bacaData(A *tabSpend, n int){
	fmt.Scan(&A[n].jumlah, &A[n].tipe)
	A[n].Time = time.Now()
}

func choice1(A tabSpend, n int){
	var i int
	fmt.Println("1.By Amount\n2.By Time")
	fmt.Scan(&i)
	if i == 1{
		fmt.Print("1.Ascending\n2.Descending")
		fmt.Scan(&i)
		switch i {
		case 1:
			insertionsortKecilBesar(&A, n)
		case 2: 
			insertionsortBesarKecil(&A, n)
		}
	} else {
		fmt.Print("1.Ascending\n2.Descending")
		fmt.Scan(&i)
		switch i {
		case 1:
			insertionsortKecilBesar2(&A, n)
			cetakData(A, n)
		case 2: 
			insertionsortBesarKecil2(&A, n)
			cetakData(A, n)
		}
	}
}

func insertionsortKecilBesar(data *tabSpend, n int) {
	var i, j int
	for i = 1; i < n; i++ {
		key := data[i]
		j = i - 1
		for j >= 0 && data[j].jumlah > key.jumlah {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
	}
}

func insertionsortBesarKecil(data *tabSpend, n int) {
	var i, j int
	for i = 1; i < n; i++ {
		key := data[i]
		j = i - 1
		for j >= 0 && data[j].jumlah < key.jumlah {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
	}
}

func insertionsortKecilBesar2(data *tabSpend, n int) {
	var i, j int
	for i = 1; i < n; i++ {
		key := data[i]
		j = i - 1
		for j >= 0 && data[j].After(key.Time) {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
	}
}

func insertionsortBesarKecil2(data *tabSpend, n int) {
	var i, j int
	for i = 1; i < n; i++ {
		key := data[i]
		j = i - 1
		for j >= 0 && data[j].Before(key.Time) {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
	}
}

func cetakData(A tabSpend, n int){
	var i int
	for i = 0; i < n; i++{
		fmt.Println(A[i].jumlah, A[i].Time.Format("2006-08-08 07:08:06"))
	}
}