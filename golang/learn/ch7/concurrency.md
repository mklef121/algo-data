## Go Concurrency

The key component of the Go concurrency model is the **goroutine**, which is the minimum executable entity in Go.

Each goroutine is executed on a single OS thread according to the instructions of the **Go scheduler**, which is responsible for the execution of goroutines.

However, goroutines cannot directly communicate with each other. Data sharing in Go is implemented using either **channels or shared memory**.


When you combine multiple `channels` and `goroutines` you can create data flows, which in Go terminology are also called **pipelines**.


### Processes, threads, and goroutines

**A process** is an OS representation of a running `program`, while a `program` is a binary file on disk that contains all the information necessary for creating an **OS process**.

The binary file is written in a specific format (ELF on Linux) and contains all the instructions the CPU is going to run as well as a plethora of other useful sections. 

That program is loaded into memory and the instructions are executed, creating a running process. 
So, a process carries with it additional resources such as memory, opened file descriptions, and user data as well as other types of resources that are obtained during runtime.


**A thread** is a smaller and lighter entity than a process. Processes consist of one or more threads that have their own flow of control and stack

**A goroutine** is the minimum Go entity that can be executed concurrently. The use of the word minimum is very important here, as goroutines are not autonomous entities like UNIX processes—goroutines live in OS threads that live in OS processes.

The good thing is that goroutines are lighter than threads, which, in turn, are lighter than processes—running thousands or hundreds of thousands of goroutines on a single machine is not a problem.



### The Go scheduler

The **OS kernel scheduler** is responsible for the execution of the threads of a program. This is controlled by the Operating system it self.

But remember Go allows creation of goroutines which are very small and run inside threads (This is not the default way the OS is programmed to work). Therefore there must be some magic happening.

The Go runtime has its own scheduler, which is responsible for the execution of the goroutines using a technique known as **m:n scheduling**


So what is **m:n scheduling** ?

A kernel thread is needed for the actual execution of code and parallelism. But it’s expensive to create, So we map **N** goroutines to **M** Kernel Thread. 

Goroutine is the Go Code, so we have full control over it. Also, it’s in the user-space so it is cheap to create.
But as **OS doesn’t know anything** about the goroutine. Every goroutine has a **state** to help Scheduler knows which goroutine to run based on goroutine state.

These goroutines have several states, they are
- Running — goroutine currently running on kernel thread.
- Runnable — goroutine waiting for kernel thread to run.
- Blocked — Goroutines waiting for some conditions (e.g. blocked on a channel, syscall, mutex, etc.)

So a **Go Runtime Scheduler** manages these goroutines at various states, by Multiplexing N Goroutine to M Kernel Thread.
