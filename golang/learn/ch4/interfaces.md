> Go is statically typed. Every variable has a static type, that is, exactly one type known and fixed at compile time: int, float32, *MyType, []byte, and so on. 

If we declare

```go
type MyInt int

var i int
var j MyInt
```

then `i` has type `int` and `j` has type `MyInt`. The variables `i` and `j` have distinct static types and, although they have the same underlying type, they cannot be assigned to one another without a conversion.

**One important category of type is interface types**

## What is the interface type

An interface type is a go data type which represents fixed sets of methods. An interface variable can store any concrete (non-interface) value as long as that value implements the interface’s methods.

Better still **interfaces** are abstract types that specify a set of methods that need to be implemented so that another type
can be considered an instance of the interface.

Example:

```go
//a pair of examples is io.Reader and io.Writer, the types Reader and Writer from the io package:

// Reader is the interface that wraps the basic Read method.
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Writer is the interface that wraps the basic Write method.
type Writer interface {
    Write(p []byte) (n int, err error)
}

// Any type that implements a Read (or Write) method with this signature is said to implement io.Reader (or io.Writer). 

var r io.Reader
r = os.Stdin
r = bufio.NewReader(r)
r = new(bytes.Buffer)
// and so on

```

An extremely important example of an interface type is the empty interface:

```go
interface{}

//It represents the empty set of methods and is satisfied by any value at all, since any value has zero or more methods.
```

NOTE: Interfaces work with methods on types (or type methods), which are like functions attached to given data types. As you already know, once you implement the required type methods of an interface, that interface is satisfied implicitly.
The `empty interface` is defined as just interface{}. As the `empty interface` has no methods, it means that it is already implemented by all data types. For a data type to satisfy an interface, it needs to implement all the type methods required by that interface



### Interface composition in golang

Lets say we have a program like the one below 

```go
package main

import "fmt"

type Article struct {
    Title string
    Author string
}

func (a Article) String() string {
    return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

func main() {
    a := Article{
        Title: "Understanding Interfaces in Go",
        Author: "Sammy Shark",
    }
    fmt.Println(a.String())
}
```

instead of always calling fmt.Println always we can create a stringer interface this

```go 


type Stringer interface {
    String() string
}

func main() {
    //Article already has a type method `String` this implements the  Stringer interface
    a := Article{
        Title: "Understanding Interfaces in Go",
        Author: "Sammy Shark",
    }
    Print(a)
}

func Print(s Stringer) {
    fmt.Println(s.String())
}
```


In case we create another struct B

```go
type Book struct {
    Title  string
    Author string
    Pages  int
}

func (b Book) String() string {
    return fmt.Sprintf("The %q book was written by %s.", b.Title, b.Author)
}

func main() {
    b := Book{
        Title:  "All About Go",
        Author: "Jenny Dolphin",
        Pages:  25,
    }
    Print(b)
}

// we can still do thus 
```

Now we can compose more than one interfaces to form a single interface thus

```go

//an interface for getting the area of a circle
type Sizer interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) String() string {
    return fmt.Sprintf("Circle {Radius: %.2f}", c.Radius)
}

func PrintArea(s Shaper) {
    fmt.Printf("area of %s is %.2f\n", s.String(), s.Area())
}

func main() {
    c := Circle{Radius: 10}
    PrintArea(c)
}

```


From the example above, you can see that the Circle struct also implements the string method as well as area method, 
we can now create another interface to describe that wider set of behavior. To do this, we’ll create an interface called Shaper.

```go
type Shaper interface {
    Sizer
    fmt.Stringer //contains the String interface
}

//thus we can have

func PrintArea(s Shaper) {
    fmt.Printf("area of %s is %.2f\n", s.String(), s.Area())
}

func main(){
   var a Shaper
   a = Circle{Radius: 10}
   PrintArea(a)
}

```

