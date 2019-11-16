package brackets_test

import (
	"github.com/andrei-punko/go-sandbox/brackets"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsExpressionValid(t *testing.T) {
	assert.Equal(t, brackets.IsExpressionValid("(4+2*3)+(4-(1-3))+[5+{-3}-(4-6)]"), true, "right expression")
	assert.Equal(t, brackets.IsExpressionValid("((1+2)"), false, "extra opened bracket")
	assert.Equal(t, brackets.IsExpressionValid("(1+2))"), false, "extra closed bracket")
	assert.Equal(t, brackets.IsExpressionValid("(1+2]"), false, "incompatible bracket types")
	assert.Equal(t, brackets.IsExpressionValid("(1+[2)]"), false, "wrong brackets order")
}

func BenchmarkIsExpressionValidSimple(b *testing.B) {
	benchmarkIsExpressionValid("(1)", b)
}

func BenchmarkIsExpressionValidComplex(b *testing.B) {
	benchmarkIsExpressionValid("(4+2*3)+({4-(1-3)})+[5+{-3}-([4]-6)]", b)
}

var result bool

// Benchmark with recommendations according to https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
func benchmarkIsExpressionValid(expression string, b *testing.B) {
	var r bool
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = brackets.IsExpressionValid(expression)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}
