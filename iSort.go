package main

import "fmt" 

func sortInsertion(b []int){
	a:=make([]int, len(b))
	copy(a,b)
	for i:=1;i<len(a);i++{
		j :=i-1
		temp:=a[i]
		for j>=0 && temp>a[j]{
			a[j+1]=a[j]
			j--			
		}
		a[j+1]=temp
	}
	fmt.Println(a)
}

func main(){
	arr:=[]int{101,8,5,4,0,1000,3,99,72,2,2}
	fmt.Println("***")
	sortInsertion(arr)
}