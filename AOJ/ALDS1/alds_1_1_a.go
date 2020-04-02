package main

import "fmt"

func insertionSort(arr []int, n int) []int {
	for i := 1; i <= n-1; i++ {
		v := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > v {
			arr[j+1] = arr[j]
			j--
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		arr = insertionSort(arr)
		fmt.Println(arr)
	}
}

//func printSort(arr []int) {
//	for i := 0; i < len(arr); i++ {
//		if i != 0 {
//			fmt.Print(" ")
//		}
//		fmt.Print(arr[i])
//	}
//	fmt.Println()
//}
//
//func insertSort(arr []int, n int) {
//	printSort(arr)
//	for i := 1; i <= n-1; i++ {
//		v := arr[i]
//		j := i - 1
//		for j >= 0 && arr[j] > v {
//			arr[j+1] = arr[j]
//			j--
//		}
//		arr[j+1] = v
//		printSort(arr)
//	}
//}
//
//func main() {
//	var n int
//	fmt.Scan(&n)
//	arr := make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&arr[i])
//	}
//	insertSort(arr, n)
//}
