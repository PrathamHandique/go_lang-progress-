package main

import (
	"fmt"
	"sync"
)

func worker(i int,wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("worker %d started\n",i)
	fmt.Printf("worker %d ended\n",i)

}
func main(){
	fmt.Println("main started")
	var wg sync.WaitGroup
	//start 3 workers go routines
	for i:=0;i<3;i++{
		wg.Add(1)
		go worker(i,&wg)
	}
	//wait for all workers to complete
	wg.Wait()
	fmt.Println("worker task completed")
}