package adt

const GSqListMaxLen = 1000

// SqList sequence list 顺序表(线性表)结构
type SqList struct {
	length int
	data   []interface{}
	maxLen int
}

func NewSqList(maxLen int) *SqList {
	l := &SqList{}
	if maxLen <= 0 {
		l.maxLen = GSqListMaxLen
	} else {
		l.maxLen = maxLen
	}
	l.data = make([]interface{}, l.maxLen)
	l.length = 0
	return l
}

func (l *SqList) Clear() {
	l.length = 0
	return
}

func (l *SqList) Destroy() {
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

func (l *SqList) Get(index int) interface{} {
	if index < 0 || index > l.length {
		return nil
	}
	res := l.data[index]
	return res
}

func (l *SqList) Find(item interface{}) int {
	for i := 0; i < l.length; i++ {
		if l.data[i] == item {
			return i
		}
	}
	return -1
}

func (l *SqList) Insert(index int, item interface{}) bool {
	if l.length >= l.maxLen {
		return false
	}
	if index < 0 || index > l.length {
		return false
	}
	if index > l.length {
		l.data = append(l.data, item)
	} else {
		newData := append(l.data[:index], item)
		l.data = append(newData, l.data[index:])
	}
	l.length += 1
	return true
}

func (l *SqList) Append(item interface{}) bool {
	if l.length >= l.maxLen {
		return false
	}
	l.data[l.length] = item
	l.length += 1
	return true
}
func (l *SqList) Del(index int) bool {
	if index < 0 || index > l.length {
		return false
	}
	l.data[index] = nil
	l.length -= 1
	return true
}
