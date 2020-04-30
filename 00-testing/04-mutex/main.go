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
	var mt sync.Mutex

	for i := 0; i < gs ; i ++ {
		go func() {
			mt.Lock()
			v := incremento
			v++
			fmt.Println(incremento)
			incremento = v
			mt.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("El valor de la variable incremento es : ", incremento)
}
