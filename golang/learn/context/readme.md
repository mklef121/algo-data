## Go Context 101

Go Context is an interface with only four functions, time(Deadline), signal(Done), exception(error), and data(Value).

```go

var ctx context.Context

// A Context carries a deadline, cancellation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
type Context interface {
    // Done returns a channel that is closed when this Context is canceled
    // or times out.
    Done() <-chan struct{}

    // Err indicates why this context was canceled, after the Done channel
    // is closed.
    Err() error

    // Deadline returns the time when this Context will be canceled, if any. Deadline returns ok==false when no deadline is
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    Value(key interface{}) interface{}
}

```

#### Understanding Context

- The problem

> In Go servers, each incoming request is handled in its own **goroutine**. Request handlers often start additional **goroutines** to access backends such as databases and RPC services. The **set of goroutines working on a request** typically needs access to request-specific values such as the identity of the end user, authorization tokens, and the request’s deadline. When a request is canceled or times out, all the goroutines working on that request should exit quickly so the system can reclaim any resources they are using.

- The solution 

> Google developed a context package that makes it easy to pass request-scoped values, cancellation signals, and deadlines across API boundaries to all the goroutines involved in handling a request.


### Using contexts 

developers don’t need to implement the Context interface directly since Go already provides two useful implementations, backgroundand todo.

The can be instanciated using the 

`context.Background()` or `context.TODO()` methods

They both return a context to use in the **parent context** thread.

### Implementation of Context

After creating the `parent context` via `context.Background()`, we can derive more sub-contexts from it by the four `With*` functions provided by the context package.

```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context
```

The `WithCancel`, `WithDeadline`, and `WithTimeout` methods all return a **CancelFunc** function that can propagate to the sub-contexts.








