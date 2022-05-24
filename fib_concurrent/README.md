          Concurrency != Parallelism

Concurrency means that the application has more than one thing to do at the same time. It's about creating multiple processes executing independently.  
Let's assume that we have a single-core system, multiple tasks need to be accomplished, but we have a constraint where, at any moment in time, only one task can be executed on the single-core available…  
In a concurrent execution model, there is context switching between the tasks. So the application is dealing with multiple tasks but not executing them together as we have only one executing core.<br />
Since there is only one core, we can't achieve parallel execution.
REMEMBER: At any moment a single physical core can only execute one process at a time. What concurrency does is it takes two problems or process and solve them one at a time but constantly(not constant) switching between them. eg: for 1ms it is running one problem, then for next 0.5ms it is running problem2, then again it swtiches to problem1
When we run an application, like the browser I am using to write this post, a process is created by the operating system for the application. The job of the process is to act like a container for all the resources the application uses and maintains as it runs. These resources include things like a memory address space, handles to files, devices and threads.
A thread is a path of execution that is scheduled by the operating system to execute the code we write in our functions against a processor. The main thread can then in turn launch more threads and those threads can launch even more threads.
A process starts out with a thread (main thread), and when the thread terminates the process too terminates. Because this main thread is the origin of application. 
when main thread creates branch of thread to run, the OS schedules thread to run available processor. It depends on type of Operating System. As the algorithm may change in each version it is dangerous to write program codes based on OS thread for concurrency.
when we use go routine, go runtime creates different branches of thread. but when main thread finishes execution then application terminates no matter how many go rountines are running. ( so we use sync.WaitGroup or channels to delay the finish line of main thread)
Goroutines are considered to be lightweight because they use little memory and resources plus their initial stack size is small. Prior to version 1.2 the stack size started at 4K and now as of version 1.4 it starts at 8K. The stack has the ability to grow as needed.
By default, the Go runtime allocates a single logical processor to execute all the goroutines that are created for our program.
Even with this single logical processor and operating system thread, hundreds of thousands of goroutines can be scheduled to run concurrently with amazing efficiency and performance. It is not recommended to add more that one logical processor
Go provides the ability to add more via the GOMAXPROCS environment variable or runtime function.
in "runtime" package, NumCPU() function returns number of available cores in the CPU.
using "runtime.GOMAXPROCS(n) func we can tell go runtime to use more than one core ( n - no of cores)
Thats it!. Now go runtime will automatically split all the go routines between these cores
However, to have true parallelism you need to run your program on a machine with multiple physical processors. If not, then the goroutines will be running concurrently against a single physical processor, even though the Go runtime is using multiple logical processors.


