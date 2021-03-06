## Algorithm
### 1. 题目
```
146. LRU 缓存机制
```
### 2. 题目描述
```
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
实现 LRUCache 类：

LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
 

进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？

 

示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4
 

提示：

1 <= capacity <= 3000
0 <= key <= 3000
0 <= value <= 104
最多调用 3 * 104 次 get 和 put

```

### 3. 解答：
```golang
type Node struct {
	key int
	val int
}

type LRUCache struct {
	capacity int
	cache map[int]*list.Element
	link *list.List
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		link:     list.New(),
	}
}


func (s *LRUCache) Get(key int) int {
	v, ok := s.cache[key]
	if ok {
		s.link.MoveToFront(v)
		return v.Value.(*Node).val
	}
	return -1
}


func (s *LRUCache) Put(key int, value int)  {
	v, ok := s.cache[key]
	if ok {
		v.Value.(*Node).val = value
		s.link.MoveToFront(v)
	} else {
		if len(s.cache) + 1 > s.capacity {
			ele := s.link.Back()
			delete(s.cache, ele.Value.(*Node).key)
			s.link.Remove(ele)
		}
		ele := s.link.PushFront(&Node{
			key: key,
			val: value,
		})
		s.cache[key] = ele
	}
}
```
### 4. 说明
使用链表+map的方式，链表用来保持访问的顺序，map用来存储数据，这个方案的缺点是内存占用有点大