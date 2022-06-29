package main 

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"math/rand"
	"time"
)


var (
	wg sync.WaitGroup
	counter int64
)

const (
	numberOfGoroutines = 4 
	taskLoad = 10
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

func main(){

	tasks := make(chan string, taskLoad)

	wg.Add(numberOfGoroutines)

	for gr := 1; gr <= numberOfGoroutines; gr++{
		go worker(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}


	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, goroutine int){
	defer wg.Done()

	for {
		task, ok := <-tasks

		if !ok {
			fmt.Printf("Worker: %d : Shutting Dwon\n", worker)
			return
		}
		fmt.Printf("Worker: %d : Started %s\n", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)			
	}

}















func Runner(baton chan int){
	var newRunner int

	runner := <-baton
	fmt.Printf("Runner %d Running With Baton\n", runner)
	
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		
		
		go Runner(baton)		
	}	
	time.Sleep(1 * time.Second)

	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return	
	}
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	baton <- newRunner
}

func player(name string, court chan int) {
	defer wg.Done()

	for{
		ball, ok := <-court
		if !ok{
			fmt.Printf("Player %s Won\n", name)
			return
		}

		n := rand.Intn(100)
		if n % 13 == 0 {
			fmt.Printf("Player %s Missed\n",name)
			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		court <- ball
	}
}


func incCouter(id int){

	defer wg.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)

		runtime.Gosched()
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

