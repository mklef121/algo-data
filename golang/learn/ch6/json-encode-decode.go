package main

import (
	"encoding/json"
	"fmt"
)

type UseAll struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"created"`
}

// Ignoring empty fields in JSON
type NoEmpty struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"creationyear,omitempty"` // convert into a JSON record without including this field if it's empty

	//imagine that you have some sensitive data on some of the fields of a Go structure that
	//you do not want to include in the JSON records. You can do that by including the "-"
	//special value in the desired json: structure tags.
	Pass string `json:"-"`
}

func main() {
	useall := UseAll{Name: "Mike", Surname: "Tsoukalos", Year: 2021}

	// Encoding a structure as a string
	t, err := json.Marshal(useall)

	if err != nil {
		fmt.Println("Error doing Json encoding", err)
	} else {
		fmt.Printf("Value %s\n", t)
	}

	//Decoding a string into a structure

	// Decoding JSON data given as a string
	str := `{"username": "M.", "surname": "Ts", "created":2020}`

	newVal := UseAll{}
	err = json.Unmarshal([]byte(str), &newVal)

	if err != nil {
		fmt.Println("\n There was an error decoding Json", err)
	} else {
		fmt.Println("\nstruct Value ", newVal)
	}
}
