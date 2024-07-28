package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func readCsv(filename string) chan []string {
	stream := make(chan []string)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Not able to open file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Not able to read records", err)
	}

	go func() {
		defer close(stream)
		for _, eachrecord := range records {
			stream <- eachrecord

		}
	}()

	return stream

}

func countFanOut(stream <-chan []string, thread int) <-chan int {
	result := make(chan int)
	count := 0
	go func() {
		defer close(result)
		for range stream {
			fmt.Printf("Thread %d is processing now \n", thread)
			count += 1
			// time.Sleep(time.Second * 2)
		}
		result <- count
	}()

	return result
}

func aggregateStreamResult(channels ...<-chan int) int {
	count := 0
	for _, ch := range channels {
		for segmentCount := range ch {
			count += segmentCount
		}
	}
	return count
}

func main() {
	startTime := time.Now()
	rowStream := readCsv("../sample_data.csv")
	channels := make([]<-chan int, 5)
	for i := 0; i < 5; i++ {
		channels[i] = countFanOut(rowStream, i)
	}
	result := aggregateStreamResult(channels...)
	fmt.Println("result", result)
	fmt.Println(time.Since(startTime))
}
