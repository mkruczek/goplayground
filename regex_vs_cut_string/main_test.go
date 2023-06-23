package regex_vs_cut_string

import (
	"bytes"
	"regexp"
	"testing"
)

var exceptionResponsePrefix = []byte("<Exception") //ex error at xml response

func isExceptionInResponse_CUT(body []byte) bool {
	comp := body[:len(exceptionResponsePrefix)]
	if bytes.Equal(exceptionResponsePrefix, comp) {
		return true
	}

	return false
}

var rgx = regexp.MustCompile("<Exception")

func isExceptionInResponse_REGEX(body []byte) bool {
	find := rgx.Find(body)
	return len(find) != 0
}

func Benchmark_CUT(b *testing.B) {

	exception := []byte("<Exception123456789012345678901234567890")
	noException := []byte("<NoException123456789012345678901234567890")

	for i := 0; i < b.N; i++ {
		isExceptionInResponse_CUT(exception)
		isExceptionInResponse_CUT(noException)
	}
}

// Benchmark_CUT-12        165617754                7.175 ns/op
// Benchmark_REGEX-12        6652236              179.3 ns/op //with precompiled regex
// Benchmark_REGEX-12         250579             4612 ns/op
func Benchmark_REGEX(b *testing.B) {

	exception := []byte("<Exception123456789012345678901234567890")
	noException := []byte("<NoException123456789012345678901234567890")

	for i := 0; i < b.N; i++ {
		isExceptionInResponse_REGEX(exception)
		isExceptionInResponse_REGEX(noException)
	}
}

func Test_CUT_REGEX(t *testing.T) {

	exception := []byte("<Exception123456789012345678901234567890")
	noException := []byte("<NoException123456789012345678901234567890")

	if !isExceptionInResponse_CUT(exception) {
		t.Errorf("CUT: exception not found")
	}

	if isExceptionInResponse_CUT(noException) {
		t.Errorf("CUT: exception found")
	}

	if !isExceptionInResponse_REGEX(exception) {
		t.Errorf("REGEX: exception not found")
	}

	if isExceptionInResponse_REGEX(noException) {
		t.Errorf("REGEX: exception found")
	}
}
