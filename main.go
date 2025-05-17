package main

import "fmt"

type tSpend struct {
	jumlah int
	tipe   string
	time   waktu
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
	var Budget bdgt
	var option, count int
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
			
		case 4:

		case 6:
			fmt.Println("Terima Kasih")
		case 5:

		}
	}
}
