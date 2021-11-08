package main

import "fmt"

// This programe illustrates composition and polymorphism as well as embedding
// an anonymous structure into an existing one to get all its fields.

type IntA interface {
	foo()
}

type IntB interface {
	bar()
	foo()
}

type IntC interface {
	IntA
	IntB
}

func processA(s IntA) {
	fmt.Printf("%T\n", s)
}

type aStruct struct {
	XX int
	YY int
}

type bStruct struct {
	AA string
	XX int
}

// Structure c has two fields
type cStruct struct {
	A aStruct
	B bStruct
}

// Structure compose gets the fields of structure a
type compose struct {
	field1  int
	aStruct // this is different from aStruct: aStruct
	hi      int
}

// Satisfying IntA
func (varC cStruct) foo() {
	// fmt.Printf("The type shit %T\n", varC)
	fmt.Println("Foo Processing", varC)
}

// Satisfying IntB
func (varC cStruct) bar() {

	fmt.Println("Bar Processing", varC)
}

/* As `cStruct` satisfies both IntA and IntB, it implicitly satisfies IntC,
which is a composition of the IntA and IntB interfaces. */
func main() {
	me := compose{
		field1:  5,
		aStruct: aStruct{},
	}

	var iC cStruct = cStruct{aStruct{120, 12}, bStruct{"-12", -12}}
	fmt.Println(me, iC)
	// iC.
	iC.A.A()
	iC.B.A()

	//The anonymose structure aStruct  definition now makes the compose struct inherit all
	// properties of aStruct. More like inheritance
	iComp := compose{123, aStruct{456, 789}, 676}
	fmt.Println(iComp.XX, iComp.YY, iComp.field1)

	processA(iC)

}

// Different structures can have methods with the same name
func (A aStruct) A() {
	fmt.Println("Function A() for A")
}
func (B bStruct) A() {
	fmt.Println("Function A() for B")
}
