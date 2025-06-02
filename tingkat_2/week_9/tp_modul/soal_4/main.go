package main

import "fmt"

const MAXNUM int = 10

type tabInt [MAXNUM]int

func main() {
    var A tabInt
    var n int

    baca(&A, &n)
    cetak(A, n)
    fmt.Println("Jumlah:", jumlah(A, n))
    fmt.Println("Rata-rata:", rata_rata(A, n))
}

func baca(A *tabInt, n *int) {
    fmt.Print("Masukkan banyak elemen (n): ")
    fmt.Scan(n)

    if *n > MAXNUM {
        *n = MAXNUM
    }

    var i int
    for i = 0; i < *n; i++ {
        fmt.Printf("Masukkan nilai ke-%d: ", i+1)
        fmt.Scan(&(*A)[i])
    }
}

func cetak(A tabInt, n int) {
    if n == 0 {
        fmt.Println("Array kosong.")
        return
    }

    fmt.Println("Nilai yang terdapat pada array: ")
    for i := 0; i < n; i++ {
        fmt.Print(A[i], " ")
    }
    fmt.Println()
}

func jumlah(A tabInt, n int) int {
    total := 0
    for i := 0; i < n; i++ {
        total += A[i]
    }
    return total
}

func rata_rata(A tabInt, n int) float64 {
    if n == 0 {
        return 0
    }
    return float64(jumlah(A, n)) / float64(n)
}