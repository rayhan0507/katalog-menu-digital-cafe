package main

import "fmt"

const NMAX = 99

type katalogKopi struct {
	namaMinuman string
	harga, id   int
	tersedia    bool
}

type tabKopi [NMAX]katalogKopi

func init() {
	kopi[0] = katalogKopi{"Americano", 20000, 1, true}
	kopi[1] = katalogKopi{"Espresso", 18000, 2, true}
	kopi[2] = katalogKopi{"Cappuccino", 25000, 3, true}
	kopi[3] = katalogKopi{"CafeLatte", 25000, 4, true}
	kopi[4] = katalogKopi{"Mocha", 28000, 5, true}
	kopi[5] = katalogKopi{"Macchiato", 27000, 6, true}
	kopi[6] = katalogKopi{"FlatWhite", 26000, 7, true}
	kopi[7] = katalogKopi{"PiccoloLatte", 24000, 8, true}
	kopi[8] = katalogKopi{"Affogato", 30000, 9, true}
	kopi[9] = katalogKopi{"CaramelMacchiato", 32000, 10, true}
	kopi[10] = katalogKopi{"VanillaLatte", 30000, 11, true}

	nonKopi[0] = katalogKopi{"MatchaLatte", 28000, 21, true}
	nonKopi[1] = katalogKopi{"ChocolateLatte", 25000, 22, true}
	nonKopi[2] = katalogKopi{"RedVelvetLatte", 28000, 23, true}
	nonKopi[3] = katalogKopi{"TaroLatte", 27000, 24, true}
	nonKopi[4] = katalogKopi{"ThaiTea", 22000, 25, true}
	nonKopi[5] = katalogKopi{"GreenTea", 20000, 26, true}
	nonKopi[6] = katalogKopi{"LemonTea", 18000, 27, true}
	nonKopi[7] = katalogKopi{"PeachTea", 22000, 28, true}
	nonKopi[8] = katalogKopi{"LycheeTea", 22000, 29, true}
	nonKopi[9] = katalogKopi{"MilkTea", 23000, 30, true}
	nonKopi[10] = katalogKopi{"StrawberryMilk", 25000, 31, true}

}

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

// buat ngurutin harga minuman
func urutkanHarga(arr *tabKopi, N int) {
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

// buat ngurutin ID
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

// nyari nama minuman
func cariNama(arr tabKopi, N int, target string) {
	var l, r, mid int
	l = 0
	r = N - 1

	for l <= r {
		mid = (l + r) / 2
		if arr[mid].namaMinuman == target {
			fmt.Printf("Ditemukan: %s - Rp%d\n", arr[mid].namaMinuman, arr[mid].harga)
			return
		} else if arr[mid].namaMinuman < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	fmt.Println("Minuman tidak ditemukan.")
}

// nyari minuman berdasarkan harga
func cariHarga(arr tabKopi, N int, target int) {
	var i int
	var ditemukan bool
	ditemukan = false
	for i = 0; i < N; i++ {
		if arr[i].harga == target {
			fmt.Printf("Ditemukan: %s - Rp%d\n", arr[i].namaMinuman, arr[i].harga)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Minuman tidak ditemukan.")
	}
}

func tampilTersedia(arr tabKopi, N int, kategori string) {
	var i int
	var ada bool
	ada = false
	fmt.Printf("Menu %s yang tersedia:\n", kategori)
	for i = 0; i < N; i++ {
		if arr[i].tersedia {
			fmt.Printf("%d. %s - Rp%d\n", arr[i].id, arr[i].namaMinuman, arr[i].harga)
			ada = true
		}
	}
	if !ada {
		fmt.Println("Tidak ada menu yang tersedia.")
	}
}

var kopi tabKopi
var nonKopi tabKopi
var jumlahKopi int = 20
var jumlahNonKopi int = 20

func main() {
	var fitur, pilih int

	for fitur != 5 {
		fmt.Print("\nFitur:\n1. Lihat menu tersedia \n2. Urutkan harga termurah\n3. Cari minuman berdasarkan nama\n4. Cari minuman berdasarkan harga\n5. Keluar\nPilih fitur (1-5): ")
		fmt.Scanln(&fitur)

		if fitur == 1 {
			fmt.Print("Pilih kategori:\n1. Kopi\n2. Non-Kopi\nPilihan: ")
			fmt.Scanln(&pilih)
			if pilih == 1 {
				// Kembalikan ke urutan ID asli sebelum ditampilkan
				urutkanID(&kopi, jumlahKopi)
				tampilTersedia(kopi, jumlahKopi, "Kopi")
			} else if pilih == 2 {
				// Kembalikan ke urutan ID asli sebelum ditampilkan
				urutkanID(&nonKopi, jumlahNonKopi)
				tampilTersedia(nonKopi, jumlahNonKopi, "Non-Kopi")
			} else {
				fmt.Println("Pilihan tidak valid.")
			}

		} else if fitur == 2 {
			fmt.Print("Pilih kategori:\n1. Kopi\n2. Non-Kopi\nPilihan: ")
			fmt.Scanln(&pilih)

			if pilih == 1 {
				urutkanHarga(&kopi, jumlahKopi)
				tampilTersedia(kopi, jumlahKopi, "Kopi")
			} else if pilih == 2 {
				urutkanHarga(&nonKopi, jumlahNonKopi)
				tampilTersedia(nonKopi, jumlahNonKopi, "Non-Kopi")
			} else {
				fmt.Println("Pilihan tidak valid.")
			}

		} else if fitur == 3 {
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

		} else if fitur == 4 {
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

		}
	}
	fmt.Println("Terima kasih!")
}
