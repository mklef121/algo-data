
> Each Go package can optionally have a private function named **init()** that is automatically executed at the beginning of execution time. `init()` runs when the package is initialized at the beginning of program execution. 

This Init function has the following characteristics

init() takes no arguments.
-  `init()` returns no values.
-  The `init()` function is optional.
-  The `init()` function is called implicitly by Go.
-  You can have an `init()` function in the main package. In that case, `init()` is executed before the `main()` function. In fact, all `init()` functions are always executed prior to the `main()` function.
-  A source file can contain multiple init() functions—these are executed in the order of declaration.
-  The `init()` function or functions of a package are executed only once, even if the package is imported multiple times.
-  Go packages can contain multiple files. Each source file can contain one or more `init()` functions.
-  `init()` function is a private function by design means that it cannot be called from outside the package in which it is contained.



### Aliasing Imported Packages

You may want to change a package name if you have a local package already named the same as a third party package you are using. When this happens, aliasing your import is the best way to handle the collision. 

`import another_name "package"`

e.g


```go
package main

import (
 fm_we "fmt"
  "math/rand"
)

func main() {
  for i := 0; i < 10; i++ {
    fm_we.Printf("%d) %d\n", i, rand.Intn(25))
  }
}
```


To import a package solely for its side-effects (initialization), use the blank identifier as explicit package name: 

`import _ "lib/math"`

In this example, the init() function will runn without importing any other functions: