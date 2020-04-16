/*Nama : Fahmi Abdillah Al Amien
  Kelas: SE-43-03
  NIM  : 1302194037
*/
package main
import (
        "fmt"
        "math"
)

type Point struct {
                  x, y    float64
}

const N int = 10000

var Points [N]Point
var numpoints int

func jarak (p1 Point, p2 Point) float64 {
    return math.Sqrt(((p1.x - p2.x) * (p1.x - p2.x) + (p1.y - p2.y) * (p1.y - p2.y)))
}

func bacaTitik() {
            var (
                p Point
            )
            numpoints = 0
            fmt.Scanln(&p.x, &p.y)
            for (numpoints < N) && !(p.x == 0 && p.y == 0) {
                Points[numpoints] = p
                numpoints++
                fmt.Scanln(&p.x, &p.y)
            }
}

func ambiltitikterdekat(p1 *Point, p2 *Point) {
    *p1 = Points[0]
    *p2 = Points[1]
    for i := 0; i < numpoints; i++ {
        for j := i + 1; j < numpoints; j++ {
            if jarak(Points[i], Points[j]) < jarak(*p1, *p2) {
              *p1 = Points[i]
              *p2 = Points[j]
            }
        }
    }
}

func main() {
    var ( 
		p1, p2   Point
	)
    fmt.Println(" ")
    fmt.Println("=======FAHMI ABDILLAH AL AMIEN========")
    fmt.Println("==============1302194037==============")
    fmt.Println(" ")
    bacaTitik()
    ambiltitikterdekat(&p1, &p2)
    fmt.Printf("Titik Terdekat adalah (%.1f,%.1f) dan (%.1f,%.1f) dengan jarak %.1f.\n",p1.x, p2.y, p2.x, p2.y, jarak(p1, p2))
}