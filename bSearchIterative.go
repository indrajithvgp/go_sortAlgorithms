package main

import "fmt"

func binarySearch(a []int, x int) int{
    start:=0
    end:=len(a)-1
    val:=-1
    for start<=end{
        mid:=start+end/2
        if a[mid]==x{
            val=mid
            break
        }else if x<a[mid]{
            end=mid-1
        }else if x>a[mid]{
            start=mid+1
        }
    }
    return val

}

func main(){
    arr:=[]int{1,2,3,4,5,6,7,8}
    fmt.Println(binarySearch(arr, 0))
}