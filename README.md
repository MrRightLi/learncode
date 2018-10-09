# learncode


# defer 调用
- 在函数return时调用
- defer 相当于一个栈,先进后出


## goroutine 切换的点

- I/O, selelct
- channel
- 等待锁
- 函数调用(有时)
- runtime.Gosched()

上述只是参考,不能保证切换,不能保证在其他地方不切换


# channel
- no-buffered channel
- buffered channel
- range
- 理论基础:Communication Sequential Process(CSP)
- 不要通过共享内存来通信;通过通信来共享内存