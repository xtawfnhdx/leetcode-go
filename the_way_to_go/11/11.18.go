package main

import "fmt"

type Any interface{}
type Car struct {
	Model       string
	Manufacture string
	BuildYear   int
}
type Cares []*Car

func main() {
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}
	allCars := Cares([]*Car{ford, bmw, merc, bmw2})
	allNeeBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacture == "BMW") && (car.BuildYear > 2010)
	})
	fmt.Println("allcars:", allCars)
	fmt.Println("newcars:", allNeeBMWs)
}

func (cs Cares) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}
func (cs Cares) FindAll(f func(car *Car) bool) Cares {
	cars := make([]*Car, 0)
	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}
func (cs Cares) Map(f func(car *Car) Any) []Any {
	result := make([]Any, len(cs))
	ix := 0
	cs.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})
	return result
}

func MakeSortAppender(manufacutres []string) (func(car *Car), map[string]Cares) {
	sortedCars := make(map[string]Cares)
	for _, m := range manufacutres {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)
	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacture]; ok {
			sortedCars[c.Manufacture] = append(sortedCars[c.Manufacture], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}
	return appender, sortedCars
}
