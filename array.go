// package main

// import "fmt"

// func main(){
// 	var size, i int
// 	fmt.Println("enter a size:")
// 	fmt.Scanln(&size)

// 	arr := make([]int, size)

// 	fmt.Println("enter elements:")
// 	for i=0;i<size;i++{
// 		fmt.Scanln(&arr[i])
// 	}
// 	arrOdd := make([]int, size)
// 	fmt.Println("odd arr is:")
// 	for i=0;i<size;i++{
// 		if arr[i] % 2 != 0{
// 			arrOdd = append([]int{arr[i]}, arrOdd...)
// 		} 
// 	}

// 	fmt.Printf("odd list = %v", arrOdd)
// }