package main

import (
	"fmt"
	"io.analogalex.collections/collections/set"
)

func main() {
	fmt.Println("Hello World!")

	var s set.OrderedSet = set.MakeBinaryTreeSet()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(2)
	s.Add(4)

	fmt.Println(s.Formatted())
	fmt.Println(s.Contains(2))
	fmt.Println(s.Contains(5))
	fmt.Println(s.Size())

	s.Remove(5)

	fmt.Println(s.Formatted())
	fmt.Println(s.Contains(2))
	fmt.Println(s.Size())

	fmt.Println(s.First())
	fmt.Println(s.Last())

	var h set.Set = set.MakeHashSet()
	h.Add(1)
	h.Add(2)
	h.Add(3)
	h.Add(2)

	fmt.Println(h.Formatted())
	fmt.Println(h.Contains(2))
	fmt.Println(h.Contains(5))
	fmt.Println(h.Size())

	h.Remove(5)

	fmt.Println(h.Formatted())
	fmt.Println(h.Contains(2))
	fmt.Println(h.Size())

}
