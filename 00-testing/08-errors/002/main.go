package main

import (
	"fmt"
)

type errorPer struct {
	info string
}

func (e errorPer) Error() string{
	return fmt.Sprintf("Hubo un nuevo error")
}


func main() {
	error := errorPer{
		"Error de Berenice",
	}

	foo(error)
}

func foo(error error)  {
	fmt.Println("recibiendo error: ", error.(errorPer).info)
}
