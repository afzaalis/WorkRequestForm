/*
Aplikasi  Work Request Form
Deskripsi: Aplikasi ini digunakan oleh bagian operasional dan pemeliharaan. Data yang diolah terdiri dari no_wrf, nama site, deskripsi_wrf, item remark, tanggal_kerja, pelanggan , status. Contoh data :  WRF001, Waris, Aktivasi Link, 10 GB, 12-10-2022, Cust001, closed. (Permintaan melakukan pekerjaan di site Waris untuk akivasi link sebesar 10 GB yang akan dikerjakan pada tanggal 12-10-2022 untuk Cust001 dan sudah dilaksanakan (closed).
Contoh nama site : bisa berupa nama kota seperti Ambon, Merauke, Kupang, Tanah Merah
Contoh deskripsi_wrf bisa berupa Instalasi perangkat, Aktifasi Link, Install modul, Dismantel perangkat, Upgrade Capacity, Degrade Capacity, Upgrade Software, Upgrade Modul
Contom item_remark menyesuaikan dengan deskripsi wrf. Misal untuk dismantle perangkat, item remark diisi dengan nama perangkat dan typenya, untuk aktifasi link, item remark diisi jumlah kapasitasnya. Item remark untuk memberikan keterangan tambahan deskripsi wrf.


Spesifikasi:
j.	Admin bisa menambahkan, menghapus atau mengubah data pada aplikasi.
k.	Pengguna bisa melihat daftar wrf terurut berdasarkan kriteria tertentu (misalnya berdasarkan nama site, deskripsi wrf, tanggal )
l.	Pengguna bisa mencari data pelanggan berdasarkan kata kunci dan kategori tertentu.
*/

package main

import (
	"fmt"
	"strings"
)

const NMAX = 2048

type user struct {
	username string
	password string
	status   string
}

type date struct {
	tanggal, bulan, tahun int
}

type WRF struct {
	no_wrf        string
	nama_site     string
	deskripsi_wrf string
	item_remark   string
	tanggal_kerja date
	pelanggan     string
	status        string
}

type users [NMAX]user
type WRFs [NMAX]WRF

func main() {
	var dataUser users
	var arrWrf WRFs

	dataUser[0] = user{username: "admin", password: "admin123", status: "admin"}
	dataUser[1] = user{username: "user", password: "user123", status: "user"}
	dataUser[2] = user{username: "afzaal", password: "afzaal", status: "user"}

	Login(dataUser, arrWrf)
}

func findUser(user, pass string, dataUser users, i *int, isFound *bool, flag string) {
	if flag == "login" {
		*i = 0
		for *i < NMAX && !*isFound {
			if dataUser[*i].username == user && dataUser[*i].password == pass {
				*isFound = true
			}
			*i++
		}
		*i -= 1
	} else if flag == "cari" {
		*i = 0
		for *i < NMAX && !*isFound {
			if dataUser[*i].username == user {
				*isFound = true
			}
			*i++
		}
		*i -= 1
	} else if flag == "ubah" {
		*i -= 1
		if *i != -1 {
			if dataUser[*i].username != "" {
				*isFound = true
			} else {
				*isFound = false
			}
		} else {
			*isFound = false
		}
	}
}

func findWRFNo(arrWrf WRFs, no_wrf string) bool {
	var isFound bool
	for i := 0; i < NMAX; i++ {
		if arrWrf[i].no_wrf == no_wrf {
			isFound = true
		}
	}
	return isFound
}

func findWRFindex(arrWrf WRFs, n int, isFound *bool) int {
	for i := 0; i < NMAX; i++ {
		if arrWrf[i].no_wrf == arrWrf[n-1].no_wrf && *isFound == false {
			*isFound = true
			return i
		}
	}

	return -1
}

func Login(dataUser users, arrWrf WRFs) {
	var user, pass string
	var isFound bool
	var n int
	fmt.Print(" __________________________________ \n")
	fmt.Print("|  ______________________________  |\n")
	fmt.Print("| |  | WELCOME !|                | |\n")
	fmt.Print("| |                              | |\n")
	fmt.Print("| |  TO                          | |\n")
	fmt.Print("| |                              | |\n")
	fmt.Print("| |   | WRF APPLICATION |        | |\n")
	fmt.Print("| |______________________________| |\n")
	fmt.Print("|__________________________________|\n")
	fmt.Print("username: ")
	fmt.Scan(&user)
	fmt.Print("Password: ")
	fmt.Scan(&pass)
	findUser(user, pass, dataUser, &n, &isFound, "login")
	for !isFound {
		fmt.Println("User tidak ditemukan")
		fmt.Print("username: ")
		fmt.Scan(&user)
		fmt.Print("Password: ")
		fmt.Scan(&pass)
		findUser(user, pass, dataUser, &n, &isFound, "login")
	}
	if dataUser[n].status == "admin" {
		fmt.Println("Go to admin`")
		UiAdmin(dataUser, arrWrf)
	} else if dataUser[n].status == "user" {
		fmt.Println("Go to user")
		uiPengguna(dataUser, arrWrf)
	}
}

func lihatDataUser(dataUser users) {
	fmt.Println("-------------------------------------------------------------------------")
	for i := 0; i < NMAX; i++ {
		if dataUser[i].username != "" && dataUser[i].password != "" && dataUser[i].status != "" {
			fmt.Println(i+1, ".	username:", dataUser[i].username)
			fmt.Println("	password:", dataUser[i].password)
			fmt.Println("	status:", dataUser[i].status)
			fmt.Println()
		}
	}
	fmt.Println("-------------------------------------------------------------------------")
}

func cariUser(dataUser users) {
	var input string
	var isFound bool
	var n int

	fmt.Println("-------------------------------------------------------------------------")
	fmt.Print("Masukkan Username: ")
	fmt.Scan(&input)
	findUser(input, " ", dataUser, &n, &isFound, "cari")
	if isFound {
		fmt.Println(n+1, ".	username:", dataUser[n].username)
		fmt.Println("	password:", dataUser[n].password)
		fmt.Println("	status:", dataUser[n].status)
		fmt.Println()
	} else {
		fmt.Println("User tidak ditemukan")
	}
	fmt.Println("-------------------------------------------------------------------------")
}

func ubahDataUser(dataUser *users) {
	var num, n int
	var isFound, usernameFound bool
	var input string
	var data user

	lihatDataUser(*dataUser)
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Print("masukkan nomor user yang ingin di ubah: ")
	fmt.Scan(&num)
	//check if the num is in valid value
	if num >= 1 && num <= NMAX {
		findUser(" ", " ", *dataUser, &num, &isFound, "ubah")

		if isFound {
			fmt.Println(num+1, ".	username:", dataUser[num].username)
			fmt.Println("	password:", dataUser[num].password)
			fmt.Println("	status:", dataUser[num].status)
			fmt.Println()

			fmt.Print("Ingin mengubah user tersebut? y/n: ")
			fmt.Scan(&input)
			if input == "y" {
				cancelled := false

				fmt.Println("Input 'cancel' untuk batalkan pengubahan")
				fmt.Println("Ubah menjadi: ")

				// Prompt for and validate the Username
				usernameFilled := false
				for !usernameFilled {
					usernameFound = false
					fmt.Print("Username: ")
					fmt.Scan(&data.username)
					findUser(data.username, " ", *dataUser, &n, &usernameFound, "cari")
					if data.username != "cancel" {
						if !usernameFound {
							usernameFilled = true
						} else {
							if dataUser[num].username == data.username {
								usernameFilled = true
							} else {
								fmt.Println("username sudah ada, masukkan username yang berbeda")
							}
						}
					} else {
						cancelled = true
						usernameFilled = true
					}
				}

				// Prompt for and validate the Password
				if !cancelled {
					passwordFilled := false
					for !passwordFilled {
						fmt.Print("Password: ")
						fmt.Scan(&data.password)
						if data.password != "cancel" {
							passwordFilled = true
						} else {
							cancelled = true
							passwordFilled = true
						}
					}
				}
				// Prompt for and validate the Status
				if !cancelled {
					statusFilled := false
					for !statusFilled {
						fmt.Print("Status: ")
						fmt.Scan(&data.status)
						if data.status != "cancel" {
							if data.status == "admin" || data.status == "user" {
								statusFilled = true
							} else {
								fmt.Println("Status yang tersedia hanya admin atau user")
							}
						} else {
							cancelled = true
							statusFilled = true
						}
					}
				}
				if cancelled {
					fmt.Println("Batal mengubah user")
				} else {
					dataUser[num].username = data.username
					dataUser[num].password = data.password
					dataUser[num].status = data.status
					fmt.Println("User berhasil di ubah")
				}
			} else {
				fmt.Println("Kembali")
			}
		} else {
			fmt.Println("User tidak ditemukan")
		}
	} else {
		fmt.Println("Nomor tidak valid")
	}
}

func tambahUser(dataUser *users) {
	var isFound bool
	var data user
	var n int

	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println("Tambah User")

	cancelled := false

	// Prompt for and validate the Username
	usernameFilled := false

	fmt.Println("Input 'cancel' untuk batalkan penambahan")
	for !usernameFilled {
		isFound = false
		fmt.Print("Username: ")
		fmt.Scan(&data.username)
		if data.username == "cancel" {
			fmt.Println("Batal membuat pengguna baru.")
			cancelled = true
			usernameFilled = true
		} else {
			findUser(data.username, " ", *dataUser, &n, &isFound, "cari")
			if isFound {
				fmt.Println("Username sudah ada. Masukkan username yang berbeda.")
			} else {
				usernameFilled = true
			}
		}
	}

	if !cancelled {
		// Prompt for and validate the Password
		passwordFilled := false
		for !passwordFilled {
			fmt.Print("Password: ")
			fmt.Scan(&data.password)
			if data.password == "cancel" {
				fmt.Println("Batal membuat pengguna baru.")
				cancelled = true
				passwordFilled = true
			} else {
				passwordFilled = true
			}
		}
	}

	if !cancelled {
		// Prompt for and validate the Status
		statusFilled := false
		for !statusFilled {
			fmt.Print("Status: ")
			fmt.Scan(&data.status)
			if data.status == "cancel" {
				fmt.Println("Batal membuat pengguna baru.")
				cancelled = true
				statusFilled = true
			} else if data.status == "admin" || data.status == "user" {
				statusFilled = true
			} else {
				fmt.Println("Status yang tersedia hanya admin atau user")
			}
		}
	}

	if !cancelled {
		// Find an empty slot to add the new user
		n = 0
		for n < NMAX-1 && dataUser[n].username != "" {
			n++
		}

		if n == NMAX-1 {
			fmt.Println("Tidak dapat menambahkan user. Kapasitas maksimum tercapai.")
		} else {
			dataUser[n].username = data.username
			dataUser[n].password = data.password
			dataUser[n].status = data.status
			fmt.Println("User berhasil ditambahkan.")
		}

		fmt.Println("-------------------------------------------------------------------------")
	}
}

func hapusUser(dataUser *users) {
	var num int
	var isFound bool
	var input string

	lihatDataUser(*dataUser)

	fmt.Println("-------------------------------------------------------------------------")
	fmt.Print("Masukkan nomor user yang ingin dihapus: ")
	fmt.Scan(&num)

	// Check if the index is within valid range
	if num >= 1 && num <= NMAX {
		findUser(" ", " ", *dataUser, &num, &isFound, "ubah")

		// Check if the user is found
		if isFound {
			fmt.Println(isFound)
			// Confirm deletion
			fmt.Println("Anda akan menghapus user berikut:")
			fmt.Println(num+1, ".	username:", dataUser[num].username)
			fmt.Println("	password:", dataUser[num].password)
			fmt.Println("	status:", dataUser[num].status)
			fmt.Print("Lanjutkan penghapusan? y/n: ")
			fmt.Scan(&input)

			if input == "y" {
				// Shift the indexes of other user data
				for i := num; i < NMAX-1; i++ {
					dataUser[i] = dataUser[i+1]
				}
				// Clear the last user data
				dataUser[NMAX-1] = user{}
				fmt.Println("User berhasil dihapus.")
			} else {
				fmt.Println("Penghapusan dibatalkan.")
			}
		} else {
			fmt.Println("User tidak ditemukan")
		}
	} else {
		fmt.Println("Nomor user tidak valid.")
	}
	fmt.Println("-------------------------------------------------------------------------")
}

func KelolaDataUser(dataUser users, arrWrf WRFs) {
	var pilihan int
	var inKelolaDataUser bool

	inKelolaDataUser = true

	for inKelolaDataUser {
		fmt.Println("=======================")
		fmt.Println("1. Lihat data User: ")
		fmt.Println("2. Cari user")
		fmt.Println("3. Ubah data user")
		fmt.Println("4. Tambah user")
		fmt.Println("5. Hapus user")
		fmt.Println("6. Kembali")
		fmt.Println("=======================")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			lihatDataUser(dataUser)
		} else if pilihan == 2 {
			cariUser(dataUser)
		} else if pilihan == 3 {
			ubahDataUser(&dataUser)
		} else if pilihan == 4 {
			tambahUser(&dataUser)
		} else if pilihan == 5 {
			hapusUser(&dataUser)
		} else if pilihan == 6 {
			inKelolaDataUser = false
			UiAdmin(dataUser, arrWrf)
		}
	}
}

func UiAdmin(dataUser users, arrWrf WRFs) {
	var choice int
	fmt.Println("=============================================")
	fmt.Println("|                                           |")
	fmt.Println("|            Selamat Datang, Admin!          |")
	fmt.Println("|                                           |")
	fmt.Println("|         Anda memiliki akses penuh          |")
	fmt.Println("|        untuk mengelola sistem ini.         |")
	fmt.Println("|                                           |")
	fmt.Println("=============================================")
	fmt.Println("1. kelola data user")
	fmt.Println("2. kelola wrf")
	fmt.Println("3. logout ")
	fmt.Print("pilihan: ")
	fmt.Scan(&choice)
	if choice == 1 {
		KelolaDataUser(dataUser, arrWrf)
	} else if choice == 2 {
		KelolaWRF(dataUser, arrWrf)
	} else if choice == 3 {
		Login(dataUser, arrWrf)
	}
}

func uiPengguna(dataUser users, arrWrf WRFs) {
	var choice int
	var newArrWrf WRFs

	for i := 0; i < NMAX; i++ {
		newArrWrf[i] = arrWrf[i]
	}

	fmt.Println("**********************************")
	fmt.Println("*                                *")
	fmt.Println("*         Selamat Datang!        *")
	fmt.Println("*                                *")
	fmt.Println("*  Terima kasih telah menggunakan *")
	fmt.Println("*    layanan kami. Nikmati       *")
	fmt.Println("*      pengalaman yang baik!     *")
	fmt.Println("*                                *")
	fmt.Println("**********************************")
	fmt.Println("1. Lihat daftar wrf")
	fmt.Println("2. Cari daftar wrf")
	fmt.Println("3. Lihat daftar wrf terurut")
	fmt.Println("4. Logout")
	fmt.Scan(&choice)
	if choice == 1 {
		lihatDaftarWrf(arrWrf, arrWrf, "user", dataUser, true)
	} else if choice == 2 {
		cariWrf(arrWrf, dataUser, "user")
	} else if choice == 3 {
		lihatSortWrf(arrWrf, &newArrWrf, dataUser, "user")
	} else if choice == 4 {
		Login(dataUser, arrWrf)
	}
}

func KelolaWRF(dataUser users, arrWrf WRFs) {
	var pilihan int
	var inkelolaWRF bool
	var newArrWrf WRFs

	for i := 0; i < NMAX; i++ {
		newArrWrf[i] = arrWrf[i]
	}

	inkelolaWRF = true

	for inkelolaWRF {
		fmt.Println("=======================")
		fmt.Println("1. Lihat Daftar wrf")
		fmt.Println("2. Cari wrf")
		fmt.Println("3. Ubah data wrf")
		fmt.Println("4. Tambah data wrf")
		fmt.Println("5. Hapus data wrf")
		fmt.Println("6. Lihat daftar wrf terurut")
		fmt.Println("7. Kembali")
		fmt.Println("=======================")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			lihatDaftarWrf(arrWrf, arrWrf, "admin", dataUser, true)
		} else if pilihan == 2 {
			cariWrf(arrWrf, dataUser, "admin")
		} else if pilihan == 3 {
			ubahDataWrf(&arrWrf, dataUser)
		} else if pilihan == 4 {
			tambahDataWrf(&arrWrf, dataUser)
		} else if pilihan == 5 {
			hapusDataWrf(&arrWrf, dataUser)
		} else if pilihan == 6 {
			lihatSortWrf(arrWrf, &newArrWrf, dataUser, "admin")
		} else if pilihan == 7 {
			inkelolaWRF = false
			UiAdmin(dataUser, arrWrf)
		}
	}

}

func lihatSortWrf(oldArrWrf WRFs, arrWrf *WRFs, dataUser users, in string) {
	var inputKategori, inputSort int

	fmt.Println("Lihat wrf")
	fmt.Println("Lihat berdasarkan: ")
	fmt.Println("1. Nomor wrf")
	fmt.Println("2. Nama site")
	fmt.Println("3. Deskripsi wrf")
	fmt.Println("4. Remark")
	fmt.Println("5. Tanggal Kerja")
	fmt.Println("6. Pelanggan")
	fmt.Println("7. Kembali")
	fmt.Print("Pilihan: ")
	fmt.Scan(&inputKategori)
	if inputKategori == 1 {
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Print("Pilihan: ")
		fmt.Scan(&inputSort)
		if inputSort == 1 {
			sortWRF(arrWrf, "no_wrf", "ascending")
		} else if inputSort == 2 {
			sortWRF(arrWrf, "no_wrf", "descending")
		}
	} else if inputKategori == 2 {
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Print("Pilihan: ")
		fmt.Scan(&inputSort)
		if inputSort == 1 {
			sortWRF(arrWrf, "nama_site", "ascending")
		} else if inputSort == 2 {
			sortWRF(arrWrf, "nama_site", "descending")
		}
	} else if inputKategori == 3 {
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Print("Pilihan: ")
		fmt.Scan(&inputSort)
		if inputSort == 1 {
			sortWRF(arrWrf, "deskripsi_wrf", "ascending")
		} else if inputSort == 2 {
			sortWRF(arrWrf, "deskripsi_wrf", "descending")
		}
	} else if inputKategori == 4 {
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Print("Pilihan: ")
		fmt.Scan(&inputSort)
		if inputSort == 1 {
			sortWRF(arrWrf, "item_remark", "ascending")
		} else if inputSort == 2 {
			sortWRF(arrWrf, "item_remark", "descending")
		}
	} else if inputKategori == 5 {
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Print("Pilihan: ")
		fmt.Scan(&inputSort)
		if inputSort == 1 {
			sortWRF(arrWrf, "tanggal_kerja", "ascending")
		} else if inputSort == 2 {
			sortWRF(arrWrf, "tanggal_kerja", "descending")
		}
	} else if inputKategori == 6 {
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Print("Pilihan: ")
		fmt.Scan(&inputSort)
		if inputSort == 1 {
			sortWRF(arrWrf, "pelanggan", "ascending")
		} else if inputSort == 2 {
			sortWRF(arrWrf, "pelanggan", "descending")
		}

	} else if inputKategori == 7 {
		if in == "admin" {
			KelolaWRF(dataUser, oldArrWrf)
		} else {
			uiPengguna(dataUser, *arrWrf)
		}
	}

	if in == "admin" {
		lihatDaftarWrf(*arrWrf, oldArrWrf, "admin", dataUser, true)
	} else {
		lihatDaftarWrf(*arrWrf, oldArrWrf, "user", dataUser, true)
	}
}

func lihatDaftarWrf(arrWrf WRFs, oldArrWrf WRFs, in string, dataUser users, Lihatonly bool) {
	fmt.Println("-------------------------------------------------------------------------")
	for i := 0; i < NMAX; i++ {
		if arrWrf[i].no_wrf != "" {
			fmt.Println(i+1, ".	no wrf:", arrWrf[i].no_wrf)
			fmt.Println("	nama site:", arrWrf[i].nama_site)
			fmt.Println("	deskripsi wrf:", arrWrf[i].deskripsi_wrf)
			fmt.Println("	item remark:", arrWrf[i].item_remark)
			fmt.Println("	tanggal kerja:", arrWrf[i].tanggal_kerja.tanggal, "-", arrWrf[i].tanggal_kerja.bulan, "-", arrWrf[i].tanggal_kerja.tahun)
			fmt.Println("	pelanggan:", arrWrf[i].pelanggan)
			fmt.Println("	status:", arrWrf[i].status)
			fmt.Println()
		}
	}
	fmt.Println("-------------------------------------------------------------------------")

	if Lihatonly == true {
		if in == "admin" {
			KelolaWRF(dataUser, oldArrWrf)
		} else {
			uiPengguna(dataUser, oldArrWrf)
		}
	}
}

func cariWrf(arrWrf WRFs, datauser users, in string) {
	var input string
	var date date
	var found bool

	fmt.Println("cari berdasarkan: ")
	fmt.Println("1. Nomor wrf")
	fmt.Println("2. Nama site")
	fmt.Println("3. Deskripsi wrf")
	fmt.Println("4. Remark")
	fmt.Println("5. Tanggal Kerja")
	fmt.Println("6. Pelanggan")
	fmt.Println("7. Status")
	fmt.Println("8. Kembali")
	fmt.Scan(&input)

	if input == "1" {
		fmt.Print("Nomor wrf: ")
		fmt.Scan(&input)
		for i := 0; i < NMAX; i++ {
			if arrWrf[i].no_wrf == input {
				fmt.Println(i+1, ".	no wrf:", arrWrf[i].no_wrf)
				fmt.Println("	nama site:", arrWrf[i].nama_site)
				fmt.Println("	deskripsi wrf:", arrWrf[i].deskripsi_wrf)
				fmt.Println("	item remark:", arrWrf[i].item_remark)
				fmt.Println("	tanggal kerja:", arrWrf[i].tanggal_kerja.tanggal, "-", arrWrf[i].tanggal_kerja.bulan, "-", arrWrf[i].tanggal_kerja.tahun)
				fmt.Println("	pelanggan:", arrWrf[i].pelanggan)
				fmt.Println("	status:", arrWrf[i].status)
				fmt.Println()
				found = true
			}
		}
	} else if input == "2" {
		fmt.Print("Nama Site: ")
		fmt.Scan(&input)
		for i := 0; i < NMAX; i++ {
			if arrWrf[i].nama_site == input {
				fmt.Println(i+1, ".	no wrf:", arrWrf[i].no_wrf)
				fmt.Println("	nama site:", arrWrf[i].nama_site)
				fmt.Println("	deskripsi wrf:", arrWrf[i].deskripsi_wrf)
				fmt.Println("	item remark:", arrWrf[i].item_remark)
				fmt.Println("	tanggal kerja:", arrWrf[i].tanggal_kerja.tanggal, "-", arrWrf[i].tanggal_kerja.bulan, "-", arrWrf[i].tanggal_kerja.tahun)
				fmt.Println("	pelanggan:", arrWrf[i].pelanggan)
				fmt.Println("	status:", arrWrf[i].status)
				fmt.Println()
				found = true
			}
		}
	} else if input == "3" {
		fmt.Print("Deskripsi wrf: ")
		fmt.Scan(&input)
		for i := 0; i < NMAX; i++ {
			if arrWrf[i].deskripsi_wrf == input {
				fmt.Println(i+1, ".	no wrf:", arrWrf[i].no_wrf)
				fmt.Println("	nama site:", arrWrf[i].nama_site)
				fmt.Println("	deskripsi wrf:", arrWrf[i].deskripsi_wrf)
				fmt.Println("	item remark:", arrWrf[i].item_remark)
				fmt.Println("	tanggal kerja:", arrWrf[i].tanggal_kerja.tanggal, "-", arrWrf[i].tanggal_kerja.bulan, "-", arrWrf[i].tanggal_kerja.tahun)
				fmt.Println("	pelanggan:", arrWrf[i].pelanggan)
				fmt.Println("	status:", arrWrf[i].status)
				fmt.Println()
				found = true
			}
		}

	} else if input == "4" {
		fmt.Print("Item remark: ")
		fmt.Scan(&input)
		for i := 0; i < NMAX; i++ {
			if arrWrf[i].item_remark == input {
				fmt.Println(i+1, ".	no wrf:", arrWrf[i].no_wrf)
				fmt.Println("	nama site:", arrWrf[i].nama_site)
				fmt.Println("	deskripsi wrf:", arrWrf[i].deskripsi_wrf)
				fmt.Println("	item remark:", arrWrf[i].item_remark)
				fmt.Println("	tanggal kerja:", arrWrf[i].tanggal_kerja.tanggal, "-", arrWrf[i].tanggal_kerja.bulan, "-", arrWrf[i].tanggal_kerja.tahun)
				fmt.Println("	pelanggan:", arrWrf[i].pelanggan)
				fmt.Println("	status:", arrWrf[i].status)
				fmt.Println()
				found = true
			}
		}

	} else if input == "5" {
		fmt.Print("Tanggal: ")
		fmt.Scan(&date.tanggal, &date.bulan, &date.tahun)
		for i := 0; i < NMAX; i++ {
			if arrWrf[i].tanggal_kerja.tanggal == date.tanggal && arrWrf[i].tanggal_kerja.bulan == date.bulan && arrWrf[i].tanggal_kerja.tahun == date.tahun {
				fmt.Println(i+1, ".	no wrf:", arrWrf[i].no_wrf)
				fmt.Println("	nama site:", arrWrf[i].nama_site)
				fmt.Println("	deskripsi wrf:", arrWrf[i].deskripsi_wrf)
				fmt.Println("	item remark:", arrWrf[i].item_remark)
				fmt.Println("	tanggal kerja:", arrWrf[i].tanggal_kerja.tanggal, "-", arrWrf[i].tanggal_kerja.bulan, "-", arrWrf[i].tanggal_kerja.tahun)
				fmt.Println("	pelanggan:", arrWrf[i].pelanggan)
				fmt.Println("	status:", arrWrf[i].status)
				fmt.Println()
				found = true
			}
		}

	} else if input == "6" {
		fmt.Print("Pelanggan: ")
		fmt.Scan(&input)
		for i := 0; i < NMAX; i++ {
			if arrWrf[i].pelanggan == input {
				fmt.Println(i+1, ".	no wrf:", arrWrf[i].no_wrf)
				fmt.Println("	nama site:", arrWrf[i].nama_site)
				fmt.Println("	deskripsi wrf:", arrWrf[i].deskripsi_wrf)
				fmt.Println("	item remark:", arrWrf[i].item_remark)
				fmt.Println("	tanggal kerja:", arrWrf[i].tanggal_kerja.tanggal, "-", arrWrf[i].tanggal_kerja.bulan, "-", arrWrf[i].tanggal_kerja.tahun)
				fmt.Println("	pelanggan:", arrWrf[i].pelanggan)
				fmt.Println("	status:", arrWrf[i].status)
				fmt.Println()
				found = true
			}
		}

	} else if input == "7" {
		fmt.Print("Status: ")
		fmt.Scan(&input)
		for i := 0; i < NMAX; i++ {
			if arrWrf[i].status == input {
				fmt.Println(i+1, ".	no wrf:", arrWrf[i].no_wrf)
				fmt.Println("	nama site:", arrWrf[i].nama_site)
				fmt.Println("	deskripsi wrf:", arrWrf[i].deskripsi_wrf)
				fmt.Println("	item remark:", arrWrf[i].item_remark)
				fmt.Println("	tanggal kerja:", arrWrf[i].tanggal_kerja.tanggal, "-", arrWrf[i].tanggal_kerja.bulan, "-", arrWrf[i].tanggal_kerja.tahun)
				fmt.Println("	pelanggan:", arrWrf[i].pelanggan)
				fmt.Println("	status:", arrWrf[i].status)
				fmt.Println()
				found = true
			}
		}
	} else {
		if in == "admin" {
			KelolaWRF(datauser, arrWrf)
		} else {
			uiPengguna(datauser, arrWrf)
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan")
	}

	if in == "admin" {
		KelolaWRF(datauser, arrWrf)
	} else {
		uiPengguna(datauser, arrWrf)
	}
}

func ubahDataWrf(arrWrf *WRFs, dataUser users) {
	var num int
	var isFound bool
	var input string
	var data WRF

	fmt.Println("ubah wrf")

	lihatDaftarWrf(*arrWrf, *arrWrf, "admin", dataUser, false)
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Print("masukkan nomor user yang ingin di ubah: ")
	fmt.Scan(&num)
	//check if the num is in valid value
	if num >= 1 && num <= NMAX {
		index := findWRFindex(*arrWrf, num, &isFound)

		if isFound {
			fmt.Println(index+1, ".	no wrf:", arrWrf[index].no_wrf)
			fmt.Println("	nama site:", arrWrf[index].nama_site)
			fmt.Println("	deskripsi wrf:", arrWrf[index].deskripsi_wrf)
			fmt.Println("	item remark:", arrWrf[index].item_remark)
			fmt.Println("	tanggal kerja:", arrWrf[index].tanggal_kerja.tanggal, "-", arrWrf[index].tanggal_kerja.bulan, "-", arrWrf[index].tanggal_kerja.tahun)
			fmt.Println("	pelanggan:", arrWrf[index].pelanggan)
			fmt.Println("	status:", arrWrf[index].status)
			fmt.Println()

			fmt.Print("Ingin mengubah wrf tersebut? y/n: ")
			fmt.Scan(&input)
			if input == "y" {
				fmt.Println("Input 'cancel' untuk batalkan pengubahan")
				fmt.Println("Ubah menjadi: ")

				// Prompt for and validate the wrf no
				cancelled := false
				no_wrfFilled := false
				for !no_wrfFilled {
					isFound = false
					fmt.Print("no wrf: ")
					fmt.Scan(&data.no_wrf)
					if data.no_wrf == "cancel" {
						fmt.Println("Batal mengubah wrf.")
						cancelled = true
						no_wrfFilled = true
					} else {
						isFound = findWRFNo(*arrWrf, data.no_wrf)
						if isFound {
							if arrWrf[index].no_wrf == data.no_wrf {
								no_wrfFilled = true
							} else {
								fmt.Println("Nomor wrf sudah ada. Masukkan nomor wrf yang berbeda.")
							}
						} else {
							no_wrfFilled = true
						}
					}
				}

				if !cancelled {
					// Prompt for and validate the nama_site
					namaSiteFilled := false
					for !namaSiteFilled {
						fmt.Print("nama site: ")
						data.nama_site = inputString()
						if data.nama_site == "cancel" {
							fmt.Println("Batal mengubah wrf.")
							cancelled = true
							namaSiteFilled = true
						} else {
							namaSiteFilled = true
						}
					}
				}

				if !cancelled {
					// Prompt for and validate the deskripsi_wrf
					deskripsiFilled := false
					for !deskripsiFilled {
						fmt.Print("deskripsi: ")
						data.deskripsi_wrf = inputString()
						if data.deskripsi_wrf == "cancel" {
							fmt.Println("Batal mengubah wrf.")
							cancelled = true
							deskripsiFilled = true
						} else {
							deskripsiFilled = true
						}
					}
				}

				if !cancelled {
					// Prompt for and validate the item_remark
					remarkFilled := false
					for !remarkFilled {
						fmt.Print("item remark: ")
						data.item_remark = inputString()
						if data.item_remark == "cancel" {
							fmt.Println("Batal mengubah wrf.")
							cancelled = true
							remarkFilled = true
						} else {
							remarkFilled = true
						}
					}
				}

				if !cancelled {
					// Prompt for and validate the tanggal_kerja
					tanggalFilled := false
					for !tanggalFilled {
						fmt.Print("tanggal kerja: ")
						fmt.Scan(&data.tanggal_kerja.tanggal, &data.tanggal_kerja.bulan, &data.tanggal_kerja.tahun)
						if !isValidDate(data.tanggal_kerja.tanggal, data.tanggal_kerja.bulan, data.tanggal_kerja.tahun) {
							fmt.Println("Pastikan tanggal yang di input benar.")
						} else {
							tanggalFilled = true
						}
					}
				}

				if !cancelled {
					// Prompt for and validate the pelanggan
					pelangganFilled := false
					for !pelangganFilled {
						fmt.Print("pelanggan: ")
						fmt.Scan(&data.pelanggan)
						if data.pelanggan == "cancel" {
							fmt.Println("Batal mengubah wrf.")
							cancelled = true
							pelangganFilled = true
						} else {
							pelangganFilled = true
						}
					}
				}

				if !cancelled {
					// Prompt for and validate the Status
					statusFilled := false
					for !statusFilled {
						fmt.Print("Status: ")
						fmt.Scan(&data.status)
						if data.status == "cancel" {
							fmt.Println("Batal mengubah wrf.")
							cancelled = true
							statusFilled = true
						} else if data.status == "open" || data.status == "closed" {
							statusFilled = true
						} else {
							fmt.Println("Status yang tersedia hanya open atau closed")
						}
					}
				}

				if !cancelled {
					arrWrf[index].no_wrf = data.no_wrf
					arrWrf[index].nama_site = data.nama_site
					arrWrf[index].deskripsi_wrf = data.deskripsi_wrf
					arrWrf[index].item_remark = data.item_remark
					arrWrf[index].tanggal_kerja = date{tanggal: data.tanggal_kerja.tanggal, bulan: data.tanggal_kerja.bulan, tahun: data.tanggal_kerja.tahun}
					arrWrf[index].pelanggan = data.pelanggan
					arrWrf[index].status = data.status
					fmt.Println("wrf berhasil di ubah.")
					fmt.Println("-------------------------------------------------------------------------")
				}
			} else {
				fmt.Println("Kembali")
			}
		} else {
			fmt.Println("Wrf tidak ditemukan")
		}
	} else {
		fmt.Println("Nomor tidak valid")
	}
}

func tambahDataWrf(arrWrf *WRFs, dataUser users) {
	fmt.Println("tambah data wrf")
	var isFound bool
	var data WRF
	var n int

	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println("Tambah User")

	// Prompt for and validate the no_wrf

	fmt.Println("Input 'cancel' untuk batalkan penambahan wrf")
	cancelled := false
	no_wrfFilled := false
	for !no_wrfFilled {
		isFound = false
		fmt.Print("no wrf: ")
		fmt.Scan(&data.no_wrf)
		if data.no_wrf == "cancel" {
			fmt.Println("Batal membuat wrf baru.")
			cancelled = true
			no_wrfFilled = true
		} else {
			isFound = findWRFNo(*arrWrf, data.no_wrf)
			if isFound {
				fmt.Println("Nomor wrf sudah ada. Masukkan nomor wrf yang berbeda.")
			} else {
				no_wrfFilled = true
			}
		}
	}

	if !cancelled {
		// Prompt for and validate the nama_site
		namaSiteFilled := false
		for !namaSiteFilled {
			fmt.Print("nama site: ")
			data.nama_site = inputString()
			if data.nama_site == "cancel" {
				fmt.Println("Batal membuat wrf baru.")
				cancelled = true
				namaSiteFilled = true
			} else {
				namaSiteFilled = true
			}
		}
	}

	if !cancelled {
		// Prompt for and validate the deskripsi_wrf
		deskripsiFilled := false
		for !deskripsiFilled {
			fmt.Print("deskripsi: ")
			data.deskripsi_wrf = inputString()
			if data.deskripsi_wrf == "cancel" {
				fmt.Println("Batal membuat wrf baru.")
				cancelled = true
				deskripsiFilled = true
			} else {
				deskripsiFilled = true
			}
		}
	}

	if !cancelled {
		// Prompt for and validate the item_remark
		remarkFilled := false
		for !remarkFilled {
			fmt.Print("item remark: ")
			data.item_remark = inputString()
			if data.item_remark == "cancel" {
				fmt.Println("Batal membuat wrf baru.")
				cancelled = true
				remarkFilled = true
			} else {
				remarkFilled = true
			}
		}
	}

	if !cancelled {
		// Prompt for and validate the tanggal_kerja
		tanggalFilled := false
		for !tanggalFilled {
			data.tanggal_kerja.tanggal, data.tanggal_kerja.bulan, data.tanggal_kerja.tahun = 0, 0, 0
			fmt.Print("tanggal kerja: ")
			fmt.Scan(&data.tanggal_kerja.tanggal, &data.tanggal_kerja.bulan, &data.tanggal_kerja.tahun)
			if !isValidDate(data.tanggal_kerja.tanggal, data.tanggal_kerja.bulan, data.tanggal_kerja.tahun) {
				fmt.Println("Pastikan tanggal yang di input benar.")
			} else {
				tanggalFilled = true
			}
		}
	}

	if !cancelled {
		// Prompt for and validate the pelanggan
		pelangganFilled := false
		for !pelangganFilled {
			fmt.Print("pelanggan: ")
			fmt.Scan(&data.pelanggan)
			if data.pelanggan == "cancel" {
				fmt.Println("Batal membuat wrf baru.")
				cancelled = true
				pelangganFilled = true
			} else {
				pelangganFilled = true
			}
		}
	}

	if !cancelled {
		// Prompt for and validate the Status
		statusFilled := false
		for !statusFilled {
			fmt.Print("Status: ")
			fmt.Scan(&data.status)
			if data.status == "cancel" {
				fmt.Println("Batal membuat wrf baru.")
				cancelled = true
				statusFilled = true
			} else if data.status == "open" || data.status == "closed" {
				statusFilled = true
			} else {
				fmt.Println("Status yang tersedia hanya open atau closed")
			}
		}
	}

	if !cancelled {
		// Find an empty slot to add the new user
		n = 0
		for n < NMAX-1 && arrWrf[n].no_wrf != "" {
			n++
		}

		if n == NMAX-1 {
			fmt.Println("Tidak dapat menambahkan wrf. Kapasitas maksimum tercapai.")
		} else {
			arrWrf[n].no_wrf = data.no_wrf
			arrWrf[n].nama_site = data.nama_site
			arrWrf[n].deskripsi_wrf = data.deskripsi_wrf
			arrWrf[n].item_remark = data.item_remark
			arrWrf[n].tanggal_kerja = date{tanggal: data.tanggal_kerja.tanggal, bulan: data.tanggal_kerja.bulan, tahun: data.tanggal_kerja.tahun}
			arrWrf[n].pelanggan = data.pelanggan
			arrWrf[n].status = data.status
			fmt.Println("wrf berhasil ditambahkan.")
		}
		fmt.Println("-------------------------------------------------------------------------")
	}

	KelolaWRF(dataUser, *arrWrf)
}

func inputString() string {
	var ascii int
	var input byte
	var theString string

	fmt.Scanf("%c", &input)
	ascii = int(input)
	for ascii != 13 {
		theString += string(rune(ascii))
		fmt.Scanf("%c", &input)
		ascii = int(input)
	}
	return strings.ReplaceAll(theString, "\n", "")
}

func isValidDate(tanggal int, bulan int, tahun int) bool {
	var tanggalBulan int = 31

	switch bulan {
	case 4, 6, 9, 11:
		tanggalBulan = 30
	case 2:
		if tahun%4 == 0 {
			if tahun%100 == 0 || tahun%400 == 0 {
				tanggalBulan = 29
			} else {
				tanggalBulan = 28
			}
		}
	}
	if tahun < 0 || tahun > 9999 || bulan < 1 || bulan > 12 || tanggal < 1 || tanggal > tanggalBulan {
		return false
	}
	return true
}

func hapusDataWrf(arrWrf *WRFs, dataUser users) {
	var num int
	var isFound bool
	var input string

	fmt.Println("hapus data wrf")
	lihatDaftarWrf(*arrWrf, *arrWrf, "admin", dataUser, false)

	fmt.Println("-------------------------------------------------------------------------")
	fmt.Print("Masukkan nomor user yang ingin dihapus: ")
	fmt.Scan(&num)

	// Check if the index is within valid range
	if num >= 1 && num <= NMAX {
		index := findWRFindex(*arrWrf, num, &isFound)

		// Check if the user is found
		if isFound {
			fmt.Println(isFound)
			// Confirm deletion
			fmt.Println(index+1, ".	no wrf:", arrWrf[index].no_wrf)
			fmt.Println("	nama site:", arrWrf[index].nama_site)
			fmt.Println("	deskripsi wrf:", arrWrf[index].deskripsi_wrf)
			fmt.Println("	item remark:", arrWrf[index].item_remark)
			fmt.Println("	tanggal kerja:", arrWrf[index].tanggal_kerja.tanggal, "-", arrWrf[index].tanggal_kerja.bulan, "-", arrWrf[index].tanggal_kerja.tahun)
			fmt.Println("	pelanggan:", arrWrf[index].pelanggan)
			fmt.Println("	status:", arrWrf[index].status)
			fmt.Println()
			fmt.Print("Lanjutkan penghapusan? y/n: ")
			fmt.Scan(&input)

			if input == "y" {
				// Shift the indexes of other wrf data
				for i := index; i < NMAX-1; i++ {
					arrWrf[i] = arrWrf[i+1]
				}
				// Clear the last wrf data
				arrWrf[NMAX-1] = WRF{}
				fmt.Println("WRF berhasil dihapus.")
			} else {
				fmt.Println("Penghapusan dibatalkan.")
			}
		} else {
			fmt.Println("WRF tidak ditemukan")
		}
	} else {
		fmt.Println("Nomor WRF tidak valid.")
	}
	fmt.Println("-------------------------------------------------------------------------")
}

func sortWRF(arrWrf *WRFs, flag string, dasar string) {
	n := NMAX

	for i := 0; i < n-1; i++ {
		minIndex := i

		if dasar == "ascending" {
			if flag == "no_wrf" {
				for j := i + 1; j < n; j++ {
					if arrWrf[j].no_wrf != "" {
						if arrWrf[j].no_wrf < arrWrf[minIndex].no_wrf {
							minIndex = j
						}
					}
				}
			} else if flag == "nama_site" {
				for j := i + 1; j < n; j++ {
					if arrWrf[j].no_wrf != "" {
						if arrWrf[j].nama_site < arrWrf[minIndex].nama_site {
							minIndex = j
						}
					}
				}
			} else if flag == "deskripsi_wrf " {
				for j := i + 1; j < n; j++ {
					if arrWrf[j].no_wrf != "" {
						if arrWrf[j].deskripsi_wrf < arrWrf[minIndex].deskripsi_wrf {
							minIndex = j
						}
					}
				}
			} else if flag == "item_remark " {
				for j := i + 1; j < n; j++ {
					if arrWrf[j].no_wrf != "" {
						if arrWrf[j].item_remark < arrWrf[minIndex].item_remark {
							minIndex = j
						}
					}
				}
			} else if flag == "tanggal_kerja" {
				for j := i + 1; j < n; j++ {
					if arrWrf[j].no_wrf != "" {
						if arrWrf[j].tanggal_kerja.tahun < arrWrf[minIndex].tanggal_kerja.tahun ||
							(arrWrf[j].tanggal_kerja.tahun == arrWrf[minIndex].tanggal_kerja.tahun && arrWrf[j].tanggal_kerja.bulan < arrWrf[minIndex].tanggal_kerja.bulan) ||
							(arrWrf[j].tanggal_kerja.tahun == arrWrf[minIndex].tanggal_kerja.tahun && arrWrf[j].tanggal_kerja.bulan == arrWrf[minIndex].tanggal_kerja.bulan && arrWrf[j].tanggal_kerja.tanggal < arrWrf[minIndex].tanggal_kerja.tanggal) {
							minIndex = j
						}
					}
				}
			} else if flag == "pelanggan" {
				for j := i + 1; j < n; j++ {
					if arrWrf[j].no_wrf != "" {
						if arrWrf[j].pelanggan < arrWrf[minIndex].pelanggan {
							minIndex = j
						}
					}
				}
			}
		} else if dasar == "descending" {
			if flag == "no_wrf" {
				for j := i + 1; j < n; j++ {
					if arrWrf[j].no_wrf > arrWrf[minIndex].no_wrf {
						minIndex = j
					}
				}
			} else if flag == "nama_site" {
				for j := i + 1; j < n; j++ {
					if arrWrf[j].nama_site > arrWrf[minIndex].nama_site {
						minIndex = j
					}
				}
			} else if flag == "deskripsi_wrf " {
				for j := i + 1; j < n; j++ {
					if arrWrf[j].deskripsi_wrf > arrWrf[minIndex].deskripsi_wrf {
						minIndex = j
					}
				}
			} else if flag == "item_remark " {
				for j := i + 1; j < n; j++ {
					if arrWrf[j].item_remark > arrWrf[minIndex].item_remark {
						minIndex = j
					}
				}
			} else if flag == "tanggal_kerja" {
				for j := i + 1; j < n; j++ {
					if arrWrf[j].tanggal_kerja.tahun > arrWrf[minIndex].tanggal_kerja.tahun ||
						(arrWrf[j].tanggal_kerja.tahun == arrWrf[minIndex].tanggal_kerja.tahun && arrWrf[j].tanggal_kerja.bulan > arrWrf[minIndex].tanggal_kerja.bulan) ||
						(arrWrf[j].tanggal_kerja.tahun == arrWrf[minIndex].tanggal_kerja.tahun && arrWrf[j].tanggal_kerja.bulan == arrWrf[minIndex].tanggal_kerja.bulan && arrWrf[j].tanggal_kerja.tanggal > arrWrf[minIndex].tanggal_kerja.tanggal) {
						minIndex = j
					}
				}
			} else if flag == "pelanggan" {
				for j := i + 1; j < n; j++ {
					if arrWrf[j].pelanggan > arrWrf[minIndex].pelanggan {
						minIndex = j
					}
				}
			}
		}

		if minIndex != i {
			arrWrf[i], arrWrf[minIndex] = arrWrf[minIndex], arrWrf[i]
		}
	}
}
