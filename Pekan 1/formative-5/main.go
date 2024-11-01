package main

import (
	"fmt"
	"strings"
)

//------------------soal 1------------------

// Function untuk menghitung luas persegi panjang
func luasPersegiPanjang(panjang, lebar int) int {
	return panjang * lebar
}

// Function untuk menghitung keliling persegi panjang
func kelilingPersegiPanjang(panjang, lebar int) int {
	return 2 * (panjang + lebar)
}

// Function untuk menghitung volume balok
func volumeBalok(panjang, lebar, tinggi int) int {
	return panjang * lebar * tinggi
}

// ------------------soal 2------------------
// Function introduce dengan predefined value/ named value
func introduce(nama, gender, pekerjaan, umur string) (description string) {
	if gender == "laki-laki" {
		description = "Pak " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	} else if gender == "perempuan" {
		description = "Bu " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	}
	return
}

// ------------------soal 3------------------
func buahFavorit(nama string, buah ...string) string {
	return fmt.Sprintf("Halo nama saya %s dan buah favorit saya adalah \"%s\"", nama, strings.Join(buah, "\", \""))
}

//------------------soal 4------------------

func main() {
	//------------------soal 1------------------

	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	//------------------soal 2------------------
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//------------------soal 3------------------

	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// Output: Halo nama saya John dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	//------------------soal 4------------------

	var dataFilm = []map[string]string{}

	// Closure function untuk menambahkan data film ke slice dataFilm
	tambahDataFilm := func(title, jam, genre, tahun string) {
		film := map[string]string{
			"genre": genre,
			"jam":   jam,
			"tahun": tahun,
			"title": title,
		}
		dataFilm = append(dataFilm, film)
	}

	// Menambahkan data film ke dataFilm
	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	// Menampilkan dataFilm
	for _, item := range dataFilm {
		fmt.Println(item)
	}

}
