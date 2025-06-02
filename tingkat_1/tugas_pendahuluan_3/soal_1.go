package main

import (
    "fmt"
)

func main1() {
    var waktuDetik int
    fmt.Print("Masukkan waktu dalam satuan detik: ")
    fmt.Scan(&waktuDetik)

    jam := waktuDetik / 4500
    sisaDetik := waktuDetik % 4500
    menit := sisaDetik / 75
    detik := sisaDetik % 75

    fmt.Printf("Maka hasil konversinya adalah %d jam, %d menit dan %d detik di Mars\n", jam, menit, detik)
}
