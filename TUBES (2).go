package main

import "fmt"

// Struct untuk menyimpan data pendapatan
type Pendapatan struct {
	Sumber   string
	Jumlah   float64
	Tanggal  string
	Kategori string
}

type ID struct {
	Username       string
	PassUser       string
	Kategori       string
	Target         int
	dataPendapatan arrPendapatan
}
type detail struct {
	nama, deskripsi       string
	jumlah                int
	tanggal, bulan, tahun int
}

const NMAX int = 50

type arrID [NMAX]ID
type arrPendapatan [NMAX]detail

func main() {
	var pilih, nArrayPengguna int
	var arrayPengguna arrID

	for {
		menu()
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			menuRegister(&arrayPengguna, &nArrayPengguna)
		case 2:
			menuLogin(&arrayPengguna, &nArrayPengguna)
		case 3:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

//fungsi untuk tampilan login
func tampilkanDashboard(user *ID) {
	var pilihan int
	for pilihan != 5 {
		fmt.Println("\n=== Aplikasi Pencatatan Pendapatan ===")
		fmt.Println("1. Tambah Pendapatan Baru")
		fmt.Println("2. Tampilkan Laporan Keuangan")
		fmt.Println("3. Tampilkan Analisis Terkini")
		fmt.Println("4. Tampilkan Progres Target")
		fmt.Println("5. Logout")
		fmt.Println("Pilih menu: (1/2/3/4/5): ")

		n := hitungJumlahPendapatan(user)

		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			tambahPendapatan(user, &n)
		case 2:
			laporan(user, &n)
		case 3:
			TampilanAnalisis(user, &n)
		case 4:
			TampilanProgres(user, &n)
		}
	}
}

func hitungJumlahPendapatan(user *ID) int {
	// hitung jumlah data pendapatan yang sudah diisi (jumlah != 0)
	count := 0
	for i := 0; i < NMAX; i++ {
		if user.dataPendapatan[i].jumlah != 0 {
			count++
		}
	}
	return count
}

//fungsi untuk tampilan utama menu
func menu() {
	fmt.Println("========== Selamat Datang =========")
	fmt.Println("======= Di Aplikasi Pelacak =======")
	fmt.Println("== Side Hustle dan Pssive Income ==")
	fmt.Println("Daftar Menu")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")
	fmt.Printf("%s ", "Pilih Menu (1/2/3): ")
}

//fungsi untuk menambahkan pengguna/register
func menuRegister(arrayPengguna *arrID, nArrayPengguna *int) {
	var namaUser, Kategori, PassUser string
	var TargetUsser int

	fmt.Println("Nama Pengguna:")
	fmt.Scan(&namaUser)

	for i := 0; i < *nArrayPengguna; i++ {
		if namaUser == arrayPengguna[i].Username {
			fmt.Println("Nama sudah dipakai user lain, masukkan nama lain")
			fmt.Scan(&namaUser)
			i = -1 // reset ulang pengecekan
		}
	}

	fmt.Println("Password Akun:")
	fmt.Println("Pastikan password terdiri angka, simbol dan huruf besar")
	fmt.Scan(&PassUser)
	fmt.Println("Profesi (Pelajar/Mahasiswa/Pekerja, dll):")
	fmt.Scan(&Kategori)
	fmt.Println("Masukan Target Pencapaian Pendapatan (Input Target Hanya Angka): ")
	fmt.Scan(&TargetUsser)

	if *nArrayPengguna < NMAX {
		arrayPengguna[*nArrayPengguna].Username = namaUser
		arrayPengguna[*nArrayPengguna].PassUser = PassUser
		arrayPengguna[*nArrayPengguna].Kategori = Kategori
		arrayPengguna[*nArrayPengguna].Target = TargetUsser
		fmt.Println("Pengguna berhasil didaftarkan!")
		(*nArrayPengguna)++
	} else {
		fmt.Println("Kapasitas Pengguna Sudah Penuh")
	}
}

func menuLogin(arrayPengguna *arrID, nArrayPengguna *int) {
	fmt.Println("Masukkan Username:")
	var nama, password string
	fmt.Scan(&nama)
	fmt.Println("Masukan Password: ")
	fmt.Scan(&password)
	var i int = 0
	var valid bool = false
	for i < *nArrayPengguna && !valid {
		if nama == arrayPengguna[i].Username && password == arrayPengguna[i].PassUser {
			fmt.Println("Selamat! Anda berhasil login.")
			fmt.Print(arrayPengguna[i].Target)
			tampilkanDashboard(&arrayPengguna[i])
			valid = true
		}
		i++
	}
	fmt.Println("Username atau Password salah.")
}

// Fungsi untuk menambahkan data pendapatan
func tambahPendapatan(array *ID, n *int) {

	fmt.Println("\n--- Tambah Data Pendapatan ---")

	fmt.Println("\nJenis Pendapatan: ")
	fmt.Println("1. Side Hustle")
	fmt.Println("2. Passive Income")
	fmt.Println("Pilih (1/2) :")

	var pilihan, jenisSideHustle, jenisPassive int
	if *n < NMAX {

		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			fmt.Println("Jenis Side Hustle")
			fmt.Println("1. Freelance")
			fmt.Println("2. Konten Kreator")
			fmt.Println("3. Jualan Online")
			fmt.Println("Pilih (1/2/3): ")
			fmt.Scan(&jenisSideHustle)

			switch jenisSideHustle {
			case 1:
				array.dataPendapatan[*n].nama = "Freelance"
				KetSideHustle(array, n)
			case 2:
				array.dataPendapatan[*n].nama = "Konten Kreator"
				KetSideHustle(array, n)
			case 3:
				array.dataPendapatan[*n].nama = "Jualan Online"
				KetSideHustle(array, n)

			}

		case 2:
			fmt.Println("Jenis Passive Income")
			fmt.Println("1. Uang Sewa Lahan")
			fmt.Println("2. Royalti")
			fmt.Println("Pilih (1/2): ")
			fmt.Scan(&jenisPassive)

			switch jenisPassive {
			case 1:
				array.dataPendapatan[*n].nama = "Uang Sewa Lahan"
				KetSideHustle(array, n)
			case 2:
				array.dataPendapatan[*n].nama = "Royalti"
				KetSideHustle(array, n)
			}

		}
		(*n)++
	} else {
		fmt.Println("Data pendapatan sudah penuh.")
	}
}

// fungsi laporan keuangan lengkap dengan pilihan laporan bulanan dan tahunan
func laporan(array *ID, n *int) {

	var bulan, tahun int

	fmt.Println("===============")
	fmt.Println("Masukkan Bulan")
	fmt.Scan(&bulan)
	fmt.Println("Masukkan Tahun")
	fmt.Scan(&tahun)
	sortPendapatanByDate(array, n)
	for i := 0; i < *n; i++ {
		if array.dataPendapatan[i].tahun == array.dataPendapatan[i].tahun && array.dataPendapatan[i].bulan == array.dataPendapatan[i].bulan {
			fmt.Printf("- %s: %d (Tanggal %d)\n", array.dataPendapatan[i].nama, array.dataPendapatan[i].jumlah, array.dataPendapatan[i].tanggal)
		}
	}
}

// Selection sort data pendapatan berdasarkan tanggal (tahun, bulan, tanggal)
func sortPendapatanByDate(array *ID, n *int) {

	var i, idx, pass int
	var temp detail
	pass = 1
	for pass < *n {
		idx = pass - 1
		i = pass
		for i < *n {
			if array.dataPendapatan[i].tahun <= array.dataPendapatan[idx].tahun && array.dataPendapatan[i].bulan <= array.dataPendapatan[idx].bulan && array.dataPendapatan[i].tanggal < array.dataPendapatan[idx].tanggal {
				idx = i
			}
			i++
		}
		temp = array.dataPendapatan[pass-1]
		array.dataPendapatan[pass-1] = array.dataPendapatan[idx]
		array.dataPendapatan[idx] = temp
		pass = pass + 1
	}
}

// Fungsi untuk menampilkan analisis (sementara kosong)
func TampilanAnalisis(array *ID, n *int) {

	fmt.Println("\n=== Berikut Hasil Analisis Sumber Pendapatan Terbesarmu ===")

	if *n == 0 {
		fmt.Println("Belum ada data pendapatan.")
		return
	}

	// Urutkan pendapatan berdasarkan jumlah (dari kecil ke besar)
	insertionSortPendapatan(array, n)

	// Data terakhir setelah sorting ascending adalah pendapatan terbesar
	terbesar := array.dataPendapatan[*n-1]

	fmt.Printf("Kategori: %s\n", terbesar.nama)
	fmt.Printf("Jumlah: %d\n", terbesar.jumlah)
	fmt.Printf("Tanggal: %d-%d-%d\n", terbesar.tanggal, terbesar.bulan, terbesar.tahun)
	fmt.Printf("Deskripsi: %s\n", terbesar.deskripsi)
}

func insertionSortPendapatan(array *ID, n *int) {
	for pass := 1; pass < *n; pass++ {
		temp := array.dataPendapatan[pass]
		i := pass
		for i > 0 && temp.jumlah < array.dataPendapatan[i-1].jumlah {
			array.dataPendapatan[i] = array.dataPendapatan[i-1]
			i--
		}
		array.dataPendapatan[i] = temp
	}
}

// Fungsi untuk menampilkan progres pencapaian target
func TampilanProgres(array *ID, n *int) {
	fmt.Printf("Target Pendapatan Anda Bulan ini : %d\n", array.Target)
	fmt.Printf("Persentase Target Tercapai Bulan ini sebesar %.2f%%\n", hitungProgress(array, n))
}

func hitungProgress(array *ID, n *int) float64 {
	var total int
	for i := 0; i < *n; i++ {
		total += array.dataPendapatan[i].jumlah
	}
	if array.Target == 0 {
		return 0
	}
	return (float64(total) / float64(array.Target)) * 100
}

func KetSideHustle(array *ID, n *int) {
	fmt.Println("Masukkan Nominal (tanpa tanda titik, contoh: 10000 bukan 10.000):")
	fmt.Scan(&array.dataPendapatan[*n].jumlah)

	fmt.Println("Tanggal Pemasukan")
	fmt.Println("Masukan Tanggal:")
	fmt.Scan(&array.dataPendapatan[*n].tanggal)

	fmt.Println("Masukan Bulan:")
	fmt.Scan(&array.dataPendapatan[*n].bulan)

	fmt.Println("Masukan Tahun:")
	fmt.Scan(&array.dataPendapatan[*n].tahun)

	fmt.Println("Deskripsi Singkat Pendapatanmu: ")
	fmt.Scan(&array.dataPendapatan[*n].deskripsi)

	fmt.Println("Pendapatan berhasil ditambahkan!")
}
