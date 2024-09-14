package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producer(id int, numbersChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		num := rand.Intn(100)
		fmt.Printf("Producer %d produced: %d\n", id, num)
		numbersChan <- num
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}

func consumer(id int, numbersChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range numbersChan {
		fmt.Printf("Consumer %d consumed: %d\n", id, num)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500))) // Simulate processing time
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	numbersChan := make(chan int, 10) // Buffered channel for communication
	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	// Launch producer goroutines
	for i := 1; i <= 3; i++ {
		producerWg.Add(1)
		go producer(i, numbersChan, &producerWg)
	}

	// Launch consumer goroutines
	for i := 1; i <= 2; i++ {
		consumerWg.Add(1)
		go consumer(i, numbersChan, &consumerWg)
	}

	// Wait for all producers to finish
	producerWg.Wait()

	// Close the channel after all producers are done
	close(numbersChan)

	// Wait for consumers to finish processing
	consumerWg.Wait()

	fmt.Println("All producers and consumers have finished.")
}
