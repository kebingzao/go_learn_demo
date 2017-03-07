package main

import (
	"fmt"
	"time"
)

func main(){
	c1 := make(chan string)

	c2 := make(chan string)



	go func() {

		time.Sleep(time.Second * 3)

		c1 <- "one"

	}()

	go func() {

		time.Sleep(time.Second * 8)

		c2 <- "two"

	}()

	fmt.Println("start==============")

	for i := 0; i < 2; i++ {

		select {

		case msg1 := <-c1:

			fmt.Println("received1", msg1)

		case msg2 := <-c2:

			fmt.Println("received2", msg2)

		}

	}
}