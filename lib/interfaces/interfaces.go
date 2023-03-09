// Golang program illustrates how
// to implement an interface
package interfaces

import "fmt"

// Creating an interface
type tank interface {

	// Methods
	Tarea() float64
	Volume() float64
}

type myvalue struct {
	radius float64
	height float64
}

// Implementing methods of
// the tank interface
func (m *myvalue) Tarea() float64 {

	return 2*m.radius*m.height +
		2*3.14*m.radius*m.radius
}

func (m *myvalue) Volume() float64 {

	return 3.14 * m.radius * m.radius * m.height
}

// Main Method
func Run() {

	// Accessing elements of
	// the tank interface

	var t tank = &myvalue{10, 10}

	var totalArea = t.Tarea()
	var volume = t.Volume()

	fmt.Println("Area of tank :", totalArea)
	fmt.Println("Volume of tank:", volume)
}
