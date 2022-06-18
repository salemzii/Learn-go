package main 

import (
	"fmt"
	"runtime"
	"sync"
)


var (
	wg sync.WaitGroup
	counter int
)


func main(){
	runtime.GOMAXPROCS(1)
	
	wg.Add(2)
	fmt.Println("Start goroutines")
	fmt.Println("Value of counter: ", counter)

	go incCouter(1)
	go incCouter(2)

	fmt.Println("waiting to finish")
	wg.Wait()
	fmt.Println("Final counter: ", counter)
}


func incCouter(id int){

	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter

		runtime.Gosched()

		value++
		counter = value
	}

}

func printPrime(prefix string){
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}