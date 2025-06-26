package main

import "fmt"

type Transport interface {
	Move()
	Stop()
}

type Car struct {
	Model string
}

func (c Car) Move() {
	fmt.Printf("Автомобиль %s едет по дороге\n", c.Model)
}

func (c Car) Stop() {
	fmt.Printf("Автомобиль %s остановился\n", c.Model)
}

type Plane struct {
	FlightNumber string
}

func (p Plane) Move() {
	fmt.Printf("Самолет рейса %s взлетел\n", p.FlightNumber)
}

func (p Plane) Stop() {
	fmt.Printf("Самолет рейса %s приземлился\n", p.FlightNumber)
}

func testTransport(t Transport) {
	t.Move()
	t.Stop()
	fmt.Println()
}

func main() {
	car := Car{Model: "Toyota Camry"}
	airplane := Plane{FlightNumber: "SU-123"}

	testTransport(car)
	testTransport(airplane)
}
