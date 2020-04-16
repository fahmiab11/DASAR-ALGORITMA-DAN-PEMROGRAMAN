/*Nama : Fahmi Abdillah Al Amien
  Kelas: SE-43-03
  NIM  : 1302194037
*/
package main

import "fmt"

func main() {
    var (
        nam     float64
        nmk     string
	)
	fmt.Println("Fahmi Abdillah Al Amien")
	fmt.Println("------1302194037-------")
	fmt.Println(" ")
	fmt.Print("Nilai Akhir Mata Kuliah: ")
	fmt.Scanln(&nam)
	if nam > 80 && nam <= 100 {
		nmk = "A"
	} else if nam > 72.5 && nam <= 80 {
		nmk = "AB"
	} else if nam > 65 && nam <= 72.5 {
		nmk = "B"
	} else if nam > 57.5 && nam <= 65 {
		nmk = "BC"
	} else if nam > 50 && nam <= 57.5 {
		nmk = "C"
	} else if nam > 40 && nam <= 50 {
		nmk = "D"
	} else if nam > 0 && nam <= 40 {
		nmk = "E"
	} else {
		nmk = "Tidak Terdefinisi"
	}
	fmt.Println("Nilai Mata Kuliah: ", nmk)
	
}