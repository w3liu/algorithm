# 线程池的缺陷
* 当发生系统调用的时候现成会被阻塞
* 如果大部分任务都会进行系统调用会让阻塞情况变多
* 线程池数量不好确定，有时候线程池数量过多，由于过多得争抢CPU资源，消费能力反而会下降

# Goroutine调度器
* Goroutine的主要概念如下：
    - G(Goroutine):即Go协程，每个go关键字都会创建一个协程
    - M(Machine):工作线程，在Go中称为Machine
    - P(Processor):处理器（Go中定义的一个概念，不是CPU）,包含运行代码的必要资源
* 可以使用runtime.GOMAXPROCS()设置P的个数，在某些IO密集型的场景下可以在一定程度上提高性能

# Goroutine调度策略
* 队列轮转
    - P维护着一个包含G的队列
    - P周期性的将G调度到M中执行
    - 执行一小段时间再将G放到队列尾部
    - 然后从队列中重新取出一个G进行调度
    - 除了每个P维护的G队列以外，还有一个全局队列
    - 每个P会周期性的查看全局队列中是否有G待运行
    - P周期性的查看全局队列，是为了防止全局队列中的G被饿死
* 系统调用
    - 一般情况下M的个数会略大于P的个数
    - 多出来的M会在G产生系统调用的时候发挥作用
    - Go提供了M的池子，需要就从池子里取，用完就放回，不够就再创建一个
    - 系统调用结束后
        1. 如果有空闲的P，获取一个P，继续执行G
        2. 如果没有空闲的P，将G加入全局队列，等待其他的P调度，M进入缓存池睡眠
* 工作量窃取
    - 多个P的G队列可能不均衡
    - 空闲的P会将其他的P中的G窃取一部分过来
    - 一般情况下，每次窃取一半
    

