package sort

func Reverse(arr1 []int) []int {
	n := len(arr1)
	for i := 0; i < n/2; i++ {
		tmp := arr1[i]
		oppIndex := n - i - 1
		arr1[i] = arr1[oppIndex]
		arr1[oppIndex] = tmp
	}
	return arr1
}
