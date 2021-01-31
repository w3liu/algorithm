## Algorithm
### 1. 题目
```
155. 最小栈
```
### 2. 题目描述
```
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。

示例:

输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
 

提示：

pop、top 和 getMin 操作总是在 非空栈 上调用。
```

### 3. 解答：
```golang

type MinStack struct {
	min int
	arr []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min: 0,
		arr: make([]int, 0),
	}
}


func (s *MinStack) Push(x int)  {
	if len(s.arr) == 0 {
		s.arr = append(s.arr, 0)
		s.min = x
	} else {
		diff := int64(x) - int64(s.min)
		s.arr = append(s.arr, int(diff))
		if s.min > x {
			s.min = x
		}
	}

}


func (s *MinStack) Pop()  {
	if len(s.arr) == 0 {
		return
	}
	diff := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	if diff < 0 {
		s.min = s.min - diff
	}
}


func (s *MinStack) Top() int {
	if len(s.arr) == 0 {
		return 0
	}
	diff := s.arr[len(s.arr)-1]
	if diff >= 0 {
		return s.min + diff
	} else {
		return s.min
	}
}


func (s *MinStack) GetMin() int {
	return s.min
}
```
### 4. 说明
1. 这个方案比较节约内存，不需要额外的空间；
2. arr数组存储的不是变量x而是x与min的差值；
3. pop min top 需要根据s.min及arr数组里的值做运算；