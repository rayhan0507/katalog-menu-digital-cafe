package main

import "fmt"

const NMAX = 20

type katalogKopi struct {
	namaMinuman string
	harga, id   int
	bahan       string
	tersedia    bool
}

type tabKopi [NMAX]katalogKopi

func init() {
	kopi[0] = katalogKopi{"Americano", 20000, 1, "Espresso,Air", true}
	kopi[1] = katalogKopi{"Espresso", 18000, 2, "KopiEspresso", true}
	kopi[2] = katalogKopi{"Cappuccino", 25000, 3, "Espresso,SusuFrothed", true}
	kopi[3] = katalogKopi{"CafeLatte", 25000, 4, "Espresso,SusuSteamed", true}
	kopi[4] = katalogKopi{"Mocha", 28000, 5, "Espresso,Coklat,Susu", true}
	kopi[5] = katalogKopi{"Macchiato", 27000, 6, "Espresso,Susu", true}
	kopi[6] = katalogKopi{"FlatWhite", 26000, 7, "Espresso,SusuMicrofoam", true}
	kopi[7] = katalogKopi{"PiccoloLatte", 24000, 8, "Espresso,Susu", true}
	kopi[8] = katalogKopi{"Affogato", 30000, 9, "Espresso,EsKrimVanilla", true}
	kopi[9] = katalogKopi{"CaramelMacchiato", 32000, 10, "Espresso,Susu,Karamel", true}
	kopi[10] = katalogKopi{"VanillaLatte", 30000, 11, "Espresso,Susu,SirupVanilla", true}

	nonKopi[0] = katalogKopi{"MatchaLatte", 28000, 21, "BubukMatcha,Susu", true}
	nonKopi[1] = katalogKopi{"ChocolateLatte", 25000, 22, "Coklat,Susu", true}
	nonKopi[2] = katalogKopi{"RedVelvetLatte", 28000, 23, "BubukRedVelvet,Susu", true}
	nonKopi[3] = katalogKopi{"TaroLatte", 27000, 24, "BubukTaro,Susu", true}
	nonKopi[4] = katalogKopi{"ThaiTea", 22000, 25, "TehThai,Susu,Gula", true}
	nonKopi[5] = katalogKopi{"GreenTea", 20000, 26, "TehHijau,Gula", true}
	nonKopi[6] = katalogKopi{"LemonTea", 18000, 27, "Teh,Lemon,Gula", true}
	nonKopi[7] = katalogKopi{"PeachTea", 22000, 28, "Teh,SirupPeach", true}
	nonKopi[8] = katalogKopi{"LycheeTea", 22000, 29, "Teh,SirupLychee", true}
	nonKopi[9] = katalogKopi{"MilkTea", 23000, 30, "Teh,Susu,Gula", true}
	nonKopi[10] = katalogKopi{"StrawberryMilk", 25000, 31, "Susu,SirupStrawberry", true}
}

var kopi tabKopi
var nonKopi tabKopi
var jumlahKopi int = 11
var jumlahNonKopi int = 11

// ===================== TAMPIL =====================

func tampilTersedia(arr tabKopi, N int, kategori string) {
	var i int
	var ada bool
	ada = false
	fmt.Printf("Menu %s yang tersedia:\n", kategori)
	for i = 0; i < N; i++ {
		if arr[i].tersedia {
			fmt.Printf("%d. %s - Rp%d (Bahan: %s)\n", arr[i].id, arr[i].namaMinuman, arr[i].harga, arr[i].bahan)
			ada = true
		}
	}
	if !ada {
		fmt.Println("Tidak ada menu yang tersedia.")
	}
}

// ===================== SORTING =====================

// Selection Sort berdasarkan nama (dipakai sebelum Binary Search)
func urutkanNama(arr *tabKopi, N int) {
	var i, j, minIndex int
	var temp katalogKopi
	for i = 0; i < N-1; i++ {
		minIndex = i
		for j = i + 1; j < N; j++ {
			if (*arr)[j].namaMinuman < (*arr)[minIndex].namaMinuman {
				minIndex = j
			}
		}
		temp = (*arr)[i]
		(*arr)[i] = (*arr)[minIndex]
		(*arr)[minIndex] = temp
	}
}

// Insertion Sort berdasarkan ID (buat ngembaliin urutan tampilan default)
func urutkanID(arr *tabKopi, N int) {
	var i, j int
	var key katalogKopi
	for i = 1; i < N; i++ {
		key = (*arr)[i]
		j = i - 1
		for j >= 0 && (*arr)[j].id > key.id {
			(*arr)[j+1] = (*arr)[j]
			j--
		}
		(*arr)[j+1] = key
	}
}

// Selection Sort berdasarkan harga (poin d - termurah ke termahal)
func urutkanHargaSelection(arr *tabKopi, N int) {
	var i, j, minIndex int
	var temp katalogKopi
	for i = 0; i < N-1; i++ {
		minIndex = i
		for j = i + 1; j < N; j++ {
			if (*arr)[j].harga < (*arr)[minIndex].harga {
				minIndex = j
			}
		}
		temp = (*arr)[i]
		(*arr)[i] = (*arr)[minIndex]
		(*arr)[minIndex] = temp
	}
}

// Insertion Sort berdasarkan harga (poin d - termurah ke termahal)
func urutkanHargaInsertion(arr *tabKopi, N int) {
	var i, j int
	var key katalogKopi
	for i = 1; i < N; i++ {
		key = (*arr)[i]
		j = i - 1
		for j >= 0 && (*arr)[j].harga > key.harga {
			(*arr)[j+1] = (*arr)[j]
			j--
		}
		(*arr)[j+1] = key
	}
}

// ===================== SEARCHING =====================

// Binary Search berdasarkan nama (array harus sudah diurutkan dulu)
func cariNama(arr tabKopi, N int, target string) {
	var l, r, mid int
	l = 0
	r = N - 1

	for l <= r {
		mid = (l + r) / 2
		if arr[mid].namaMinuman == target {
			fmt.Printf("Ditemukan: %s - Rp%d (Bahan: %s)\n", arr[mid].namaMinuman, arr[mid].harga, arr[mid].bahan)
			return
		} else if arr[mid].namaMinuman < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	fmt.Println("Minuman tidak ditemukan.")
}

// Sequential Search berdasarkan harga
func cariHarga(arr tabKopi, N int, target int) {
	var i int
	var ditemukan bool
	ditemukan = false
	for i = 0; i < N; i++ {
		if arr[i].harga == target {
			fmt.Printf("Ditemukan: %s - Rp%d (Bahan: %s)\n", arr[i].namaMinuman, arr[i].harga, arr[i].bahan)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Minuman tidak ditemukan.")
	}
}

// ===================== CRUD MENU (poin a & b) =====================

func tambahMenu(arr *tabKopi, N *int, nama string, harga int, bahan string) {
	if *N >= NMAX {
		fmt.Println("Kapasitas menu sudah penuh, tidak bisa menambah lagi.")
		return
	}
	var i, maxID int
	maxID = 0
	for i = 0; i < *N; i++ {
		if (*arr)[i].id > maxID {
			maxID = (*arr)[i].id
		}
	}
	(*arr)[*N] = katalogKopi{nama, harga, maxID + 1, bahan, true}
	*N++
	fmt.Printf("Menu '%s' berhasil ditambahkan dengan ID %d.\n", nama, maxID+1)
}

func ubahMenu(arr *tabKopi, N int, target string) {
	var i int
	var ketemu bool
	ketemu = false
	for i = 0; i < N; i++ {
		if (*arr)[i].namaMinuman == target {
			ketemu = true
			var hargaBaru, statusBaru int
			var bahanBaru string

			fmt.Print("Masukkan harga baru: Rp")
			fmt.Scanln(&hargaBaru)
			fmt.Print("Masukkan komposisi bahan baru (pisah dengan koma, tanpa spasi): ")
			fmt.Scanln(&bahanBaru)
			fmt.Print("Status tersedia? (1=Ya, 0=Tidak): ")
			fmt.Scanln(&statusBaru)

			(*arr)[i].harga = hargaBaru
			(*arr)[i].bahan = bahanBaru
			(*arr)[i].tersedia = statusBaru == 1
			fmt.Println("Menu berhasil diubah.")
			break
		}
	}
	if !ketemu {
		fmt.Println("Menu tidak ditemukan.")
	}
}

func hapusMenu(arr *tabKopi, N *int, target string) {
	var i, idx int
	idx = -1
	for i = 0; i < *N; i++ {
		if (*arr)[i].namaMinuman == target {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println("Menu tidak ditemukan.")
		return
	}
	for i = idx; i < *N-1; i++ {
		(*arr)[i] = (*arr)[i+1]
	}
	*N--
	fmt.Println("Menu berhasil dihapus.")
}

// ===================== STATISTIK (poin e) =====================

func tampilStatistik(arr tabKopi, N int, kategori string) {
	var i, jumlahTersedia, totalHarga int
	jumlahTersedia = 0
	totalHarga = 0
	for i = 0; i < N; i++ {
		if arr[i].tersedia {
			jumlahTersedia++
			totalHarga += arr[i].harga
		}
	}
	fmt.Printf("\nStatistik kategori %s:\n", kategori)
	fmt.Printf("Jumlah menu tersedia : %d\n", jumlahTersedia)
	if jumlahTersedia > 0 {
		fmt.Printf("Rata-rata harga      : Rp%.2f\n", float64(totalHarga)/float64(jumlahTersedia))
	} else {
		fmt.Println("Rata-rata harga      : -")
	}
}

// ===================== MAIN =====================

func main() {
	var fitur, pilih int

	for fitur != 9 {
		fmt.Print("\nFitur:\n1. Lihat menu tersedia\n2. Tambah menu baru\n3. Ubah data menu\n4. Hapus menu\n5. Urutkan harga termurah-termahal\n6. Cari minuman berdasarkan nama\n7. Cari minuman berdasarkan harga\n8. Statistik menu\n9. Keluar\nPilih fitur (1-9): ")
		fmt.Scanln(&fitur)

		if fitur == 1 {
			fmt.Print("Pilih kategori:\n1. Kopi\n2. Non-Kopi\nPilihan: ")
			fmt.Scanln(&pilih)
			if pilih == 1 {
				urutkanID(&kopi, jumlahKopi)
				tampilTersedia(kopi, jumlahKopi, "Kopi")
			} else if pilih == 2 {
				urutkanID(&nonKopi, jumlahNonKopi)
				tampilTersedia(nonKopi, jumlahNonKopi, "Non-Kopi")
			} else {
				fmt.Println("Pilihan tidak valid.")
			}

		} else if fitur == 2 {
			var nama, bahan string
			var harga int
			fmt.Print("Pilih kategori:\n1. Kopi\n2. Non-Kopi\nPilihan: ")
			fmt.Scanln(&pilih)
			fmt.Print("Nama menu baru: ")
			fmt.Scanln(&nama)
			fmt.Print("Harga: Rp")
			fmt.Scanln(&harga)
			fmt.Print("Komposisi bahan (pisah koma, tanpa spasi): ")
			fmt.Scanln(&bahan)

			if pilih == 1 {
				tambahMenu(&kopi, &jumlahKopi, nama, harga, bahan)
			} else if pilih == 2 {
				tambahMenu(&nonKopi, &jumlahNonKopi, nama, harga, bahan)
			} else {
				fmt.Println("Pilihan tidak valid.")
			}

		} else if fitur == 3 {
			var target string
			fmt.Print("Pilih kategori:\n1. Kopi\n2. Non-Kopi\nPilihan: ")
			fmt.Scanln(&pilih)
			fmt.Print("Nama menu yang diubah: ")
			fmt.Scanln(&target)

			if pilih == 1 {
				ubahMenu(&kopi, jumlahKopi, target)
			} else if pilih == 2 {
				ubahMenu(&nonKopi, jumlahNonKopi, target)
			} else {
				fmt.Println("Pilihan tidak valid.")
			}

		} else if fitur == 4 {
			var target string
			fmt.Print("Pilih kategori:\n1. Kopi\n2. Non-Kopi\nPilihan: ")
			fmt.Scanln(&pilih)
			fmt.Print("Nama menu yang dihapus: ")
			fmt.Scanln(&target)

			if pilih == 1 {
				hapusMenu(&kopi, &jumlahKopi, target)
			} else if pilih == 2 {
				hapusMenu(&nonKopi, &jumlahNonKopi, target)
			} else {
				fmt.Println("Pilihan tidak valid.")
			}

		} else if fitur == 5 {
			var metode int
			fmt.Print("Pilih kategori:\n1. Kopi\n2. Non-Kopi\nPilihan: ")
			fmt.Scanln(&pilih)
			fmt.Print("Metode pengurutan:\n1. Selection Sort\n2. Insertion Sort\nPilihan: ")
			fmt.Scanln(&metode)

			if pilih == 1 {
				if metode == 1 {
					urutkanHargaSelection(&kopi, jumlahKopi)
				} else if metode == 2 {
					urutkanHargaInsertion(&kopi, jumlahKopi)
				} else {
					fmt.Println("Metode tidak valid.")
					continue
				}
				tampilTersedia(kopi, jumlahKopi, "Kopi")
			} else if pilih == 2 {
				if metode == 1 {
					urutkanHargaSelection(&nonKopi, jumlahNonKopi)
				} else if metode == 2 {
					urutkanHargaInsertion(&nonKopi, jumlahNonKopi)
				} else {
					fmt.Println("Metode tidak valid.")
					continue
				}
				tampilTersedia(nonKopi, jumlahNonKopi, "Non-Kopi")
			} else {
				fmt.Println("Pilihan tidak valid.")
			}

		} else if fitur == 6 {
			var target string
			fmt.Print("Pilih kategori:\n1. Kopi\n2. Non-Kopi\nPilihan: ")
			fmt.Scanln(&pilih)
			fmt.Print("Masukkan nama minuman: ")
			fmt.Scanln(&target)

			if pilih == 1 {
				urutkanNama(&kopi, jumlahKopi)
				cariNama(kopi, jumlahKopi, target)
			} else if pilih == 2 {
				urutkanNama(&nonKopi, jumlahNonKopi)
				cariNama(nonKopi, jumlahNonKopi, target)
			} else {
				fmt.Println("Pilihan tidak valid.")
			}

		} else if fitur == 7 {
			var targetHarga int
			fmt.Print("Pilih kategori:\n1. Kopi\n2. Non-Kopi\nPilihan: ")
			fmt.Scanln(&pilih)
			fmt.Print("Masukkan harga yang dicari: Rp")
			fmt.Scanln(&targetHarga)
			if pilih == 1 {
				cariHarga(kopi, jumlahKopi, targetHarga)
			} else if pilih == 2 {
				cariHarga(nonKopi, jumlahNonKopi, targetHarga)
			} else {
				fmt.Println("Pilihan tidak valid.")
			}

		} else if fitur == 8 {
			tampilStatistik(kopi, jumlahKopi, "Kopi")
			tampilStatistik(nonKopi, jumlahNonKopi, "Non-Kopi")
		}
	}
	fmt.Println("Terima kasih!")
}
