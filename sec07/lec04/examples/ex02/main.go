// Section 07 - Lecture 04 : Channels and Goroutines
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

func main() {
	// non-concurrent data producer and consumer
	// ------
	// d := make(chan string, 5)
	// producer(1, d)
	// consumer(d)

	// concurrent data producer and consumer
	// ------
	// d := make(chan string, 5)
	// go producer(1, d)
	// consumer(d)

	// *0 capacity channel* with concurrent data producer and consumer
	// ------
	d := make(chan string)
	go producer(1, d)
	consumer(d)
}
func consumer(in chan string) {
	count := 0
	for v := range in {
		count++
		fmt.Printf("Consumer got: %v\n", v)
	}

	if count == 0 {
		fmt.Println("No data received")
		return
	}

	fmt.Printf("Processed %v items\n", count)
}
func producer(id int, out chan string) {
	n := r.Int() % cap(out)
	for i := 0; i < n; i++ {
		out <- fmt.Sprintf("Producer: %v, item: %v", id, i+1)
	}
	close(out)
}
