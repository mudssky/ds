//Package bst 二叉查找树的实现
package bst

import "fmt"

// 因为interface比较大小的时候需要确定类型，而且需要类型一致，valueType用于存储interface{}的类型方便修改
type valueType int

// BinNode 二叉查找树的节点
type BinNode struct {
	value int
	left  *BinNode
	right *BinNode
	// root用于表示节点属于哪个BST实例，删除的时候用于判断
	// root *BST
}

// BST 二叉查找树的结构体
type BST struct {
	root *BinNode
	len  int
}

// Init 二叉查找树BST的初始化，或者清空BST
func (b *BST) Init() *BST {
	b.root = nil
	b.len = 0
	return b
}

// New 返回一个初始化好的BST
func New() *BST {
	return new(BST).Init()
}

// Empty 通过判断根节点是否为nil，判断二叉树是否为空
func (b *BST) Empty() bool {
	if b.root == nil {
		return true
	}
	return false
}

// 由于不同interface{}的比较没有合适的方法，所以取消interface{} 用固定类型
// 感觉要做的话要先用type-switch判断类型，然后进行类型转换再比较，就算真的实现，性能损失也很大，不如直接用固定类型
// 还有一个问题是go 中type定义的类型，比如  type kkk int 和int不是同一个类型这样的设定。
// 比较两个interface接口的大小，过程中会先进行类型断言，再报错，如果出错，会触发panic
// 大于的情况，返回1，小于的情况返回-1，相等返回0
/*func compareInterface(v1 interface{}, v2 interface{}) int {
	v1t := v1.(type)
	v2t := v2.(type)
	if v1t != v2t {
		panic("v1 and v2 is different type")
	}
	if v1t(v1) > v2t(v2) {
		return -1
	} else if v1t(v1) < v2t(v2) {
		return 1
	} else {
		return 0
	}
}*/

// insert 插入一个节点进入二叉查找树
// 如果二叉树原本为空树，那么根节点用这个节点替换，否则插入到合适位置
// 插入需要满足二叉查找树的性质,每个节点的值比左子女大,比右子女小
// 因为要支持相同大小的元素,所以我这边实现如果一个插入节点的值大于等于节点的值,那么放右边.
// 毕竟相同元素也不会影响搜索，只是如果要搜索到全部的相同元素，需要搜一个删一个
func insert(node *BinNode, at *BinNode) {
	/*
		if compareInterface(node.value, at.value) == -1 {
			if at.left
			insert(node, at.left)
		} else {
			insert(node, at.right)
		}*/
	if at.left == nil && node.value < at.value {
		at.left = &BinNode{value: node.value, left: nil, right: nil}
		return
	}
	if at.right == nil && node.value >= at.value {
		at.right = &BinNode{value: node.value, left: nil, right: nil}
		return
	}

	if node.value < at.value {
		insert(node, at.left)
	} else {
		insert(node, at.right)
	}

}

func (b *BST) insert(node *BinNode) {
	if b.Empty() {
		b.root = node
	} else {
		insert(node, b.root)
	}
	b.len++
}

// 插入一个指定值到BST中，是对b.insert(&BinNode{value: v, left: nil, right: nil})的封装
func (b *BST) insertValue(v int) {
	b.insert(&BinNode{value: v, left: nil, right: nil})
	b.len++
}

// Add 添加一个值到BST
func (b *BST) Add(v int) {
	b.insertValue(v)
}

// 辅助进行递归操作的函数
func search(node *BinNode, v int) *BinNode {

	if node == nil || node.value == v {
		return node
	}
	if s1 := search(node.left, v); s1 != nil {
		return s1
	}
	if s2 := search(node.right, v); s2 != nil {
		return s2
	}
	return nil
}

// Search 返回值等于v的节点
// 如果没有找到，返回空
func (b *BST) Search(v int) *BinNode {
	if b.Empty() {
		return nil
	}
	return search(b.root, v)
}

// 用于递归调用函数，打印当前节点的值
// 字符串order指定遍历使用的顺序
func traverse(node *BinNode, order string) {
	switch order {
	case "LVR":
		if node != nil {
			traverse(node.left, "LVR")
			fmt.Print(node.value, " ")
			traverse(node.right, "LVR")
		}

	case "VLR":
		if node != nil {
			fmt.Print(node.value, " ")
			traverse(node.left, "VLR")
			traverse(node.right, "VLR")
		}
	case "LRV":
		if node != nil {
			traverse(node.left, "LRV")
			traverse(node.right, "LRV")
			fmt.Print(node.value, " ")
		}
	}
}

// TraverseMid  中序遍历二叉查找树
func (b *BST) TraverseMid() {
	traverse(b.root, "LVR")
	fmt.Println()
}

// TraverseFront 前序遍历二叉查找树
func (b *BST) TraverseFront() {
	traverse(b.root, "VLR")
	fmt.Println()
}

// TraverseBack 后序遍历二叉查找树
func (b *BST) TraverseBack() {
	traverse(b.root, "LRV")
	fmt.Println()
}

// graph 用于递归在指定位置输出二叉树的节点，躺着输出
// 执行流程如下
// 1.输出当前节点的右子树，整颗右子树的深度depth+1
// 2.输出depth个数的\t调整输出位置
// 3.输出当前节点的数据
// 4.输出连线，左右子树都有输出< .只有左子树输出\   只有右子树输出/
// 5.换行，再输出左子树
func graph(node *BinNode, depth int) {
	if node == nil {
		return
	}
	graph(node.right, depth+1)
	for i := 0; i < depth; i++ {
		fmt.Print("\t")
	}
	fmt.Print(node.value)
	if node.left != nil && node.right != nil {
		fmt.Print("<")
	} else if node.left != nil && node.right == nil {
		fmt.Print("\\")
	} else if node.left == nil && node.right != nil {
		fmt.Print("/")
	}
	fmt.Println()
	graph(node.left, depth+1)
}

// Graph BST图形输出
func (b *BST) Graph() {
	graph(b.root, 0)
}

func existNode(root *BinNode, node *BinNode) *BinNode {
	if root == nil || root == node {
		return root
	}
	if s1 := existNode(root.left, node); s1 != nil {
		return s1
	}
	if s2 := existNode(root.right, node); s2 != nil {
		return s2
	}
	return nil
}

// ExistNode 判断一个节点是否属于BST，如果存在返回true，不存在返回false
// 不能输入nil进行判断，没有相应的处理逻辑
func (b *BST) ExistNode(node *BinNode) bool {
	if existNode(b.root, node) != nil {
		return true
	}
	return false
}

// 不考虑根节点这种特殊情况，找到一般节点的父节点，如果找不到，返回nil
// 如果已经确定节点node属于root下的节点，那么返回空证明root=node
func findParent(root *BinNode, node *BinNode) *BinNode {
	if root == nil || root.left == node || root.right == node {
		return root
	}
	if s1 := findParent(root.left, node); s1 != nil {
		return s1
	}
	if s2 := findParent(root.right, node); s2 != nil {
		return s2
	}
	return nil
}

// 递归的删除操作，隐藏内部细节，用户只需调用大写的Remove方法
func remove(root *BinNode, node *BinNode) {
	// 情况1，2,左右节点至少有一个为空
	if node.left == nil || node.right == nil {
		// 情况一，node为叶节点，找到父节点，将父节点指向node的指针置为nil
		if node.left == nil && node.right == nil {
			parent := findParent(root, node)
			if parent.right == node {
				parent.right = nil
			} else {
				parent.left = nil
			}
			// 情况2，node左子节点不为空,父节点执行node的指针，指向node的左子节点
		} else if node.left != nil {
			parent := findParent(root, node)
			if parent.right == node {
				parent.right = node.left
			} else {
				parent.left = node.left
			}
			//  情况2，node右子节点不为空，父节点指向node的指针，指向node的右子节点
		} else if node.right != nil {
			parent := findParent(root, node)
			if parent.right == node {
				parent.right = node.right
			} else {
				parent.left = node.right
			}
		}
	} else {
		// 情况3 node左右子节点均不为空，找到node最左子节点的值，代替node的值，在对最左子节点，进行删除操作
		tmp := node.left
		for tmp.left != nil {
			tmp = tmp.left
		}
		node.value = tmp.value
		remove(node, tmp)
	}
}

// Remove 删除二叉树节点node
// 先把node是根节点的情况单独拿出来处理，因为根节点的异常之处在于，他没有父节点
// 分3中情况进行讨论
// 1.node是一个叶子节点
// 2.node有一个子女
// 3.node有两个子女
// 如果是1，只需把双亲节点的对应指针置为nil
// 如果是2，无论node是左节点不为空还是右节点不为空，只要把不为空的子女节点和node的父节点连接上即可
// 综上 1、2其实可以合为一种情况，即node的左右节点不全为空
// 如果是3，那么我们找到右侧节点，因为右侧都比node值要大，我们可以找出右边最小的最接近node，即不断往left走，直到下一个是nil
// 把node设为这个值，然后对这个节点执行删除操作。递归调用方法
func (b *BST) Remove(node *BinNode) {
	if b.root == node {
		// 只有一个根节点的情况,也就是没有子女节点时
		if b.len <= 1 {
			b.Init()
		} else {
			// 左右子节点均不为空
			if b.root.left != nil && b.root.right != nil {
				tmp := node.left
				for tmp.left != nil {
					tmp = tmp.left
				}
				node.value = tmp.value
				fmt.Println(tmp)
				remove(b.root, tmp)
				// 左节点不为空，右节点为空,根节点左移
			} else if b.root.left != nil {
				b.root = b.root.left
				// 右节点不为空，左节点为空，根节点右移
			} else if b.root.right != nil {
				b.root = b.root.right
			}

		}
		return
	}
	if b.ExistNode(node) {
		remove(b.root, node)
	} else {
		panic("the node given is not belong the BST")
	}
}
