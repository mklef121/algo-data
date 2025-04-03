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
    numbs := strconv.FormatInt(int64(N), 2)
    var (
        count int
        max int
        startCount bool
    )

    for i := 0; i < len(numbs); i++{
        val := string(numbs[i])
        
        if val == "1" {
            if count > max {
                max = count
            }
            count = 0
            startCount = true
        }else if val == "0" && startCount {
            count++
        }
    }

    return max
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

