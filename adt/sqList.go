package adt

// SqList 顺序表类型(线性表)
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
}

func (l *SqList) Destroy() {
	l.data = nil
	l.length = 0
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
