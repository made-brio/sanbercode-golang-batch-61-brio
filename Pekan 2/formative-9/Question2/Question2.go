package Question3

import (
	"fmt"
	"strings"
)

// =======================Soal 2=======================
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

// Interface untuk menampilkan informasi phone
type DisplayInfo interface {
	ShowInfo() string
}

// Method untuk menampilkan informasi phone dalam bentuk string
func (p Phone) ShowInfo() string {
	return fmt.Sprintf(
		"name  : %s\nbrand : %s\nyear  : %d\ncolors: %s",
		p.Name, p.Brand, p.Year, strings.Join(p.Colors, ", "),
	)
}
