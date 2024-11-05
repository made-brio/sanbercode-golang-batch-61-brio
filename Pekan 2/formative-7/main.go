package main

import (
	"fmt"
)

// //=======================Soal 1=======================
type Buah struct {
	Nama       string
	Warna      string
	AdaBijinya bool
	Harga      int
}

// =======================Soal 2=======================
type Segitiga struct {
	alas, tinggi int
}

type Persegi struct {
	sisi int
}

type PersegiPanjang struct {
	panjang, lebar int
}

// Method untuk menghitung luas Segitiga
func (s Segitiga) Luas() int {
	return (s.alas * s.tinggi) / 2
}

// Method untuk menghitung luas Persegi
func (p Persegi) Luas() int {
	return p.sisi * p.sisi
}

// Method untuk menghitung luas Persegi Panjang
func (pp PersegiPanjang) Luas() int {
	return pp.panjang * pp.lebar
}

// =======================Soal 3=======================
type phone struct {
	name, brand string
	year        int
	colors      []string
}

// Method untuk menambahkan warna ke property Colors
func (p *phone) AddColor(color string) {
	p.colors = append(p.colors, color)
}

// =======================Soal 4=======================
// Struct Movie dengan beberapa property
type Movie struct {
	Title    string
	Genre    string
	Duration int
	Year     int
}

// Function untuk menambahkan data movie ke dalam slice dataFilm
func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]Movie) {
	movie := Movie{
		Title:    title,
		Duration: duration,
		Genre:    genre,
		Year:     year,
	}
	*dataFilm = append(*dataFilm, movie)
}

func main() {
	//=======================Soal 1=======================
	// Membuat beberapa objek buah sesuai data yang diberikan
	nanas := Buah{
		Nama:       "Nanas",
		Warna:      "Kuning",
		AdaBijinya: false,
		Harga:      9000,
	}

	jeruk := Buah{
		Nama:       "Jeruk",
		Warna:      "Oranye",
		AdaBijinya: true,
		Harga:      8000,
	}

	semangka := Buah{
		Nama:       "Semangka",
		Warna:      "Hijau & Merah",
		AdaBijinya: true,
		Harga:      10000,
	}

	pisang := Buah{
		Nama:       "Pisang",
		Warna:      "Kuning",
		AdaBijinya: false,
		Harga:      5000,
	}

	// Menampilkan data buah-buahan
	fruits := []Buah{nanas, jeruk, semangka, pisang}
	for i, buah := range fruits {
		if buah.AdaBijinya {
			fmt.Printf("%d. Nama: %s\n   Warna: %s\n   Ada Bijinya: Ada\n   Harga: %d\n",
				i+1, buah.Nama, buah.Warna, buah.Harga)
		} else {
			fmt.Printf("%d. Nama: %s\n   Warna: %s\n   Ada Bijinya: Tidak\n   Harga: %d\n",
				i+1, buah.Nama, buah.Warna, buah.Harga)
		}

	}

	//=======================Soal 2=======================
	// Membuat objek untuk setiap bentuk
	segitiga := Segitiga{alas: 10, tinggi: 5}
	persegi := Persegi{sisi: 4}
	persegiPanjang := PersegiPanjang{panjang: 8, lebar: 3}

	// Mencetak luas masing-masing bentuk
	fmt.Printf("Luas Segitiga: %d\n", segitiga.Luas())
	fmt.Printf("Luas Persegi: %d\n", persegi.Luas())
	fmt.Printf("Luas Persegi Panjang: %d\n", persegiPanjang.Luas())

	//=======================Soal 3=======================
	// Membuat object phone
	phone1 := phone{
		name:   "Samsung A55",
		brand:  "Samsung",
		year:   2024,
		colors: []string{"Black", "White"},
	}

	// Menambahkan warna baru
	phone1.AddColor("Blue")
	phone1.AddColor("Red")

	// Menampilkan hasil setelah penambahan warna
	fmt.Printf("Phone: %s\nBrand: %s\nYear: %d\nColors: %v\n",
		phone1.name, phone1.brand, phone1.year, phone1.colors)

	//=======================Soal 4=======================

	// Slice untuk menyimpan data film
	var dataFilm = []Movie{}

	// Menambahkan data film
	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	// Menampilkan data film
	for i, film := range dataFilm {
		fmt.Printf("%d. Title: %s\n   Genre: %s\n   Duration: %d minutes\n   Year: %d\n",
			i+1, film.Title, film.Genre, film.Duration, film.Year)
	}

}
