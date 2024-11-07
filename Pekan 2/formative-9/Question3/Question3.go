package Question3

import "fmt"

// =======================Soal 3=======================
// Fungsi luasPersegi dengan return tipe interface kosong
func LuasPersegi(sisi int, showText bool) interface{} {
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
