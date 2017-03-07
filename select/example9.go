package main

import (
	"fmt"
	"time"
)
// 可以尝试去掉default case，看看打印结果
func main(){
	unbufChan := make(chan int)
	sign := make(chan byte, 2)

	go func(){
		for i := 0; i < 10; i++ {
			select {
			case unbufChan <- i:
			case unbufChan <- i + 10:

			}
			fmt.Printf("The %d select is selected\n",i)
			time.Sleep(time.Second)
		}
		close(unbufChan)
		fmt.Println("The channel is closed.")
		sign <- 0
	}()

	go func(){
		loop:
		for {
			select {
			case e, ok := <-unbufChan:
				if !ok {
					fmt.Println("Closed channel.")
					break loop
				}
				fmt.Printf("e: %d\n",e)
				time.Sleep(5 * time.Second)
			}
		}
		sign <- 1
	}()
	<- sign
	<- sign
}