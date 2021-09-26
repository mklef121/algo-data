
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

### About slice length and capacity

1. **Capacity :** The capacity shows how much a slice can be expanded without the need to allocate more memory and change the underlying array. Although after slice creation the capacity of a slice is handled by Go, a developer can define the capacity of a slice at creation time using the make() function.

2. **Length :** The


### Joining an array/slice to a slice

The `append` function takes multiple arguments of same type to be added to the initial array. So go provides the `...` operator to split a slice or array into individual individual.

```go
	theSlice := make([]int, 2, 4)
	fmt.Println(theSlice)
	theSlice[0] = 3
	theSlice[1] = 5

    addingSlice := []int{-98, -12, -563, -4}
    theSlice = append(theSlice, addingSlice...)
```

### Selecting a part of a slice

In Go you select a part of a slice by defining two indexes, the first one is the beginning of the selection whereas the second one is the end of the selection, without including the element at that index, separated by `:` .

```go
//Select the cmd arguments except the one at index 0
arguments := os.Args[1:]

addingSlice := []int{-98, -12, -563, -4}

//selects -98, -12, -563 to a new slice
selected := addingSlice[0:2]

aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// First 5 elements
fmt.Println(aSlice[0:5])
// First 5 elements
fmt.Println(aSlice[:5])

// Elements at indexes 2,3,4
// Capacity will be 10-2
t = aSlice[2:5:10]
fmt.Println(len(t), cap(t))
```

> Setting the correct capacity of a slice, if known in advance, will make your programs faster because Go will not have to allocate a new underlying array and have all the data copied over.


### Byte slices

**Byte :** The byte type in Go is just an alias for **uint8**. This is a number that has 8 bits of storage. 8 bits have 256 possible combinations of "off" and "on," uint8 has 256 possible integer values from 0 to 255.

What is special is that Go uses byte slices for performing file I/O operations because they allow you to determine with precision the amount of data you want to read or write to a file. This happens because bytes are a universal unit among computer systems.

```go

// Byte slice
	b := make([]byte, 12)
    // b now points to a different memory location than before, which is where "Byte slice €" is stored.
	b = []byte("Byte slice €")
    // The cap and length of this will be bigger than expected since there is a unicode `€`
    fmt.Println("Byte slice: ", b, "cap: ", cap(b), "length: ", len(b))

    //convert and print byte slice contents as text
	aStr := string(b)
	fmt.Println(aStr)

```

### Deleting an element from a slice

There is no default function for deleting an element from a slice, which means that if you need to delete an element from a slice, you must write your own code. 

#### We presents two techniques for deleting an element from a splice

- The first technique virtually divides the original slice into two slices, split at the index of the element that needs to be deleted. Neither of the two slices includes the element that is going to be deleted. After that, we concatenate these two slices and creates a new one. 
    ```go

        aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
        index := 4
        aSlice = append(aSlice[:index], aSlice[index+1:]...)
	    fmt.Println("After 1st deletion:", aSlice)
    ```
- The second technique copies the last element at the place of the element that is going to be deleted and creates a new slice by excluding the last element from the original slice.
    ```go

        theSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
        // Replace element at index i with last element
        theSlice[index] = theSlice[len(theSlice)-1]
        // Remove last element
        theSlice = theSlice[:len(theSlice)-1]
        fmt.Println("After 2nd deletion:", theSlice)
        //This seems to be more optimal but does not take into consideration the maintenance of values at each index
    ```


### How slices are connected to arrays
Behind the scenes, each slice is implemented using an **underlying array**. The length of the underlying array is the same as the **capacity of the slice** and there exist pointers that connect the slice elements to the appropriate array elements.

 Go allows you to reference an array or a part of an array using a slice. However, when the capacity of the slice changes, the connection to the array ceases to exist!