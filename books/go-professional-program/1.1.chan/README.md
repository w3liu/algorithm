# chan数据结构

```go
type hchan struct { 
    qcount uint // 当前队列中剩余元素个数 
    dataqsiz uint // 环形队列长度，即可以存放的元素个数 
    buf unsafe.Pointer // 环形队列指针 
    elemsize uint16 // 每个元素的大小 
    closed uint32 // 标识关闭状态 
    elemtype *_type // 元素类型 
    sendx uint // 队列下标，指示元素写入时存放到队列中的位置 
    recvx uint // 队列下标，指示元素从队列的该位置读出
    recvq waitq // 等待读消息的goroutine队列 
    sendq waitq // 等待写消息的goroutine队列 
    lock mutex // 互斥锁，chan不允许并发读写
 }
```

# 锁
一个channel同时仅允许被一个goroutine读写

# channel 读写
## 写数据
    1. 如果等待接收队列recvq不为空，说明缓冲区中没有数据或者没有缓冲区，此时直接从recvq取出G,并把数据 写入，最后把该G唤醒，结束发送过程； 
    2. 如果缓冲区中有空余位置，将数据写入缓冲区，结束发送过程； 
    3. 如果缓冲区中没有空余位置，将待发送数据写入G，将当前G加入sendq，进入睡眠，等待被读goroutine唤 醒；
## 读数据
    1. 如果等待发送队列sendq不为空，且没有缓冲区，直接从sendq中取出G，把G中数据读出，最后把G唤醒，结束 读取过程； 
    2. 如果等待发送队列sendq不为空，此时说明缓冲区已满，从缓冲区中首部读出数据，把G中数据写入缓冲区尾 部，把G唤醒，结束读取过程； 
    3. 如果缓冲区中有数据，则从缓冲区取出数据，结束读取过程； 
    4. 将当前goroutine加入recvq，进入睡眠，等待被写goroutine唤醒；
## 关闭
    1. 关闭channel会把recvq中的G全部唤醒，本应该写入G的数据位置为nil。
    2. 把sendq中的G全部唤醒，但这些G会panic。
    3. 其他会panic得场景：
        * 关闭值为nil的channel
        * 关闭已经被关闭的channel
        * 向已经关闭的channel写数据
# 常见用法
## 单项channel

