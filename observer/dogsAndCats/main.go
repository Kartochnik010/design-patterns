package main

import (
	"log"
	"os"
)

type pet interface {
	pet()
	makeNoise()
}

type animal struct {
	l    *log.Logger
	name string
}

func (p *animal) pet() {
	p.l.Println("Recieved pets")
}

func (p *animal) makeNoise() {
	p.l.Println(p.name + "'s pet noises")
}

func main() {
	lg := log.New(os.Stderr, "OBSERVER: ", log.Lshortfile)
	cnt := &animal{name: "Joseph", l: lg}

	cnt.pet()
	cnt.pet()
	cnt.makeNoise()

}
