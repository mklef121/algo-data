package main

import "fmt"

//https://golang.org/doc/faq go frequently asked question
func main() {

	indices := []string{"Come", "go", "follow", "string"}

	res := array_pop(&indices)

	fmt.Println(res, indices)
}

func array_pop(arg *[]string) interface{} {
	if len(*arg) == 0 {
		return nil
	}

	lastIndex := len(*arg) - 1
	lastVal := (*arg)[lastIndex]

	*arg = (*arg)[:lastIndex]

	return lastVal
}

func array_unshift(arr *[]interface{}, elements ...interface{}) int {
	*arr = append(*arr, elements...)
	return len(*arr)
}

func array_splice(arr *[]interface{}, offset int, length int, mixed ...interface{}) {

}

// https://stackoverflow.com/questions/13544374/idiomatic-slice-splice-in-go
//
//splice
