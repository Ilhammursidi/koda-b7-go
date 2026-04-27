package main

import (
	"fmt"

	"github.com/ilhammursidi/koda-b7-go/internals/minitask1"
	"github.com/ilhammursidi/koda-b7-go/internals/minitask2"
	"github.com/ilhammursidi/koda-b7-go/internals/minitask3"
	"github.com/ilhammursidi/koda-b7-go/internals/minitask4"
)

func main() {
	for {
		fmt.Println("===============MENU================")
		fmt.Println("1. Luas dan Keliling Lingkaran")
		fmt.Println("2. Segigita")
		fmt.Println("3. Append Slice")
		fmt.Println("4. User")
		fmt.Println("0. Keluar")
		fmt.Println("===================================")

		var s uint
		fmt.Print("Pilih : ")
		fmt.Scan(&s)
		if s != 1 && s != 2 && s != 3 && s != 4 && s != 0 {
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
			} else if o != 0 && o != 1 && o != 2 && o != 3 && o != 4 {
				fmt.Println("Invalid Input")
				return
			}

		} else if s == 2 {
			// minitask 2
			var n int
			fmt.Print("Input angka : ")
			fmt.Scan(&n)
			minitask2.Triangle(n)
			fmt.Println("===================================")

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
		} else if s == 0 {
			fmt.Println("Terima Kasih")
			return
		}

	}

}
