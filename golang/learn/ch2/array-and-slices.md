
## Arrays

Arrays in Go have the following characteristics and limitations:

- When defining an array variable, you must define its size. Otherwise, you should put `[...]` in the array declaration and let the Go compiler find out the length for you. 

    ```go 
     known_length_array := [4]string{"Zero", "One", "Two", "Three"} 
     compiler_length_array := [...]string{"Zero", "One", "Two", "Three"}

     a_slice := []string{"Zero", "One", "Two", "Three"}
    ```
- You cannot change the size of an array after you have created it.
- When you pass an array to a function, what is happening is that Go creates a copy of that array and passes that copy to that function—therefore any changes you make to an array inside a function are lost when the function returns.

## Slices

Slices in Go are more powerful than arrays mainly because they are dynamic, which means that they can grow or shrink after creation if needed.

> Additionally, any changes you make to a slice inside a function also affect the original slice. Also the slice value does not include its elements, just a pointer to the underlying array.

** a slice ** value is a header that contains a pointer to an underlying array where the elements are actually stored, the length of the array, and its capacity.

```go
//The interface that defines the structure of a slice
type SliceHeader struct {
    Data uintptr
    Len int
    Cap int 
}

// creating slices

aSlice := []float64{1.2, 3.2, -4.5}

//If you do not want to initialize a slice, then using make() is better and faster
emptySlice := make([]float64, 3)

//Both slices and arrays can have many dimensions. This is what we call associative arrays in php 
twoDEmptySlice := make([][]int, 2)
twoD := [][]int{{1, 2, 3}, {4, 5, 6}}

//Find the length of an array or slice
length := len(twoD)

```

A side effect of passing the slice header is that it is faster to pass a slice to a function because Go does not need to make a copy of the slice and its elements, just the slice header.