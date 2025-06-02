package main

import "fmt"

func main1() {
    var hasil, pangkat, bilangan, i int

    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&bilangan)
    fmt.Print("Masukkan pangkat: ")
    fmt.Scan(&pangkat)

    hasil = 1
    for i = 0; i < pangkat; i++ {
        hasil *= bilangan
    }

    fmt.Printf("Hasil dari %d pangkat %d adalah %d\n", bilangan, pangkat, hasil)
}

