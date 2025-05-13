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
	namaUsser string
	PassUser  string
	Kategori  string
	Target    int
	masukan   arrPendapatan
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
func tampilkanDashboard() {
	fmt.Println("\n=== Aplikasi Pencatatan Pendapatan ===")
	fmt.Println("1. Tambah Pendapatan Baru")
	fmt.Println("2. Tampilkan Laporan Keuangan")
	fmt.Println("3. Tampilkan Analisis Terkini")
	fmt.Println("4. Tampilkan Progres Target")
	fmt.Println("5. Logout")
	fmt.Println("Pilih menu: (1/2/3/4/5): ")

	var pilihan, n int
	var array ID
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		tambahPendapatan(&array, &n)
	case 2:
		laporan(&array, &n)
	case 3:
		TampilanAnalisis()
	case 4:
		TampilanProgres(&array, &n)
	case 5:
		return
	}

}

//fungsi untuk tampilan login
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
	var namaUsser, Kategori, PassUser string
	var TargetUsser int

	fmt.Println("Nama Pengguna:")
	fmt.Scan(&namaUsser)
	fmt.Println("Password Akun:")
	fmt.Println("Pastikan password terdiri angka, simbol dan huruf besar")
	fmt.Scan(&PassUser)
	fmt.Println("Profesi (Pelajar/Mahasiswa/Pekerja, dll):")
	fmt.Scan(&Kategori)
	fmt.Println("Masukan Target Pencapaian Pendapatan (Input Target Hanya Angka): ")
	fmt.Scan(&TargetUsser)

	if *nArrayPengguna < NMAX {
		arrayPengguna[*nArrayPengguna].namaUsser = namaUsser
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

	for i := 0; i < *nArrayPengguna; i++ {
		if nama == arrayPengguna[i].namaUsser && password == arrayPengguna[i].PassUser {
			fmt.Println("Selamat! Anda berhasil login.")
			tampilkanDashboard()
			return // login berhasil, keluar dari fungsi
		}
	}
	fmt.Println("Username atau Password salah.") // hanya tampil jika tidak ada yang cocok
}

// Fungsi untuk menambahkan data pendapatan
func tambahPendapatan(array *ID, n *int) {

	fmt.Println("\n--- Tambah Data Pendapatan ---")

	fmt.Println("Jenis Pendapatan: ")
	fmt.Println("1. Side Hustle")
	fmt.Println("2. Passive Income")
	fmt.Println("Pilih (1/2) :")

	var pilihan, jenisSideHustle, jenisPassive int
	for *n = 0; *n < NMAX; (*n)++ {
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
				array.masukan[*n].nama = "Freelance"
				KetSideHustle(array, n)
			case 2:
				array.masukan[*n].nama = "Konten Kreator"
				KetSideHustle(array, n)
			case 3:
				array.masukan[*n].nama = "Jualan Online"
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
				array.masukan[*n].nama = "Uang Sewa Lahan"
				KetSideHustle(array, n)
			case 2:
				array.masukan[*n].nama = "Royalti"
				KetSideHustle(array, n)
			}

			tampilkanDashboard()
		}

	}

}

// fungsi untuk laporan keuangan
func laporan(array *ID, n *int) {

	var pilihanlap int
	fmt.Println("===============")
	fmt.Println("1. Tampilkan Laporan Bulanan")
	fmt.Println("2. Tampilkan Laporan Tahunan")
	fmt.Println("Pilih (1/2): ")
	fmt.Scan(&pilihanlap)

	switch pilihanlap {
	case 1:
		lapBulanan()
	case 2:
		lapTahunan()
	}
}

func lapBulanan() {
	fmt.Println("Berikut adalah laporan bulanan anda: ")

}

func lapTahunan() {
	fmt.Println("Berikut adalah laporan tahunan anda: ")

}

func TampilanAnalisis() {

}

func TampilanProgres(array *ID, n *int) {
	fmt.Printf("Target Pendapatan Anda Bulan ini : %d", array.Target)
	fmt.Printf("Persentase Target Tercapai Bulan ini sebesar %f", hitungProgress(array, n))
}
func hitungProgress(array *ID, n *int) float64 {
	var total int
	for i := 0; i < *n; i++ {
		total += array.masukan[i].jumlah
	}
	return (float64(total) / float64(*n)) * 100
}
func KetSideHustle(array *ID, n *int) {
	fmt.Println("Masukan nominal:")
	fmt.Scan(&array.masukan[*n].jumlah)

	fmt.Println("Tanggal Pemasukan")
	fmt.Println("masukan Tanggal:")
	fmt.Scan(&array.masukan[*n].tanggal)

	fmt.Println("masukan bulan:")
	fmt.Scan(&array.masukan[*n].bulan)

	fmt.Println("masukan tahun:")
	fmt.Scan(&array.masukan[*n].tahun)

	fmt.Println("Deskripsi: ")
	fmt.Scan(&array.masukan[*n].deskripsi)

	fmt.Println("Pendapatan berhasil ditambahkan!")

}
