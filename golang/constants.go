package main

import "fmt"

// We have two types of data we need to cache: books and CDs. Both use the ID,

const GlobalLimit = 100

// Create a MaxCacheSize that is 10 times the global limit size:
const MaxCacheSize int = 10 * GlobalLimit

// 6. Create our cache prefixes:
const (
	CacheKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

var cache map[string]string

func main() {
	//The make method initializes or creates the empty cache with storage setthe cache
	cache = make(map[string]string)
	isbn1 := "1234-5678"
	SetBook(isbn1, "Get Ready To Go")
	SetCD(isbn1, "Get Ready To Go Audio Book")

	fmt.Println("Book :", GetBook(isbn1))
	fmt.Println("CD :", GetCD(isbn1))
	fmt.Println(cache)

	const (
		first = iota
		second
		// When a constant’s type and value is not declared, it will get it from the previous constant.
		// Here above, second and  third gets its type and value from Pi.
		third
	)

	fmt.Println(first, second, third)

	const Tau = 3.14 * 2

	// implicitly converted to a float64,
	// because of the assignment to a "runtime" variable
	// so, a type is needed for this "context"
	pi := Tau / 2

	{
		level := "Nest 1"
		fmt.Println("Block end:", level, pi)
	}

}

func cacheGet(key string) string {
	return cache[key]
}

func cacheSet(key, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val
}

// Create a function to get a book from the cache:
func GetBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

func SetBook(isbn string, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

// Create a function to get CD data from the cache:
func GetCD(sku string) string {
	return cacheGet(CacheKeyCD + sku)
}

func SetCD(sku string, name string) {
	cacheSet(CacheKeyCD+sku, name)
}
