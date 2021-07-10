package main

import "fmt"

func main() {
	maxHeap := &MaxHeap{}
	fmt.Println(maxHeap)

	buildHeap := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	for _, val := range buildHeap {
		maxHeap.Insert(val)
		fmt.Print(maxHeap)
		fmt.Print("\n")
	}
}
