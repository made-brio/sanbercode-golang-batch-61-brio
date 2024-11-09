package main

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

func main() {
	//========================Soal 1=======================
	defer cetakKalimat("Golang Backend Development", 2021)
	fmt.Println("Kalimat ini dieksekusi terlebih dahulu")

	//========================Soal 2=======================
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	//========================Soal 3=======================
	angka := 1
	defer cetakAngka(&angka)

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	//========================Soal 4=======================
	var phonesSoal4 = []string{}

	tambahData(&phonesSoal4, "Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo")

	sort.Strings(phonesSoal4)

	tampilkanData(phonesSoal4)

	//========================Soal 5=======================
	var phonesSoal5 = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

	sort.Strings(phonesSoal5)

	var wg sync.WaitGroup
	wg.Add(len(phonesSoal5))

	for i, phone := range phonesSoal5 {
		go func(i int, phone string) {
			defer wg.Done()
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Printf("%d. %s\n", i+1, phone)
		}(i, phone)
	}

	wg.Wait()

	//========================Soal 6=======================
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}
	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}
}

// =========================Soal 1=======================
func cetakKalimat(kalimat string, tahun int) {
	fmt.Printf("%s %d\n", kalimat, tahun)
}

// =========================Soal 2=======================
func kelilingSegitigaSamaSisi(sisi int, tampilkanKalimat bool) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	if sisi == 0 {
		err := errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		if !tampilkanKalimat {
			panic(err)
		}
		return "", err
	}

	keliling := sisi * 3
	if tampilkanKalimat {
		return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling), nil
	}
	return fmt.Sprintf("%d", keliling), nil
}

// =========================Soal 3=======================
func tambahAngka(n int, angka *int) {
	*angka += n
}

func cetakAngka(angka *int) {
	fmt.Printf("Total angka: %d\n", *angka)
}

// =========================Soal 4=======================
func tambahData(phones *[]string, data ...string) {
	*phones = append(*phones, data...)
}

func tampilkanData(phones []string) {
	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}
}

// =========================Soal 6=======================
func getMovies(ch chan string, movies ...string) {
	for _, movie := range movies {
		ch <- movie
	}
	close(ch)
}
