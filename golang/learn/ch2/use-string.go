package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello there!")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THERE"))

	f("%s\n", s.Title("tHis wiLL be A title!"))

	// The strings.EqualFold() function compares two strings without considering their `case `
	// and returns true when they are the same and false otherwise.

	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAlis"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAli"))

	// The strings.Index() function checks whether the string of the second parameter can be found in
	// the string that is given as the first parameter and returns the index where it was found for the first time.
	f("Index: %v\n", s.Index("Mihalis", "ha")) //2
	f("Index: %v\n", s.Index("Mihalis", "Ha")) //- 1

	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "IS"))

	//strings.Fields() function splits the given string around one or more white space
	//characters as defined by the unicode.IsSpace() function and returns a slice of
	//substrings found in the input string
	t := s.Fields("This is a string!")
	f("Fields: %v\n", len(t))
	t = s.Fields("ThisIs a\tstring!") // here \t is a white space character too

	fmt.Println(t)
	f("Fields: %v\n", len(t))

	/*
		The strings.Split() function allows you to split the given string according to
		the desired separator string—the strings.Split() function returns a string slice.
		Using "" as the second parameter of strings.Split() allows you to process a string character by character.
	*/
	f("Split: %s\n", s.Split("abcd efg", ""))

	/*
		The strings.Replace() function takes four parameters. The first parameter is the string
		that you want to process. The second parameter contains the string that,
		if found, will be replaced by the third parameter of strings.Replace().
		The last parameter is the maximum number of replacements that are allowed to happen.
		If that parameter has a negative value, then there is no limit to the number of
		replacements that can take place.
	*/
	f("Replace-1: %s\n", s.Replace("abcd efg", "", "_", -1))
	f("Replace-2: %s\n", s.Replace("abcd efg", "", "_", 4))
	f("Replace-3 %s\n", s.Replace("abcd efg", "", "_", 2))

	/*
		The strings.SplitAfter() function splits its first parameter string into substrings
		 based on the separator string that is given as the second parameter to the function.
		 The separator string is included in the returned slice.
	*/
	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++")) // returns [123++ 432++ ] -> empty string is the last

	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))

	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join: %s\n", s.Join(lines, "+++"))

}
