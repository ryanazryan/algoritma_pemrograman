package main

import (
    "fmt"
)

func main2() {
    var a, b, c float64
    fmt.Print("Masukkan tiga bilangan: ")
    fmt.Scan(&a, &b, &c)

    if a > 0 && b > 0 && c > 0 {
        if a+b > c && a+c > b && b+c > a {
            fmt.Printf("%.1f, %.1f, dan %.1f segitiga? true\n", a, b, c)
            fmt.Println("Ya iyalah")
        } else {
            fmt.Printf("%.1f, %.1f, dan %.1f segitiga? false\n", a, b, c)
            fmt.Println("Salah satu sisi terlalu panjang dibandingkan yang lain")
        }
    } else {
        fmt.Printf("%.1f, %.1f, dan %.1f segitiga? false\n", a, b, c)
        fmt.Println("Masa sisi segitiga negatif")
    }
}
