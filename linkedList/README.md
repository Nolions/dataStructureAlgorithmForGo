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

## Reverse

***反轉整個鏈結列表***

step1. 新增`mid`、`last`兩個節點，分別作為**中間**與**結尾**節點用

step2. 將`last`節點指向`mid`節點

step3. 將最前面節點傳遞給`mid`節點

step4. 最前面節點往前一個節點

step5. 反轉原本最前面節點與中間節點順序關係

step6. 透過迴圈方式不段執行`step2`~`step5`步驟，直到最前面節點為空

```go
func Reverse() *Node {
    var mid, last *Node

    for Head != nil {
        last = mid       // 中間節點傳遞給結尾節點
        mid = Head       // 最前面節點傳遞給中間節點
        Head = Head.next // 最前面節點往前一個節點
        mid.next = last  // 反轉原本最前面節點與中間節點順序關係
    }

    return mid
}
```

## Length

***計算鏈結長度***

使用迴圈，並透過判斷節點的鏈結指向事不是空方式來計算鏈結列表長度

```go
func Size(node *Node) (size int) {
	for node != nil {
		size++
		node = node.next
	}

	return size
}
```

## Remove by value

***移除鏈結中特定值的節點***

step1. 新建一個`find`節點，以便紀錄之後找到節點

step2. 透過迴圈方式尋找符合目標的節點

step3. 檢查`step2`是否有找到符合目標的值，如果沒有直接結束

step4. 新建一個新節點`temp`，並指向`find`節點的鏈結

step5. 將`find`節點的鏈結指向`temp`節點的鏈結

> 在`step 2`中需要額外去檢查符合目標的值是不是開頭的節點

```go
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

    temp := findNode.next
	findNode.next = temp.next
}
```

## Merge

鏈結列表的合併

鏈結列表的合併與`在鏈結最後新增一筆節點`方法相似

step1. 找到鏈結列表中最後一個節點

step2. 將最後一個節點的鏈結欄位指向剛剛建立的新節點

```go
func Merge(node *Node) {
    if Head == nil {
        Head = node
        return
    }
	
	if node == nil {
	    return 	
    }
	
    final := Head
    
	for final.next != nil {
        final = final.next
    }
    final.next = node
}
```
