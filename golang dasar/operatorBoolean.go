package main

import "fmt"

func main() {
	// && or || or !
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir = nilaiAkhir >= 80
	var lulusAbsensi = absensi >= 80

	var lulus = lulusAbsensi && lulusNilaiAkhir
	fmt.Println(lulus)
}
