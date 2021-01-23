package main

import "fmt"

func bubbleSort(n []int){
	arr:=n
	temp:=0
	for i:=0;i<len(n);i++{
		for j:=0;j<len(n)-i-1;j++{
			if arr[j]>arr[j+1]{
				temp=arr[j]
				arr[j]=arr[j+1]
				arr[j+1]=temp
			}

			fmt.Println(arr)
		}
		fmt.Println("*********", i)
	}

}

func main(){
	arr:=[]int{34,5,4,56}
	bubbleSort(arr)
}