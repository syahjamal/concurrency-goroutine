package main

import (
	"fmt"
	"time"
)

func printAngka() {
	for i := 0; i < 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println(i)
	}
}

func printText() {
	for i := 0; i < 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("text ", i)
	}
}

func main() {
	start := time.Now() //cel kecepatan proses

	go printAngka()
	go printText()

	time.Sleep(1000 * time.Millisecond)

	fmt.Println(time.Since(start)) //cek kecepatan proses
}
