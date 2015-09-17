package main
import "fmt"
import "math"

type shape interface {
area() float64
perimeter() float64
}

type Circle struct {
x,y,r float64
  }

  type Rectangle struct {
    x1,y1,x2,y2 float64
  }
func (r *Rectangle) recarea() float64 {
l := r.x2
w := r.y2
return l * w
}

func (c *Circle) circarea() float64 {
return math.Pi * c.r * c.r

}

func (r *Rectangle) recperi() float64 {
l := r.x2
w := r.y2
return 2 *(l + w)
}

func (c *Circle) cirperi() float64 {
return math.Pi * 2 * c.r

}

func main() {
c := Circle{0, 0, 3}
r := Rectangle{0,0,10,30}
fmt.Println("Area of Circle is :" ,c.circarea())
fmt.Println("Area of Rectangle is :" ,r.recarea())
fmt.Println("Perimeter of Circle is :" ,c.cirperi())
fmt.Println("Perimeter of Rectangle is :" ,r.recperi())
}
