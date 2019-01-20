package main

import (
	"fmt"
	"log"
)

func main() {

	for i := 0; i <= 10; i++ {
        index, result := fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
		log.Printf("index: %d, result: %d", index, result)
    }
}

func fibonacci(n int, index_ int) (index int, res int) {

	if n <= 1 {
        res = 1
    } else {
		index_ += 1
		index, res = fibonacci(n-1, index_) + fibonacci(n-2, index_)
    }
    return index, res
}
