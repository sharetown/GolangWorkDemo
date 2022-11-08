package main

import (
	"fmt"
	"sync"
	"time"
)

func add(count *int, wg *sync.WaitGroup) {

    for i := 0; i < 1000; i++ {
        *count = *count + 1
        time.Sleep(1)
    }
            wg.Done()

}

func main() {
    var wg sync.WaitGroup
    count := 0
    wg.Add(3)
    go add(&count, &wg)
    go add(&count, &wg)
    go add(&count, &wg)

    wg.Wait()
    fmt.Println("count 的值为：", count)
}
