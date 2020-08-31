# tips
    * 遍历过程中可以适情况放弃接收index或value，可以一定程度上提升性能
    * 遍历channel时，如果channel中没有数据，可能会阻塞
    * 尽量避免遍历过程中修改原数据
# 总结
    * for-range的实现实际上是C风格的for循环
    * 使用index,value接收range返回值会发生一次数据拷贝