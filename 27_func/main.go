package main

import "fmt"

type myStruct struct {
	Name string
	ID   int
}

func (a myStruct) valueReceiver() {
	a.Name = a.Name + a.Name
	a.ID = a.ID * 2

}

func (a *myStruct) pointerReceiver() {
	a.Name = a.Name + a.Name
	a.ID = a.ID * 2
}
func main() {
	anInt := 12
	aStr := "Hello"
	aStruct := myStruct{Name: "Nui", ID: 1234}

	fmt.Printf("anInt = %+v, aStr = \"%+v\", aStruct = %+v\n",
		anInt, aStr, aStruct)
	changeVar(&anInt, &aStr, &aStruct)
	fmt.Printf("anInt = %+v, aStr = \"%+v\", aStruct = %+v\n",
		anInt, aStr, aStruct)

	if d, err := divide(5.0, 3.0); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("%.4f\n", d)
	}

	// Anonymous function
	func() {
		fmt.Println("Here is in an anonymous function")
	}()

	// Value Receiver and Pointer Receiver
	bStruct := myStruct{Name: "Test", ID: 100}
	fmt.Printf("bStruct = %+v\n", bStruct)

	bStruct.pointerReceiver()
	fmt.Printf("bStruct = %+v\n", bStruct)

	bStruct.valueReceiver()
	fmt.Printf("bStruct = %+v\n", bStruct)

}

func changeVar(anInt *int, aStr *string, aStruct *myStruct) {
	*anInt++
	*aStr = "Hi"
	aStruct.ID = 5678
	aStruct.Name = "Sam"
	fmt.Printf("anInt = %+v, aStr = \"%+v\", aStruct = %+v\n",
		anInt, aStr, aStruct)

}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Can not divide by zero")
	}
	return a / b, nil
}
