package main

import (
	"errors"
	"fmt"
)

type Car struct {
	Model string
}

type Cars []Car

var myCars Cars

func (cars *Cars) Add(car Car) {
	myCars = append(myCars, car)
}

func (cars *Cars) Find(model string) (*Car, error) {
	for _, car := range *cars {
		if car.Model == model {
			return &car, nil
		}
	}
	return nil, errors.New("car not found")
}

func main() {
	//imperative
	var found bool
	carToLookFor := "Blazer"
	cars := []string{"Accord", "IS", "Blazer"}
	for _, car := range cars {
		if car == carToLookFor {
			found = true
		}
	}

	fmt.Printf("Found? %v", found)

	//functional way of accomplishing the same task:
	//fmt.Printf("Found %v", cars.contains("Blazer"))

	//object oriented programming
	accord := &Car{"Accord"}
	is250 := &Car{"IS"}
	blazer := &Car{"Blazer"}

	carsOP := []*Car{accord, is250, blazer}
	carToLookForOP := is250
	found = false
	for _, car := range carsOP {
		if car == carToLookForOP {
			found = true
		}
	}

	fmt.Printf("Found? %v", found)

	myCars.Add(Car{"IS"})
	myCars.Add(Car{"HighLander"})
	myCars.Add(Car{"Blazer"})
	car, err := myCars.Find("HighLander")
	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println("Found", car)
	}
}

//object oriented programming
