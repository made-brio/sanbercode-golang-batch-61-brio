package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type NilaiMahasiswa struct {
	ID          uint   `json:"id"`
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	Nilai       uint   `json:"nilai"`
	IndeksNilai string `json:"indeks_nilai"`
}

var (
	nilaiMahasiswaList      = []NilaiMahasiswa{}
	idCounter          uint = 1
	mu                 sync.Mutex
)

func main() {
	http.HandleFunc("/nilai_mahasiswa", basicAuth(nilaiMahasiswaHandler))
	log.Println("Server is running on port 8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}

// Basic authentication middleware
func basicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != "brioamarta" || password != "brioamarta" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}
		handler(w, r)
	}
}

// Handler for POST and GET requests
func nilaiMahasiswaHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		handlePost(w, r)
	case "GET":
		handleGet(w)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Handle POST request to add a new NilaiMahasiswa
func handlePost(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Nama       string `json:"nama"`
		MataKuliah string `json:"mata_kuliah"`
		Nilai      uint   `json:"nilai"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if input.Nilai > 100 {
		http.Error(w, "Nilai must be less than or equal to 100", http.StatusBadRequest)
		return
	}

	// Determine IndeksNilai based on Nilai
	indeksNilai := calculateIndeks(input.Nilai)

	mu.Lock()
	defer mu.Unlock()

	nilaiMahasiswa := NilaiMahasiswa{
		ID:          idCounter,
		Nama:        input.Nama,
		MataKuliah:  input.MataKuliah,
		Nilai:       input.Nilai,
		IndeksNilai: indeksNilai,
	}
	nilaiMahasiswaList = append(nilaiMahasiswaList, nilaiMahasiswa)
	idCounter++

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nilaiMahasiswa)
}

// Handle GET request to retrieve all NilaiMahasiswa
func handleGet(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nilaiMahasiswaList)
}

// Function to calculate IndeksNilai based on Nilai
func calculateIndeks(nilai uint) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70:
		return "B"
	case nilai >= 60:
		return "C"
	case nilai >= 50:
		return "D"
	default:
		return "E"
	}
}
