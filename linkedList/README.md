# Linked List 鏈結

## Append

**在鏈結最後新增一筆節點**

step1. 建立一個節點

step2. 找到鏈結列表中最後一個節點

step3. 將最後一個節點的鏈結欄位指向剛剛建立的新節點

```go
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
```
