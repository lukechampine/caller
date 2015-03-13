# caller #

caller is a simple utility package that provides nicely-formatting strings describing the callstack of a function invocation. Functionally, it serves the same purpose as `runtime.Caller`, but its output is easier to work with (one compact string, instead of four different types). Another convenience is that caller reports the name of the invoking function directly, instead of forcing you to look the name up manually using `runtime.FuncForPC`.

## Example usage ##

```go
// $GOPATH/src/github.com/username/example/example.go
package main

import "github.com/lukechampine/caller"

func foo() {
	for _, call := range caller.Trace(2) {
		println("\t" + call)
	}
}

func main() {
	println(caller.Get())
	println(caller.At(1))
	foo()
}
```

Output:
```
main (runtime/proc.c:247)
main (example/example.go:14)
	foo (example/example.go:7)
	main (example/example.go:15)
```
