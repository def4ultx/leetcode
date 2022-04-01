package main

import (
	"fmt"
	"strconv"
)

func main() {
	// calculate("3+2*2")
	// calculate(" 3/2 ")
	// calculate(" 3+5 / 2")
	// calculate("22-3*5")
	calculate("1-1+1-11+12")
	// calculate("0-22222")
}

type Operator string

type Value struct {
	value int
}

func (v Value) evaluate() int {
	return v.value
}

type Node interface {
	evaluate() int
}

type OperatorNode struct {
	left, right Node
	op          Operator
}

func (o OperatorNode) evaluate() int {
	switch o.op {
	case "+":
		// fmt.Println("+++")
		// fmt.Println(o.left.evaluate())
		// fmt.Println(o.right.evaluate())
		return o.left.evaluate() + o.right.evaluate()
	case "-":
		return o.left.evaluate() - o.right.evaluate()
	case "*":
		// fmt.Println("***")
		// fmt.Println(o.left.evaluate())
		// fmt.Println(o.right.evaluate())
		// fmt.Println("fin ***")
		return o.left.evaluate() * o.right.evaluate()
	case "/":
		return o.left.evaluate() / o.right.evaluate()
	}
	return 0
}

func calculate(s string) int {
	var root *OperatorNode
	var head *OperatorNode

	var current string
	for _, r := range s {
		// fmt.Println(root)
		if r == ' ' {
			continue
		} else if r >= '0' && r <= '9' {
			current += string(r)
		} else if r == '+' || r == '*' {

			// fmt.Println(current, string(r))

			var left int
			if head != nil {
				if head.right != nil {
					left = head.right.evaluate()
				} else {
					left, _ = strconv.Atoi(current)
				}
				current = ""
			} else {
				left, _ = strconv.Atoi(current)
				current = ""
			}
			op := &OperatorNode{
				left:  &Value{left},
				right: nil,
				op:    Operator(r),
			}
			if root == nil {
				root, head = op, op
			} else {
				head.right = op
				head = op
			}
		} else if r == '-' || r == '/' {
			var number int
			number, _ = strconv.Atoi(current)
			var x Node
			if root == nil {
				x = &Value{number}
			} else {
				x = root
			}
			op := &OperatorNode{
				left:  x,
				right: nil,
				op:    Operator(r),
			}
			root, head = op, op
			current = ""
		}
	}
	// add current to right node
	if head != nil {
		number, _ := strconv.Atoi(current)
		head.right = &Value{number}
	}
	fmt.Println(root)
	fmt.Println(root.left.evaluate())
	fmt.Println(root.right.evaluate())

	fmt.Println(head)
	fmt.Println(head.left.evaluate())
	fmt.Println(head.right.evaluate())

	fmt.Println(root.evaluate())

	if root == nil && current == "" {
		return 0
	}
	if root == nil && current != "" {
		number, _ := strconv.Atoi(current)
		return number
	}
	return 0
}

func isOperator(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/'
}
