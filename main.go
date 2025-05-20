package main

import "fmt"
import "time"

type tSpend struct {
	jumlah int
	tipe   string
	time.Time
	name string
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
	var option, idx, i int
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
		fmt.Println("5.Spend Chart")
		fmt.Println("6.Exit")
		fmt.Print("-------------------------------------------------\n")
		fmt.Scan(&option)
		switch option {
		case 1:
			fmt.Print("Input Balance:")
			fmt.Scan(&Budget.tots)
		case 2:
			if Budget.trip == 0 {
				if Budget.tots == 0 {
					fmt.Println("No Balance")
				} else {
					fmt.Print("Set a budget:")
					fmt.Scan(&Budget.trip)
					fmt.Println("Where to?")
					fmt.Scan(&tripName[i])
					if Budget.trip > Budget.tots {
						fmt.Println("Insufficient Balance")
						Budget.trip = 0
					} else {
						Budget.tots -= Budget.trip
					}
				}
			} else {
				Budget.tots += Budget.trip
				Budget.trip = 0
				i++
			}
		case 3:
			if Budget.trip == 0 {
				fmt.Println("No active trip budget. Start a trip first.")
			} else {
				sPend(&data, &idx, &Budget.trip, i)
			}

		case 4:
			hIst(data, idx)
		case 6:
			fmt.Println("Terima Kasih!")
		case 5:
			fmt.Print("-------------------------------------------------\n")
			cHart(data, idx, "Pembelian")
			fmt.Println()
			cHart(data, idx, "Konsumsi")
			fmt.Println()
			cHart(data, idx, "Transportasi")
			fmt.Println()
			cHart(data, idx, "Lainnya")
			fmt.Println()
		default:
			fmt.Print("Not a valid command")

		}
	}
}

func sPend(A *tabSpend, n *int, B *int, i int) {
	fmt.Print("Input Your Spending And Spending Type: ")
	bacaData(A, *n, i)
	if A[*n].jumlah > *B {
		fmt.Println("Insufficient Budget")
		A[*n].jumlah = 0
	} else {
		*B = *B - A[*n].jumlah
		*n += 1
	}
}

func hIst(A tabSpend, n int) {
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
		choice1(A, n)
	case 2:
		var kategori string
		var idx, i, katIdx int
		fmt.Println("Categories:\n 1.Pembelian\n2.Konsumsi\n3.Transportasi\n4.Lainnya\n")
		fmt.Print("-------------------------------------------------\n")
		fmt.Print("Input the category you want to search: ")
		fmt.Scan(&katIdx)
		switch katIdx {
		case 1:
			kategori = "Pembelian"
		case 2:
			kategori = "Konsumsi"
		case 3:
			kategori = "Transportasi"
		default:
			kategori = "Lainnya"
		}
		insertionsortKategori(&A, n)
		idx = binarySearchType(A, n, kategori)
		if idx == -1 {
			fmt.Println("The category you are looking for was not found.")
		} else {
			fmt.Printf("Spending in the %s category: \n", kategori)
			for i = idx; i < n && A[i].tipe == kategori; i++ {
				fmt.Println(A[i].jumlah, A[i].Time.Format("2006-01-02 15:04:05"), A[i].tipe, A[i].name)
			}
		}
	}
}

func bacaData(A *tabSpend, n int, i int) {
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
	A[n].name = tripName[i]
}

func choice1(A tabSpend, n int) {
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

func insertionsortKategori(data *tabSpend, n int) {
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].tipe > key.tipe {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func binarySearchType(data tabSpend, n int, target string) int {
	var low, high, mid int
	low = 0
	high = n - 1
	for low <= high {
		mid = (low + high) / 2
		if data[mid].tipe == target {
			for mid > 0 && data[mid-1].tipe == target {
				mid--
			}
			return mid
		} else if data[mid].tipe < target {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return -1
}

func cetakData(A tabSpend, n int) {
	var i, total int
	fmt.Print("-------------------------------------------------\n")
	for i = 0; i < n; i++ {
		fmt.Println(A[i].jumlah, A[i].Time.Format("2006-01-02 15:04:05"), A[i].tipe, A[i].name)
		total += A[i].jumlah
	}
	fmt.Println()
	fmt.Println("Pembelian:", toTal(A, n, "Pembelian"))
	fmt.Println("Konsumsi:", toTal(A, n, "Konsumsi"))
	fmt.Println("Transportasi:", toTal(A, n, "Transportasi"))
	fmt.Println("Lainnya:", toTal(A, n, "Lainnya"))
	fmt.Println("Total Pengeluaran:", total)

}

func toTal(A tabSpend, n int, tipe string) int {
	var total, i int
	for i = 0; i < n; i++ {
		if A[i].tipe == tipe {
			total += A[i].jumlah
		}
	}
	return total
}

func cHart(A tabSpend, n int, tipe string) {
	var i int
	var length int
	length = toTal(A, n, tipe)/100000
	for i = 0; i < length; i++{
		if i == 0 {
			fmt.Printf("%12s |", tipe)
		}
		if A[i].tipe == tipe{
			fmt.Print("-")
		}
	}
}