# STACK 堆疊

資料採進先後出(後進先出)進行存取

並且可以透過陣列(Array)與鏈結列表(Linked List)的方式來實作，這邊使用陣列(Array)來實作

## Pop

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
