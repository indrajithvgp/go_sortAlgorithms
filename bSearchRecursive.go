package main

import "fmt"

func binarySearch(array []int, numberToFind int, start int, end int) int{
    if start > end {
		return -1
	}
	middle := (start + end) / 2
	if array[middle] == numberToFind{
		return middle
	} else if array[middle] < numberToFind {
		return binarySearch(array, numberToFind, middle + 1, end)
	} else {
		return binarySearch(array, numberToFind, start, middle - 1)
    }    

}

func main(){
	numbers := []int{2,5,6,7,88,99,122,133,344,3446}
	toFind := 3446
	findNumber := binarySearch(numbers, toFind, 0, len(numbers) - 1)
	fmt.Println(findNumber)
}
