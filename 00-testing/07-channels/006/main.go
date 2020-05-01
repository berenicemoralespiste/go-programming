package main

import "fmt"

func main() {

	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}

		}()
	}

	for v := 0; v <100; v++{
		fmt.Println(<-c)
	}
}
