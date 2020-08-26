package main

import (
	"fmt"
	"time"
)

func printSalam(text string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(text)
	}
}

func main() {
	start := time.Now() //cel kecepatan proses

	go printSalam("Selamat Datang !") //goroutine ditandai dengan code go di awal func
	printSalam("Selamat Jalan")

	fmt.Println(time.Since(start)) //cek kecepatan proses
}
