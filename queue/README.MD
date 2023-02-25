# QUEUE 佇列

資料採先進先出(後進後出)進行存取

與堆疊(stack)一樣都可以使用陣列或鏈結列表實現，這邊一樣使用陣列來進行實作

## Enqueue

```go
func (p *People) Enqueue(data Data) {
	p.list = append(p.list, data)
}
```

## Dequeue

```go
func (p *People) Dequeue() Data {
	temp := p.list[0]
	p.list = p.list[1:] // 移除最開頭(index=0)的值

	return temp
}
```

> 實作上與`堆疊(stack)`大同小異，不同的地方只有在取得資料時貯列是從最開頭地方取得，並從列表移除掉該元素
