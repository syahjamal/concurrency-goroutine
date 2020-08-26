package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//Add (Berapa go routin)
	//Wait (Menunggu go routine)
	//Done
	wg.Add(2)
	go printText("Salam", &wg)
	go printText("Hallo", &wg)

	wg.Wait()
}

func printText(text string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
	}
	wg.Done()
}

//===== Wait Go routine with sync waitgroup =====
/*Casenya memiliki channel yang berulang 5 kali dan di dalam main
channelnya berulang tanpa batas*/
