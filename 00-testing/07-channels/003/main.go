package main

import (
	"fmt"
)

func main() {
	salir := make(chan int)
	c := gen(salir)
	recibir(c, salir)

	fmt.Println("A punto de finalizar.")
}

func recibir(c, c2 <-chan int) {
	for {
		select {
		case message1 := <-c:
			fmt.Println("received message1 ", message1)
		case <-c2:
			fmt.Println("received exit")
			return
		}

	}

}

func gen(salir chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		salir <- 1
		close(c)
	}()

	return c
}
