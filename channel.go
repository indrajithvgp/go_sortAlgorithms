package main

import "fmt"

func main(){

	owner := func()<-chan int{
		ch:=make(chan int)
		go func(chan int){
			defer close(ch)
			for i:=1;i<=3;i++{
				ch <- i
			}
		}(ch)
		return ch
	}

	consumer := func(ch <-chan int){
		for v:=range ch{
			fmt.Println("Got it*****", v)
		}
	}

	ch:=owner()
	consumer(ch)
}