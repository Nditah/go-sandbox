package brackets_test

import (
	"github.com/stretchr/testify/assert"
	"go-sandbox/brackets"
	"testing"
)

func TestIsExpressionValid(t *testing.T) {
	assert.Equal(t, brackets.IsExpressionValid("(4+2*3)+(4-(1-3))+[5+{-3}-(4-6)]"), true, "right expression")
	assert.Equal(t, brackets.IsExpressionValid("((1+2)"), false, "extra opened bracket")
	assert.Equal(t, brackets.IsExpressionValid("(1+2))"), false, "extra closed bracket")
	assert.Equal(t, brackets.IsExpressionValid("(1+2]"), false, "incompatible bracket types")
	assert.Equal(t, brackets.IsExpressionValid("(1+[2)]"), false, "wrong brackets order")
}
