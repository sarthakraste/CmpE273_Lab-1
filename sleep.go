package main
import "time"
import "fmt"
func main() {
  c1 := make(chan string)
  c2 := make(chan string)




 go func() {
   for {
     c1 <- "Sleep 1"
     time.Sleep(time.Second *1)
   }
 } ()

 go func() {
   for {
     c2 <- "Sleep 2"
     time.Sleep(time.Second *2)
   }
 } ()



 go func() {
   for {
     select {
     case msg1 := <- c1 :
       fmt.Println(msg1)
     case msg2 := <- c2 :
         fmt.Println(msg2)
      case <- time.After(time.Second) :
  fmt.Println("Timeout")

       }
}
     }()

     var input string
     fmt.Scanln(&input)
   }
