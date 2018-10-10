package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(ii int) {
			for  {
				a[ii]++
				runtime.Gosched() // 交出 goroutine 控制权
				// 如果不交出控制权,其他 goroutine 没有执行的机会
			}
		}(i) // go run -race xxx.go
	}

	time.Sleep(time.Millisecond)
	fmt.Println(a) // 在这里依然存在 race 思考分析一下吧
}
