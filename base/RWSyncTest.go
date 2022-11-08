package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	lock := &sync.RWMutex{}
	lock.Lock()
	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func(i int,wg *sync.WaitGroup) {
			wg.Done()
			fmt.Printf("第 %d 个协程准备开始... \n", i)
			lock.RLock()
			fmt.Printf("第 %d 个协程获得读锁, sleep 3s 后，释放锁\n", i)
			time.Sleep(3*time.Second)
			lock.RUnlock()

		}(i,&wg)
	}

	wg.Wait()

	fmt.Println("准备释放写锁，读锁不再阻塞")
	// 写锁一释放，读锁就自由了
	lock.Unlock()

	// 由于会等到读锁全部释放，才能获得写锁
	// 因为这里一定会在上面 4 个协程全部完成才能往下走
	lock.Lock()
	fmt.Println("程序退出...")
	lock.Unlock()
}
