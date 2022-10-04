package main

import (
	"fmt"
	"log"
	"os"
)

type Stringer interface {
	Print()
	GetString() string
}

type text struct {
	s string
}

func (t text) Print() {
	fmt.Println(t.s)
}

func (t text) GetString() string {
	return t.s
}

func loggingDecorator(e Stringer, l *log.Logger) Stringer {
	l.Println("---", e.GetString())
	return e
}

func main() {

	myLogger := log.New(os.Stdout, "New  : ", 3)
	t := text{s: "inside text"}
	t.Print()

	tD := loggingDecorator(t, myLogger)
	tD.Print()

}
