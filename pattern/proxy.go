package pattern

import "fmt"

type Driven interface {
	Drive()
}

type Car struct{}

type Driver struct {
	Age int
}

func (c *Car) Drive() {
	fmt.Println("Car being driven")
}

type CarProxy struct {
	car    Car
	driver *Driver
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 18 {
		c.car.Drive()
	} else {
		fmt.Printf("Driver is not old enough to drive")
	}
}

func TestCarDrive() {
	driver := Driver{
		16,
	}

	c := CarProxy{car: Car{}, driver: &driver}

	c.Drive()
}
