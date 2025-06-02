package main

import "fmt"

func main3() {
    var n int
    fmt.Print("Masukkan jumlah baris segitiga: ")
    fmt.Scan(&n)

    segitigaBintang(n)
}

func segitigaBintang(n int) {
    for i := 1; i <= n; i++ {
        row := ""
        for j := 0; j < i; j++ {
            row += "*"
        }
        fmt.Println(row)
    }
}
