package main

import (
	"fmt"
	"math"
	"strings"
)

// //=======================Soal 1=======================
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

// Struktur dan interface untuk Bangun Ruang
type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * float64((b.panjang*b.lebar)+(b.panjang*b.tinggi)+(b.lebar*b.tinggi))
}

// =======================Soal 2=======================
type phone struct {
	name, brand string
	year        int
	colors      []string
}

// Interface untuk menampilkan informasi phone
type displayInfo interface {
	showInfo() string
}

// Method untuk menampilkan informasi phone dalam bentuk string
func (p phone) showInfo() string {
	return fmt.Sprintf(
		"name  : %s\nbrand : %s\nyear  : %d\ncolors: %s",
		p.name, p.brand, p.year, strings.Join(p.colors, ", "),
	)
}

// =======================Soal 3=======================
// Fungsi luasPersegi dengan return tipe interface kosong
func luasPersegi(sisi int, showText bool) interface{} {
	if sisi == 0 {
		if showText {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}

	luas := sisi * sisi
	if showText {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	}
	return luas
}

func main() {
	//=======================Soal 1=======================
	// Bangun Datar
	var segitiga hitungBangunDatar = segitigaSamaSisi{alas: 6, tinggi: 4}
	var persegi hitungBangunDatar = persegiPanjang{panjang: 5, lebar: 3}

	fmt.Println("Bangun Datar:")
	fmt.Printf("Segitiga Sama Sisi - Luas: %d, Keliling: %d\n", segitiga.luas(), segitiga.keliling())
	fmt.Printf("Persegi Panjang - Luas: %d, Keliling: %d\n", persegi.luas(), persegi.keliling())

	// Bangun Ruang
	var tabung hitungBangunRuang = tabung{jariJari: 3, tinggi: 10}
	var balok hitungBangunRuang = balok{panjang: 4, lebar: 5, tinggi: 6}

	fmt.Println("Bangun Ruang:")
	fmt.Printf("Tabung - Volume: %.2f, Luas Permukaan: %.2f\n", tabung.volume(), tabung.luasPermukaan())
	fmt.Printf("Balok - Volume: %.2f, Luas Permukaan: %.2f\n", balok.volume(), balok.luasPermukaan())

	//=======================Soal 2=======================

	// Input data phone
	myPhone := phone{
		name:   "Samsung Galaxy Phone 20",
		brand:  "Samsung Galaxy Phone 20",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	// Menampilkan informasi phone menggunakan method showInfo
	fmt.Println(myPhone.showInfo())

	//=======================Soal 3=======================
	fmt.Println(luasPersegi(4, true))  // Output: "luas persegi dengan sisi 4 cm adalah 16 cm"
	fmt.Println(luasPersegi(8, false)) // Output: 64
	fmt.Println(luasPersegi(0, true))  // Output: "Maaf anda belum menginput sisi dari persegi"
	fmt.Println(luasPersegi(0, false)) // Output: <nil>

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
