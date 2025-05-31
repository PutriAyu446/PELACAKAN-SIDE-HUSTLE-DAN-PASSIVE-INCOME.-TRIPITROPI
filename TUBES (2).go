package main

import "fmt"

type ID struct {
	Username       string
	PassUser       string
	Target         int
	dataPendapatan arrPendapatan
}

type detail struct {
	kategori, deskripsi   string
	jumlah                int
	tanggal, bulan, tahun int
}

const NMAX int = 50

type arrID [NMAX]ID
type arrPendapatan [NMAX]detail

func main() {
	// I.S: Program siap dieksekusi, menampilkan menu utama.
	// F.S: Program berhenti setelah pengguna memilih opsi 'Exit'.
	var pilih, nArrayPengguna int
	var arrayPengguna arrID

	for {
		menu()
		fmt.Scan(&pilih)
		//apabila user menginput pilihan yang tidak valid maka, user akan dimintakan untuk menginput kembali pilihan yang valid
		for pilih > 3 || pilih <= 0 {
			fmt.Println("\nPilihan Tidak Valid!")
			fmt.Print(">> Silahkan Pilih Kembali Menu Yang Tersedia << \n")
			fmt.Println("\n>> Daftar Menu:")
			fmt.Println("1. Register")
			fmt.Println("2. Login")
			fmt.Println("3. Exit")
			fmt.Print("\nPilih Menu (1/2/3): ")
			fmt.Scan(&pilih)
		}

		switch pilih {
		case 1:
			menuRegister(&arrayPengguna, &nArrayPengguna)
		case 2:
			menuLogin(&arrayPengguna, &nArrayPengguna)
		case 3:
			fmt.Println("\n-------------------------------------------------")
			fmt.Println("     Terima kasih telah menggunakan aplikasi     ")
			fmt.Println("-------------------------------------------------\n")
			return
		}
	}
}

// menu menampilkan pilihan menu utama aplikasi.
// I.S: Layar konsol siap untuk menampilkan menu.
// F.S: Menu utama aplikasi ditampilkan di konsol, menunggu input pilihan pengguna.
func menu() {
	fmt.Println("\n----------------------------------")
	fmt.Println("          Selamat Datang          ")
	fmt.Println("       Di Aplikasi Pelacak        ")
	fmt.Println("  Side Hustle dan Passive Income  ")
	fmt.Println("----------------------------------\n")
	fmt.Println(">> Daftar Menu:")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")
	fmt.Printf("%s ", "\nPilih Menu (1/2/3): ")
}

// menuRegister memungkinkan pengguna baru untuk mendaftar akun.
// I.S: `arrayPengguna` berisi data pengguna yang sudah terdaftar, `nArrayPengguna` adalah jumlah pengguna saat ini.
// F.S: Jika pendaftaran berhasil, pengguna baru ditambahkan ke `arrayPengguna`, `nArrayPengguna` bertambah satu.
//      Jika username sudah ada atau kapasitas penuh, pesan error ditampilkan.
func menuRegister(arrayPengguna *arrID, nArrayPengguna *int) {
	var namaUser, PassUser string
	var TargetUser int

	if *nArrayPengguna < NMAX {
		fmt.Println("\n---------------------------------------")
		fmt.Println("              REGISTER AKUN              ")
		fmt.Println("---------------------------------------\n")
		fmt.Print("Username (Tanpa Menggunakan Spasi): ")
		fmt.Scan(&namaUser)

		for i := 0; i < *nArrayPengguna; i++ {
			if namaUser == arrayPengguna[i].Username {
				fmt.Println("Nama Sudah Digunakan User Lain!\n")
				fmt.Print(" >> Silahkan Register Kembali << \n")
				fmt.Print("Masukan Nama Pengguna Lainnya: ")
				fmt.Scan(&namaUser)
				i = -1 // reset ulang pengecekan
			}
		}

		fmt.Print("Password Akun (Password Terdiri Angka, Simbol Dan Huruf Besar): ")
		fmt.Scan(&PassUser)
		fmt.Print("Masukan Target Pencapaian Pendapatan (Input Target Hanya Angka): ")
		fmt.Scan(&TargetUser)
		for TargetUser < 1000 || TargetUser <= 0 {
			fmt.Println("Target Tidak Valid!\n")
			fmt.Print(" >> Silahkan Input Target Kembali << \n")
			fmt.Print("Masukan Target Pencapaian Pendapatan (Input Target Hanya Angka): ")
			fmt.Scan(&TargetUser)
		}

		arrayPengguna[*nArrayPengguna].Username = namaUser
		arrayPengguna[*nArrayPengguna].PassUser = PassUser
		arrayPengguna[*nArrayPengguna].Target = TargetUser
		fmt.Println("\n Selamat Pengguna Berhasil Didaftarkan!")
		(*nArrayPengguna)++

	} else {
		fmt.Println("Kapasitas Pengguna Sudah Penuh")
	}
}

// menuLogin memungkinkan pengguna yang sudah terdaftar untuk masuk ke akun mereka.
// I.S: `arrayPengguna` berisi data pengguna yang sudah terdaftar, `nArrayPengguna` adalah jumlah pengguna saat ini.
// F.S: Jika username dan password cocok, pengguna berhasil login dan dialihkan ke dashboard.
//      Jika tidak cocok, pesan error ditampilkan dan pengguna diminta mencoba lagi.
func menuLogin(arrayPengguna *arrID, nArrayPengguna *int) {
	var nama, password string
	loginBerhasil := false

	fmt.Println("\n---------------------------------------")
	fmt.Println("           LOGIN KE AKUN ANDA          ")
	fmt.Println("---------------------------------------")

	for !loginBerhasil {
		fmt.Print("\nMasukkan Username: ")
		fmt.Scan(&nama)
		fmt.Print("Password Akun (Password Terdiri Angka, Simbol Dan Huruf Besar): ")
		fmt.Scan(&password)

		for i := 0; i < *nArrayPengguna; i++ {
			if nama == arrayPengguna[i].Username && password == arrayPengguna[i].PassUser {
				fmt.Println("\nSelamat! Anda berhasil login.")
				tampilkanDashboard(&arrayPengguna[i])
				loginBerhasil = true
				return
			}
		}

		fmt.Println("\nUsername atau Password salah!")
		fmt.Println(">> Silakan coba login kembali <<")
	}
}

// tampilkanDashboard menampilkan menu utama setelah pengguna berhasil login.
// I.S: Pengguna sudah login, `user` berisi data akun pengguna yang sedang aktif.
// F.S: Pengguna memilih opsi di dashboard atau logout.
func tampilkanDashboard(user *ID) {
	var pilihan int

	for {
		fmt.Println("\n-----------------------------------------")
		fmt.Println("     APLIKASI PENCATATAN PENDAPATAN      ")
		fmt.Println("-----------------------------------------\n")
		fmt.Println(">> Daftar Menu")
		fmt.Println("1. Tambah Pendapatan Baru")
		fmt.Println("2. Tampilkan Laporan Keuangan")
		fmt.Println("3. Tampilkan Analisis Terkini")
		fmt.Println("4. Tampilkan Progres Target")
		fmt.Println("5. Logout")
		fmt.Print("\nPilih Menu (1/2/3/4/5): ")

		n := hitungJumlahPendapatan(user) // Mendapatkan jumlah data pendapatan terkini

		fmt.Scan(&pilihan)
		//apabila user menginput pilihan yang tidak valid maka, user akan dimintakan untuk menginput kembali pilihan yang valid
		for pilihan > 5 || pilihan <= 0 {
			fmt.Println("\nPilihan Tidak Valid!")
			fmt.Print(">> Silahkan Pilih Kembali Menu Yang Tersedia << \n")
			fmt.Println("\n>> Daftar Menu")
			fmt.Println("1. Tambah Pendapatan Baru")
			fmt.Println("2. Tampilkan Laporan Keuangan")
			fmt.Println("3. Tampilkan Analisis Terkini")
			fmt.Println("4. Tampilkan Progres Target")
			fmt.Println("5. Logout")
			fmt.Print("\nPilih Menu (1/2/3/4/5): ")
			fmt.Scan(&pilihan)
		}
		switch pilihan {
		case 1:
			tambahPendapatan(user, &n)
		case 2:
			laporan(user, &n)
		case 3:
			TampilanAnalisis(user, &n)
		case 4:
			TampilanProgres(user, &n)
		case 5:
			return
		}
	}
}

// hitungJumlahPendapatan menghitung jumlah data pendapatan yang sudah tercatat untuk pengguna.
// I.S: `user` berisi data akun pengguna, termasuk array `dataPendapatan`.
// F.S: Mengembalikan integer yang merepresentasikan jumlah entri pendapatan yang tidak kosong (jumlah != 0).
func hitungJumlahPendapatan(user *ID) int {
	var count, i int
	count = 0
	for i = 0; i < NMAX; i++ {
		if user.dataPendapatan[i].jumlah != 0 {
			count++
		}
	}
	return count
}

// tambahPendapatan memungkinkan pengguna untuk menambahkan data pendapatan baru.
// I.S: `array` adalah pointer ke struct ID pengguna, `n` adalah pointer ke jumlah data pendapatan saat ini.
// F.S: Data pendapatan baru ditambahkan ke `array.dataPendapatan` jika dikonfirmasi, dan `*n` bertambah.
//      Pengguna memiliki opsi untuk mengedit atau menghapus data sebelum disimpan final.
func tambahPendapatan(array *ID, n *int) {
	var pilihan int

	for {
		fmt.Println("\n---------------------------------------")
		fmt.Println("          TAMBAH DATA PENDAPATAN         ")
		fmt.Println("---------------------------------------\n")
		fmt.Println(">> Jenis Pendapatan:")
		fmt.Println("1. Side Hustle")
		fmt.Println("2. Passive Income")
		fmt.Print("\nPilih Jenis Pendapatan (1/2): ")
		fmt.Scan(&pilihan)

		for pilihan > 2 || pilihan <= 0 {
			fmt.Println("\nPilihan Tidak Valid!")
			fmt.Print(">> Silahkan Pilih Kembali Jenis Pendapatan Yang Tersedia << \n")
			fmt.Println("\n>> Jenis Pendapatan:")
			fmt.Println("1. Side Hustle")
			fmt.Println("2. Passive Income")
			fmt.Print("\nPilih Jenis Pendapatan (1/2): ")
			fmt.Scan(&pilihan)
		}

		switch pilihan {
		case 1:
			jenisSideHustle(array, n)
			isiData(array, n)
			KonfirmData(array, n)
			return // selesai, keluar dari fungsi
		case 2:
			jenisPassiveIncome(array, n)
			isiData(array, n)
			KonfirmData(array, n)
			return // selesai, keluar dari fungsi
		}
	}
}

// jenisPassiveIncome memungkinkan pengguna untuk memilih kategori pendapatan pasif.
// I.S: `array` adalah pointer ke struct ID pengguna, `n` adalah pointer ke jumlah data pendapatan saat ini.
// F.S: Kategori pendapatan pasif yang dipilih oleh pengguna ("Uang Sewa Lahan" atau "Royalti") disimpan di `array.dataPendapatan[*n].kategori`.
func jenisPassiveIncome(array *ID, n *int) {
	var jenisPassive int
	fmt.Println("\n>> Jenis Passive Income:")
	fmt.Println("1. Uang Sewa Lahan")
	fmt.Println("2. Royalti")
	fmt.Print("\nPilih Jenis Passive Income (1/2): ")
	fmt.Scan(&jenisPassive)

	//apabila user menginput pilihan yang tidak valid maka, user akan dimintakan untuk menginput kembali pilihan yang valid
	for jenisPassive > 2 || jenisPassive <= 0 {
		fmt.Println("\nPilihan Tidak Valid!")
		fmt.Print(">> Silahkan Pilih Kembali Jenis Passive Income Yang Tersedia << \n")
		fmt.Println("\n>> Jenis Passive Income:")
		fmt.Println("1. Uang Sewa Lahan")
		fmt.Println("2. Royalti")
		fmt.Print("\nPilih Jenis Passive Income (1/2): ")
		fmt.Scan(&jenisPassive)
	}

	switch jenisPassive {
	case 1:
		array.dataPendapatan[*n].kategori = "Uang Sewa Lahan"
	case 2:
		array.dataPendapatan[*n].kategori = "Royalti"
	}
}

// jenisSideHustle memungkinkan pengguna untuk memilih kategori side hustle.
// I.S: `array` adalah pointer ke struct ID pengguna, `n` adalah pointer ke jumlah data pendapatan saat ini.
// F.S: Kategori side hustle yang dipilih oleh pengguna ("Freelance", "Konten Kreator", atau "Jualan Online") disimpan di `array.dataPendapatan[*n].kategori`.
func jenisSideHustle(array *ID, n *int) {
	var jenisSideHustle int
	fmt.Println("\n>> Jenis Side Hustle:")
	fmt.Println("1. Freelance")
	fmt.Println("2. Konten Kreator")
	fmt.Println("3. Jualan Online")
	fmt.Print("\nPilih Jenis Side Hustle (1/2/3): ")
	fmt.Scan(&jenisSideHustle)

	//apabila user menginput pilihan yang tidak valid maka, user akan dimintakan untuk menginput kembali pilihan yang valid
	for jenisSideHustle > 3 || jenisSideHustle <= 0 {
		fmt.Println("\nPilihan Tidak Valid!")
		fmt.Print(">> Silahkan Pilih Kembali Jenis Side Hustle Yang Tersedia << \n")
		fmt.Println("\n>> Jenis Side Hustle:")
		fmt.Println("1. Freelance")
		fmt.Println("2. Konten Kreator")
		fmt.Println("3. Jualan Online")
		fmt.Print("\nPilih Jenis Side Hustle (1/2/3): ")
		fmt.Scan(&jenisSideHustle)
	}

	switch jenisSideHustle {
	case 1:
		array.dataPendapatan[*n].kategori = "Freelance"
	case 2:
		array.dataPendapatan[*n].kategori = "Konten Kreator"
	case 3:
		array.dataPendapatan[*n].kategori = "Jualan Online"
	}
}

// KonfirmData memungkinkan pengguna untuk mengonfirmasi, mengedit, atau menghapus data pendapatan yang baru diinput.
// I.S: `array` adalah pointer ke struct ID pengguna, `n` adalah pointer ke jumlah data pendapatan saat ini.
// F.S: Jika pengguna mengonfirmasi, data pendapatan baru disimpan dan `*n` bertambah. Jika pengguna memilih untuk mengedit, fungsi `editPendapatan` dipanggil. Jika pengguna memilih untuk menghapus, fungsi `hapusPendapatan` dipanggil dan fungsi ini akan keluar Konfirmasi data
func KonfirmData(array *ID, n *int) {
	// Konfirmasi data
	sign := 1
	for sign == 1 {
		fmt.Println("\n>> Apakah data sudah benar?")
		fmt.Println("1. Ya")
		fmt.Println("2. Edit")
		fmt.Println("3. Hapus")
		fmt.Print("\nPilih (1/2/3): ")

		var konfirmasi int
		fmt.Scan(&konfirmasi)

		//apabila user menginput pilihan yang tidak valid maka, user akan dimintakan untuk menginput kembali pilihan yang valid
		for konfirmasi > 3 || konfirmasi <= 0 {
			fmt.Println("\nPilihan Tidak Valid!")
			fmt.Print(">> Silahkan Pilih Kembali Menu Yang Tersedia <<\n ")
			fmt.Println("\n>> Apakah data sudah benar?")
			fmt.Println("1. Ya")
			fmt.Println("2. Edit")
			fmt.Println("3. Hapus")
			fmt.Print("\nPilih (1/2/3): ")
			fmt.Scan(&konfirmasi)
		}

		switch konfirmasi {
		case 1:
			fmt.Println(">> Data Berhasil Disimpan << ")
			(*n)++
			sign = -1
		case 2:
			editPendapatan(array, *n) // Mengedit data yang baru diinput (ada di indeks *n)
		case 3:
			// Perlu diperhatikan: jika hapus di sini, *n tidak di-increment
			hapusPendapatan(array, n) // Menghapus data yang baru diinput
			return                    // Keluar dari fungsi tambahPendapatan
		}
	}
}

// isiData meminta dan menyimpan detail informasi pendapatan (nominal, tanggal, deskripsi).
// I.S: `array.dataPendapatan[*n].kategori` sudah terisi, `n` adalah pointer ke indeks data pendapatan saat ini.
// F.S: Detail pendapatan (jumlah, tanggal, bulan, tahun, deskripsi) diisi ke `array.dataPendapatan[*n]`.
func isiData(array *ID, n *int) {
	fmt.Println("\n------- Mohon Lengkapi Informasi Berikut -------")

	fmt.Print("\n> Masukkan Nominal (tanpa tanda titik, contoh: 10000 bukan 10.000): ")
	fmt.Scan(&array.dataPendapatan[*n].jumlah)
	for array.dataPendapatan[*n].jumlah < 1000 {
		fmt.Println("Nominal Uang Tidak Valid!\n")
		fmt.Print(" >> Silahkan Input Nominal Kembali << \n")
		fmt.Print("Masukkan Nominal Baru: ")
		fmt.Scan(&array.dataPendapatan[*n].jumlah)
	}

	fmt.Print("> Masukan Tahun : ")
	fmt.Scan(&array.dataPendapatan[*n].tahun)
	for array.dataPendapatan[*n].tahun < 1000 || array.dataPendapatan[*n].tahun >= 3000 {
		fmt.Println("Tahun Tidak Valid!")
		fmt.Print(" \n>> Silahkan Input Tahun Kembali << \n")
		fmt.Print("Masukkan Tahun Baru: ")
		fmt.Scan(&array.dataPendapatan[*n].tahun)
	}

	fmt.Print("> Masukan Bulan (1-12): ")
	fmt.Scan(&array.dataPendapatan[*n].bulan)
	for array.dataPendapatan[*n].bulan > 12 || array.dataPendapatan[*n].bulan <= 0 {
		fmt.Println("Bulan Tidak Valid!")
		fmt.Print("\n >> Silahkan Input Bulan Kembali << \n")
		fmt.Print("Masukkan Bulan Baru: ")
		fmt.Scan(&array.dataPendapatan[*n].bulan)
	}

	fmt.Print("> Tanggal Pemasukan: ")
	fmt.Scan(&array.dataPendapatan[*n].tanggal)
	bulan := array.dataPendapatan[*n].bulan
	if bulan == 1 || bulan == 3 || bulan == 5 || bulan == 7 || bulan == 8 || bulan == 10 || bulan == 12 {
		for array.dataPendapatan[*n].tanggal <= 0 || array.dataPendapatan[*n].tanggal > 31 {
			fmt.Println("Tanggal Tidak Valid!")
			fmt.Print(" \n>> Silahkan Input Tanggal Kembali << \n")
			fmt.Print("Masukkan Tanggal Baru: ")
			fmt.Scan(&array.dataPendapatan[*n].tanggal)
		}
	} else if bulan == 2 {
		for array.dataPendapatan[*n].tanggal <= 0 || array.dataPendapatan[*n].tanggal > 29 {
			fmt.Println("Tanggal Tidak Valid!")
			fmt.Print(" \n>> Silahkan Input Tanggal Kembali << \n")
			fmt.Print("Masukkan Tanggal Baru: ")
			fmt.Scan(&array.dataPendapatan[*n].tanggal)
		}
	} else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		for array.dataPendapatan[*n].tanggal <= 0 || array.dataPendapatan[*n].tanggal > 30 {
			fmt.Println("Tanggal Tidak Valid!")
			fmt.Print(" \n>> Silahkan Input Tanggal Kembali << \n")
			fmt.Print("Masukkan Tanggal Baru: ")
			fmt.Scan(&array.dataPendapatan[*n].tanggal)
		}
	}

	fmt.Print("> Deskripsi Singkat Pendapatanmu (Gunakan '_' sebagai spasi, contoh: Editor Video:")
	fmt.Scan(&array.dataPendapatan[*n].deskripsi)

	fmt.Println("\nPendapatan berhasil ditambahkan!")
}

// editPendapatan memungkinkan pengguna untuk mengubah detail data pendapatan.
// I.S: `array` adalah pointer ke struct ID pengguna, `index` adalah indeks data pendapatan yang akan diedit.
// F.S: Data pendapatan pada `array.dataPendapatan[index]` diperbarui sesuai input pengguna.
func editPendapatan(array *ID, index int) {
	fmt.Println("\n--- Edit Data Pendapatan ---")

	fmt.Print("Masukkan Nominal Baru: ")
	fmt.Scan(&array.dataPendapatan[index].jumlah)
	for array.dataPendapatan[index].jumlah < 1000 {
		fmt.Println("Nominal Uang Tidak Valid!\n")
		fmt.Print(" >> Silahkan Input Nominal Kembali << \n")
		fmt.Print("Masukkan Nominal Baru: ")
		fmt.Scan(&array.dataPendapatan[index].jumlah)
	}

	fmt.Println("Masukkan Tahun Baru: ")
	fmt.Scan(&array.dataPendapatan[index].tahun)
	for array.dataPendapatan[index].tahun < 1000 || array.dataPendapatan[index].tahun >= 3000 {
		fmt.Println("Tahun Tidak Valid!")
		fmt.Print(" \n>> Silahkan Input Tahun Kembali << \n")
		fmt.Print("Masukkan Tahun Baru: ")
		fmt.Scan(&array.dataPendapatan[index].tahun)
	}

	fmt.Print("Masukkan Bulan Baru: ")
	fmt.Scan(&array.dataPendapatan[index].bulan)
	for array.dataPendapatan[index].bulan > 12 || array.dataPendapatan[index].bulan <= 0 {
		fmt.Println("Bulan Tidak Valid!")
		fmt.Print("\n >> Silahkan Input Bulan Kembali << \n")
		fmt.Print("Masukkan Bulan Baru: ")
		fmt.Scan(&array.dataPendapatan[index].bulan)
	}

	fmt.Print("Masukkan Tanggal Baru: ")
	fmt.Scan(&array.dataPendapatan[index].tanggal)
	bulan := array.dataPendapatan[index].bulan
	if bulan == 1 || bulan == 3 || bulan == 5 || bulan == 7 || bulan == 8 || bulan == 10 || bulan == 12 {
		for array.dataPendapatan[index].tanggal <= 0 || array.dataPendapatan[index].tanggal > 31 {
			fmt.Println("Tanggal Tidak Valid!")
			fmt.Print(" \n>> Silahkan Input Tanggal Kembali << \n")
			fmt.Print("Masukkan Tanggal Baru: ")
			fmt.Scan(&array.dataPendapatan[index].tanggal)
		}
	} else if bulan == 2 {
		for array.dataPendapatan[index].tanggal <= 0 || array.dataPendapatan[index].tanggal > 29 {
			fmt.Println("Tanggal Tidak Valid!")
			fmt.Print(" \n>> Silahkan Input Tanggal Kembali << \n")
			fmt.Print("Masukkan Tanggal Baru: ")
			fmt.Scan(&array.dataPendapatan[index].tanggal)
		}
	} else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		for array.dataPendapatan[index].tanggal <= 0 || array.dataPendapatan[index].tanggal > 30 {
			fmt.Println("Tanggal Tidak Valid!")
			fmt.Print(" \n>> Silahkan Input Tanggal Kembali << \n")
			fmt.Print("Masukkan Tanggal Baru: ")
			fmt.Scan(&array.dataPendapatan[index].tanggal)
		}
	}

	fmt.Print("Masukkan Deskripsi Baru (Gunakan '_' sebagai spasi, contoh: 'Editor Video'):  ")
	fmt.Scan(&array.dataPendapatan[index].deskripsi)

	fmt.Println("\nData Anda Berhasil Diperbarui!")

}

// hapusPendapatan menghapus data pendapatan pada indeks tertentu.
// I.S: `array` adalah pointer ke struct ID pengguna, `n` adalah pointer ke jumlah data pendapatan saat ini.
// F.S: Data pada indeks terakhir (yang baru saja diinput/dikirim) dihapus dan `*n` berkurang satu.
func hapusPendapatan(array *ID, n *int) {
	// Menggeser semua elemen ke kiri dari index terakhir.
	// Catatan: Implementasi ini menghapus elemen terakhir yang 'diisi'
	// atau menggeser data yang baru diinput.
	if *n <= 0 {
		fmt.Println(">> Tidak ada data yang dapat dihapus.")
		return
	} else {
		for i := *n - 1; i < NMAX-1; i++ {
			array.dataPendapatan[i] = array.dataPendapatan[i+1]
		}
		fmt.Println(">> Data Telah Dihapus ")
		(*n)--
	}

}

// laporan menampilkan menu pilihan untuk melihat laporan keuangan.
// I.S: `array` berisi data pengguna, `n` adalah pointer ke jumlah data pendapatan.
// F.S: Pengguna memilih untuk melihat laporan berdasarkan kategori atau bulan, dan fungsi terkait dipanggil.
func laporan(array *ID, n *int) {

	var pilihan int

	fmt.Println("\n---------------------------------------")
	fmt.Println("            LAPORAN PENDAPATAN           ")
	fmt.Println("---------------------------------------\n")
	fmt.Println("\nIngin Ditampilkan Berdasarkan Apa?       ")
	fmt.Println("1. Berdasarkan Kategori                   ")
	fmt.Println("2. Berdasarkan Bulan                      ")
	fmt.Print("\nPilih (1/2): ")
	fmt.Scan(&pilihan)

	for pilihan > 2 || pilihan <= 0 {
		fmt.Println("\nPilihan Tidak Valid!")
		fmt.Print(">> Silahkan Pilih Kembali Menu Yang Tersedia << \n")
		fmt.Println("\nIngin Ditampilkan Berdasarkan Apa?       ")
		fmt.Println("1. Berdasarkan Kategori                   ")
		fmt.Println("2. Berdasarkan Bulan                      ")
		fmt.Print("\nPilih (1/2): ")
		fmt.Scan(&pilihan)
	}

	switch pilihan {
	case 1:
		cariBerdasarkanKategori(array, n)
	case 2:
		cariBerdasarkanBulan(array, n)
	}

}

// cariBerdasarkanKategori mencari dan menampilkan data pendapatan berdasarkan kategori yang dipilih.
// I.S: `array` berisi data pengguna, `n` adalah pointer ke jumlah data pendapatan.
// F.S: Data pendapatan yang sesuai dengan kategori yang dipilih ditampilkan.
//      Array pendapatan diurutkan berdasarkan kategori sebelum pencarian biner.
func cariBerdasarkanKategori(array *ID, n *int) {
	var j int
	var kategoriPilihan int // Menggunakan int untuk pilihan kategori
	var namaKategori string // Variabel untuk menyimpan nama kategori string

	fmt.Println("\n>> Daftar Kategori: ")
	fmt.Println("1. Freelance")
	fmt.Println("2. Konten Kreator")
	fmt.Println("3. Jualan Online")
	fmt.Println("4. Uang Sewa Lahan")
	fmt.Println("5. Royalti")
	fmt.Print("\nPilih Kategori (1/2/3/4/5): ")
	fmt.Scan(&kategoriPilihan) // Scan ke int

	for kategoriPilihan > 5 || kategoriPilihan <= 0 {
		fmt.Println("\nPilihan Tidak Valid!")
		fmt.Print(">> Silahkan Pilih Kembali Kategori Yang Tersedia << \n")
		fmt.Println("\n>> Daftar Kategori: ")
		fmt.Println("1. Freelance")
		fmt.Println("2. Konten Kreator")
		fmt.Println("3. Jualan Online")
		fmt.Println("4. Uang Sewa Lahan")
		fmt.Println("5. Royalti")
		fmt.Print("\nPilih Kategori (1/2/3/4/5): ")
		fmt.Scan(&kategoriPilihan)
	}

	switch kategoriPilihan { // Gunakan int untuk switch
	case 1:
		namaKategori = "Freelance"
	case 2:
		namaKategori = "Konten Kreator"
	case 3:
		namaKategori = "Jualan Online"
	case 4:
		namaKategori = "Uang Sewa Lahan"
	case 5:
		namaKategori = "Royalti"
	}

	// Mengurutkan array berdasarkan kategori agar binary search bisa bekerja
	sortPendapatanByKategori(array, n) // Mengirim namaKategori

	indexAwal := binarySearchIndexAwal(array, *n, namaKategori) // Mengirim namaKategori

	if indexAwal == -1 {
		fmt.Println("Belum ada pendapatan yang tercatat di kategori ini. Silakan tambahkan jika ada")
		return
	} else {
		//Sequential Search (lanjutan dari binary search untuk menemukan semua)
		fmt.Println("\n>> Data Pendapatan Kategori Ini:")
		var i int = indexAwal
		for i < *n {
			if array.dataPendapatan[i].kategori == namaKategori { // Bandingkan dengan namaKategori
				j++
				fmt.Printf("%d. %s: %d (%d-%d-%d)\n", j, array.dataPendapatan[i].kategori, array.dataPendapatan[i].jumlah, array.dataPendapatan[i].tanggal, array.dataPendapatan[i].bulan, array.dataPendapatan[i].tahun)
			}
			i++
		}
	}
}

// cariBerdasarkanBulan mencari dan menampilkan data pendapatan untuk bulan tertentu,
// serta memberikan rekomendasi sumber pendapatan terbesar di bulan tersebut.
// I.S: `array` berisi data pengguna, `n` adalah pointer ke jumlah data pendapatan.
// F.S: Data pendapatan untuk bulan yang dipilih ditampilkan, beserta total dan rekomendasi kategori terbesar.
func cariBerdasarkanBulan(array *ID, n *int) {
	var bulan int
	fmt.Print("\nMasukkan bulan (1-12): ")
	fmt.Scan(&bulan)

	for bulan > 12 || bulan <= 0 {
		fmt.Println("\nPilihan Tidak Valid!")
		fmt.Print(">> Silahkan Input Bulan Kembali << ")
		fmt.Print("\nMasukkan bulan (1-12): ")
		fmt.Scan(&bulan)
	}

	fmt.Printf("\n>> Data Pendapatan pada Bulan %d:\n", bulan)
	adaData := false
	totalPendapatanBulanIni := 0

	// Variabel untuk menyimpan kategori unik dan totalnya secara paralel
	var kategoriUnik [NMAX]string
	var totalPerKategoriUnik [NMAX]int
	var jumlahKategoriUnik int = 0

	// Loop pertama: Menampilkan data, menghitung total bulanan, dan mengidentifikasi kategori unik
	var i int = 0
	for i < *n {
		if array.dataPendapatan[i].bulan == bulan {
			fmt.Printf("%d. %s - %d (%d-%d-%d): %s\n",
				i+1,
				array.dataPendapatan[i].kategori,
				array.dataPendapatan[i].jumlah,
				array.dataPendapatan[i].tanggal,
				array.dataPendapatan[i].bulan,
				array.dataPendapatan[i].tahun,
				array.dataPendapatan[i].deskripsi)
			adaData = true
			totalPendapatanBulanIni += array.dataPendapatan[i].jumlah

			kategoriDitemukanDiUnik := false
			var idxKategoriUnik int = 0
			var j int = 0
			for j < jumlahKategoriUnik {
				if kategoriUnik[j] == array.dataPendapatan[i].kategori {
					kategoriDitemukanDiUnik = true
					idxKategoriUnik = j // Simpan indeks kategori yang ditemukan
				}
				j++
			}

			if !kategoriDitemukanDiUnik && jumlahKategoriUnik < NMAX {
				// Kategori baru, tambahkan ke daftar unik
				kategoriUnik[jumlahKategoriUnik] = array.dataPendapatan[i].kategori
				totalPerKategoriUnik[jumlahKategoriUnik] = array.dataPendapatan[i].jumlah // Inisialisasi total
				jumlahKategoriUnik++
			} else if kategoriDitemukanDiUnik {
				// Kategori sudah ada, tambahkan jumlahnya ke total yang sesuai
				totalPerKategoriUnik[idxKategoriUnik] += array.dataPendapatan[i].jumlah
			}
		}
		i++
	}
	fmt.Printf("\nTotal Pendapatan Bulan %d: %d\n", bulan, totalPendapatanBulanIni)

	if !adaData {
		fmt.Println("Belum ada data pendapatan untuk bulan ini.")
		return
	}

	// Cari kategori dengan pendapatan terbesar dari daftar kategori unik
	kategoriTerbesar := ""
	maxJumlahKategori := -1 // Inisialisasi dengan -1 agar nominal 0 pun bisa jadi yang terbesar

	var k int = 0
	for k < jumlahKategoriUnik {
		if totalPerKategoriUnik[k] > maxJumlahKategori {
			maxJumlahKategori = totalPerKategoriUnik[k]
			kategoriTerbesar = kategoriUnik[k]
		}
		k++
	}

	fmt.Println("\n--- Rekomendasi Analisis Pendapatan ---")
	if kategoriTerbesar != "" && maxJumlahKategori >= 0 {
		fmt.Printf("Sumber pendapatan **terbesar** Anda di bulan ini adalah **%s** dengan total **%d**.\n", kategoriTerbesar, maxJumlahKategori)
		fmt.Println("Pertimbangkan untuk lebih fokus atau mengembangkan area ini untuk meningkatkan pendapatan!")
	} else {
		fmt.Println("Tidak dapat menentukan sumber pendapatan terbesar. Mungkin tidak ada pendapatan, atau data terlalu sedikit.")
	}
}

// binarySearchIndexAwal mencari indeks awal dari kategori tertentu dalam array pendapatan yang sudah diurutkan.
// I.S: `array` berisi data pengguna, `n` adalah jumlah data, `kategori` adalah string kategori yang dicari.
// F.S: Mengembalikan indeks pertama kemunculan kategori, atau -1 jika tidak ditemukan.
func binarySearchIndexAwal(array *ID, n int, kategori string) int {
	kiri := 0
	kanan := n - 1
	hasil := -1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		data := array.dataPendapatan[tengah]

		if data.kategori == kategori {
			hasil = tengah
			kanan = tengah - 1 // Terus cari ke kiri untuk menemukan indeks paling awal
		} else if data.kategori < kategori {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return hasil
}

// sortPendapatanByKategori mengurutkan array data pendapatan berdasarkan kategori menggunakan Selection Sort.
// I.S: `array` berisi data pengguna, `n` adalah pointer ke jumlah data pendapatan
// F.S: Array `array.dataPendapatan` diurutkan berdasarkan  kategori secara ascending.
func sortPendapatanByKategori(array *ID, n *int) {
	var i, min, pass int
	var temp detail
	pass = 1
	for pass < *n {
		min = pass - 1
		i = pass
		for i < *n {
			if array.dataPendapatan[i].kategori < array.dataPendapatan[min].kategori {
				min = i
			}
			i++
		}
		// Tukar elemen
		temp = array.dataPendapatan[pass-1]
		array.dataPendapatan[pass-1] = array.dataPendapatan[min]
		array.dataPendapatan[min] = temp
		pass = pass + 1
	}
}

// TampilanAnalisis menampilkan analisis pendapatan terbesar dan terkecil.
// I.S: `array` berisi data pengguna, `n` adalah pointer ke jumlah data pendapatan.
// F.S: Menampilkan kategori, jumlah, tanggal, dan deskripsi pendapatan terbesar dan terkecil.
func TampilanAnalisis(array *ID, n *int) {
	fmt.Println("\n---------------------------------------")
	fmt.Println("         ANALISIS SUMBER PENDAPTAN         ")
	fmt.Println("--------------------------------------- ")
	fmt.Println("Berikut Hasil Analisis Sumber Pendapatanmu")

	if *n == 0 {
		fmt.Println("Belum ada data pendapatan.")
		return
	}

	// Urutkan pendapatan berdasarkan jumlah (dari kecil ke besar)
	insertionSortPendapatan(array, n)

	// Data terakhir setelah insertion sorting ascending adalah pendapatan terbesar
	terbesar := array.dataPendapatan[*n-1]

	// Data pertama setelah insetion sorting ascending adalah pendapatan terkecil
	terkecil := array.dataPendapatan[0]

	fmt.Println("\n>> Pendapatan Terbesar:")
	fmt.Printf("Kategori: %s\n", terbesar.kategori)
	fmt.Printf("Jumlah: %d\n", terbesar.jumlah)
	fmt.Printf("Tanggal: %d-%d-%d\n", terbesar.tanggal, terbesar.bulan, terbesar.tahun)
	fmt.Printf("Deskripsi: %s\n", terbesar.deskripsi)

	fmt.Println("\n>> Pendapatan Terkecil:")
	fmt.Printf("Kategori: %s\n", terkecil.kategori)
	fmt.Printf("Jumlah: %d\n", terkecil.jumlah)
	fmt.Printf("Tanggal: %d-%d-%d\n", terkecil.tanggal, terkecil.bulan, terkecil.tahun)
	fmt.Printf("Deskripsi: %s\n", terkecil.deskripsi)
}

// insertionSortPendapatan mengurutkan array data pendapatan berdasarkan jumlah secara ascending menggunakan Insertion Sort.
// I.S: `array` berisi data pengguna, `n` adalah pointer ke jumlah data pendapatan.
// F.S: Array `array.dataPendapatan` diurutkan berdasarkan `jumlah` secara ascending.
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

// TampilanProgres menampilkan progres pencapaian target pendapatan pengguna.
// I.S: `array` berisi data pengguna (termasuk Target dan dataPendapatan), `n` adalah pointer ke jumlah data pendapatan.
// F.S: Menampilkan target pendapatan, persentase yang sudah tercapai, dan pesan motivasi.
func TampilanProgres(array *ID, n *int) {

	fmt.Println("\n---------------------------------------")
	fmt.Println("            PROGRES PENDAPATAN           ")
	fmt.Println("---------------------------------------\n")
	fmt.Printf("Target Pendapatan Anda : %d\n", array.Target)

	progres := hitungProgress(array, n)
	fmt.Printf("Persentase Target Tercapai :  %.2f%%\n", progres)

	if progres >= 100 {
		fmt.Println("Selamat! Pemasukan Anda telah melebihi target!")
	} else {
		fmt.Println("Semangat! Tetap konsisten untuk mencapai target.")
	}

	selectionSortPendapatan(array, n)
	fmt.Println("\n--- Riwayat Lengkap Pendapatan Anda ---")
	if *n == 0 {
		fmt.Println("Belum ada data pendapatan yang tercatat.")
	} else {
		var no int = 1
		var i int = 0
		for i < *n {
			// Hanya tampilkan data yang benar-benar ada (jumlah tidak nol),  hanya data valid yang dicetak.
			if array.dataPendapatan[i].jumlah != 0 {
				fmt.Printf("%d. Kategori: %s, Jumlah: %d, Tanggal: %d-%d-%d, Deskripsi: %s\n",
					no,
					array.dataPendapatan[i].kategori,
					array.dataPendapatan[i].jumlah,
					array.dataPendapatan[i].tanggal,
					array.dataPendapatan[i].bulan,
					array.dataPendapatan[i].tahun,
					array.dataPendapatan[i].deskripsi)
				no++
			}
			i++
		}
	}
}

// selectionSortPendapatan mengurutkan array data pendapatan berdasarkan jumlah secara descending menggunakan Selection Sort.
// I.S: `array` berisi data pengguna, `n` adalah pointer ke jumlah data pendapatan.
// F.S: Array `array.dataPendapatan` diurutkan berdasarkan `jumlah` secara descending.
func selectionSortPendapatan(array *ID, n *int) {
	var pass, idx, i int
	var temp detail

	pass = 0
	for pass < *n-1 {
		idx = pass
		i = pass + 1
		for i < *n {
			if array.dataPendapatan[i].jumlah > array.dataPendapatan[idx].jumlah { // Mengubah > untuk descending
				idx = i
			}
			i++
		}
		// Tukar elemen
		temp = array.dataPendapatan[pass]
		array.dataPendapatan[pass] = array.dataPendapatan[idx]
		array.dataPendapatan[idx] = temp
		pass++
	}
}

// hitungProgress menghitung persentase progres total pendapatan terhadap target yang ditetapkan.
// I.S: `array` berisi data pengguna (termasuk Target dan dataPendapatan), `n` adalah pointer ke jumlah data pendapatan.
// F.S: Mengembalikan nilai float64 yang merepresentasikan persentase progres. Mengembalikan 0 jika target 0.
func hitungProgress(array *ID, n *int) float64 {
	var total int
	var i int = 0
	for i < *n {
		total += array.dataPendapatan[i].jumlah
		i++
	}
	if array.Target == 0 {
		return 0
	}
	return (float64(total) / float64(array.Target)) * 100
}
