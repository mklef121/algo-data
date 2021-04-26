// Golang struct allows us to create user-defined types by combining with other types.

/*
syntax**
type identifier struct{
  field1 data_type
  field2 data_type
  field3 data_type
}
*/
package main

import "fmt"

type teleframe struct {
	length  int
	breadth int
	color   string

	geometry struct {
		area      int
		perimeter int
	}
}

func main() {

	type rectangle struct {
		length  float64
		breadth float64
		color   string
	}

	fmt.Println(rectangle{10.5, 25.10, "red"})

	//Initialize a struct
	var rect teleframe // dot notation
	rect.length = 10
	rect.breadth = 20
	rect.color = "Green"

	rect.geometry.area = rect.length * rect.breadth
	rect.geometry.perimeter = 2 * (rect.length + rect.breadth)

	fmt.Println(rect)
	fmt.Println("Area:\t", rect.geometry.area)
	fmt.Println("Perimeter:", rect.geometry.perimeter)

	var rect1 = rectangle{10, 20, "Green"}
	fmt.Println(rect1)

	var rect2 = rectangle{length: 34, breadth: 23, color: "Black"}
	fmt.Println(rect2)

	rect3 := new(rectangle) // rect3 is a pointer to an instance of rectangle
	rect3.length = 10
	rect3.breadth = 20
	rect3.color = "Purple"
	fmt.Println(*rect3)

	type Salary struct {
		Basic, HRA, TA float64 //All keys of this struct will inherit the float64 type from the last guy
	}
	type Employee struct {
		FirstName, LastName, Email string
		Age                        int
		MonthlySalary              []Salary
	}

	emp := Employee{
		FirstName: "Miracle",
		LastName:  "Nwabueze",
		Email:     "mklef121@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{ //We can add the struct type name here
				Basic: 150000,
				HRA:   10000,
				TA:    90,
			},
			{ //Or we can decide not to add struct description
				Basic: 200000,
				HRA:   5000,
				TA:    10,
			},
		},
	}

	fmt.Println(emp, emp.MonthlySalary[0].Basic)

}
