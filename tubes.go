// kurang tambah data nama no hp, tampil assessment, cari id secq, binary , urutkan data, laporan


package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Assessment struct {
	id      int
	tanggal int
	skor    [5]int
}

var dataPengguna [100]Assessment // kapasitas maksimum 100 pengguna
var jumlahPengguna int = 0
var idOtomatis int = 1001

// Daftar pertanyaan kuisioner
var semuaPertanyaan = []string{
	"1. Apakah kamu merasa cemas akhir-akhir ini?",
	"2. Apakah kamu kesulitan tidur atau istirahat?",
	"3. Apakah kamu merasa lelah secara emosional?",
	"4. Apakah kamu bisa fokus saat bekerja atau belajar?",
	"5. Apakah kamu merasa berharga dan dihargai?",
	"6. Apakah kamu mengalami perubahan suasana hati secara drastis?",
	"7. Apakah kamu menghindari interaksi sosial?",
	"8. Apakah kamu merasa kehilangan motivasi?",
	"9. Apakah kamu sulit mengendalikan emosi?",
	"10. Apakah kamu merasa tertekan dengan tuntutan hidup?",
}

// Rekomendasi berdasarkan skor
var rekomendasiBagus = []string{"bagus 1", "bagus 2", "bagus 3"} // belum final (harus rubah)
var rekomendasiCukup = []string{"cukup 1", "cukup 2", "cukup 3"}
var rekomendasiKurang = []string{"kurang 1", "kurang 2", "kurang 3"}

func rekomendasi(nilai int) string {
	if nilai > 20 {
		return rekomendasiBagus[rand.Intn(len(rekomendasiBagus))]
	} else if nilai >= 15 {
		return rekomendasiCukup[rand.Intn(len(rekomendasiCukup))]
	}
	return rekomendasiKurang[rand.Intn(len(rekomendasiKurang))]
}

func tambahData() {
	var skorUser [5]int

	// Buat indeks acak untuk memilih 5 pertanyaan random
	indeks := rand.Perm(len(semuaPertanyaan))[:5]

	fmt.Println("Isi kuisioner dengan skala 1-5 (1 = sangat tidak setuju, 5 = sangat setuju):")
	for i := 0; i < 5; i++ {
		var skor int
		for {
			fmt.Println(semuaPertanyaan[indeks[i]])
			fmt.Print("Jawaban: ")
			fmt.Scanln(&skor)
			if skor >= 1 && skor <= 5 {
				break
			}
			fmt.Println("Jawaban harus antara 1 dan 5")
		}
		skorUser[i] = skor
	}

	tanggalAcak := 20250500 + rand.Intn(30) // id belum final (harus di ubah)

	dataPengguna[jumlahPengguna] = Assessment{
		id:      idOtomatis,
		tanggal: tanggalAcak,
		skor:    skorUser,
	}

	fmt.Printf("Data berhasil ditambahkan (ID: %d)\n", idOtomatis)

	total := hitungTotal(skorUser)
	fmt.Println("Total Skor:", total)
	fmt.Println("Rekomendasi:", rekomendasi(total))

	idOtomatis++
	jumlahPengguna++
}

func hitungTotal(skor [5]int) int {
	total := 0
	for i := 0; i < 5; i++ {
		total += skor[i]
	}
	return total
}

func main() {
	rand.Seed(time.Now().UnixNano())
	tambahData()
}
