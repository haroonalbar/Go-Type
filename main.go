package main

import "fmt"

// Here in Index function there is a type parameter of type T that has a constraint of comparable 
// so it can take any comparable type like int, float, string , struct , etc
func Index[T comparable](arr []T,item T) int{
  for i:= 0 ; i< len(arr); i++{
    if arr[i] == item {
      return i
    }
  }
  return -1
}

func main(){
  //using string
  s:=[]string{"bro","what's","up"}
  fmt.Println(Index[string](s,"up"))

  //using int
  n:=[]int{1,34,5,334,32}
  fmt.Println(Index[int](n,32))

  //not in
  fmt.Println(Index[int](n,0))

}
