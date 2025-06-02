package main

import "fmt"

func main2(){
	type individu struct {
		nama string
		target float64
		pendapatan float64
		sisaUang float64
	}

	var person individu
	fmt.Scan(&person.nama, &person.target, &person.pendapatan, &person.sisaUang)

	targetTabung := (person.target / 100) * person.pendapatan
	ditabung := person.pendapatan - person.sisaUang

	if ditabung == targetTabung {
		fmt.Println(person.nama + ", kamu berhasil menabung sesuai target, " + fmt.Sprintf("%.0f", person.target) + "% dari pendapatan kamu")
	} else if ditabung > targetTabung{
		fmt.Println(person.nama + ", kamu berhasil menabung melebihi target, " + fmt.Sprintf("%.0f", person.target) + "% dari pendapatan kamu")
	} else {
		fmt.Println(person.nama + ", kamu gagal menabung sesuai target, " + fmt.Sprintf("%.0f", person.target) + "% dari pendapatan kamu")
	}
}