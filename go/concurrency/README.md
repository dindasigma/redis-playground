# Concurrency in Go

## Preliminary
### Process
Process is a program in execution. When we execute our program, it will become a process that performs all the tasks mentioned in the program.

When a program is loaded into memory, and it becomes a process, it can be divided into four sections:

![process](https://miro.medium.com/max/572/1*pplsGMeRKFcc0IHr1j3YwA.jpeg)

- Stack: contains the temporary data such as method/function param, return address, and local variable.
- Heap: dynamically allocated memory to a process during it's run time.
- Text: include current activity presented by the value of the program counter and the processor's registers' contents.
- Data: contains global and state variable

### Stack and Heap
Go implements an optimization called escape analysis. If any references to a value do not escape in the function, it will be stored on the stack.

Go also implements a garbage collector to free up memory space used in the heap that is no longer used. Stack does not require a garbage collector because it’s just a push/pop operation.

```
func​ ​stayOnStack​() ​user​ {
    ​// in the stack frame, create a value and initialize it.
    u := user{name: ​"Dinda"​}

    ​// return the value, pass back up to the main stack frame.
    ​return​ u
}

func escapeToHeap() *user {
    u:= user{name: "Dinda"}

    // we do not return the value itself but the address of u, so value cannot be put inside the stack frame but put out on the heap
    return &u
}
```

From [the official FAQ](https://golang.org/doc/faq#stack_or_heap):
> Go compilers will allocate variables that are local to a function in that function’s stack frame. However, if the compiler cannot prove that the variable is not referenced after the function returns, then the compiler must allocate the variable on the garbage-collected heap to avoid dangling pointer errors. Also, if a local variable is very large, it might make more sense to store it on the heap rather than the stack.


### Thread

Thread is the flow of execution through the process code. There is two implements of the thread:
1. User level thread: user manages threads
2. Kernel level thread: operating system manages threads acting on the kernel, an operating system core.

In single-threaded processing, each thread belongs to exactly one core. But remember that OS can multitask with OS Scheduler (preemptive scheduler).

![single-threaded vs multi-threaded process](https://miro.medium.com/max/1400/1*U3WUG21SOB1XPVj3djkZjw.jpeg)

Today, CPU can handle program multithread with multicore, and most program nowadays is multithread application.

## Concurrency vs. Parallelism

- Concurrency is about dealing with lots of things at once, like in 1 core.
- Parallelism is about doing lots of things at once, like in multicore processing.


## Goroutine
Every time go program starts up, it looks to see how many cores are available, then it creates a logical process. Go using OS scheduler as the preemptive scheduler.

When dealing with Concurrency in Go, we use goroutine. While Process is a program in execution and threads is a paths of execution, goroutine is a path of execution.

So one thread can have multiple goroutines but remember only one goroutine can execute at a single time. To handle this, go come with Go’s Scheduler.

Go's Scheduler is a logical process that runs in user mode or calls cooperating scheduler. The goal is: less thread, much work.

If we have 3 Goroutines execute in one thread against one core, go's scheduler will decide what goroutine will run and what will be put to the queue. What scheduler is doing is managing the execution of these 3 Goroutines.

Concurrency is about managing a lot of things at once. We can see that is what the go scheduler is doing. It manages the execution of these 3 goroutines against one thread for one core.

So concurrent program is one that can be parallelized. If we want to do something in parallel, like have 2 goroutines running at the same time, then we need to create another core that has another thread. 
And we don't need to be concerned about what happened at the CPU level, whether something is running on multiple cores or not, go, and os will take care of it.

## Real-World Example
Main function is goroutine. In real-world, goroutine often used in:
- Application tasks: running web servers, db connection pools, daemons, API polling, data processing queues
- Request/event tasks: handling incoming HTTP requests, executing expensive subtasks (ex. multiple network calls) to fulfill a request, publishing a new message to message queue
- Fire & Forget tasks: logging, alerting, metrics

## Exercises:
- [WaitGroup](https://github.com/dindasigma/my-playground/blob/master/go/concurrency/exercises/waitgroup.go)
- [Mutexes](https://github.com/dindasigma/my-playground/blob/master/go/concurrency/exercises/mutexes.go)
- [Channel](https://github.com/dindasigma/my-playground/blob/master/go/concurrency/exercises/unbuffered_channel.go)
- [Buffered Channel](https://github.com/dindasigma/my-playground/blob/master/go/concurrency/exercises/buffered_channel.go)
- [Unbidirectional Channel](https://github.com/dindasigma/my-playground/blob/master/go/concurrency/exercises/unbidirectional_channel.go)
- [Closing Channel](https://github.com/dindasigma/my-playground/blob/master/go/concurrency/exercises/closing_channel.go)

## References
[How Operating Systems Work](https://medium.com/cracking-the-data-science-interview/how-operating-systems-work-10-concepts-you-should-know-as-a-developer-8d63bb38331f)

[The Ultimate Go Study Guide](https://github.com/ardanlabs/gotraining-studyguide)