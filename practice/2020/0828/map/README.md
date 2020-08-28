# 1.map数据结构

```go
type hmap struct {
    count int // 当前保存的元素个数
    B uint8 // 指示bucket数组大小
    buckets unsafe.Pointer // bucket数组指针，数组的大小为2^B
}
```

# 2.bucket数据结构

```go
type bmap struct {
    tophash [8]uint8 // 存储哈希值的高8位
    data byte[1] // key value数据：key/key/key.../value/value/value...
    overflow *bmap // 溢出bucket的地址
}
```

每个bucket可以存着8个键值对

# 3.哈希冲突

Go使用链地址法来解决键冲突，由于每bucket可以存8个键值对，超过8个键值对就会再创建一个键值对，用类似链表的方式将bucket连接起来

# 4.负载因子

负载因子 = 键数量/bucket数量

Go的负载因子达到6.5的时候才会触发rehash

# 5.渐进式扩容

    * 扩容的前提条件
        - 负载因子 > 6.5时
        - overflow > 2^15时