package main

import "fmt"

func main() {
	//hello world
	fmt.Println("hello, world!")



	//string
	var nama = "dimas" //string menggunakan 2 tanda kutip
	fmt.Println(nama)


	//const
	const kota = "tokyo"
	fmt.Println(kota)


	//var
	var ini = "example" //variable diawali dengan kata var lalu diikuti dengan nama variabel lalu isi variabel
	fmt.Println(ini)

	//perbandingan
	var a = 10
	var b = 20
	var c = a == b
	fmt.Println(c)

	var x = 20
	var y = 10
	var z = x != y
	fmt.Println(z)


	//operasi matematika

	i := 20
	i++
	fmt.Println(i)

	i += 20
	fmt.Println(i)



	//type declaration
	type str string
	type num int

	var name str = "dimas"
	var nomor num = 10
	fmt.Println(name)
	fmt.Println(nomor)



	//boolean

	var menikah = true
	fmt.Println("sudah menikah?", menikah)
	//conversion
	var jumlah = "sepuluh"
	var e = jumlah[0]
	var ab = string(e)
	fmt.Println(jumlah)
	fmt.Println(ab)

	//type data number

	var nilai32 int32 = 1201201
	fmt.Println(nilai32)
}

