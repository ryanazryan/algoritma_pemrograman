package main

import (
    "fmt"
    "strings"
)

func main3() {
    var op1, op2 float64
    var operator string

    fmt.Print("Masukkan operan1 operator operan2 (contoh: 1.5 + 5.6): ")
    fmt.Scanf("%f %s %f", &op1, &operator, &op2)

    operator = strings.TrimSpace(operator)

    switch operator {
    case "+":
        fmt.Printf("Hasil: %g\n", op1+op2)
    case "-":
        fmt.Printf("Hasil: %g\n", op1-op2)
    case "*":
        fmt.Printf("Hasil: %g\n", op1*op2)
    case "/":
        if op2 != 0 {
            fmt.Printf("Hasil: %g\n", op1/op2)
        } else {
            fmt.Println("Error: Tidak dapat membagi dengan nol.")
        }
    default:
        fmt.Println("Error: Operator tidak valid.")
    }
}
