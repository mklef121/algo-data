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