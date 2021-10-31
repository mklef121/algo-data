package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

type Record struct {
	Field1 string
	Field2 string
	Field3 Secret
}

func main() {
	data := Record{"data", "datum", Secret{Username: "Miracle", Password: "mklef"}}

	//Gats the value of the variable
	ab := reflect.ValueOf(data)
	fmt.Println(ab)
	fmt.Println("String value:", ab.String())

	iType := ab.Type()
	fmt.Printf("i Type: %s\n", iType)
	fmt.Printf("The %d fields of %s are\n", ab.NumField(), iType)

	for i := 0; i < ab.NumField(); i++ {
		//This three will output something like
		// Field2  with type: string       and value _datum_
		fmt.Printf("\t%s ", iType.Field(i).Name)
		fmt.Printf("\twith type: %s ", ab.Field(i).Type())
		fmt.Printf("\tand value _%v_\n", ab.Field(i).Interface())

		// Check whether there are other structures embedded in Record
		k := reflect.TypeOf(ab.Field(i).Interface()).Kind()
		// Need to convert it to string in order to compare it
		if k.String() == "struct" {
			fmt.Println(ab.Field(i).Type())
		}

		// Same as before but using the internal value
		if k == reflect.Struct {
			fmt.Println(ab.Field(i).Type())
		}

	}

	fmt.Println("\n\n New AltMain function")

	AltMain()
}

type Ttype struct {
	F1 int
	F2 string
	F3 float64
}

func AltMain() {
	A := Ttype{1, "F2", 3.0}
	fmt.Println("A:", A)
	r := reflect.ValueOf(&A).Elem()

	//With the use of Elem() and a pointer to variable A, variable A can be modified if needed.
	fmt.Println("String value:", r.String())
	typeOfA := r.Type()
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		tOfA := typeOfA.Field(i).Name
		fmt.Printf("%d: %s %s = %v\n", i, tOfA, f.Type(), f.Interface())

		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		if k == reflect.Int {
			r.Field(i).SetInt(-100)
		} else if k == reflect.String {
			r.Field(i).SetString("Changed!")
		}

	}

	fmt.Println("A:", A)
}
