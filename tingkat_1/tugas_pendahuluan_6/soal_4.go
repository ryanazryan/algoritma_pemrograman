package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan nilai n: ")
    fmt.Scan(&n)

    telkomUniversity(n)
}

func telkomUniversity(n int) {
    for i := 1; i <= n; i++ {
        fmt.Printf("%d Telkom University %d\n", i, n-i+1)
    }
}
