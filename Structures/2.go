package main

import "fmt"

type Engine struct {
	Power    int
	Volume   float64
	FuelType string
}

type Car struct {
	Brand  string
	Model  string
	Year   int
	Engine Engine
}

func NewCar(brand, model string, year int, power int, volume float64, fuelType string) Car {
	return Car{
		Brand:  brand,
		Model:  model,
		Year:   year,
		Engine: Engine{Power: power, Volume: volume, FuelType: fuelType},
	}
}

func (c Car) PrintInfo() {
	fmt.Printf("Автомобиль: %s %s\n", c.Brand, c.Model)
	fmt.Printf("Год выпуска: %d\n", c.Year)
	fmt.Printf("Двигатель:\n")
	fmt.Printf("  Мощность: %d л.с.\n", c.Engine.Power)
	fmt.Printf("  Объем: %.1f л\n", c.Engine.Volume)
	fmt.Printf("  Топливо: %s\n", c.Engine.FuelType)
}

func main() {
	car := NewCar("Toyota", "Camry", 2022, 249, 3.5, "бензин")

	fmt.Println("Информация об автомобиле:")
	car.PrintInfo()

	car.Engine.Power = 300
	fmt.Println("\nПосле тюнинга:")
	car.PrintInfo()
}
