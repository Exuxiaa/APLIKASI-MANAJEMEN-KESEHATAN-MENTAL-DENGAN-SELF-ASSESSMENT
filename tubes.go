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
	tanggal string
	skor    [5]int
	rekom   string
}

var data [100]Assessment
var jumlahData int = 0
var idOtomatis int = 1001

var pertanyaan = [20]string{
	"Apakah kamu merasa bahagia dan puas dengan hidupmu?",
	"Apakah kamu mampu mengelola stres sehari-hari dengan baik?",
	"Apakah kamu sering merasa tenang dan damai secara emosional?",
	"Apakah kamu optimis tentang masa depanmu?",
	"Apakah kamu nyaman dalam berinteraksi sosial dengan orang lain?",
	"Apakah kamu bisa mengatasi tantangan atau kesulitan dengan baik?",
	"Apakah kamu merasa termotivasi untuk melakukan hal-hal yang penting bagimu?",
	"Apakah kamu percaya diri dalam mengambil keputusan?",
	"Apakah kamu merasa dihargai dan dicintai oleh orang-orang di sekitarmu?",
	"Apakah kamu memiliki kualitas tidur yang baik dalam beberapa minggu terakhir?",
	"Apakah kamu sering merasa cemas atau khawatir tanpa alasan yang jelas?",
	"Apakah kamu kehilangan minat atau energi untuk melakukan aktivitas sehari-hari?",
	"Apakah kamu merasa bisa fokus dan berkonsentrasi dengan baik?",
	"Apakah kamu merasa bahwa hidupmu memiliki makna dan tujuan?",
	"Apakah kamu nyaman dalam mengungkapkan emosi dan perasaan kepada orang lain?",
	"Apakah kamu merasa kesepian atau terisolasi dari orang-orang di sekitarmu?",
	"Apakah kamu mampu mengendalikan pikiran negatif?",
	"Apakah kamu merasa terbebani oleh tanggung jawab atau tekanan hidup?",
	"Apakah kamu puas dengan hubungan sosial dan emosional yang kamu miliki?",
	"Apakah kamu bisa menikmati momen kecil dalam hidup tanpa gangguan pikiran negatif?",
}

var rekomendasiBagus = [3]string{"Pertahankan kondisi positifmu.", "Kamu dalam kondisi mental yang baik!", "Terus lakukan hal-hal yang kamu sukai."}
var rekomendasiCukup = [3]string{"Luangkan waktu untuk dirimu sendiri.", "Coba lakukan aktivitas relaksasi.", "Pertimbangkan berbicara dengan teman."}
var rekomendasiKurang = [3]string{"Pertimbangkan untuk berkonsultasi dengan profesional.", "Jangan ragu untuk mencari bantuan.", "Kamu tidak sendiri, cari dukungan."}

func tambahData() {
	var nama, hp string
	var skor [5]int
	var indeksTerpilih [5]int
	var i, id int
	var tanggal, hasilRekom string

	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan No HP: ")
	fmt.Scanln(&hp)

	// Validasi ID
	for {
		fmt.Print("Masukkan ID: ")
		fmt.Scanln(&id)

		duplikat := false
		for i = 0; i < jumlahData; i++ {
			if data[i].id == id && (data[i].nama != nama || data[i].hp != hp) {
				fmt.Println("ID sudah digunakan oleh pengguna lain. Silakan masukkan ID yang berbeda.")
				duplikat = true
				break
			}
		}

		if !duplikat {
			break
		}
	}

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

	tanggal = time.Now().Format("2006-01-02")
	total := hitungTotal(skor)
	hasilRekom = rekomendasi(total) // Simpan hasil rekomendasi

	data[jumlahData] = Assessment{
		id:      id,
		nama:    nama,
		hp:      hp,
		tanggal: tanggal,
		skor:    skor,
		rekom:   hasilRekom,
	}
	jumlahData++

	fmt.Println("Data berhasil ditambahkan!")
	fmt.Println("Total Skor:", total)
	fmt.Println("Rekomendasi:", hasilRekom)
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

func editData() {
	var idCari, i int
	var namaBaru, hpBaru string

	fmt.Print("Masukkan ID yang ingin diedit: ")
	fmt.Scanln(&idCari)

	var ditemukan bool = false

	for i = 0; i < jumlahData; i++ {
		if data[i].id == idCari {
			if !ditemukan {
				fmt.Print("Masukkan Nama Baru: ")
				fmt.Scanln(&namaBaru)
				fmt.Print("Masukkan No HP Baru: ")
				fmt.Scanln(&hpBaru)
			}
			data[i].nama = namaBaru
			data[i].hp = hpBaru
			ditemukan = true
		}
	}

	if ditemukan {
		fmt.Println("Semua data dengan ID tersebut berhasil diperbarui!")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func hapusData() {
	var idCari, i, j int

	fmt.Print("Masukkan ID yang ingin dihapus: ")
	fmt.Scanln(&idCari)

	for i = 0; i < jumlahData; i++ {
		if data[i].id == idCari {
			for j = i; j < jumlahData-1; j++ {
				data[j] = data[j+1]
			}
			jumlahData--
			fmt.Println("Data berhasil dihapus.")
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func tampilkanAssessment(a Assessment) {
	var i int
	fmt.Println("---------------")
	fmt.Println("ID:      ", a.id)
	fmt.Println("Nama:    ", a.nama)
	fmt.Println("Tanggal: ", a.tanggal)
	fmt.Print("Skor:    ")

	for i = 0; i < 5; i++ {
		fmt.Print(a.skor[i], " ")
	}
	total := hitungTotal(a.skor)
	fmt.Println("\nTotal Skor:", total)
	fmt.Println("Rekomendasi:", a.rekom) // Gunakan rekomendasi yang tersimpan
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
		for j >= 0 && data[j].tanggal > temp.tanggal { // terlama ke terbaru
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

func urutkanSkorTotalSelection() {
	var i, j, minIdx int

	for i = 0; i < jumlahData-1; i++ {
		minIdx = i
		for j = i + 1; j < jumlahData; j++ {
			if hitungTotal(data[j].skor) < hitungTotal(data[minIdx].skor) {
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

func tampilkanLaporan(nama, hp string) {
	var count, i, rata, jumlahSkor int

	fmt.Println("Laporan untuk:", nama)

	// Tampilkan 5 data terakhir dari belakang ke depan
	count = 0
	for i = jumlahData - 1; i >= 0 && count < 5; i-- {
		if data[i].nama == nama && data[i].hp == hp {
			tampilkanAssessment(data[i])
			count++
		}
	}

	// Hitung rata-rata skor dari data yang tanggalnya dalam 30 hari terakhir
	count = 0
	jumlahSkor = 0
	sekarang := time.Now()

	for i = 0; i < jumlahData; i++ {
		if data[i].nama == nama && data[i].hp == hp {
			waktuData, err := time.Parse("2006-01-02", data[i].tanggal)
			if err == nil {
				selisih := sekarang.Sub(waktuData).Hours() / 24
				if selisih <= 30 {
					jumlahSkor += hitungTotal(data[i].skor)
					count++
				}
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
		fmt.Println("4. Urutkan berdasarkan Tanggal (Insertion Sort)")
		fmt.Println("5. Urutkan berdasarkan Skor (Selection Sort)")
		fmt.Println("6. Tampilkan Laporan Pengguna")
		fmt.Println("7. Edit Data Assessment")
		fmt.Println("8. Hapus Data Assessment")
		fmt.Println("9. Keluar")
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
			urutkanSkorTotalSelection()
			fmt.Println("Data berhasil diurutkan berdasarkan skor (Selection Sort).")
		case 6:
			var nama, hp string
			fmt.Print("Masukkan Nama: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan No HP: ")
			fmt.Scanln(&hp)
			tampilkanLaporan(nama, hp)
		case 7:
			editData()
		case 8:
			hapusData()
		case 9:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
