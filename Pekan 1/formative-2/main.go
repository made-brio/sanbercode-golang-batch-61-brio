package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//------------------soal 1------------------
	var varBootcamp, varDigital, varSkill, varSanbercode, varGolang = "Bootcamp", "Digital", "Skill", "Sanbercode", "Golang"
	sentence := strings.Join([]string{varBootcamp, varDigital, varSkill, varSanbercode, varGolang}, " ")
	fmt.Println(sentence)

	//------------------soal 2------------------
	halo := "Halo Dunia"
	halo = strings.Replace(halo, "Dunia", "Golang", 1)
	fmt.Println(halo)

	//------------------soal 3------------------
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	//Mengubah awalan kataKedua
	kataKedua = strings.ToUpper(kataKedua[:1]) + kataKedua[1:]

	//Mengubah akhiran kataKetiga
	kataKetiga = kataKetiga[:6] + strings.ToUpper(string(kataKetiga[6]))

	//Membuat kata keempat uppercase
	kataKeempat = strings.ToUpper(kataKeempat)

	fmt.Println(kataPertama, kataKedua, kataKetiga, kataKeempat)

	//------------------soal 4------------------
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	//Mengubah tipe data string ke integer
	intPertama, _ := strconv.Atoi(angkaPertama)
	intKedua, _ := strconv.Atoi(angkaKedua)
	intKetiga, _ := strconv.Atoi(angkaKetiga)
	intKeempat, _ := strconv.Atoi(angkaKeempat)

	//Menjumlahkan integer
	result := intPertama + intKedua + intKetiga + intKeempat
	fmt.Println(result)

	//------------------soal 5------------------
	kalimat := "halo halo bandung"
	angka := 2021

	//Mengubah kata "halo" mennjadi "Hi"
	kalimat = strings.ReplaceAll(kalimat, "halo", "Hi")

	fmt.Printf("\"%s\" - %d\n", kalimat, angka)
}
