package main

import "fmt"

// Generic Type
type List[T any] struct {
	arr *[]T
	val T
}

// Here in Index function there is a type parameter of type T that has a constraint of comparable
// so it can take any comparable type like int, float, string , struct , etc
func Index[T comparable](arr []T, item T) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			return i
		}
	}
	return -1
}

func main() {
	//using string
	s := []string{"bro", "what's", "up"}
	fmt.Println(Index[string](s, "up"))

	l := List[string]{&s, "bro"}
	fmt.Println(Index[string](*l.arr, l.val))

	//using int
	n := []int{1, 34, 5, 334, 32}
	fmt.Println(Index[int](n, 32))

	i := List[int]{&n, 5}
	fmt.Println(Index[int](*i.arr, i.val))

	//not in
	fmt.Println(Index[int](n, 0))

}
