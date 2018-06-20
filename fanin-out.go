package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

type Generator struct {
	generator_id int
	channel      chan int
	iteractions  int
}

func main() {
	//tracing start
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	//tracing stop
	startTime := time.Now()
	gen := make(chan int, 10)
	done := make(chan bool)
	counter := make(chan int)
	//var wg sync.WaitGroup
	iteractions := 50
	worker_generators := 80
	readers := 1
	gen_count := make(chan int, worker_generators*iteractions)
	go pools(worker_generators, iteractions, gen, gen_count)
	go reader_pool(readers, done, gen, counter)
	<-done
	fmt.Printf("We got %d from %d iteractions", <-counter, worker_generators*iteractions)
	time.Sleep(1 * time.Millisecond)
	fmt.Printf("Lenght of CHANNEL %d ", len(gen))
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("########\ntotal time taken \n########", diff.Seconds(), "seconds")
}

func reader_pool(readers int, done chan bool, gen chan int, counter chan int) {
	for i := 1; i <= readers; i++ {
		go reader(done, gen, counter)
	}
}

func pools(worker_generators int, iteractions int, gen chan int, gen_count chan int) {
	var wg sync.WaitGroup
	//var m sync.Mutex
	for i := 1; i <= worker_generators; i++ {
		wg.Add(1)
		generator := Generator{generator_id: i, channel: gen, iteractions: iteractions}
		go generator.generator(&wg, gen_count)
	}
	wg.Wait()
	fmt.Printf("########\nGenerators finished: %d\n########", len(gen_count))
	close(gen)
	close(gen_count)
}
func (gen Generator) generator(wg *sync.WaitGroup, gen_count chan int) {
	defer wg.Done()
	for z := 1; z <= gen.iteractions; z++ {
		gen.channel <- z
		gen_count <- z
		fmt.Printf("Generator %d has wrote %d. Counter by now %d\n", gen.generator_id, z, len(gen_count))
	}
}

func reader(done chan bool, gen chan int, counter chan int) {
	a := 0
	for items := range gen {
		//fmt.Println("Reader: Lenght of channel from generators", len(gen))
		a++
		fmt.Println("read value: ", items)
		fmt.Println("Have read by now", a)
	}
	time.Sleep(100 * time.Millisecond)
	done <- true
	counter <- a
}
