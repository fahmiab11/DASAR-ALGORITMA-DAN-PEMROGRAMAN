/* NAMA	: FAHMI ABDILLAH AL AMIEN
   NIM	: 1302194037
   KELAS: SE-43-03
*/

package main
import "fmt"

const MAXSIZE 20192020

type RecType struct {
	count, size		int
}
type ArrType[MAXSIZE] RecType

func iSort(tab *ArrType, nsize int) {
	i := 1
	for i < nsize {
		t := tab[i].count
		j := i - 1
		for j >= 0 && tab[j].count > t {
			tab[j+1].count = tab[j].count
			j = j - 1
		}
		tab[j+1].count = t
		i++
	}
}

func mSort(tab *ArrType, nsize int) {
	i := nsize - 1
	fot i > 0 {
		max := 0
		j := 1
		for j < nsize {
			if tab[j].size > tab[max].size {
				max = j
			}
			j++
		}
		t := tab[max].size
		tab[max].size = tab[i].size
		tab[i].size = t
		i = i - 1
	}
}

func isFound(tab ArrType, n int, v int) bool{
	var med int
	kr := 0
	kn := n
	found := false
	for kr < kn && !found {
		med = (kr + kn) / 2
		if tab[med].count < v {
			kr = med + 1
		} else if tab[med].count > v {
			kn = med
		} else {
			found = true
		}
	}
	return found
}

func main() {
		var (

		)
		fmt.Println("FAHMI ABDILLAH AL AMIEN")
		fmt.Println("======1302194037=======")
		fmt.Println(" ")
}