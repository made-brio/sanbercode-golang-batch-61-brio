package main

import (
	"fmt"
)

// Struct Product dengan atribut Name, Price, dan Stock
type Product struct {
	Name  string
	Price float64
	Stock int
}

// Fungsi untuk memperbarui harga produk
func (p *Product) updatePrice(newPrice float64) {
	p.Price = newPrice
}

// Fungsi untuk mengurangi stok produk saat terjual
func (p *Product) updateStock(sold int) {
	if sold > p.Stock {
		fmt.Println("Stok tidak cukup")
	} else {
		p.Stock -= sold
	}
}

func main() {
	//1. Membuat contoh produk
	product1 := Product{Name: "Laptop", Price: 15000.0, Stock: 10}
	product2 := Product{Name: "Headphone", Price: 500.0, Stock: 25}

	// Menampilkan data awal produk
	fmt.Printf("Produk 1: %+v\n", product1)
	fmt.Printf("Produk 2: %+v\n", product2)

	//2. Memperbarui harga produk
	product1.updatePrice(14000.0)
	product2.updatePrice(450.0)

	//3. Menampilkan data harga produk
	fmt.Printf("Harga %s: %f\n", product1.Name, product1.Price)
	fmt.Printf("Harga %s: %f\n", product2.Name, product2.Price)

	//4. Memperbarui stok produk
	product1.updateStock(3)
	product2.updateStock(5)

	//5. Menampilkan data stok produk
	fmt.Printf("stok %s: %d\n", product1.Name, product1.Stock)
	fmt.Printf("stok %s: %d\n", product2.Name, product2.Stock)
}
