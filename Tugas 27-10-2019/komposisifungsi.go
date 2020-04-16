/*Nama : Fahmi Abdillah Al Amien
  Kelas: SE-43-03
  NIM  : 1302194037
*/
package main
import "fmt"
func main() {
    var ( 
		x         float64
	)
    fmt.Println(" ")
    fmt.Println("=======FAHMI ABDILLAH AL AMIEN========")
    fmt.Println("==============1302194037==============")
    fmt.Println(" ")
    fmt.Print("Masukan Nilai x = ")
    fmt.Scanln(&x)
    fmt.Println("f(", x, ") =", f(x))
    fmt.Println("g(", x, ") =", g(x))
    fmt.Println("h(", x, ") =", h(x))
    fmt.Println("(fogoh)", "(", x, ") =", fogoh(x))
}

func f(x float64) float64 {
  return x * x
}

func g(x float64) float64 {
  return x - 2
}

func h(x float64) float64 {
  return x + 1
}

func fogoh(x float64) float64 {
  return (x - 1) * (x - 1)
}