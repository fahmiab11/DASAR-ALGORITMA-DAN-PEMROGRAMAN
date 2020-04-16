/*Nama : Fahmi Abdillah Al Amien
  Kelas: SE-43-03
  NIM  : 1302194037
*/
package main
import "fmt"
func main() {
    var ( 
		jumlah, total, rumus float64
		penyebut, a 		 int
	)
	
    fmt.Println(" ")
    fmt.Println("=======FAHMI ABDILLAH AL AMIEN========")
    fmt.Println("==============1302194037==============")
	fmt.Println(" ")
	fmt.Println("N Suku Pertama : ")
	fmt.Scan(&a)

	penyebut = -1
	rumus = -1
	jumlah = 0
	total = 0

	for i := 1; i <= a;i++ {
		penyebut = penyebut + 2
		rumus = rumus * (-1)
		jumlah = rumus / float64 (penyebut)
		total = total + jumlah
	}
	fmt.Printf("Hasil PI : %.7f", total*4)
}
