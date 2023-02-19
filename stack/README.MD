# STACK 堆疊

資料採後進先出進行存取

並且可以透過陣列(Array)與鏈結列表(Linked List)的方式來實作

## 透過Array實現

### Pop

```go
func (s *Stack) Pop() stack.Data {
	tmp := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]

	return tmp
}
```

## Push

```go
func (s *Stack) Push(data stack.Data) {
	s.list = append(s.list, data)
}
```

## 透過鏈結實現