/*Nama : Fahmi Abdillah Al Amien
  Kelas: SE-43-03
  NIM  : 1302194037
*/
package main
import "fmt"
func main() {
    var ( 
        GanGen      int
        penyebut    float64
        s1, s2      float64
		sel1, sel2  float64
	)
    penyebut = 1
    GanGen = 1
    sel1 = 1
    sel2 = 1

    fmt.Println(" ")
    fmt.Println("=======FAHMI ABDILLAH AL AMIEN========")
    fmt.Println("==============1302194037==============")
    fmt.Println(" ")

    for sel1*4 >= 0.00001 && sel2*4 >= 0.00001 {
        s1 = s2
        if GanGen%2 == 0 {
        	s2 -= (1 / penyebut)
        } else {
        	s2 += (1 / penyebut)
        }
        sel1 = s1 - s2
        sel2 = s2 - s1
        if sel1 < 0 {
        sel1 *= -1
        }
        if sel2 < 0 {
        	sel2 *= -1
        }
        penyebut += 2
        GanGen++
    }
    fmt.Println("Hasil PI : ", s1*4)
    fmt.Println("Hasil PI : ", s2*4)
    fmt.Println("Pada i ke : ", GanGen)
}