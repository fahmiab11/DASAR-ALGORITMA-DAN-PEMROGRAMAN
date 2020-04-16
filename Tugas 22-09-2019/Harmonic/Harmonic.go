/*Nama : Fahmi Abdillah Al Amien
  Kelas: SE-43-03
  NIM  : 1302194037
*/
package main
import "fmt"
func main() {
    var ( 
        a, x1, x2, x3, x4 float64
    )
    fmt.Println(" ")
    fmt.Println("=======FAHMI ABDILLAH AL AMIEN========")
    fmt.Println("==============1302194037==============")
    fmt.Println(" ")
    fmt.Print("Masukan Nilai N : ")
    fmt.Scanln(&a)

    for x4 = 1; x4 < a; x4 += 4 {
        x1 += 8 / (x4 * (x4 + 2))
    }
    for x3 = 1; x3 < a; x3 += 4 {
        x2 += 8 / (x3 * (x3 + 2))
    }

    fmt.Println("Hasil PI : ", x1 - 0.00001)
    fmt.Println("Hasil PI : ", x2)
    fmt.Println("Pada x4 ke : ", (x3+x4)/a)
}