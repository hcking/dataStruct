package adt

// SqList sequence list 顺序表(线性表)结构
type SqList struct {
	length int
	data   []interface{}
}

func NewSqList() (l *SqList) {
	l = &SqList{}
	l.length = 0
	return l
}

func (l *SqList) Clear() {
	l.length = 0
	return
}

func (l *SqList) Destroy() {
	l.data = nil
	l.length = 0
	return
}

func (l *SqList) Len() int {
	return l.length
}

func (l *SqList) IsEmpty() bool {
	if l.length == 0 {
		return true
	}
	return false
}

func (l *SqList) Get(i int) (res interface{}) {
	if i < 1 || i > l.length {
		return nil
	}
	res = l.data[i-1]
	return res
}

func (l *SqList) Find(item interface{}) (index int) {
	for i := 0; i < l.Len(); i++ {
		if l.data[i] == item {
			return i
		}
	}
	return -1
}

func (l *SqList) Set(index int, item interface{}) bool {
	if index < 0 || index > l.Len()+1 {
		return false
	}
	if index > l.Len() {
		l.data = append(l.data, item)
	} else {
		newData := append(l.data[:index], item)
		l.data = append(newData, l.data[index:])
	}
	l.length += 1
	return true
}
