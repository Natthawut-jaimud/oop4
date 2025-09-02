package main

import "fmt"

type (
	Car interface {
		Drive()
	}

	Sedan struct {
		Brand string
		Model string
		Color string
		Price int
		Size  string
	}

	Pickup struct {
		Brand string
		Model string
		Color string
	}
)

func (s Sedan) Drive() {
	fmt.Println("Driving the sedan car:", s.Brand, s.Model, s.Color, "Price:", s.Price, "Size:", s.Size)
}

func (p Pickup) Drive() {
	fmt.Println("Driving the pickup car:", p.Brand, p.Model, p.Color)
}

func main() {
	sedan := Sedan{
		Brand: "Natthawut",
		Model: "Natthawut",
		Color: "Silver",
		Price: 2500000,
		Size:  "Medium",
	}

	pickup := Pickup{
		Brand: "Natthawut",
		Model: "WorkPro",
		Color: "Black",
	}

	sedan.Drive()
	pickup.Drive()
}
