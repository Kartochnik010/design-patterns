package main

import (
	"fmt"
	"time"
)

type animal interface {
	makeNoise()
}

type duck struct{}

func (d duck) makeNoise() {
	fmt.Println(time.Now().Format(time.Stamp), "Cuak Cuak")
}

type cat struct{}

func (c cat) makeNoise() {
	fmt.Println(time.Now().Format(time.Stamp), "Meow")
}

type dog struct{}

func (d dog) makeNoise() {
	fmt.Println(time.Now().Format(time.Stamp), "Bark Bark")
}

type animals struct {
	dog
	cat
	duck
}

func main() {
	a := animals{}
	a.dog.makeNoise()
	a.cat.makeNoise()
	a.duck.makeNoise()
}
