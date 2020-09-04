# channel
* 适用于主协程等待所有子协程退出后再继续后续流程
* 优点：实现简单，直观
* 缺点：需要大量创建协程时，就需要相同数量的channel，对子协程派生出来的协程不方便控制

# waitGroup
* Add()操作必须遭遇Wait()，否则会panic
* Add()设置的值必须与实际等待的goroutine个数一直，否则会panic

# context
* 可以控制多级的goroutine
* Context仅仅是一个接口定义，根据实现的不同，可以衍生出不同的context类型
* cancelCtx实现了Context接口，通过WithCancel()创建cancelCtx实例
* timerCtx实现了Context接口，通过WithDeadline()和WithTimeout()创建timerCtx实例
* valueCtx实现了Context接口，通过WithValue()创建valueCtx实例
* 三种context实例可互为父节点，从而可以组合成不同的应用形式
