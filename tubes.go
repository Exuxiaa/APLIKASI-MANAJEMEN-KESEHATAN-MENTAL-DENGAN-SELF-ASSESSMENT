package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Assessment struct {
	id      int
	nama    string
	hp      string
	tanggal int
	skor    [5]int
}

var data [100]Assessment
var jumlahData int = 0
var idOtomatis int = 1001

var pertanyaan = [10]string{
	"1. Apakah kamu merasa cemas akhir-akhir ini?",
	"2. Apakah kamu kesulitan tidur atau istirahat?",
	"3. Apakah kamu merasa lelah secara emosional?",
	"4. Apakah kamu bisa fokus saat bekerja atau belajar?",
	"5. Apakah kamu merasa berharga dan dihargai?",
	"6. Apakah kamu sering merasa sedih tanpa sebab?",
	"7. Apakah kamu kehilangan minat pada hal yang biasa kamu sukai?",
	"8. Apakah kamu merasa mudah marah akhir-akhir ini?",
	"9. Apakah kamu mengalami perubahan nafsu makan?",
	"10. Apakah kamu mengalami kesulitan dalam mengambil keputusan?",
}

var rekomendasiBagus = [3]string{"Pertahankan kondisi positifmu.", "Kamu dalam kondisi mental yang baik!", "Terus lakukan hal-hal yang kamu sukai."}
var rekomendasiCukup = [3]string{"Luangkan waktu untuk dirimu sendiri.", "Coba lakukan aktivitas relaksasi.", "Pertimbangkan berbicara dengan teman."}
var rekomendasiKurang = [3]string{"Pertimbangkan untuk berkonsultasi dengan profesional.", "Jangan ragu untuk mencari bantuan.", "Kamu tidak sendiri, cari dukungan."}

func tambahData() {
	var nama, hp string
	var skor [5]int
	var indeksTerpilih [5]int
	var i, tanggal, total int

	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan No HP: ")
	fmt.Scanln(&hp)

	fmt.Println("Jawab pertanyaan berikut (1-5):")

	digunakan := [10]bool{}
	for i = 0; i < 5; {
		indeks := rand.Intn(10)
		if !digunakan[indeks] {
			digunakan[indeks] = true
			indeksTerpilih[i] = indeks
			i++
		}
	}

	for i = 0; i < 5; i++ {
		var jawaban int
		for {
			fmt.Println(pertanyaan[indeksTerpilih[i]])
			fmt.Print("Jawaban (1-5): ")
			fmt.Scanln(&jawaban)
			if jawaban >= 1 && jawaban <= 5 {
				break
			}
			fmt.Println("Input tidak valid.")
		}
		skor[i] = jawaban
	}

	tanggal = time.Now().Year()*10000 + int(time.Now().Month())*100 + time.Now().Day()

	data[jumlahData] = Assessment{id: idOtomatis, nama: nama, hp: hp, tanggal: tanggal, skor: skor}
	jumlahData++
	fmt.Println("Data berhasil ditambahkan!")
	total = hitungTotal(skor)
	fmt.Println("Total Skor:", total, idOtomatis)
	fmt.Println("Rekomendasi:", rekomendasi(total))
	idOtomatis++
}

func hitungTotal(skor [5]int) int {
	var total, i int

	total = 0
	for i = 0; i < 5; i++ {
		total += skor[i]
	}
	return total
}

func rekomendasi(nilai int) string {
	if nilai > 20 {
		return rekomendasiBagus[rand.Intn(3)]
	} else if nilai >= 15 {
		return rekomendasiCukup[rand.Intn(3)]
	}
	return rekomendasiKurang[rand.Intn(3)]
}

func tampilkanAssessment(a Assessment) {
	var i, total int

	fmt.Println("---------------")
	fmt.Println("ID:      ", a.id)
	fmt.Println("Nama:    ", a.nama)
	fmt.Println("Tanggal: ", a.tanggal)
	fmt.Print("Skor:    ")
	for i = 0; i < 5; i++ {
		fmt.Print(a.skor[i], " ")
	}
	total = hitungTotal(a.skor)
	fmt.Println("\nTotal Skor:", total)
	fmt.Println("Rekomendasi:", rekomendasi(total))
	fmt.Println("---------------")
}

func cariIDSequential() {
	var idCari, i int
	fmt.Print("Masukkan ID yang dicari: ")
	fmt.Scanln(&idCari)
	for i = 0; i < jumlahData; i++ {
		if data[i].id == idCari {
			tampilkanAssessment(data[i])
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

// binary search
func binarySearch(idCari int) int {
	var kanan, kiri, tengah int

	kiri, kanan = 0, jumlahData-1
	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if data[tengah].id == idCari {
			return tengah
		} else if data[tengah].id < idCari {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func urutkanDataIDSelection() {
	var i, j, minIdx int

	for i = 0; i < jumlahData-1; i++ {
		minIdx = i
		for j = i + 1; j < jumlahData; j++ {
			if data[j].id < data[minIdx].id {
				minIdx = j
			}
		}
		if minIdx != i {
			temp := data[i]
			data[i] = data[minIdx]
			data[minIdx] = temp
		}
	}
}

func cariIDBinary() {
	var idCari, idx int

	fmt.Print("Masukkan ID yang dicari (Binary Search): ")
	fmt.Scanln(&idCari)
	urutkanDataIDSelection()
	idx = binarySearch(idCari)
	if idx != -1 {
		tampilkanAssessment(data[idx])
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func urutkanTanggalInsertion() {
	var i, j int

	for i = 1; i < jumlahData; i++ {
		temp := data[i]
		j = i - 1
		for j >= 0 && data[j].tanggal > temp.tanggal {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

func tampilkanLaporan(nama, hp string) {
	var count, hariIni, i, rata, jumlahSkor int

	count = 0
	jumlahSkor = 0
	hariIni = time.Now().Year()*10000 + int(time.Now().Month())*100 + time.Now().Day()

	fmt.Println("Laporan untuk:", nama)

	for i = jumlahData - 1; i >= 0 && count < 5; i-- {
		if data[i].nama == nama && data[i].hp == hp {
			tampilkanAssessment(data[i])
			count++
		}
	}

	count = 0
	for i = 0; i < jumlahData; i++ {
		if data[i].nama == nama && data[i].hp == hp {
			if hariIni-data[i].tanggal <= 30 {
				jumlahSkor += hitungTotal(data[i].skor)
				count++
			}
		}
	}
	if count > 0 {
		rata = jumlahSkor / count
		fmt.Println("Rata-rata skor 30 hari terakhir:", rata)
	} else {
		fmt.Println("Tidak ada data dalam 30 hari terakhir.")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		fmt.Println("\n=== Menu Self-Assessment ===")
		fmt.Println("1. Tambah Data Assessment")
		fmt.Println("2. Cari (Sequential Search)")
		fmt.Println("3. Cari (Binary Search)")
		fmt.Println("4. Urutkan berdasarkan Tanggal")
		fmt.Println("5. Tampilkan Laporan Pengguna")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")

		var pilih int

		fmt.Scanln(&pilih)
		switch pilih {
		case 1:
			tambahData()
		case 2:
			cariIDSequential()
		case 3:
			cariIDBinary()
		case 4:
			urutkanTanggalInsertion()
			fmt.Println("Data berhasil diurutkan berdasarkan tanggal.")
		case 5:
			var nama, hp string
			fmt.Print("Masukkan Nama: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan No HP: ")
			fmt.Scanln(&hp)
			tampilkanLaporan(nama, hp)
		case 6:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
