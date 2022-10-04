package main

import "fmt"

type Cupcake struct {
	Name  string
	Price int
}

func (p *Cupcake) GetName() string {
	return p.Name
}

func (p *Cupcake) GetPrice() int {
	return p.Price
}

// cupcake with chocolate
type AddChocolate struct {
	cupcake Cupcake
}

func (p *AddChocolate) GetName() string {
	return p.cupcake.GetName() + " with chocolate"
}

func (p *AddChocolate) GetPrice() int {
	return p.cupcake.GetPrice() + 7
}

// cupcake with sprinkles
type AddSprinkles struct {
	cupcake AddChocolate
}

func (p *AddSprinkles) GetName() string {
	return p.cupcake.GetName() + " and sprinkles"
}

func (p *AddSprinkles) GetPrice() int {
	return p.cupcake.GetPrice() + 7
}

func main() {

	commonCupcake := &Cupcake{
		Name:  "common cupcake",
		Price: 10,
	}
	cupcakeWithChocolate := &AddChocolate{
		cupcake: *commonCupcake,
	}
	cupcakeWithChocolateAndSprinkles := &AddSprinkles{
		cupcake: *cupcakeWithChocolate,
	}

	fmt.Println(commonCupcake.GetName(), commonCupcake.GetPrice())
	fmt.Println(cupcakeWithChocolate.GetName(), cupcakeWithChocolate.GetPrice())
	fmt.Println(cupcakeWithChocolateAndSprinkles.GetName(), cupcakeWithChocolateAndSprinkles.GetPrice())
}
