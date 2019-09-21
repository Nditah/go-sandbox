package brackets

import "github.com/golang-collections/collections/stack"

func IsExpressionValid(s string) bool {
	brackets := stack.New()
	bracketsMapping := map[byte]byte{'(': ')', '[': ']', '{': '}'}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			expectedBracket := bracketsMapping[s[i]]
			brackets.Push(expectedBracket)
		}
		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if brackets.Len() == 0 || brackets.Peek() != s[i] {
				return false
			}
			brackets.Pop()
		}
	}
	return brackets.Len() == 0
}
