package fibonacci

//import "fmt"

//fmt.Println(fibonacci(4))

//}

func Fibonacci (x uint) uint {
if x==0 {
  return 0
  }
if x==1  {
  return 1
} else  {
  return Fibonacci(x-1) + Fibonacci(x-2)
}
}
