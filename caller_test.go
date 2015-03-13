package caller

import (
	"strings"
	"testing"
)

func TestMyName(t *testing.T) {
	at := At(1)
	atExpected := "TestMyName (caller/caller_test.go:9)"
	if at != atExpected {
		t.Errorf("At produced bad result: expected %s, got %s", atExpected, at)
	}
	got := Get()
	gotExpected := "tRunner (testing/testing.go:422)" // TODO: brittle
	if got != gotExpected {
		t.Errorf("Get produced bad result: expected %s, got %s", gotExpected, got)
	}
	trace := strings.Join(Trace(2), ", ")
	traceExpected := "TestMyName (caller/caller_test.go:19), tRunner (testing/testing.go:422)"
	if trace != traceExpected {
		t.Errorf("Get produced bad result: expected %s, got %s", traceExpected, trace)
	}
	var lambda string
	func() {
		lambda = At(1)
	}()
	lambdaExpected := "func·001 (caller/caller_test.go:26)"
	if lambda != lambdaExpected {
		t.Errorf("At produced bad result: expected %s, got %s", lambdaExpected, lambda)
	}
}
