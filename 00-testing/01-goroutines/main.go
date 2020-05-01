package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Printf("Numero de CPU al inicio: %v\n", runtime.NumCPU())
	fmt.Printf("Numero de go routines al inicio : %v\n", runtime.NumGoroutine())
	wg.Add(2)
	go routine1()
	go routine2()

	fmt.Printf("Numero de CPU medio : %v\n", runtime.NumCPU())
	fmt.Printf("Numero de go routines medio : %v\n", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Se Finalizo")
	fmt.Printf("Numero de CPU : %v\n", runtime.NumCPU())
	fmt.Printf("Numero de go routines : %v\n", runtime.NumGoroutine())

}
func routine1() {
	fmt.Println("Go routine 1")
	wg.Done()
}

func routine2() {
	fmt.Println("Go routine 2")
	wg.Done()
}
