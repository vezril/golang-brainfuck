package main

import (
	"fmt"
	"github.com/vezril/go-cookiejar/collections/stack"
)

func addToStack(s *stack.Stack, m int, n int) {
    for i := m; i < n; i++ {
        s.Push(i)
    }
}

func main() {

    s := stack.New()

    for i := 0; i < 5; i++ {
        go addToStack(s, 0+i, 5*i)
    }

    fmt.Println(s)

    // Pop out the stack contents and display them
    for !s.Empty() {
        fmt.Println(s.Pop())
    }
}
