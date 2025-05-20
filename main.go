package main

import "fmt"
import "time"

type tSpend struct {
	jumlah int
	tipe   string
	time.Time
	line int
}

type bdgt struct {
	tots int
	trip int
}

const NMAX int = 50
var tripName [NMAX]string
type tabSpend [NMAX]tSpend

func main() {
	var data tabSpend
	var Budget bdgt
	var option, count, idx, tripNum int
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
		fmt.Print("-------------------------------------------------\n")
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
					fmt.Println("Where to?")
					fmt.Scan(&tripName[tripNum])
					if Budget.trip > Budget.tots {
						fmt.Println("Insufficient Balance")
						Budget.trip = 0
					} else {
						Budget.tots -= Budget.trip
						count++
					}
				}
			} else {
				data[count-1].line = idx
				Budget.tots += Budget.trip
				Budget.trip = 0
				tripNum++
				count++
			}
		case 3:
			sPend(&data, &idx, &Budget.trip)
			idx++
		case 4:
			hIst(data, idx, tripNum)
		case 6:
			fmt.Println("Terima Kasih!")
		case 5:

		default:
			fmt.Print("Not a valid command")

		}
	}
}

func sPend(A *tabSpend, n *int, B *int) {
	fmt.Print("Input Your Spending And Spending Type: ")
	bacaData(A, *n)
	if A[*n].jumlah > *B {
		fmt.Println("Insufficient Budget")
		A[*n].jumlah = 0
		*n -= 1
	} else {
		*B = *B - A[*n].jumlah
	}
}

func hIst(A tabSpend, n int, x int) {
	var opsi int
	fmt.Print("-------------------------------------------------\n")
	fmt.Print("Spending History:\n")
	fmt.Println()
	cetakData(A, n)
	fmt.Print("-------------------------------------------------\n")
	fmt.Print("Filter:\n")
	fmt.Println("1.Sort")
	fmt.Println("2.Search")
	fmt.Println("3.Exit")
	fmt.Scan(&opsi)
	switch opsi {
	case 1:
		choice1(A, n, x)
	}
}

func bacaData(A *tabSpend, n int) {
	var typeIndex int
	fmt.Scan(&A[n].jumlah, &typeIndex)
	switch typeIndex {
	case 1:
		A[n].tipe = "Pembelian"
	case 2:
		A[n].tipe = "Konsumsi"
	case 3:
		A[n].tipe = "Transportasi"
	case 4:
		A[n].tipe = "Lainnya"
	}
	A[n].Time = time.Now()
}

func choice1(A tabSpend, n int, x int) {
	var i int
	fmt.Print("-------------------------------------------------\n")
	fmt.Println("1.By Amount\n2.By Time")
	fmt.Print("-------------------------------------------------\n")
	fmt.Scan(&i)
	if i == 1 {
		fmt.Print("-------------------------------------------------\n")
		fmt.Print("1.Ascending\n2.Descending\n")
		fmt.Print("-------------------------------------------------\n")
		fmt.Scan(&i)
		switch i {
		case 1:
			insertionsortKecilBesar(&A, n)
			fmt.Print("Spending History:\n")
			fmt.Println()
			cetakData(A, n)
		case 2:
			insertionsortBesarKecil(&A, n)
			fmt.Print("Spending History:\n")
			fmt.Println()
			cetakData(A, n)
		default:
			fmt.Print("Not an option")
		}
	} else {
		fmt.Print("-------------------------------------------------\n")
		fmt.Print("1.Ascending\n2.Descending\n")
		fmt.Print("-------------------------------------------------\n")
		fmt.Scan(&i)
		switch i {
		case 1:
			insertionsortKecilBesar2(&A, n)
			fmt.Print("Spending History:\n")
			fmt.Println()
			cetakData(A, n)
		case 2:
			insertionsortBesarKecil2(&A, n)
			fmt.Print("Spending History:\n")
			fmt.Println()
			cetakData(A, n)
		default:
			fmt.Print("Not an option")
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

func cetakData(A tabSpend, n int) {
	var i, j, total int
		for i = 0; i < n; i++ {
			if i == A[j].line {
				fmt.Println()
				fmt.Printf("TRIP %s\n", tripName[j])
				j++
			}
			fmt.Println(A[i].jumlah, A[i].Time.Format("2006-01-02 15:04:05"), A[i].tipe)
		}
	
	for i = 0; i < n; i++ {
		total += A[i].jumlah
	}
	fmt.Println()
	fmt.Println("Total Pengeluaran:", total)
	
}
