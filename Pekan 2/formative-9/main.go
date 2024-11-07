package main

import (
	"fmt"
	. "formative9/Question1"
	. "formative9/Question2"
	. "formative9/Question3"
)

func main() {
	//=======================Soal 1=======================
	// Bangun Datar
	var segitiga HitungBangunDatar = SegitigaSamaSisi{Alas: 6, Tinggi: 4}
	var persegi HitungBangunDatar = PersegiPanjang{Panjang: 5, Lebar: 3}

	fmt.Println("Bangun Datar:")
	fmt.Printf("Segitiga Sama Sisi - Luas: %d, Keliling: %d\n", segitiga.Luas(), segitiga.Keliling())
	fmt.Printf("Persegi Panjang - Luas: %d, Keliling: %d\n", persegi.Luas(), persegi.Keliling())

	// Bangun Ruang
	var tabung HitungBangunRuang = Tabung{JariJari: 3, Tinggi: 10}
	var balok HitungBangunRuang = Balok{Panjang: 4, Lebar: 5, Tinggi: 6}

	fmt.Println("Bangun Ruang:")
	fmt.Printf("Tabung - Volume: %.2f, Luas Permukaan: %.2f\n", tabung.Volume(), tabung.LuasPermukaan())
	fmt.Printf("Balok - Volume: %.2f, Luas Permukaan: %.2f\n", balok.Volume(), balok.LuasPermukaan())

	//=======================Soal 2=======================

	// Input data phone
	myPhone := Phone{
		Name:   "Samsung Galaxy Phone 20",
		Brand:  "Samsung Galaxy Phone 20",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	// Menampilkan informasi phone menggunakan method showInfo
	fmt.Println(myPhone.ShowInfo())

	//=======================Soal 3=======================
	fmt.Println(LuasPersegi(4, true))  // Output: "luas persegi dengan sisi 4 cm adalah 16 cm"
	fmt.Println(LuasPersegi(8, false)) // Output: 64
	fmt.Println(LuasPersegi(0, true))  // Output: "Maaf anda belum menginput sisi dari persegi"
	fmt.Println(LuasPersegi(0, false)) // Output: <nil>

	//=======================Soal 4=======================

	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// Melakukan type assertion
	prefixStr := prefix.(string)
	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)

	// Menggabungkan seluruh angka dan menghitung total
	total := 0
	output := prefixStr
	for i, num := range angkaPertama {
		output += fmt.Sprintf("%d", num)
		total += num
		if i < len(angkaPertama)-1 || len(angkaKedua) > 0 {
			output += " + "
		}
	}
	for i, num := range angkaKedua {
		output += fmt.Sprintf("%d", num)
		total += num
		if i < len(angkaKedua)-1 {
			output += " + "
		}
	}

	// Menambahkan hasil akhir
	output += fmt.Sprintf(" = %d", total)

	// Menampilkan output
	fmt.Println(output)

}
