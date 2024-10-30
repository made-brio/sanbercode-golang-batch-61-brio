package main

import (
	"fmt"
	"strconv"
)

func main() {
	//------------------soal 1------------------
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int

	//konversi variable
	intP, _ := strconv.Atoi(panjangPersegiPanjang)
	intL, _ := strconv.Atoi(lebarPersegiPanjang)

	intA, _ := strconv.Atoi(alasSegitiga)
	intT, _ := strconv.Atoi(tinggiSegitiga)

	//lakukan perhitungan dengan formula
	luasPersegiPanjang = intP * intL
	kelilingPersegiPanjang = 2 * (intP + intL)
	luasSegitiga = (intA * intT) / 2

	//tampilkan hasil perhitungan
	fmt.Println("Luas Persegi Panjang : ", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang : ", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga : ", luasSegitiga)

	//------------------soal 2------------------
	var nilaiJohn = 80
	var nilaiDoe = 50

	var inxJohn string
	var inxDoe string

	//Menghitung index nilai John
	switch {
	case nilaiJohn >= 80:
		inxJohn = "A"
	case nilaiJohn >= 70:
		inxJohn = "B"
	case nilaiJohn >= 60:
		inxJohn = "C"
	case nilaiJohn >= 50:
		inxJohn = "D"
	default:
		inxJohn = "E"
	}

	//Menghitung index nilai Doe
	switch {
	case nilaiDoe >= 80:
		inxDoe = "A"
	case nilaiDoe >= 70:
		inxDoe = "B"
	case nilaiDoe >= 60:
		inxDoe = "C"
	case nilaiDoe >= 50:
		inxDoe = "D"
	default:
		inxDoe = "E"
	}

	fmt.Println("Index Nilai John:", inxJohn)
	fmt.Println("Index Nilai Doe:", inxDoe)

	//------------------soal 3------------------
	var tanggal = 12
	var bulan = 1
	var tahun = 2003

	months := []string{"", "Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}

	if bulan >= 1 && bulan <= 12 {
		fmt.Printf("%d %s %d\n", tanggal, months[bulan], tahun)
	} else {
		fmt.Println("Reperesentasi bulan hanya angka 1-12")
	}

	//------------------soal 4------------------
	switch {
	case (tahun >= 1944 && tahun <= 1964):
		fmt.Println("Baby boomer")
	case (tahun >= 1965 && tahun <= 1979):
		fmt.Println("Generasi X")
	case (tahun >= 1980 && tahun <= 1994):
		fmt.Println("Generasi Y (Millenials)")
	case (tahun >= 1995 && tahun <= 2015):
		fmt.Println("Generasi Z")
	default:
		fmt.Println("Tidak termasuk dalam generasi terdefinisi")
	}
}
