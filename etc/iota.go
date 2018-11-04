package main

import "fmt"

const (
	KB uint64 = iota + 1
	MB
)

type Point struct {
	X, Y int
}

func main() {
	var s1 []int

	var s2 = []int{
		1,
		2,
		3,
	}

	s1 = append(s1, 100)
	s1 = append(s1, 200)
	s1 = append(s1, 300)
	s1 = append(s1, 400)
	s1 = append(s1, 700)
	fmt.Println(s2)
	fmt.Printf("%v, Cap: %v, Len: %v", s1, cap(s1), len(s1))

	s2 = append(s2, s1...)
	fmt.Println(s2)

	var s12 [][]int

	s12 = append(s12, s2)
	fmt.Println(s12)

	slice3 := make([]int, 10)
	fmt.Println(slice3)

	a := []int{1, 2, 3}
	fmt.Println(a)
	s18 := a[:]
	fmt.Println(s18)

	fmt.Println(Signum(-1))

	x := 1

	switch x {
	case 0:
		fmt.Println(x)
	case 1:
		fmt.Println(x)
	default:
		fmt.Println(x)
	}
}

func Signum(x int) string {
	switch {
	case x > 0:
		return "> 0"
	default:
		return "is 0"
	case x < 0:
		return "< 0"
	}
}
