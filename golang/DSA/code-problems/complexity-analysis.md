### Complexity Analysis Problems and Solutions

- BinaryGap

https://app.codility.com/programmers/lessons/1-iterations/binary_gap/start/

```go
package solution

// you can also use imports, for example:
import "strconv"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
    // Implement your solution here
    binaryString := strconv.FormatInt(int64(N), 2)
    inprogress := false
    largest := 0
    count := 0

    for i := 0; i < len(binaryString); i++ {
        if string(binaryString[i]) == "1" {
            if inprogress {
                inprogress = false

                if count > largest {
                    largest = count
                }
                count = 0
            }
        }else {
            // current item is "0"
           inprogress = true
           count++

        }
    }

    return largest
}

```

- CyclicRotation

https://app.codility.com/programmers/lessons/2-arrays/cyclic_rotation/start/

```go
package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, K int) []int {
    // Implement your solution here
    length := len(A)
    if length == 0 || length == 1 {
        return A
    }

    for j := 0; j < K; j++ {
        last := A[length - 1]
        for i := length - 1; i > 0; i-- {
            A[i] = A[i-1]
        }
        A[0] = last
    }

    return A
    
}
```

