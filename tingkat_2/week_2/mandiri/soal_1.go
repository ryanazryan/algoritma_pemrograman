package main

import "fmt"

func main() {
    type Tim struct {
        cabangLomba string
        anggota1    string
        anggota2    string
        anggota3    string
        skorTotal   float64
        kelulusan   string
		idTim string
    }

    type Juri struct {
        inisial      string
        skorDiberikan float64
    }

    var tim Tim
    var juri1, juri2, juri3 Juri
	
    fmt.Scan(&tim.anggota1)
    fmt.Scan(&tim.anggota2)
    fmt.Scan(&tim.anggota3)
    fmt.Scan(&tim.cabangLomba)

	if tim.cabangLomba == "CP" {
		tim.idTim = "111"
	} else if tim.cabangLomba == "CTF" {
		tim.idTim = "112"
	} else if tim.cabangLomba == "ITB" {
		tim.idTim = "113"
	}

    fmt.Scan(&juri1.inisial)
    fmt.Scan(&juri1.skorDiberikan)
    tim.skorTotal += juri1.skorDiberikan

    fmt.Scan(&juri2.inisial)
    fmt.Scan(&juri2.skorDiberikan)
    tim.skorTotal += juri2.skorDiberikan

    fmt.Scan(&juri3.inisial)
    fmt.Scan(&juri3.skorDiberikan)
    tim.skorTotal += juri3.skorDiberikan

    rataRata := tim.skorTotal / 3

    if rataRata >= 80 {
        tim.kelulusan = "Lolos"
    } else {
        tim.kelulusan = "Tidak Lolos"
    }

    fmt.Println("DATA TIM: ", tim.idTim, tim.anggota1, tim.anggota2, tim.anggota3)
    fmt.Println("Anggota:")
    fmt.Println(tim.anggota1)
    fmt.Println(tim.anggota2)
    fmt.Println(tim.anggota3)
    fmt.Printf("Dinilai oleh: %s %s %s\n", juri1.inisial, juri2.inisial, juri3.inisial)
    fmt.Printf("Lolos dengan nilai: %.2f\n", rataRata)
}