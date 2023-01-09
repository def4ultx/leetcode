package main

func main() {
	isValid("()")
}

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				return false
			}
			item := stack[len(stack)-1]
			if !check(item, ch) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func check(a, b rune) bool {
	switch a {
	case '(':
		return b == ')'
	case '{':
		return b == '}'
	case '[':
		return b == ']'
	}
	return false
}
