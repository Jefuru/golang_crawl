// package main

// import "fmt"

// func main(){
// 	var unit, fixChange, amount, total float64
// 	fmt.Println("Plz enter unit:")
// 	fmt.Scanln(&unit)

// 	if unit < 50 {
// 		amount = unit*2.6
// 		fixChange = 25
// 	} else if unit <= 100 {
// 		amount = 130 + (unit - 50)*3.25
// 		fixChange = 35
// 	} else if unit <= 200 {
// 		amount = 130 + 162.50 + (unit - 100)*5.26
// 		fixChange = 45
// 	} else {
// 		amount = 130 + 162.50 + 526 + (unit - 200)*7.75
// 		fixChange = 55
// 	}
// 	total = amount + fixChange
// 	fmt.Println("You should pay ", total, " dollars for bill.")
// }