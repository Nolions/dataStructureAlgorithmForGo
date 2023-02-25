package queue

type Data interface {
}

type Queue interface {
	Enqueue(data Data)
	Dequeue() Data
	IsEmpty() bool
}

type People struct {
	list []Data
}

func New(size int) *People {
	d := new(People)
	d.list = make([]Data, 0, size)

	return d
}

func (p *People) Enqueue(data Data) {
	p.list = append(p.list, data)
}

func (p *People) Dequeue() Data {
	temp := p.list[0]
	p.list = p.list[1:] // 移除最開頭(index=0)的值

	return temp
}

func (p *People) IsEmpty() bool {
	return len(p.list) == 0
}
