/*Nama  : Fahmi Abdillah Al Amien
  Kelas : SE-43-03
  NIM   : 1302194037
*/

package main
import "fmt"

type RecType struct {
      f1  string
      f2  int
      f3  float64
}

const n = 2019

type ArrType [n]RecType

var data ArrType

func rmax(data *ArrType) float64 {
    var (
        max float64
    )

    for i := 0; i < n; i++ {
        data[i].f1 = string('A' - 1 + i)
        data[i].f2 = i + 2
        data[i].f3 = float64(i) + 2.5 + 10.0 - 5.4
    }

    max = data[0].f3
    for i := 0; i < n; i++ {
        if data[i].f3 > max {
            max = data[i].f3
        }
    }

    fmt.Println(data)
    fmt.Println("Nilai Max :",max)

    return max
}

func imin(data *ArrType) int {
    min := 9999
    for i := 0; i < n; i++ {
        if data[i].f2 < min {
            min = data[i].f2
        }
    }
    fmt.Println("Nilai Min :",min)
    return min
}

func found(data *ArrType, key string) bool {
    var found bool
    for i := 0; i < n; i++ {
        if data[i].f1 == key {
            fmt.Println(" ")
            fmt.Println("Hasil Pencarian")
            fmt.Println("=-=-=-=-=-=-=-=")
            fmt.Println(data[i])
            found = true
        }
    }
    return found
}

func pos(data *ArrType, key string) int {
    low := 0
    high := len(data) - 1

    for low <= high {
        median := (low + high) / 2

        if data[median].f1 < key {
            low = median + 1
        } else if data[median].f1 > key {
            high = median - 1
        } else {
            return median
        }
    }

    return -1

}

func main() {
    var (
        data ArrType
        key   string
    )

/*  inputData(&data) */
    rmax(&data)
    imin(&data)
    found(&data,key)
    pos(&data,key)
    fmt.Println("FAHMI ABDILLAH AL AMIEN")
    fmt.Println("-------1302194037------")
    fmt.Println(" ")
    fmt.Print("Masukkan key String : ")
    fmt.Scanln(&key)
    fmt.Println("Key Ada didalam Array ? :", found(&data, key))
	fmt.Println("POS : ", pos(&data, key))
	fmt.Println("=-=-=-=-=-=-=-=")
    
}