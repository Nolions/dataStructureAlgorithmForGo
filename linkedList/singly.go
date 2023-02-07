package linkedList

import "fmt"

type Node struct {
	value int
	next  *Node
}

var Head *Node = nil

// NewNode 建立鏈結
func NewNode(value int, next *Node) *Node {
	var n Node
	n.value = value
	n.next = next
	return &n
}

// Append 在鏈結最後新增一筆節點
func Append(data int) {
	node := NewNode(data, nil)
	if Head == nil {
		Head = node
		return
	}
	temp := Head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = node

}

// Prepend 在鏈結最前面新增一筆節點
func Prepend(data int) {
	node := NewNode(data, nil)
	if Head == nil {
		Head = node
		return
	}
	temp := Head
	Head = node
	Head.next = temp
}

// Reverse 反轉整個鏈結節點
func Reverse() *Node {
	var mid, last *Node

	for Head != nil {
		last = mid       // 中間節點傳遞給結尾節點
		mid = Head       // 最前面節點傳遞給中間節點
		Head = Head.next // 最前面節點往前一個節點
		mid.next = last  // 反轉原本最前面節點與中間節點順序關西
	}

	return mid
}

// RemoveVal 移除某個節點
func RemoveVal(data int) {
	var findNode *Node = nil

	for p := Head; p.next != nil; p = p.next {
		if p.value == data {
			Head = Head.next
			return
		}

		if p.next.value == data {
			findNode = p
			break
		}
	}

	if findNode == nil {
		return
	}

	q := findNode.next
	findNode.next = q.next
}

// Print 列印出整的鏈結節點
func print(node *Node) {
	temp := node
	for temp != nil {
		fmt.Printf("%d ", temp.value)
		temp = temp.next
	}
	fmt.Println(" ")
}

// TODO 鏈結由小到大排序

// TODO 合併兩個鏈結

// TODO 新增一個節點至從小到大排訊鏈結

// Size 計算鏈結長度
func Size(node *Node) (size int) {
	for node != nil {
		size++
		node = node.next
	}

	return size
}
