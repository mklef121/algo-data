package arrayFuncs

import "fmt"

func main() {
	arr := []string{"Hello"}
	// fmt.Println(arr[4:])

	poped, arr1 := pop(arr)

	fmt.Println(poped, arr1)

	shifted, arr2 := shift(arr)
	fmt.Println(shifted, arr2)
	// sec_arr := arr[1:]

	reversed := stringReverse("Miracle")
	fmt.Println(reversed)

}

func pop(list []string) (string, []string) {
	pop, list := list[len(list)-1], list[:len(list)-1]
	return pop, list
}

func shift(list []string) (string, []string) {
	shift := list[0]

	if len(list) > 1 {
		list = list[1:]
	} else {
		list = []string{}
	}
	return shift, list
}

func unshift(list []string) {

}

func stringReverse(str string) string {

	strRune := []rune(str)
	if len(strRune) < 2 {
		return str
	}
	length := len(strRune)
	fromEnd := length - 1
	fromBegining := 0
	var newRune []rune = make([]rune, length)

	for fromBegining < length {
		newRune[fromBegining] = strRune[fromEnd]
		fromBegining++
		fromEnd--
	}

	return string(newRune)
}

// Filter returns a new slice holding only
// the elements of s that satisfy fn()
func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

//effective slices and arays golang
//https://blog.golang.org/slices-intro
