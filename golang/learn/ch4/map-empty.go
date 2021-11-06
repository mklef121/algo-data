package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var JSONrecord = `{
    "Flag": true,
    "MyArray": ["a","b","c",7,9],
    "Entity": {
      "a1": "b1",
      "a2": "b2",
      "Value": -456,
      "Null": null
	},
    "Message": "Hello Go!"
  }`

func main() {
	if len(os.Args) == 1 {
		fmt.Println("*** Using default JSON record.")
	} else {
		JSONrecord = os.Args[1]
	}

	JSONMap := make(map[string]interface{})
	//json.Unmarshal() processes JSON data and converts it into a Go value.
	err := json.Unmarshal([]byte(JSONrecord), &JSONMap)

	if err != nil {
		fmt.Println(err)
		return
	}

	exploreMap(JSONMap)

	fmt.Println("\n\n Starting typeSwitch \n\n")
	typeSwitch(JSONMap)
}

func typeSwitch(m map[string]interface{}) {
	for key, value := range m {
		switch value := value.(type) {
		case string:
			fmt.Println("Is a string!", key, value)
		case float64:
			fmt.Println("Is a float64!", key, value)
		case bool:
			fmt.Println("Is a Boolean!", key, value)
		case map[string]interface{}:
			fmt.Println("Is a map!", key, value)
			typeSwitch(value)
		default:
			fmt.Printf("...Is %v: %T!\n", key, value)

		}
	}

	return
}

func exploreMap(m map[string]interface{}) {
	for key, value := range m {
		moreMap, ok := value.(map[string]interface{})

		// If it is a map, explore deeper
		if ok {
			fmt.Printf("{\"%v\": \n", key)
			exploreMap(moreMap)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v: %v\n", key, value)
		}
	}
}
