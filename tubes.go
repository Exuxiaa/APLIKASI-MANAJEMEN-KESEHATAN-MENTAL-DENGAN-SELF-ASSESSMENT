package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Assessment struct {
	id      int
	tanggal int
	skor    []int
}

var dataPengguna []Assessment
var idOtomatis int = 1001

// Daftar pertanyaan kuisioner
var pertanyaan = []string{
	"1. Apakah kamu merasa cemas akhir-akhir ini?",
	"2. Apakah kamu kesulitan tidur atau istirahat?",
	"3. Apakah kamu merasa lelah secara emosional?",
	"4. Apakah kamu bisa fokus saat bekerja atau belajar?",
	"5. Apakah kamu merasa berharga dan dihargai?",
}

// Rekomendasi berdasarkan skor
var rekomendasiBagus = []string{
	"bagus 1",
	"bagus 2",
	"bagus 3",
}

var rekomendasiCukup = []string{
	"cukup 1",
	"cukup 2",
	"cukup 3",
}

var rekomendasiKurang = []string{
	"kurang 1",
	"kurang 2",
	"kurang 3",
}

// Fungsi untuk menampilkan jawaban (rekomendasi) acak
func rekomendasi(nilai int) string {
	if nilai > 20 {
		return rekomendasiBagus[rand.Intn(len(rekomendasiBagus))]
	} else if nilai >= 15 {
		return rekomendasiCukup[rand.Intn(len(rekomendasiCukup))]
	}
	return rekomendasiKurang[rand.Intn(len(rekomendasiKurang))]
}

// Tambah data dari input user
func tambahData() {
	var jawabanUser []int
	fmt.Println("Isi kuisioner dengan skala 1-5 (1 = sangat tidak setuju, 5 = sangat setuju):")
	for i := 0; i < len(pertanyaan); i++ {
		var skor int
		for {
			fmt.Println(pertanyaan[i])
			fmt.Print("Jawaban: ")
			fmt.Scanln(&skor)
			if skor >= 1 && skor <= 5 {
				break
			}
			fmt.Println("Jawaban harus antara 1 dan 5")
		}
		jawabanUser = append(jawabanUser, skor)
	}

	tanggalAcak := 20250500 + rand.Intn(30)

	dataPengguna = append(dataPengguna, Assessment{
		id:      idOtomatis,
		tanggal: tanggalAcak,
		skor:    jawabanUser,
	})
	fmt.Printf("Data berhasil ditambahkan (ID: %d)\n", idOtomatis)

	total := hitungTotal(jawabanUser)
	fmt.Println("Total Skor:", total)
	fmt.Println("Rekomendasi:", rekomendasi(total))

	idOtomatis++
}

func hitungTotal(skor []int) int {
	total := 0
	for _, v := range skor {
		total += v
	}
	return total
}

// Fungsi-fungsi lain seperti ubahData(), hapusData(), sort, dsb tetap sama…
func main() {
	rand.Seed(time.Now().UnixNano()) // agar hasil random (tidak selalu sama)

	tambahData() // hanya input manual dari user
}
