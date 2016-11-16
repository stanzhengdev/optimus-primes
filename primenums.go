package main

import (
	"fmt"
)

func generatePrime(count int) (arr []int) {
	/***
	  Generates a number up to a counte and returns a map
	  **/
	arr = make([]int, count)
	ch := make(chan int)
	go generate(ch)
	for i := 0; i < count; i++ {
		prime := <-ch
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
		fmt.Println(prime)
		arr[i] = prime
	}
	return arr
}

func main() {
	fmt.Println(generatePrime(10))
}

func generate(ch chan int) {
	/** Generates numbers sequentially starting from 2
	 **/
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in, out chan int, prime int) {
	/**
	  Determines if the number is prime by checking if is divisible by itself
	  **/
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
