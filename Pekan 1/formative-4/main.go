package main

import (
	"fmt"
)

func main() {
	//------------------soal 1------------------

	for i := 1; i <= 20; i++ {
		if i%2 != 0 && i%3 == 0 {
			fmt.Println(i, "- I Love Coding")
		} else if i%2 == 0 {
			fmt.Println(i, "- Berkualitas")
		} else {
			fmt.Println(i, "- Santai")
		}
	}

	//------------------soal 2------------------

	for i := 0; i < 7; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
	//------------------soal 3------------------
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	// Ambil elemen dari indeks 2 hingga akhir array
	result := kalimat[2:]
	fmt.Println(result)

	//------------------soal 4------------------

	var sayuran = []string{}
	// Menambahkan data ke dalam slice sayuran
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	// Menampilkan data sayuran dengan nomor urut
	for i, v := range sayuran {
		fmt.Printf("%d. %s\n", i+1, v)
	}
	//------------------soal 5------------------

	satuan := map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	volume := 1
	// Menampilkan setiap elemen dalam map dan menghitung volume balok
	for k, v := range satuan {
		fmt.Printf("%s= %d\n", k, v)
		volume *= v
	}
	fmt.Printf("volume balok = %d\n", volume)
}
