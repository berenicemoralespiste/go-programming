package main

import (
	"errors"
	"fmt"
	"log"
)

type raizError struct {
	lat  string
	long string
	err  error
}

func (re raizError) Error() string {
	return fmt.Sprintf("error matemático: %v %v %v", re.lat, re.long, re.err)
}

func main() {
	_, err := raiz(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func raiz(f float64) (float64, error) {
	if f < 0 {
		errorRaiz := raizError{
			lat: "no se",
			long: "no se",
			err: errors.New(fmt.Sprint("Error no se puede sacar la raiz cuadrada de un numero negativo")),
		}
		return 0, errorRaiz.err
	}
	return 42, nil
}
