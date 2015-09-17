package fibonacci
import "testing"
func Testfibonacci(t *testing.T) {

var v uint
v = Fibonacci(4)
if v!=3 {
t.Error("Expected 6765, got",v)
}
}
