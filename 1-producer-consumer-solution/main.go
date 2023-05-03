//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(tweetingInput chan<- Tweet, stream Stream) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(tweetingInput)
			return
		}
		tweetingInput <- *tweet // check without *
	}
}

func consumer(tweets <-chan Tweet, wg *sync.WaitGroup) {
	for t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	stream := GetMockStream()
	length := len(stream.tweets)
	wg.Add(length)

	// Producer
	ch := make(chan Tweet, length)

	go producer(ch, stream)

	// Consumer
	go consumer(ch, &wg)

	wg.Wait()

	fmt.Printf("Process took %s\n", time.Since(start))
}
