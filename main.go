package main

import "fmt"

const NMAX = 99

type katalogKopi struct {
	namaMinuman string
	harga, id   int
	tersedia    bool
}

type tabKopi [NMAX]katalogKopi
type keranjang [NMAX]katalogKopi

var kopi tabKopi
var nonKopi tabKopi
var jumlahKopi int = 20
var jumlahNonKopi int = 20

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

func inputKopiByID(T *keranjang, jumlah int, katalog tabKopi, nKatalog int) {
	var i, id int
	i = 0
	for i < jumlah {
		fmt.Printf("Masukkan ID minuman ke-%d: ", i+1)
		fmt.Scan(&id)
		var j int
		var ditemukan bool
		ditemukan = false
		for j = 0; j < nKatalog; j++ {
			if katalog[j].id == id {
				T[i] = katalog[j]
				fmt.Printf("Ditambahkan: %s\n", katalog[j].namaMinuman)
				ditemukan = true
				break
			}
		}
		if !ditemukan {
			fmt.Println("ID tidak ditemukan, coba lagi!")
		} else {
			i++
		}
	}
}

func tampilMenuKopi(arr tabKopi, N int) {
	var i int
	fmt.Println("Menu Kopi:")
	for i = 0; i < N; i++ {
		fmt.Printf("%d. %s - Rp%d\n", arr[i].id, arr[i].namaMinuman, arr[i].harga)
	}
}

func tampilMenuNonKopi(arr tabKopi, N int) {
	var i int
	fmt.Println("Menu Non Kopi:")
	for i = 0; i < N; i++ {
		fmt.Printf("%d. %s - Rp%d\n", arr[i].id, arr[i].namaMinuman, arr[i].harga)
	}
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

func cariNama(T tabKopi, N int, target string) {
	var i int
	var ditemukan bool
	ditemukan = false
	for i = 0; i < N; i++ {
		if T[i].namaMinuman == target {
			fmt.Printf("Ditemukan: %s - Rp%d\n", T[i].namaMinuman, T[i].harga)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Minuman tidak ditemukan")
	}
}

func cariHarga(T tabKopi, N int, t int) {
	var l, r, mid int
	l = 0
	r = N - 1
	mid = (l + r) / 2

	for l <= r {
		if T[mid].harga == t {
			fmt.Printf("Harga minuman ditemukan: %s - Rp%d\n", T[mid].namaMinuman, T[mid].harga)
			return
		} else if T[mid].harga < t {
			l = mid + 1
		} else {
			r = mid - 1
		}
		mid = (l + r) / 2
	}
	fmt.Println("Harga minuman tidak ditemukan")
}

func tampilKeranjang(T keranjang, N int) {
	var i int
	var total int
	fmt.Println("=== Keranjang Belanja ===")
	for i = 0; i < N; i++ {
		fmt.Printf("%s - Rp%d\n", T[i].namaMinuman, T[i].harga)
		total += T[i].harga
	}
	fmt.Printf("Total: Rp%d\n", total)
}

func main() {
	var pilih, fitur int
	var jumlahBeli int
	var arr keranjang
	var nKeranjang int

	for fitur != 7 {
		fmt.Print("\nFitur:\n1. Beli menu\n2. Cari berdasarkan nama\n3. Cari berdasarkan harga\n4. Urutkan harga termurah\n5. Lihat keranjang\n6. Hapus Menu\n7. Keluar\nPilih fitur (1-7): ")
		fmt.Scanln(&fitur)

		if fitur == 1 {
			fmt.Print("\n1. Kopi\n2. Non-Kopi\nPilih kategori: ")
			fmt.Scanln(&pilih)

			switch pilih {
			case 1:
				tampilMenuKopi(kopi, jumlahKopi)
				fmt.Print("Masukan jumlah kopi yang mau anda beli: ")
				fmt.Scan(&jumlahBeli)
				inputKopiByID(&arr, jumlahBeli, kopi, jumlahKopi)
				nKeranjang += jumlahBeli

			case 2:
				tampilMenuNonKopi(nonKopi, jumlahNonKopi)
				fmt.Print("Masukan jumlah non-kopi yang mau anda beli: ")
				fmt.Scan(&jumlahBeli)
				inputKopiByID(&arr, jumlahBeli, nonKopi, jumlahNonKopi)
				nKeranjang += jumlahBeli

			default:
				fmt.Println("Pilihan tidak valid")
			}

		} else if fitur == 2 {
			var target string
			fmt.Print("Masukan nama minuman yang ingin anda cari: ")
			fmt.Scan(&target)
			cariNama(kopi, jumlahKopi, target)
			cariNama(nonKopi, jumlahNonKopi, target)

		} else if fitur == 3 {
			var targetHarga int
			fmt.Print("Masukan harga minuman yang ingin anda cari: ")
			fmt.Scan(&targetHarga)
			cariHarga(kopi, jumlahKopi, targetHarga)
			cariHarga(nonKopi, jumlahNonKopi, targetHarga)

		} else if fitur == 4 {
			urutkanHarga(&kopi, jumlahKopi)
			urutkanHarga(&nonKopi, jumlahNonKopi)
			fmt.Print("Menu setelah diurutkan berdasarkan harga termurah:\n")
			tampilMenuKopi(kopi, jumlahKopi)
			tampilMenuNonKopi(nonKopi, jumlahNonKopi)

		} else if fitur == 5 {
			tampilKeranjang(arr, nKeranjang)
		}
	}
	fmt.Println("Terima kasih!")
}
