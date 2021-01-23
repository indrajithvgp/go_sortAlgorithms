package main

import "fmt"

func selectionSort(a []int)[]int{
    for i:=0;i<len(a)-1;i++{
        for j:=i+1;j<len(a)-1;j++{
            if a[i]>a[j]{
                temp:=a[j]
                a[j]=a[i]
                a[i]=temp
            }
            fmt.Println(a)
        }
        fmt.Println("**")
    }
    return a 
}

func main() {
    arr := []int{23,33, 101, 4,5,1,0,44}
    fmt.Println(selectionSort(arr))
}