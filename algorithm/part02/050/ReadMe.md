## Algorithm
### 1. 题目
```
173. 二叉搜索树迭代器
```
### 2. 题目描述
```
实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。

调用 next() 将返回二叉搜索树中的下一个最小的数。

 

示例：



BSTIterator iterator = new BSTIterator(root);
iterator.next();    // 返回 3
iterator.next();    // 返回 7
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 9
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 15
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 20
iterator.hasNext(); // 返回 false
 

提示：

next() 和 hasNext() 操作的时间复杂度是 O(1)，并使用 O(h) 内存，其中 h 是树的高度。
你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 中至少存在一个下一个最小的数。
```

### 3. 解答：
```golang
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	queue *list.List
}

func Constructor(root *TreeNode) BSTIterator {
	queue := list.New()
	midOrder(root, queue)
	return BSTIterator{
		queue: queue,
	}
}

func midOrder(node *TreeNode, queue *list.List) {
	for node != nil {
		queue.PushFront(node)
		node = node.Left
	}
}

func (bst *BSTIterator) Next() int {
	elem := bst.queue.Front()
	bst.queue.Remove(elem)
	node := elem.Value.(*TreeNode)
	if node.Right != nil {
		midOrder(node.Right, bst.queue)
	}
	return node.Val
}

func (bst *BSTIterator) HasNext() bool {
	return bst.queue.Len() > 0
}
```
### 4. 说明
1. 题目要求空间复杂度O(h)，则可以想到需要用栈来实现一个递归算法，因此我们需要显式管理一个栈，并利用这个栈实现中序遍历。在这个栈里，栈顶元素始终是节点值最小的元素。
2. 初始时，需要将当前节点的左节点依次入栈，完成中序遍历“左根右”中的“左”。
3. 此时，对于hasNext()方法，只需判断栈是否为空。
4. 对于next()方法，栈顶元素即为下一个最小的元素，将栈顶元素出栈并返回其val即可。但是，会有以下两种情况：栈顶元素的节点没有右子树；栈顶元素的节点有右子树。如果栈顶元素的节点没有右子树，则不用进行额外操作；如果栈顶元素的节点有右子树，则需要将对其右子节点再进行中序遍历，即将其右子节点的左子节点依次入栈，进而完成对树整体的中序遍历。
5. 将当前节点的左子节点依次入栈，是通过一个自定义的方法midOrder()实现。