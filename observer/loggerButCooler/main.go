package main

import (
	"fmt"
	"log"
	"os"
)

type countee interface {
	Count() int
	Move()
}

type counter struct {
	l   *log.Logger
	cnt int
}

func (p *counter) Move() {
	p.l.Println("moved for", 1)
	p.cnt++
}

func (p *counter) Count() int {
	return p.cnt
}

func main() {
	lg := log.New(os.Stderr, "OBSERVER: ", log.Lshortfile)
	cnt := &counter{cnt: 2, l: lg}
	fmt.Println(cnt.Count())
	cnt.Move()
	cnt.Move()
	cnt.Move()
	fmt.Println(cnt.Count())
}
