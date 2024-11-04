package main

import (
	"fmt"
	"math"
)

func main() {
	//soal 1
	var luasLingkaran float64
	var kelilingLingkaran float64

	changeCircle(&luasLingkaran, &kelilingLingkaran, 6)
	fmt.Println("luas: ", luasLingkaran)
	fmt.Println("keliling: ", kelilingLingkaran)

	//soal 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//soal 3
	var buah = []string{}

	addFruit(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")

	for i, j := range buah {
		fmt.Printf("%d.  %s\n", i+1, j)
	}

	//soal4
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	// Menampilkan data film
	for i, film := range dataFilm {
		fmt.Printf("%d. title: %s\n   duration: %s\n   genre: %s\n   year: %s\n",
			i+1, film["title"], film["duration"], film["genre"], film["year"])
	}
}

// soal 1
func changeCircle(area *float64, circumference *float64, radius float64) {
	*circumference = 2 * math.Pi * radius
	*area = math.Pi * radius * radius
}

// soal 2
func introduce(kalimat *string, nama string, gender string, pekerjaan string, umur string) {
	if gender == "laki-laki" {
		*kalimat = "Pak " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	} else if gender == "perempuan" {
		*kalimat = "Bu " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	}
	return
}

// soal 3
func addFruit(keranjang *[]string, fruit ...string) {
	*keranjang = append(*keranjang, fruit...)
}

// soal 4
func tambahDataFilm(title string, duration string, genre string, year string, film *[]map[string]string) {
	newFilm := map[string]string{
		"title":    title,
		"duration": duration,
		"genre":    genre,
		"year":     year,
	}
	*film = append(*film, newFilm)
}
