### What

Imagine a language that offers rapid development (Ruby, Python, PHP) and performance (C/C++). Go is that
language it bridges rapid development and performance.


### Why

* Fast - concurrency(use all cores), runs machine code
* Statically Typed - fewer runtime errors, compiler checks
* Safe - first-class error handling idioms, strong types
* Portable - compiles to single executable, easy to deploy
* Productive - tool + standard library + small = fun to write


### Programming Paradigms

* Imperative - Procedural - Series of steps e.g. _PHP_(multi), _Python_(multi), __Go__
* Declarative - Functional - Function is first-class citizen, immutable e.g. _Perl_(multi), _Erlang_(concurrent), _Haskell_
* Structured - Object Oriented - Model the world as objects e.g. _Ruby_(multi), _Java_(multi), _C#_(multi)


### Laws

* Gordon Moore - number of transistors doubles every 2yrs, “free lunch"
* Gene Amdal - parallelized system runs as fast as its critical part, “weakest link”


### Concurrency

* Concurrency - multiple processes time slices of a shared resource e.g. cpu
* Parallelism - multiple processes with own resource e.g. cpu

Hence Parallelism is a subset of Concurrency


### Concurrency Models

* Processor Threads - application has process with multiple threads which share resources communicate via locks
* Events - emitter, event object, receiver
* Callbacks and Promises - callback(return), promise(async)
* Communicating Sequential Processes - actor, message


### Demos

* Hello World - to describe program parts. use Go Playground
* Basic Server - show basic server and route handling
* Letters - show goroutines
* Relay - show goroutines with channels
* Curl - simple complete program using libraries


### Use cases

* Web Services
* APIs
* System Programming - A2A applications


### In the Wild

* Buffalo - new innovative web framework
* Docker - containers … containers … containers


### Resources

Excellent Website:- Documents(Little Go Book, Effective Go), Tutorials(Go Tour, Go Playground), Talks, Wiki


### References

* Go Getting Started (Pluralsight)
* On Track with Golang (Codeschool)
* Go in Action (Manning)
* Go Web Programming (Manning)
