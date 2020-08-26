package main

import "fmt"

//Declare global itemschannel
var itemsChannel = make(chan string)

func main() {
	items := [7]string{"batu", "harta", "kerang", "harta", "batu", "kerang"}

	go menyelam(items)
	go membersihkan()
}

func menyelam(items [7]string) {
	for _, item := range items {
		if item == "harta" {
			fmt.Println("Berhasil mendapatkan " + item)
			itemsChannel <- item
		}
	}
}

func membersihkan() {
	for rawItem := range itemsChannel {
		fmt.Println("berhasil membersihkan " + rawItem)
	}
}

//=====Konsep Channel=====
//Kita mau menyelam mencari harta karun,
//terselip diantara banyak barang
//Yang bertugas di kapal
//1.Dicari, 2. Dibersihkan, 3.Disimpan oleh penyelam
