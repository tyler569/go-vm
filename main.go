/*
 * Copyright 2015, Tyler Philbrick
 * See COPYING for license information
 */

package main

import (
	_ "errors"
	"fmt"
)

type Stack []int

func (s *Stack) Put(elem int) {
	(*s) = append((*s), elem)
}
func (s *Stack) Pop() int {
	d := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return d
}
func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}
func (s *Stack) Swap(elem int) {
	(*s)[len(*s)-1] = elem
}

const (
	HALT = iota

	POP_PRINT

	UNARY_NEGATE

	BINARY_ADD
	BINARY_MULT

	PUSH
)

type Code []int

var code Code = Code {
	PUSH, 5,
	PUSH, 10,
	PUSH, 2,
	UNARY_NEGATE,
	BINARY_MULT,
	BINARY_ADD,
	UNARY_NEGATE,
	POP_PRINT,
	HALT,
}

func (c Code) Run(debug bool) {
	stack := make(Stack, 0)

	for i := 0; code[i] != HALT; i++ {
		switch code[i] {
		case POP_PRINT:
			fmt.Println(stack.Pop())
		case UNARY_NEGATE:
			stack.Put(-stack.Pop())
		case BINARY_ADD:
			stack.Swap(stack.Pop() + stack.Peek())
		case BINARY_MULT:
			stack.Swap(stack.Pop() * stack.Peek())
		case PUSH:
			i++
			stack.Put(code[i])
		}
		if debug {
			fmt.Println(stack)
		}
	}
}

func main() {
	code.Run(true)
}
