## Fearless Concurrency

Rust is designed to make **concurrent and parallel programming** both safe and efficient. As modern systems increasingly rely on multiple processors, it's crucial to write code that can run tasks independently or simultaneously. However, doing this safely has traditionally been challenging and error-prone.

Rust changes this by using its **ownership system** and **strong type checking** to eliminate many concurrency bugs **at compile time** instead of runtime. This means you catch potential data races and thread safety issues early—before your program even runs. As a result, Rust allows developers to write robust, concurrent code with confidence. This concept is known as **fearless concurrency**.

Unlike many languages that impose strict patterns for handling concurrency (like message-passing in Erlang), Rust gives you **flexibility**. It offers multiple tools—such as threads, channels, mutexes, and atomic types—allowing you to choose the best approach for your specific use case without compromising performance or control.


### Using Threads to Run Code Simultaneously

In most modern operating systems, a **program runs inside a process**, and the OS can manage **multiple processes** at once. Within a single program, you can divide work into **threads**—independent units of execution that can run concurrently. For example, a web server might use multiple threads to handle several incoming requests simultaneously.

Running tasks in parallel using threads can boost performance, but it also introduces complexity. Since threads can execute at the same time, there’s **no guaranteed order** of execution, which can cause problems such as:

- **Race conditions** – multiple threads access shared data in conflicting ways.
- **Deadlocks** – threads wait indefinitely for resources held by each other.
- **Hard-to-reproduce bugs** – timing-based issues that only happen in specific situations.

Rust uses the **1:1 threading model**, meaning each Rust thread maps directly to an **operating system thread**, leveraging the OS’s built-in thread management APIs.

```rust
use std::thread;
use std::time::Duration;

fn main() {
    thread::spawn(|| {
        for i in 1..10 {
            println!("hi number {i} from the spawned thread!");
            thread::sleep(Duration::from_millis(1));
        }
    });

    for i in 1..5 {
        println!("hi number {i} from the main thread!");
        thread::sleep(Duration::from_millis(1));
    }
}
```
To create a new thread, we call the thread::spawn function and pass it a closure containing the code we want to run in the new thread. 

> Note that when the main thread of a Rust program completes, all spawned threads are shut down, whether or not they have finished running. 

To fix the problem of the main thread killing other threads when it's done with it's execution, let's look at the code below

```rust
use std::thread;
use std::time::Duration;

fn main() {
    let handle = thread::spawn(|| {
        for i in 1..10 {
            println!("hi number {i} from the spawned thread!");
            thread::sleep(Duration::from_millis(1));
        }
    });

    for i in 1..5 {
        println!("hi number {i} from the main thread!");
        thread::sleep(Duration::from_millis(1));
    }

    handle.join().unwrap();
}
```
The return type of `thread::spawn` is `JoinHandle`. A `JoinHandle` is an owned value that, when we call the join method on it, will wait for its thread to finish. 
Calling `join` on the handle blocks the thread currently running until the thread represented by the handle terminates. *Blocking* a thread means that thread is prevented from performing work or exiting

### Using move Closures with Threads
We’ll often use the move keyword with closures passed to thread::spawn because the closure will then take ownership of the values it uses from the environment, thus transferring ownership of those values from one thread to another.

```rust
let v = vec![1, 2, 3];

let handle = thread::spawn(|| {
    println!("Here's a vector: {v:?}");
});
```

The code above will not compile because the rust compiler does not know if the thread will outlive the borrowed value. To fix this, we can use the move closure

```rust
 let v = vec![1, 2, 3];

let handle = thread::spawn(move || {
    println!("Here's a vector: {v:?}");
});
```

By adding the move keyword before the closure, we force the closure to take ownership of the values it’s using rather than allowing Rust to infer that it should borrow the values


### Using Message Passing to Transfer Data Between Threads


A widely adopted method for achieving safe concurrency is **message passing**, where threads **communicate by sending data** to each other, rather than sharing memory directly. This idea is captured well by the Go programming motto:

> “Don’t communicate by sharing memory; share memory by communicating.”

Rust supports this model through its standard library’s **channel system**, which allows **data to be sent safely between threads**.

We create a new channel using the `mpsc::channel` function; `mpsc` stands for multiple producer, single consumer. 

> In short, the way Rust’s standard library implements channels means a channel can have multiple sending ends that produce values but only one receiving end that consumes those values.

`let (tx, rx) = mpsc::channel();`: The mpsc::channel() function returns a tuple returning a transmitter and receiver respectively.

```rust
use std::sync::mpsc;
use std::thread;

fn main() {
    let (tx, rx) = mpsc::channel();

    thread::spawn(move || {
        let val = String::from("hi");
        tx.send(val).unwrap();
    });
    // recv function blocks and waits until it receives a message
    //  When the transmitter closes, recv will return an error to signal that no more values will be coming.
    // there is a try_recv alternative which does not block but can be called to check data in the channel
    let received = rx.recv().unwrap();
    println!("Got: {received}");

    // we can also use an iterator to read values from a receiver
     for received in rx {
        println!("Got: {received}");
    }
}
```

#### Creating Multiple Producers by Cloning the Transmitter
Earlier we mentioned that **mpsc** was an acronym for multiple producer, single consumer. We can do so by cloning the transmitter

```rust
    // --snip--

    let (tx, rx) = mpsc::channel();

    let tx1 = tx.clone();
    thread::spawn(move || {
        let vals = vec![
            String::from("hi"),
            String::from("from"),
            String::from("the"),
            String::from("thread"),
        ];

        for val in vals {
            tx1.send(val).unwrap();
            thread::sleep(Duration::from_secs(1));
        }
    });

    thread::spawn(move || {
        let vals = vec![
            String::from("more"),
            String::from("messages"),
            String::from("for"),
            String::from("you"),
        ];

        for val in vals {
            tx.send(val).unwrap();
            thread::sleep(Duration::from_secs(1));
        }
    });

    for received in rx {
        println!("Got: {received}");
    }

```


### Shared-State Concurrency

**Message passing** is a popular and safe approach to concurrency, but it’s not the only option. Another common method is **shared memory**, where multiple threads access the same data directly.

The Go language’s guidance — *“Do not communicate by sharing memory”* — highlights the potential risks of this approach. But what does it actually mean to communicate **by** sharing memory?

**Comparing the Two Models**

- **Message Passing** resembles **single ownership**: when a value is sent through a channel, it moves from one thread to another, and the sender no longer uses it.
- **Shared Memory** is like **multiple ownership**: many threads can access (and potentially mutate) the same memory at once.

This can lead to **data races, inconsistent state, or difficult-to-debug concurrency bugs**, which is why message-passing advocates caution against shared memory unless it's handled with great care.

#### Using Mutexes to Allow Access to Data from One Thread at a Time
Here’s a clearer, tighter rewrite of your explanation with a better summary:

---

### **Understanding Mutex in Rust**

A **Mutex** (short for *mutual exclusion*) ensures that **only one thread can access data at a time**. It works through a locking mechanism—before a thread can read or write the data, it must first acquire the **lock**. Once finished, it must **release the lock** so other threads can access the data.

This lock acts as a **gatekeeper**, ensuring exclusive access and preventing simultaneous modification, which could lead to inconsistent or corrupt state.

**Why Mutexes Can Be Tricky**

In many languages, forgetting to acquire or release a lock can cause **deadlocks** or **resource contention**—serious bugs that are hard to debug.

But in **Rust**, the ownership model and type system ensure that:

- You **must** acquire the lock before accessing the data.
- The lock is **automatically released** when it goes out of scope.

This makes using `Mutex<T>` in Rust **safe and ergonomic**, reducing the risk of classic concurrency bugs.

```rust
use std::sync::Mutex;

fn main() {
    let m = Mutex::new(5);

    {
        // To access the data inside the mutex, we use the lock method to acquire the lock. 
        let mut num = m.lock().unwrap();
        *num = 6;
    }

    println!("m = {m:?}");
}
```
The call to lock would fail if another thread holding the lock panicked.

`Arc<T>` is similar to `Rc<T>`, but it’s designed for **safe use in concurrent environments**. The “A” stands for **atomic**, meaning it uses **atomic reference counting** to safely manage shared ownership across multiple threads.

Atomic operations ensure that reference counts are updated correctly, even when accessed simultaneously from different threads.

You might wonder: *Why isn’t everything just atomic and thread-safe by default?*  
The answer is **performance**. Atomic operations are **slower** than non-atomic ones because they involve extra steps to guarantee thread safety.

So, Rust lets you choose:  
- Use `Rc<T>` for **single-threaded** scenarios (faster).
- Use `Arc<T>` when you need to **share data across threads** (safe, but slightly slower).

```rust
use std::sync::{Arc, Mutex};
use std::thread;

fn main() {
    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];

    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            let mut num = counter.lock().unwrap();

            *num += 1;
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }

    println!("Result: {}", *counter.lock().unwrap());
}
```

### **Concurrency in Rust: `Send` and `Sync` Traits**

Rust's core language includes very few built-in concurrency features — most are provided by the standard library. However, two key traits built into the language itself are essential for safe concurrency: `Send` and `Sync`.


#### **`Send`: Ownership Transfer Across Threads**

The `Send` marker trait indicates that a type's ownership can safely be transferred to another thread. Most Rust types are `Send` by default.

However, **`Rc<T>` is *not* `Send`**, because sharing an `Rc` across threads could lead to simultaneous mutations of the reference count, which is not thread-safe.


#### **`Sync`: Shared Access Across Threads**

The `Sync` marker trait allows a type to be **safely referenced from multiple threads**. A type `T` is `Sync` if `&T` (an immutable reference) can be safely shared between threads.

Again, **`Rc<T>` is *not* `Sync`**, and neither is `RefCell<T>`, because their internal mutation logic isn’t thread-safe. In contrast, **`Mutex<T>` is `Sync`**, making it safe for shared access across threads.

#### **Unsafe Manual Implementation**

You **shouldn't manually implement `Send` or `Sync`** unless you're writing unsafe code and know what you're doing. Rust will auto-implement these traits for you if your type is composed of parts that are already `Send` or `Sync`.







