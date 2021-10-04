package main

import "fmt"

//Struct definition
type Entry struct {
	Name    string
	Surname string
	Year    int
}

func main() {

	//The type keyword allows you to define new data types or create aliases for existing ones

	//Note

	type myInt int

	var me myInt
	//is not equal to
	var her int

	//  here myInt and int are different types even though they have the sane foundation

	me = 45
	her = 98

	fmt.Println(me, her)

	//Two ways to create a structure

	// 1- Initializing it to a variable
	// 2- Creating a pointer to it

	s1 := zeroS()
	p1 := zeroPtoS()
	fmt.Println("s1:", s1, "p1:", *p1)

	s2 := initS("Mihalis", "Tsoukalos", 2020)
	p2 := initPtoS("Mihalis", "Tsoukalos", 2020)
	fmt.Println("s2:", s2, "p2:", *p2)

	fmt.Println("Year:", s1.Year, s2.Year, (*p1).Year, p2.Year)

	// hero := &[]string{"tell", "maga"}

	// fmt.Println((*hero)[0])

	pS := new(Entry)
	fmt.Println("New pS:", pS)
}

// Initialized by Go
func zeroS() Entry {
	return Entry{}
}

// Initialized by Go - returns pointer
func zeroPtoS() *Entry {
	return &Entry{}
}

// Initialized by the user - returns pointer
func initPtoS(N, S string, Y int) *Entry {
	if Y < 2000 {
		return &Entry{Name: N, Surname: S, Year: 2000}
	}

	return &Entry{Name: N, Surname: S, Year: Y}
}

// Initialized by the user
func initS(N, S string, Y int) Entry {
	if Y < 2000 {
		return Entry{Name: N, Surname: S, Year: 2000}
	}
	return Entry{Name: N, Surname: S, Year: Y}
}
