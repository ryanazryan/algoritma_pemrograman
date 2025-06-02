package main

import "fmt"
import "strings"
import "time"

func main4(){
	var nama, bulanLahir string

	fmt.Print("Masukkan nama: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan bulan lahir: ")
	fmt.Scan(&bulanLahir)

	bulanIni := time.Now().Month().String()

	namaBulanIndonesia := map[string] string{
		"January" : "Januari",
		"February" : "Februari",
		"Maret" : "Maret",
		"April" : "April",
		"Mei" : "Mei",
		"June" : "Juni",
		"July" : "Juli",
		"September" : "September",
		"October" : "Oktober",
		"November" : "November",
		"December" : "Desember",
	}

	bulanIniIndonesia := namaBulanIndonesia[ bulanIni]

	if strings.ToLower(bulanLahir) == strings.ToLower(bulanIniIndonesia){
		fmt.Printf("%s berulang tahun pada bulan ini? true\n", nama)
	} else {
		fmt.Printf("%s berulang tahun pada bulan ini? false\n", nama)
	}

	
}