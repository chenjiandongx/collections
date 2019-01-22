package collections

type avlNode struct {
	h     int
	value int
	left  *avlNode
	right *avlNode
}

type AVLTree struct {
	tree *avlNode
}

// 生成 AVL 树
func NewAVLTree() *AVLTree {
	return &AVLTree{&avlNode{h: -2}}
}

// 插入节点
func (a *AVLTree) Insert(v int) {
	a.tree = insert(v, a.tree)
}

// 搜索节点
func (a *AVLTree) Search(v int) bool {
	return a.tree.search(v)
}

// 删除节点
func (a *AVLTree) Delete(v int) bool {
	if a.tree.search(v) {
		a.tree.delete(v)
		return true
	}
	return false
}

// 获取所有节点中的最大值
func (a *AVLTree) GetMaxValue() int {
	return a.tree.maxNode().value
}

// 获取所有节点中的最小值
func (a *AVLTree) GetMinValue() int {
	return a.tree.minNode().value
}

// 返回排序后所有值
func (a *AVLTree) AllValues() []int {
	return a.tree.values()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(v int, t *avlNode) *avlNode {
	if t == nil {
		return &avlNode{value: v}
	}
	if t.h == -2 {
		t.value = v
		t.h = 0
		return t
	}

	cmp := v - t.value
	if cmp > 0 {
		// 将节点插入到右子树中
		t.right = insert(v, t.right)
	} else if cmp < 0 {
		// 将节点插入到左子树中
		t.left = insert(v, t.left)
	}
	// 维持树平衡
	t = t.keepBalance(v)
	t.h = max(t.left.height(), t.right.height()) + 1
	return t
}

func (t *avlNode) search(v int) bool {
	if t == nil {
		return false
	}
	cmp := v - t.value
	if cmp > 0 {
		// 如果 v 大于当前节点值，继续从右子树中寻找
		return t.right.search(v)
	} else if cmp < 0 {
		// 如果 v 小于当前节点值，继续从左子树中寻找
		return t.left.search(v)
	} else {
		// 相等则表示找到
		return true
	}
}

func (t *avlNode) delete(v int) *avlNode {
	if t == nil {
		return t
	}
	cmp := v - t.value
	if cmp > 0 {
		// 如果 v 大于当前节点值，继续从右子树中删除
		t.right = t.right.delete(v)
	} else if cmp < 0 {
		// 如果 v 小于当前节点值，继续从左子树中删除
		t.left = t.left.delete(v)
	} else {
		// 找到 v
		if t.left != nil && t.right != nil {
			// 如果该节点既有左子树又有右子树
			// 使用右子树中的最小节点取代删除节点，然后删除右子树中的最小节点
			t.value = t.right.minNode().value
			t.right = t.right.delete(t.value)
		} else if t.left != nil {
			// 如果只有左子树，则直接删除节点
			t = t.left
		} else {
			// 只有右子树或空树
			t = t.right
		}
	}

	if t != nil {
		t.h = max(t.left.height(), t.right.height()) + 1
		t = t.keepBalance(v)
	}
	return t
}

func (t *avlNode) minNode() *avlNode {
	if t == nil {
		return nil
	}
	// 整棵树的最左边节点就是值最小的节点
	if t.left == nil {
		return t
	} else {
		return t.left.minNode()
	}
}

func (t *avlNode) maxNode() *avlNode {
	if t == nil {
		return nil
	}
	// 整棵树的最右边节点就是值最大的节点
	if t.right == nil {
		return t
	} else {
		return t.right.maxNode()
	}
}

/*
左左情况：右旋
		*
	   *
	  *
*/
func (t *avlNode) llRotate() *avlNode {
	node := t.left
	t.left = node.right
	node.right = t

	node.h = max(node.left.height(), node.right.height()) + 1
	t.h = max(t.left.height(), t.right.height()) + 1
	return node
}

/*
右右情况：左旋
		*
	     *
	      *
*/
func (t *avlNode) rrRotate() *avlNode {
	node := t.right
	t.right = node.left
	node.left = t

	node.h = max(node.left.height(), node.right.height()) + 1
	t.h = max(t.left.height(), t.right.height()) + 1
	return node
}

/*
左右情况：先左旋 后右旋
		*
	   *
	    *
*/
func (t *avlNode) lrRotate() *avlNode {
	t.left = t.left.rrRotate()
	return t.llRotate()
}

/*
右左情况：先右旋 后左旋
		*
	     *
        *
*/
func (t *avlNode) rlRotate() *avlNode {
	t.right = t.right.llRotate()
	return t.rrRotate()
}

func (t *avlNode) keepBalance(v int) *avlNode {
	// 左子树失衡
	if t.left.height()-t.right.height() == 2 {
		if v-t.left.value < 0 {
			// 当插入的节点在失衡节点的左子树的左子树中，直接右旋
			t = t.llRotate()
		} else {
			// 当插入的节点在失衡节点的左子树的右子树中，先左旋后右旋
			t = t.lrRotate()
		}
	} else if t.right.height()-t.left.height() == 2 {
		if t.right.right.height() > t.right.left.height() {
			// 当插入的节点在失衡节点的右子树的右子树中，直接左旋
			t = t.rrRotate()
		} else {
			// 当插入的节点在失衡节点的右子树的左子树中，先右旋后左旋
			t = t.rlRotate()
		}
	}
	// 调整树高度
	t.h = max(t.left.height(), t.right.height()) + 1
	return t
}

func (t *avlNode) height() int {
	if t != nil {
		return t.h
	}
	return -1
}

// 中序遍历按顺序获取所有值
func appendValue(values []int, t *avlNode) []int {
	if t != nil {
		values = appendValue(values, t.left)
		values = append(values, t.value)
		values = appendValue(values, t.right)
	}
	return values
}

func (t *avlNode) values() []int {
	values := make([]int, 0)
	return appendValue(values, t)
}
