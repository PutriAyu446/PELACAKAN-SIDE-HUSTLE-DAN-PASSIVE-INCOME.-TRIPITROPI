# Aplikasi Pelacak Side Hustle dan Passive Income
Aplikasi ini dirancang untuk membantu Anda mencatat, mengelola, dan menganalisis pendapatan dari Side Hustle dan Passive Income. Dengan fitur registrasi dan login, setiap pengguna dapat melacak progres finansialnya secara individual.

# Fitur Utama
## 1. Manajemen Data Pendapatan: Tambah, Edit, dan Hapus
Pengguna dapat mengategorikan pendapatan mereka sebagai Side Hustle atau Passive Income. Sistem memfasilitasi pemilihan sumber pendapatan yang lebih spesifik:
Side Hustle mencakup Freelance, Konten Kreator, dan Jualan Online.
Passive Income meliputi Uang Sewa Lahan dan Royalti. Setiap pemasukan dilengkapi dengan detail penting: jumlah pendapatan (nominal), tanggal (tanggal, bulan, tahun), dan deskripsi singkat untuk memudahkan identifikasi.
Fungsi isiData bertanggung jawab untuk menerima input detail dan melakukan validasi untuk memastikan data yang dimasukkan valid, seperti nominal tidak di bawah Rp1.000, tahun antara 1000 dan 2999, serta bulan (1-12) dan tanggal yang sesuai dengan kalender.

## a. Konfirmasi dan Modifikasi Data (KonfirmData):
Setelah proses input detail pendapatan, pengguna disajikan dengan opsi konfirmasi untuk meninjau keakuratan data. Apabila ditemukan ketidaksesuaian atau kesalahan, pengguna diberikan fasilitas untuk langsung mengedit atau menghapus entri tersebut sebelum data disimpan secara permanen. Mekanisme ini dirancang untuk meminimalisir potensi kesalahan data pada tahap awal.

## b. Pengeditan Data Pendapatan (editPendapatan):
Fitur ini memungkinkan pengguna untuk memodifikasi catatan pendapatan yang sudah ada. Pengguna dapat memperbarui nilai nominal, bulan, tanggal, tahun, dan deskripsi dari suatu entri. Fungsi editPendapatan menerapkan standar validasi input yang sama ketatnya dengan proses penambahan data, guna menjaga integritas data.

## c. Penghapusan Data Pendapatan (hapusPendapatan):
Untuk data pendapatan yang dianggap tidak relevan atau mengandung kesalahan, fitur penghapusan disediakan. Fungsi hapusPendapatan secara efektif menghapus catatan pendapatan yang terakhir diinput atau yang dipilih untuk dihapus, serta secara otomatis menyesuaikan total jumlah data pendapatan yang tersimpan.

## 2. Laporan Keuangan Sederhana
Fitur ini membantu pengguna mendapatkan gambaran umum tentang pola pendapatan mereka. Laporan dapat dilihat dalam dua cara:

## a. Berdasarkan Kategori (cariBerdasarkanKategori):
Untuk mempercepat pencarian, sistem terlebih dahulu mengurutkan data pendapatan berdasarkan kategori menggunakan algoritma Selection Sort (sortPendapatanByKategori). Setelah data diurutkan, selanjutnya menggunakan Binary Search (binarySearchIndexAwal) untuk menemukan indeks pertama dari kategori yang dicari. Kemudian, Sequential Search dilanjutkan dari indeks tersebut untuk menampilkan semua data yang sesuai dengan kategori yang dipilih. Kombinasi ini dirancang untuk efisiensi pencarian.

## b. Berdasarkan Bulan (cariBerdasarkanBulan):
Pengguna dapat memasukkan bulan (1-12) untuk melihat laporan pendapatan bulanan. Selanjutnya, sistem akan menampilkan semua pemasukan yang tercatat pada bulan tersebut, lengkap dengan kategori, jumlah, tanggal, dan deskripsi. Selain itu, fitur ini juga menghitung total pendapatan untuk bulan yang dipilih dan memberikan rekomendasi optimasi pemasukan dengan mengidentifikasi sumber pendapatan terbesar pada bulan tersebut. Ini membantu pengguna memahami di mana letak potensi pendapatan terbesar mereka.

## 3. Analisis Terkini (TampilanAnalisis)
Sistem dapat menampilkan pendapatan terbesar dan terkecil berdasarkan jumlahnya. Data diurutkan secara ascending (dari kecil ke besar) menggunakan algoritma Insertion Sort (insertionSortPendapatan). Setelah pengurutan, pendapatan terkecil akan berada di awal array, dan pendapatan terbesar akan berada di akhir. Informasi yang ditampilkan mencakup kategori, jumlah, tanggal, dan deskripsi dari pendapatan terbesar dan terkecil.

# Fitur Tambahan
## 1. Registrasi dan Login Pengguna (menuRegister, menuLogin)
Aplikasi menyediakan fitur registrasi akun baru dengan validasi untuk memastikan bahwa setiap nama pengguna bersifat unik dan tidak digunakan oleh pengguna lain. Proses login berfungsi untuk mengautentikasi pengguna yang telah terdaftar, sehingga setiap pengguna dapat mengakses dan mengelola data pendapatannya secara terpisah. Dengan demikian, aplikasi mendukung penggunaan oleh banyak pengguna secara individual dan aman.

## 2. Analisis Target Pengguna (TampilanProgres, hitungProgress)
Pada fitur ini, pengguna dapat menetapkan target pendapatan yang ingin dicapai. Aplikasi akan menghitung persentase pencapaian terhadap target tersebut berdasarkan total pendapatan yang telah dicatat oleh pengguna. Selain itu, aplikasi juga menampilkan riwayat pendapatan secara lengkap, yang diurutkan dari jumlah terbesar ke terkecil menggunakan algoritma Selection Sort melalui fungsi selectionSortPendapatan.
