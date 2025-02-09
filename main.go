package main

import (
	"fmt"
	"math/rand"
	"time"
	"runtime"
	"sync"
	"context"
)

func main() {
	// waitForResult()
	// fanOut()
	// waitForTask()
	// pooling()
	// fanOutSemaphore()
	// fanOutBounded()
	// drop()
	cancellation()
}

func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("employee: sent signal")
	}()

	p := <-ch
	fmt.Println("manager: recv'd signal: ", p)

	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------------------------------------")
}

func fanOut() {
	emps := 2000
	ch := make(chan string, emps) // buffered channel, which can hold 2000 signals without blocking the sender

	for e := 0; e < emps; e++ {
		go func(emp int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "paper"
			fmt.Println("employee: sent signal: ", emp)
		}(e)
	}

	for emps > 0 {
		p := <-ch
		emps--
		fmt.Println(p)
		fmt.Println("manager: recv'd signal: ", emps)
	}

	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------------------------------------")
}

// the goroutine that is waiting for the signal from manager
func waitForTask() {
	ch := make(chan string)

	go func() {
		p := <-ch
		fmt.Println("manager: recv'd signal: ", p)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "paper"
	fmt.Println("employee: sent signal")

	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------------------------------------")
}

func pooling() {
	ch := make(chan string)

	g := runtime.NumCPU()
	for e := 0; e < g; e++ {
		go func(emp int) {
			// Pooling until the channel is closed
			for p := range ch {
				fmt.Println("employee: ", emp, " recv'd signal: ", p)
			}
			fmt.Println("employee: ", emp, " recv'd signal: shutdown signal")
		}(e)
	}

	// Pushes 100 signals to the channel
	const work = 100
	for w := 0; w < work; w++ {
		ch <- "paper " + string(w)
		fmt.Println("manager: sent signal: ", w)
	}

	close(ch)
	fmt.Println("manager: sent showdown signal")
	fmt.Println("---------------------------------------------------------------------------------------")
}

func fanOutSemaphore() {
	emps := 2000
	ch := make(chan string, emps)
	
	g := runtime.NumCPU()
	sem := make(chan bool, g)

	for e := 0; e < emps; e++ {
		go func(emp int) {
			sem <- true
			{
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
				ch <- "paper"
				fmt.Println("employee: sent signal: ", emp)
			}
			<-sem
		}(e)
	}

	for emps > 0 {
		p := <-ch
		emps--
		fmt.Println(p)
		fmt.Println("manager: recv'd signal: ", emps)
	}

	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------------------------------------")
}

// Similar to Java's ThreadPoolExecutor
func fanOutBounded() {
	work := []string{"paper", "paper", "paper", "paper", "paper", 2000: "paper"}

	g := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g)

	for e := 0; e < g; e++ {
		go func(emp int) {
			defer wg.Done()
			for p := range ch {
				fmt.Println("employee", emp, ": recv'd signal: ", p)
			}
			fmt.Println("employee", emp, ": recv'd shutdown signal")
		}(e)
	}

	for _, w := range work {
		ch <- w
	}
	close(ch)
	wg.Wait()

	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------------------------------------")
}

func drop() {
	const cap = 100
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
			fmt.Println("employee: recv'd signal: ", p)
		}
	}()

	const work = 2000
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper":
			fmt.Println("manager: sent signal: ", w)
		default:
			fmt.Println("manager: dropped data: ", w)
		}
	}

	close(ch)
	fmt.Println("manager, sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------------------------------------")
}

func cancellation() {
	duration := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		// greater than 300ms will cause the work to be cancelled
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		ch <- "paper"
	}()

	select {
	case d := <-ch:
		fmt.Println("work complete", d)

	case <-ctx.Done(): // starts the clock on the context with timeout
		fmt.Println("work cancelled")
	}

	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------------------------------------")
}