package main

import (
	"fmt"
	"sync"

	"github.com/ilhammursidi/koda-b7-go/internals/goroutine"
	"github.com/ilhammursidi/koda-b7-go/internals/minitask1"
	"github.com/ilhammursidi/koda-b7-go/internals/minitask2"
	"github.com/ilhammursidi/koda-b7-go/internals/minitask3"
	"github.com/ilhammursidi/koda-b7-go/internals/minitask4"
	"github.com/ilhammursidi/koda-b7-go/internals/minitask6"
	"github.com/ilhammursidi/koda-b7-go/internals/minitask7"
	"github.com/ilhammursidi/koda-b7-go/internals/minitask8"
)

func main() {
	for {
		fmt.Println("===============MENU================")
		fmt.Println("1. Luas dan Keliling Lingkaran")
		fmt.Println("2. Segigita")
		fmt.Println("3. Append Slice")
		fmt.Println("4. User")
		fmt.Println("5. Open and Read File")
		fmt.Println("6. Struct Person")
		fmt.Println("7. Pembayaran")
		fmt.Println("8. CoffeShop")
		fmt.Println("9. Bangun Pagi")
		fmt.Println("10. Papan Tulis")
		fmt.Println("11. KuadratLoop")
		fmt.Println("0. Keluar")
		fmt.Println("===================================")

		var s uint
		fmt.Print("Pilih : ")
		fmt.Scan(&s)
		if s != 1 && s != 2 && s != 3 && s != 4 && s != 5 && s != 6 && s != 7 && s != 8 && s != 9 && s != 10 && s != 11 && s != 0 {
			fmt.Println("Input Tidak Sesuai")
			return
		} else if s == 1 {
			var o int
			fmt.Println("==========MENU==========")
			fmt.Println("1. Hitung Luas lingkaran")
			fmt.Println("2. Hitung Lingkaran")
			fmt.Println("3. Hitung Luas dan Keliling Lingkaran")
			fmt.Println("0. Kembali")
			fmt.Print("Pilih Opsi:")
			fmt.Scan(&o)
			if o == 1 {
				// minitask 1
				var t float32
				fmt.Print("Input angka untuk Luas lingkaran: ")
				fmt.Scan(&t)
				luas := minitask1.Luas(t)
				fmt.Printf("Luas lingkaran = %.2f\n", luas)
				fmt.Println("===================================")
			} else if o == 2 {
				var x float32
				fmt.Print("Input angka untuk Keliling lingkaran: ")
				fmt.Scan(&x)
				keliling := minitask1.Keliling(x)
				fmt.Printf("Keliling = %.2f\n", keliling)
				fmt.Println("===================================")

			} else if o == 3 {

				var j float32
				fmt.Print("Input angka untuk Luas dan Keliling lingkaran: ")
				fmt.Scan(&j)
				luasLingkaran, kelilingLingkaran := minitask1.LuasDanKeliling(j)
				fmt.Printf("Luas %.2f, keliling %.2f\n", luasLingkaran, kelilingLingkaran)
				fmt.Println("===================================")

			} else if o == 0 {
				return
			} else if o != 0 && o != 1 && o != 2 && o != 3 {
				fmt.Println("Invalid Input")
			}

		} else if s == 2 {
			// minitask 2
			var n int
			fmt.Print("Input angka : ")
			fmt.Scan(&n)
			minitask2.Triangle(n)
			fmt.Println("===================================")
			goroutine.BangunPagi()
		} else if s == 3 {
			// minitask 3
			var p int
			fmt.Print("Input untuk slice : ")
			fmt.Scan(&p)
			minitask3.AddNumber(p)
			fmt.Println("===================================")

		} else if s == 4 {
			// minitask 4
			ilham := minitask4.User{
				// id:        1,
				Name:      "Ilham Mursidi",
				Image:     "alim",
				Email:     "imursidi37@gmail.com",
				Umur:      26,
				Phone:     "082189971471",
				IsMarried: false,
				Education: minitask4.Education{Name: "SMKN 1 Minasatene", Major: "Teknik Komputer dan Jaringan"},
			}
			fmt.Println(ilham)
		} else if s == 5 {
			var input string
			fmt.Print("Ketik file path: ")
			fmt.Scan(&input)
			minitask6.OpenFile(input)

		} else if s == 6 {
			ilham := minitask7.NewUser("ilham", "bogor", "0888")

			ilham.Cetak()
			ilham.Greet(ilham.Name)
			ilham.SetName("apa")
			ilham.Greet(ilham.Name)

		} else if s == 7 {
			bank := &minitask8.Bank{BankName: "BCA"}
			online := &minitask8.Online{EbankName: "Dana"}
			fiktif := &minitask8.FictivePayment{}

			prices := []int{15000, 20000, 75000}

			fmt.Println("via Bank")
			minitask8.MethodPay(bank, prices)
			fmt.Println()

			fmt.Println("via Online")
			minitask8.MethodPay(online, prices)
			fmt.Println()

			fiktifPrices := []int{30000, -50000, 120000, 0, 80000}

			fmt.Println("Pembayaran Fiktiv")
			for _, price := range fiktifPrices {
				err := fiktif.Pay(price)
				if err != nil {
					fmt.Printf("[ERROR] Gagal menyimpan pembayaran %d: %v\n", price, err)
				}
			}

			fmt.Println()
			fmt.Println("Data Pembayaran Fiktif")
			fmt.Println(fiktif.GetRecords())

		} else if s == 8 {
			goroutine.CoffeShop()
			fmt.Println("===================================")

		} else if s == 9 {
			goroutine.BangunPagi()
			fmt.Println("===================================")

		} else if s == 10 {
			board := make(chan string)
			var wg sync.WaitGroup
			messages := []struct {
				name    string
				message string
			}{
				{"Ayah", "Jangan lupa beli sayur"},
				{"Ibu", "Siapkan PR anak"},
				{"Adik", "Aku mau main game"},
				{"Kakak", "Nonton film malam ini"},
			}
			go goroutine.WhiteBoard(board)
			for _, m := range messages {
				wg.Add(1)
				go goroutine.Sender(board, &wg, m.name, m.message)
			}
			wg.Wait()
			close(board)
		} else if s == 11 {
			var n int
			fmt.Print("Input Angka :")
			fmt.Scan(&n)
			N := goroutine.GetNumber(n)
			Genap := goroutine.Genap(N)
			Hasil := goroutine.Kuadrat(Genap)
			for result := range Hasil {
				fmt.Println(result)
			}

		} else if s == 0 {
			fmt.Println("Terima Kasih")
			fmt.Println("===================================")

			return
		}

	}
}
