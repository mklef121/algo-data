package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"reflect"
	"strconv"
	"time"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

var MIN = 97
var MAX = 122

func main() {
	arguments := os.Args

	// fmt.Println(data, cap(data), len(data), "hia")
	if len(arguments) == 1 {
		// path.Base
		// The argument is like /var/folders/vj/xsf0z1495_915__5r226_phh0000gp/T/go-build2798432677/b001/exe/phonebook
		// So Base funtion gives us the last member which is phonebook
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	if inArray(arguments[1], []string{"search", "list"}) == -1 {
		fmt.Println("Passed argument muse be `search` or `list`")
		return
	}

	// data = append(data, Entry{"Mihalis", "Tsoukalos", "2109416471"})
	// data = append(data, Entry{"Mary", "Doe", "2109416871"})
	// data = append(data, Entry{"John", "Black", "2109416123"})

	SEED := time.Now().Unix()
	rand.Seed(SEED)

	data = populate(100, data)

	// fmt.Println(data, "hia")
	switch arguments[1] {

	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Telephone")
			return
		}
		result := search(arguments[2])

		if result == nil {
			fmt.Println("Entry not found for:", arguments[2])
			return
		}

		fmt.Println(*result)

	case "list":
		list()
	}
}

func search(key string) *Entry {

	for _, person := range data {

		if person.Tel == key {
			return &person
		}
	}

	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func inArray(val interface{}, array interface{}) (index int) {
	values := reflect.ValueOf(array)

	if reflect.TypeOf(array).Kind() == reflect.Slice || values.Len() > 0 {
		for i := 0; i < values.Len(); i++ {
			if reflect.DeepEqual(val, values.Index(i).Interface()) {
				return i
			}
		}
	}

	return -1
}
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(len int64) string {
	temp := ""
	// startChar := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

func populate(count int, data []Entry) []Entry {
	for i := 0; i < count; i++ {
		name := getString(4)
		surname := getString(5)

		tel := strconv.Itoa(random(100, 199))

		data = append(data, Entry{Name: name, Surname: surname, Tel: tel})
	}
	// fmt.Println(data, "fool")
	return data
}
