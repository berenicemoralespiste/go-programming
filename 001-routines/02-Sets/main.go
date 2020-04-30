package main

import "fmt"

type persona struct {
	nombre string
}

func (p *persona) hablar() {
	fmt.Println("Está hablando: ", p.nombre)
}
type humano interface {
	hablar()
}
func diAlgo(h humano) {
	h.hablar()
}

func main() {
	p1 := persona{
		nombre:   "Berenice",
	}
	diAlgo(&p1)

}
