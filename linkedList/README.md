# Linked List 鏈結

## Append

***在鏈結最後新增一筆節點***

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

## Prepend

***在鏈結最前面新增一筆節點***

step1. 建立一個`new`節點

step2. 新增一個`temp`節點來暫存最`head`(最開頭)的節點

step3. 將`new`節點的鏈結指向`temp`節點

step4. 將`head`節點指向`new`節點

```go
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
```
