package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex
var wg sync.WaitGroup

//下面这段代码是有错误的，要求更改部分代码，实现0~9的打印（最好在错误的地方，注释并标明错误原因）

func main() {
	over := make(chan bool, 1) //创建bool类型管道，没给cap

	for i := 0; i < 10; i++ {
		wg.Add(1) //让for循坏等待

		go func() {
			lock.Lock() //加锁防止资源竞争
			fmt.Println(i)
			lock.Unlock() //打印完解锁
			wg.Done()
		}()
		if i == 9 {
			over <- true
		}
		wg.Wait() //结束
	}
	for {
		_, ok := <-over
		if !ok {
			break
		}
	} //<-over无效，主程序直接跑完
	fmt.Println("over!!!")
}
