package main

func main() {
	type NoKtp string

	var ktpAgung NoKtp = "123456786542"
	var contoh string = "2456789034423"
	var contohKtp = NoKtp(contoh)

	println(ktpAgung)
	println(contohKtp)

}