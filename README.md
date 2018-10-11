# learncode


# 协程 Coroutine
- 轻量级 "线程"
- 非抢占式多任务处理,由协程主动交出控制权
- 编译器/解释器/虚拟机层面的多任务
- 多个协程可能在一个或者多个线程上运行


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

## select 
以下描述了 select 语句的语法：
- case都必须是一个通信
- hannel表达式都会被求值
- 发送的表达式都会被求值
如果任意某个通信可以进行，它就执行；其他被忽略。当然如果某个通信一直有值,前面的值会被重置,所以我们看到:如果发送方比消耗方快会导致数据丢失
- 如果有多个case都可以运行，Select会随机公平地选出一个执行。
  否则:
- - default子句，则执行该语句。
- - 如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。

> 小结:
- SELECT的使用
- 定时器的使用
- 在 Select 中使用 Nil Channel