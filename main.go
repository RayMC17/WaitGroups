package main

import (
	"fmt"
	"sync"
)
func one(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("hola")
}

func two( wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Ni hao")

}

	func three( wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Hello")

}
func main(){
	var wg sync.WaitGroup
	wg.Add(3)
	go one(&wg)
	go two(&wg)
	go three(&wg)



	//lets Delay

	wg.Wait()
}