package main

import "fmt"

const NMAX int = 10

type tDrakor struct {
    judul   string
    rating  float64
    episode, durasi int
}

type tabDrakor [NMAX]tDrakor

func bacaDataDrakor(A *tabDrakor, n *int) {
    fmt.Scan(n)
    if *n > NMAX {
        *n = NMAX
    }
    for i := 0; i < *n; i++ {
        fmt.Scan(&A[i].judul, &A[i].rating, &A[i].episode, &A[i].durasi)
    }
}

func cetakDataDrakor(A tabDrakor, n int) {
    fmt.Printf("%20s %6s %6s %6s\n", "Judul", "Rating", "Jum Ep", "Durasi")
    for i := 0; i < n; i++ {
        fmt.Printf("%20s %6.1f %6d %6d\n", A[i].judul, A[i].rating, A[i].episode, A[i].durasi)
    }
}

func urutMenurun(A *tabDrakor, n int) {
    for pass := 0; pass < n-1; pass++ {
        maxIdx := pass
        for i := pass + 1; i < n; i++ {
            if A[i].rating > A[maxIdx].rating {
                maxIdx = i
            }
        }
        A[pass], A[maxIdx] = A[maxIdx], A[pass]
    }
}

func urutMenaik(A *tabDrakor, n int) {
    for pass := 0; pass < n-1; pass++ {
        minIdx := pass
        for i := pass + 1; i < n; i++ {
            if A[i].durasi < A[minIdx].durasi {
                minIdx = i
            }
        }
        A[pass], A[minIdx] = A[minIdx], A[pass]
    }
}

func main() {
    var drakor tabDrakor
    var nDrakor int

    bacaDataDrakor(&drakor, &nDrakor)

    fmt.Println("Data sebelum diurutkan:")
    cetakDataDrakor(drakor, nDrakor)

    fmt.Println("\nData setelah diurutkan menurun berdasarkan rating:")
    urutMenurun(&drakor, nDrakor)
    cetakDataDrakor(drakor, nDrakor)

    fmt.Println("\nData setelah diurutkan menaik berdasarkan durasi:")
    urutMenaik(&drakor, nDrakor)
    cetakDataDrakor(drakor, nDrakor)
}