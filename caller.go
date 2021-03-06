package caller

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

var sep = string(filepath.Separator)

// determine gopath by reading the stack trace of this file
var gopath = func() string {
	_, path, _, ok := runtime.Caller(1)
	if !ok {
		panic("gopath lookup failed")
	}
	i := strings.Index(path, "github.com/lukechampine/caller")
	if i == -1 {
		panic("sentinel function moved: " + path)
	}
	return path[:i]
}()

// determine goroot by reading the stack trace of a stdlib function that takes
// a function
var goroot = func() string {
	var path string
	var ok bool
	strings.TrimFunc("foo", func(rune) bool {
		_, path, _, ok = runtime.Caller(1)
		return true
	})
	if !ok {
		panic("goroot lookup failed")
	} else if !strings.HasSuffix(path, "strings/strings.go") {
		panic("sentinel function moved: " + path)
	}
	return strings.TrimSuffix(path, "strings/strings.go")
}()

// At returns a string containing the function, file, and line number of a
// calling function. depth indicates the depth of the call stack to report; a
// depth of 1 reports the caller of At, while a depth of 2 indicates the
// caller of that function. The format is:
//
//   (recv).Func (pkg/folder/file.go:line)
//
// At will panic if this information cannot be determined.
func At(depth int) string {
	pc, path, line, ok := runtime.Caller(depth)
	if !ok {
		panic("function lookup failed")
	}
	// lookup and trim the function name
	fnName := filepath.Base(runtime.FuncForPC(pc).Name())
	fnName = fnName[strings.Index(fnName, ".")+1:]

	// get folder/file by trimming the appropriate prefix
	var file string
	file = path
	// if strings.HasPrefix(path, goroot) {
	// 	// stdlib: trim $GOROOT/src/
	// 	file = strings.TrimPrefix(path, goroot)
	// 	file = strings.SplitN(file, sep, 1)[0]
	// } else if strings.HasPrefix(path, gopath) {
	// 	// standard: trim $GOPATH/host/username/
	// 	file = strings.TrimPrefix(path, gopath)
	// 	file = strings.SplitN(file, sep, 3)[2]
	// }

	return fmt.Sprintf("%s (%s:%d)", fnName, file, line)
}

// Me returns a formatted call string describing of the invocation of Me. It
// is functionally equivalent to At(1).
func Me() string {
	return At(2) // Me -> Caller of Me
}

// Get returns a formatted call string of the invocation of the caller of Get.
// It is functionally equivalent to At(2).
func Get() string {
	return At(3) // Get -> Caller of Get -> Caller of caller
}

// Trace returns a listing of the callstack in the same format as At, to the
// given traversal depth. Note that unlike At, depth is not absolute, but
// relative to Trace; Trace(1)[0] will describe the invocation of Trace.
func Trace(depth int) []string {
	trace := make([]string, depth)
	for i := range trace {
		trace[i] = At(i + 2)
	}
	return trace
}

// Print prints the result of calling Trace(depth), formatted for readability.
func Print(depth int) {
	t := Trace(depth)
	t = t[1:] // remove Print call
	fmt.Println(strings.Join(t, "\n\t"))
}
