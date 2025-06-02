package main

import "fmt"

func main2() {
    var jumlahBilangan, total, bilangan int

    fmt.Print("Masukkan jumlah bilangan dalam barisan: ")
    fmt.Scan(&jumlahBilangan)

    total = 0
    fmt.Println("Masukkan bilangan satu per satu:")
    for i := 0; i < jumlahBilangan; i++ {
        fmt.Scan(&bilangan)
        total += bilangan
    }

    fmt.Printf("Jumlah total dari bilangan dalam barisan: %d\n", total)
}