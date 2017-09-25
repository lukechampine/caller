package caller

import (
	"strings"
	"testing"
)

type foo struct{}

func (f *foo) bar() string {
	return Me()
}

func TestCaller(t *testing.T) {
	at := At(1)
	atExpected := "TestCaller (caller/caller_test.go:15)"
	if at != atExpected {
		t.Errorf("At produced bad result: expected %s, got %s", atExpected, at)
	}

	got := Get()
	gotExpected := "tRunner (testing/testing.go:746)" // TODO: brittle
	if got != gotExpected {
		t.Errorf("Get produced bad result: expected %s, got %s", gotExpected, got)
	}

	trace := strings.Join(Trace(2), ", ")
	traceExpected := "TestCaller (caller/caller_test.go:27), tRunner (testing/testing.go:746)"
	if trace != traceExpected {
		t.Errorf("Get produced bad result: expected %s, got %s", traceExpected, trace)
	}

	var lambda string
	func() {
		lambda = Me()
	}()
	lambdaExpected := "TestCaller.func1 (caller/caller_test.go:35)"
	if lambda != lambdaExpected {
		t.Errorf("At produced bad result: expected %s, got %s", lambdaExpected, lambda)
	}

	f := new(foo)
	recv := f.bar()
	recvExpected := "(*foo).bar (caller/caller_test.go:11)"
	if recv != recvExpected {
		t.Errorf("At produced bad result: expected %s, got %s", recvExpected, recv)
	}
}
