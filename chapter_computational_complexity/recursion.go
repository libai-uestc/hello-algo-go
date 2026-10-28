package chaptercomputationalcomplexity

import (
	"container/list"
	"fmt"
)

func recur(n int) int {
	if n == 1 {
		return 1
	}
	res := recur(n - 1)
	return n + res
}

func forLoopRecur(n int) int {
	stack := list.New()
	res := 0
	for i := n; i > 0; i-- {
		stack.PushBack(i)
		fmt.Println("当前stack里的元素是：", listToSlice(stack))
	}
	for stack.Len() != 0 {
		res += stack.Back().Value.(int)
		stack.Remove(stack.Back())
		fmt.Println("当前stack里的元素是：", listToSlice(stack))
	}
	return res
}

func listToSlice(l *list.List) []any {
	s := make([]any, 0, l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		s = append(s, e.Value)
	}
	return s
}

func tailRecur(n int, res int) int {
	if n == 0 {
		return res
	}
	return tailRecur(n-1, res+n)
}

func fib(n int) int {
	if n == 1 || n == 2 {
		return n - 1
	}
	res := fib(n-1) + fib(n-2)
	return res
}
