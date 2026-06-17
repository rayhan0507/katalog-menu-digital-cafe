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
	kopi[3] = katalogKopi{"Cafe Latte", 25000, 4, true}
	kopi[4] = katalogKopi{"Mocha", 28000, 5, true}
	kopi[5] = katalogKopi{"Macchiato", 27000, 6, true}
	kopi[6] = katalogKopi{"Flat White", 26000, 7, true}
	kopi[7] = katalogKopi{"Piccolo Latte", 24000, 8, true}
	kopi[8] = katalogKopi{"Affogato", 30000, 9, true}
	kopi[9] = katalogKopi{"Caramel Macchiato", 32000, 10, true}
	kopi[10] = katalogKopi{"Vanilla Latte", 30000, 11, true}
	kopi[11] = katalogKopi{"Hazelnut Latte", 30000, 12, true}
	kopi[12] = katalogKopi{"Irish Coffee", 35000, 13, true}
	kopi[13] = katalogKopi{"Vietnam Drip Coffee", 28000, 14, true}
	kopi[14] = katalogKopi{"Cold Brew", 26000, 15, true}
	kopi[15] = katalogKopi{"Nitro Cold Brew", 32000, 16, true}
	kopi[16] = katalogKopi{"Coffee Frappe", 33000, 17, true}
	kopi[17] = katalogKopi{"Butterscotch Latte", 31000, 18, true}
	kopi[18] = katalogKopi{"Spanish Latte", 32000, 19, true}
	kopi[19] = katalogKopi{"Kopi Susu Gula Aren", 25000, 20, true}

	nonKopi[0] = katalogKopi{"Matcha Latte", 28000, 21, true}
	nonKopi[1] = katalogKopi{"Chocolate Latte", 25000, 22, true}
	nonKopi[2] = katalogKopi{"Red Velvet Latte", 28000, 23, true}
	nonKopi[3] = katalogKopi{"Taro Latte", 27000, 24, true}
	nonKopi[4] = katalogKopi{"Thai Tea", 22000, 25, true}
	nonKopi[5] = katalogKopi{"Green Tea", 20000, 26, true}
	nonKopi[6] = katalogKopi{"Lemon Tea", 18000, 27, true}
	nonKopi[7] = katalogKopi{"Peach Tea", 22000, 28, true}
	nonKopi[8] = katalogKopi{"Lychee Tea", 22000, 29, true}
	nonKopi[9] = katalogKopi{"Milk Tea", 23000, 30, true}
	nonKopi[10] = katalogKopi{"Strawberry Milk", 25000, 31, true}
	nonKopi[11] = katalogKopi{"Banana Milk", 25000, 32, true}
	nonKopi[12] = katalogKopi{"Mango Smoothie", 30000, 33, true}
	nonKopi[13] = katalogKopi{"Avocado Smoothie", 32000, 34, true}
	nonKopi[14] = katalogKopi{"Oreo Milkshake", 30000, 35, true}
	nonKopi[15] = katalogKopi{"Vanilla Milkshake", 28000, 36, true}
	nonKopi[16] = katalogKopi{"Chocolate Milkshake", 30000, 37, true}
	nonKopi[17] = katalogKopi{"Yakult Lychee", 24000, 38, true}
	nonKopi[18] = katalogKopi{"Yakult Strawberry", 24000, 39, true}
	nonKopi[19] = katalogKopi{"Mineral Water", 10000, 40, true}
}

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

// Fungsi untuk mengembalikan urutan berdasarkan ID asli menu
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

func cariNama(arr tabKopi, N int, target string) {
	var i int
	var ditemukan bool
	ditemukan = false
	for i = 0; i < N; i++ {
		if arr[i].namaMinuman == target {
			fmt.Printf("Ditemukan: %s - Rp%d\n", arr[i].namaMinuman, arr[i].harga)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Minuman tidak ditemukan.")
	}
}

func cariHarga(arr tabKopi, N int, target int) {
	var l, r, mid int
	l = 0
	r = N - 1
	for l <= r {
		mid = (l + r) / 2
		if arr[mid].harga == target {
			fmt.Printf("Ditemukan: %s - Rp%d\n", arr[mid].namaMinuman, arr[mid].harga)
			return
		} else if arr[mid].harga < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	fmt.Println("Minuman dengan harga tersebut tidak ditemukan.")
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
				cariNama(kopi, jumlahKopi, target)
			} else if pilih == 2 {
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
				urutkanHarga(&kopi, jumlahKopi)
				cariHarga(kopi, jumlahKopi, targetHarga)
			} else if pilih == 2 {
				urutkanHarga(&nonKopi, jumlahNonKopi)
				cariHarga(nonKopi, jumlahNonKopi, targetHarga)
			} else {
				fmt.Println("Pilihan tidak valid.")
			}

		}
	}
	fmt.Println("Terima kasih!")
}
