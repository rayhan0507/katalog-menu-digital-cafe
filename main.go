package main

import "fmt"

const NMAX = 99

type katalogKopi struct {
	namaMinuman string
	harga, id   int
	tersedia    bool
}

var kopi = [20]katalogKopi{
	{"Americano", 20000, 1, true},
	{"Espresso", 18000, 2, true},
	{"Cappuccino", 25000, 3, true},
	{"Cafe Latte", 25000, 4, true},
	{"Mocha", 28000, 5, true},
	{"Macchiato", 27000, 6, true},
	{"Flat White", 26000, 7, true},
	{"Piccolo Latte", 24000, 8, true},
	{"Affogato", 30000, 9, true},
	{"Caramel Macchiato", 32000, 10, true},
	{"Vanilla Latte", 30000, 11, true},
	{"Hazelnut Latte", 30000, 12, true},
	{"Irish Coffee", 35000, 13, true},
	{"Vietnam Drip Coffee", 28000, 14, true},
	{"Cold Brew", 26000, 15, true},
	{"Nitro Cold Brew", 32000, 16, true},
	{"Coffee Frappe", 33000, 17, true},
	{"Butterscotch Latte", 31000, 18, true},
	{"Spanish Latte", 32000, 19, true},
	{"Kopi Susu Gula Aren", 25000, 20, true},
}

var nonKopi = [20]katalogKopi{
	{"Matcha Latte", 28000, 21, true},
	{"Chocolate Latte", 25000, 22, true},
	{"Red Velvet Latte", 28000, 23, true},
	{"Taro Latte", 27000, 24, true},
	{"Thai Tea", 22000, 25, true},
	{"Green Tea", 20000, 26, true},
	{"Lemon Tea", 18000, 27, true},
	{"Peach Tea", 22000, 28, true},
	{"Lychee Tea", 22000, 29, true},
	{"Milk Tea", 23000, 30, true},
	{"Strawberry Milk", 25000, 31, true},
	{"Banana Milk", 25000, 32, true},
	{"Mango Smoothie", 30000, 33, true},
	{"Avocado Smoothie", 32000, 34, true},
	{"Oreo Milkshake", 30000, 35, true},
	{"Vanilla Milkshake", 28000, 36, true},
	{"Chocolate Milkshake", 30000, 37, true},
	{"Yakult Lychee", 24000, 38, true},
	{"Yakult Strawberry", 24000, 39, true},
	{"Mineral Water", 10000, 40, true},
}

type tabKopi []katalogKopi

func inputKopi(T tabKopi, N int) {

}

func tampilMenuKopi(arr tabKopi, N int) {
	var i int
	for i = 0; i < N; i++ {
		fmt.Printf("%d. %s - Rp%d\n", arr[i].id, arr[i].namaMinuman, arr[i].harga)
	}
}

func tampilMenuNonKopi(arr tabKopi, N int) {
	var i int
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
	var l, r, mid int
	l = 0
	r = N - 1
	mid = (l + r) / 2

	for l <= r {
		if T[mid].namaMinuman == target {
			fmt.Printf("Minuman ditemukan: %s - Rp%d\n", T[mid].namaMinuman, T[mid].harga)
			return
		} else if T[mid].namaMinuman < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
		mid = (l + r) / 2
	}
	fmt.Println("Minuman tidak ditemukan")
}

func cariHarga(T tabKopi, N int, t int) {
	var i int
	for i = 0; i < N; i++ {
		if T[i].harga == t {
			fmt.Printf("%s - Rp%d\n", T[i].namaMinuman, T[i].harga)
		}
	}
}

func main() {
	var n, pilih, fitur int
	var jumlahKopi, jumlahNonkopi int
	var arr tabKopi
	fmt.Print("fitur: 1.Beli menu\n2. cari berdasarkan nama \n3. cari berdasarkan harga\n4. urutkan harga termurah\n5. Update Menu\n6.Hapus Menu\n Pilih fitur yang ingin anda gunakan:(1-6)\n")
	if fitur == 1 {
		fmt.Print("Menu cafe kami:")
		fmt.Print("\n1. Kopi\n2. Non-Kopi\nPilih menu yang ingin anda lihat:\n")
		fmt.Scanln(&pilih)
		for true {
			if pilih == 1 {
				tampilMenuKopi(arr, jumlahKopi)
				fmt.Print("Masukan jumlah kopi yang mau anda beli:\n")
				inputKopi(arr, n)
			} else if pilih == 2 {
				tampilMenuNonKopi(arr, jumlahNonkopi)
				fmt.Print("Masukan jumlah non-kopi yang mau anda beli:\n")
				inputKopi(arr, n)
			} else {
				break
			}
		}
	} else if fitur == 2 {
		var target string
		fmt.Print("Masukan nama minuman yang ingin anda cari:")
		fmt.Scan(&target)
		cariNama(arr, n, target)
	} else if fitur == 3 {
		var targetHarga int
		fmt.Scan(&targetHarga)
		cariHarga(arr, n, targetHarga)
		fmt.Print("Masukan harga minuman yang ingin anda cari:")

	} else if fitur == 4 {
		urutkanHarga(&arr, n)
		fmt.Print("Menu minuman setelah diurutkan berdasarkan harga termurah:\n")
		tampilMenuKopi(arr, jumlahKopi)
		tampilMenuNonKopi(arr, jumlahNonkopi)
	}
}
