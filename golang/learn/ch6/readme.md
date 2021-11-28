> This chapter teaches you about systems programming in Go. Systems programming involves working with files and directories, process control, signal handling, network programming, system files, configuration files, and file input and output (I/O).


Remember that UNIX considers everything, even a printer or your mouse, as a
file. UNIX uses **file descriptors**, which are **positive integer** values, as an internal representation for accessing open files, which is much prettier than using long paths.

### stdin, stdout, and stderr
Every UNIX operating system has three files open all the time for its processes.
So, by default, all UNIX systems support three special and standard filenames: `/dev/stdin`, `/dev/stdout`, and `/dev/stderr`
which can also be accessed using file descriptors **0, 1, and 2**, respectively.
These three file descriptors are also called standard input, standard output, and standard error, respectively.


Go uses `os.Stdin` for accessing standard input, `os.Stdout` for accessing standard output, and `os.Stderr` for accessing standard error.


### UNIX processes

**A process** is an execution environment that contains instructions, user data and system data parts, and other types of resources that are obtained during runtime.

**A program** is a binary file that contains instructions and data that are used for initializing the instruction and user data parts of a process. 

Each running UNIX process is uniquely identified by an **unsigned integer**, which is called the process ID of the process.

There are three process categories: **user processes**, **daemon processes**, and **kernel processes**. 

- User processes run in user space and usually have no special access rights.
- Daemon processes are programs that can be found in the user space and run in the background without the need for a terminal
- Kernel processes are executed in kernel space only and can fully access all kernel data structures.


### Handling UNIX signals

A signal is a message which can be sent to a running process **asynchronously**.
Signals can be initiated by programs, users, or administrators.

For instance `ps aux` lists all the available 

Then the `kill` command is used to send signals to a process

e.g
The proper method of telling the Internet Daemon (inetd) to re-read its configuration file is to send it a SIGHUP signal.

`kill -SIGHUP 4140`

UNIX signal handling requires the use of **Go channels** that are used exclusively for this task.
The concurrency model of Go requires the use of goroutines and channels.

### Concurrency and signal handling in Go

A **goroutine** is the smallest executable Go entity.

A **channel** in Go is a mechanism that among other things allows goroutines to communicate and exchange data.

> Some signals cannot be caught, and the operating system cannot ignore them. So, the **SIGKILL** and **SIGSTOP** signals cannot be blocked, caught, or ignored and the reason for this is that they allow privileged users as well as the UNIX kernel to terminate any process they desire.


Look up [Here to see a list of signals](http://www.math.stonybrook.edu/~ccc/dfc/dfc/signals.html)

