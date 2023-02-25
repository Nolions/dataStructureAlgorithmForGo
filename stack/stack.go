package stack

type Data interface {
}

type Stack interface {
	Push(Data)
	Pop() Data
	IsEmpty() bool
}

type Book struct {
	list []Data
}

func New(size int) *Book {
	d := new(Book)
	d.list = make([]Data, 0, size)

	return d
}

func (s *Book) Push(data Data) {
	s.list = append(s.list, data)
}

func (s *Book) Pop() Data {
	tmp := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]

	return tmp
}

func (s *Book) IsEmpty() bool {
	return len(s.list) == 0
}
