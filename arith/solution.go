package arith

// 两数之和
func topic1(nums []int, target int) []int {
	_len := len(nums)
	for i := 0; i < _len-1; i++ {
		for j := i + 1; j < _len; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	panic("not result")
}

// 两数之和 map版
func topic1Map(nums []int, target int) []int {
	_len := len(nums)
	var m map[int]int
	m = make(map[int]int, _len)
	m[nums[0]] = 0
	for i := 1; i < _len; i++ {
		val, ok := m[target-nums[i]]
		if ok {
			return []int{val, i}
		}
		m[nums[i]] = i
	}
	panic("not result")
}

// 两数相加
func topic2() {
	return
}
