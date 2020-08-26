package main

import (
	"fmt"
	"time"
)

//Declare global itemschannel
var itemsChannel = make(chan string) //channel 1

var cleanedItemsChannel = make(chan string) //channel 2

func main() {
	items := [7]string{"batu", "harta", "kerang", "harta", "batu", "kerang"}

	go menyelam(items)
	go membersihkan()
	go menyimpan()

	time.Sleep(500 * time.Millisecond)
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
		cleanedItemsChannel <- "hartaBersih"
	}
}

func menyimpan() {
	for cleanedItem := range cleanedItemsChannel {
		fmt.Println("Berhasil menyimpan " + cleanedItem)
	}
}

//=====Konsep Channel=====
//Kita mau menyelam mencari harta karun,
//terselip diantara banyak barang
//Yang bertugas di kapal
//1.Dicari, 2. Dibersihkan, 3.Disimpan oleh penyelam
