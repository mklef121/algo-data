package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sort"
)

/*** The Different CSV row formats **/

// Format 1
type Format1 struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

// Format 2
type Format2 struct {
	Name       string
	Surname    string
	Areacode   string
	Tel        string
	LastAccess string
}

type Book1 []Format1
type Book2 []Format2

// sorting interface for both Book1 and Book2
func (a Book1) Len() int {
	return len(a)
}

// First based on surname. If they have the same
// surname take into account the name.
func (a Book1) Less(i, j int) bool {
	if a[i].Surname == a[j].Surname {
		return a[i].Name < a[j].Name
	}
	return a[i].Surname < a[j].Surname
}

func (a Book1) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Implement sort.Interface for Book2
func (a Book2) Len() int {
	return len(a)
}

// First based on areacode. If they have the same
// areacode take into account the surname.
func (a Book2) Less(i, j int) bool {
	if a[i].Areacode == a[j].Areacode {
		return a[i].Surname < a[j].Surname
	}
	return a[i].Areacode < a[j].Areacode
}

func (a Book2) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {

}

var data1 = Book1{}
var data2 = Book2{}

func readCSVFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}

	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		return err
	}

	var firstLine bool = true
	var format1 = true

	for _, line := range lines {

		if firstLine {
			if len(line) == 4 {
				format1 = true
			} else if len(line) == 5 {
				format1 = false
			} else {
				return errors.New("Unknown File Format!")
			}

			firstLine = false
		}

		if format1 {
			temp1 := Format1{
				Name:       line[0],
				Surname:    line[1],
				Tel:        line[2],
				LastAccess: line[3],
			}

			data1 = append(data1, temp1)
		} else {
			temp2 := Format2{
				Name:       line[0],
				Surname:    line[1],
				Areacode:   line[2],
				Tel:        line[3],
				LastAccess: line[4],
			}

			data2 = append(data2, temp2)
		}
	}

	return nil
}

func sortData(data interface{}) {
	// type switch
	switch T := data.(type) {
	case Book1:
		sort.Sort(T)
		list(T)
	case Book2:
		sort.Sort(T)
		list(T)
	default:
		fmt.Printf("Not supported type: %T\n", T)
	}
}

func list(d interface{}) {

	switch T := d.(type) {
	case Book1:
		for _, data := range T {
			fmt.Println(data)
		}
		break
	case Book2:
		for _, data := range T {
			fmt.Println(data)
		}
		break
	default:
		fmt.Printf("Not supported type: %T\n", T)
	}
}
