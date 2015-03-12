package caller

import (
	"fmt"
	"runtime"
	"strings"
)

// At returns a string containing the function, file, and line number of a
// calling function. depth indicates the depth of the call stack to report; a
// depth of 1 reports the caller of At, while a depth of 2 indicates the
// caller of that function. The format is:
//
//   (recv).Func (pkg/folder/file.go:line)
//
// At will panic if this information cannot be determined.
func At(depth int) string {
	pc, filepath, line, ok := runtime.Caller(depth)
	if !ok {
		panic("function lookup failed")
	}
	// lookup full function listing
	pcname := runtime.FuncForPC(pc).Name()
	// trim it down to just the pkg and function name
	pkgfunc := strings.Split(pcname[strings.LastIndex(pcname, "/")+1:], ".")
	pkg, fnName := pkgfunc[0], pkgfunc[1]

	// get the folder/file from the filepath by trimming up to the first
	// occurance of pkg
	file := filepath[strings.LastIndex(filepath, "src")+3:]
	file = file[strings.Index(file, pkg):]

	return fmt.Sprintf("%s (%s:%d)", fnName, file, line)
}

// Get returns the formatted call string of the function that called the
// caller of Get. It is functionally equivalent to At(2).
func Get() string {
	return At(3) // Get -> Caller of Get -> Caller of caller
}

// Trace returns a listing of the callstack in the same format as At, to the
// given traversal depth. Note that unlike At, depth is not an absolute depth,
// but relative to Trace; Trace(1) will return the function that called Trace.
func Trace(depth int) []string {
	trace := make([]string, depth)
	for i := range trace {
		trace[i] = At(i + 2)
	}
	return trace
}