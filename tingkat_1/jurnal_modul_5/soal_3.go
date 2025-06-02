package main

import "fmt"

func main3() {
    var jumlahHari, jarak, totalJarak int

    fmt.Print("Masukkan jumlah hari pelari berlatih: ")
    fmt.Scan(&jumlahHari)

    totalJarak = 0
    fmt.Println("Masukkan jarak yang ditempuh setiap hari (dalam kilometer):")
    for i := 0; i < jumlahHari; i++ {
        fmt.Scan(&jarak)
        totalJarak += jarak
    }

    fmt.Printf("Total jarak yang ditempuh: %d kilometer\n", totalJarak)
}