package main

import (
	"fmt"
	"math"
	"net/http"
)

// Constants for radius and height of the cylinder
const (
	radius = 7.0
	height = 10.0
)

// Function to calculate the base area of the cylinder
func luasAlas(r float64) float64 {
	return math.Pi * r * r
}

// Function to calculate the circumference of the cylinder base
func kelilingAlas(r float64) float64 {
	return 2 * math.Pi * r
}

// Function to calculate the volume of the cylinder
func volumeTabung(r, t float64) float64 {
	return math.Pi * r * r * t
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Calculate results using the functions
	luas := luasAlas(radius)
	keliling := kelilingAlas(radius)
	volume := volumeTabung(radius, height)

	// Display results on the web server
	fmt.Fprintf(w, "jariJari : %.0f, tinggi: %.0f, volume : %.2f, luas alas: %.2f, keliling alas: %.2f", radius, height, volume, luas, keliling)
}

func main() {
	// Set handler for the root URL
	http.HandleFunc("/", handler)

	// Start the server at localhost:8085
	fmt.Println("Server berjalan di http://localhost:8085")
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}
