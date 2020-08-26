package main

import (
	"fmt"
)

func main() {
	chann := make(chan string) //channel

	go runChannel("Salam", chann)

	for msg := range chann {
		fmt.Println(msg)
	}
}

func runChannel(text string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- text
	}

	close(c)
}

//===== Menutup Channel =====
/*Casenya memiliki channel yang berulang 5 kali dan di dalam main
channelnya berulang tanpa batas*/
