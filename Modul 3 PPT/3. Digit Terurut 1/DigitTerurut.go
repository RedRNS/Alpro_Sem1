package main

import "fmt"

func main() {
	var bil int
	var ratusan, puluhan, satuan int
	var urut bool

	fmt.Print("Masukan sebuah bilangan ratusan : ")
	fmt.Scanln(&bil)

	ratusan = bil / 100 % 10
	puluhan = bil / 10 % 10
	satuan = bil % 10

	urut = (ratusan > puluhan && puluhan > satuan && ratusan > satuan) ||
		(satuan > puluhan && puluhan > ratusan && satuan > ratusan)

	fmt.Println("Urut :", urut)
}
