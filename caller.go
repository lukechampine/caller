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
	fnName := runtime.FuncForPC(pc).Name()
	// trim it down to just the function name
	fnName = fnName[strings.Index(fnName, "/")+1:]
	fnName = fnName[strings.Index(fnName, "/")+1:]
	fnName = fnName[strings.Index(fnName, ".")+1:]

	// get folder/file by trimming the appropriate prefix
	var file string
	if strings.HasPrefix(filepath, runtime.GOROOT()) {
		// stdlib
		file = strings.TrimPrefix(filepath, runtime.GOROOT()+"/src/pkg/")
	} else {
		// trim the host + username
		file = filepath[strings.LastIndex(filepath, "src/")+4:]
		file = file[strings.Index(file, "/")+1:]
		file = file[strings.Index(file, "/")+1:]
	}

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
