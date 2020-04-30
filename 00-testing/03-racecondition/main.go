package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	incremento := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs ; i ++ {
		go func() {
			v := incremento
			git runtime.Gosched()
			v++
			fmt.Println(incremento)
			incremento = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("El valor de la variable incremento es : ", incremento)
}
