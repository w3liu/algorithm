package sort

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > tmp {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = tmp
	}
}

func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		var min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min > i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var l, r int
	result := make([]int, 0, len(left)+len(right))
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	var head, tail = 0, len(arr) - 1
	var mid, i = arr[0], 1
	for head < tail {
		if arr[i] > mid {
			arr[tail], arr[i] = arr[i], arr[tail]
			tail--
		} else {
			arr[head], arr[i] = arr[i], arr[head]
			head++
			i++
		}
	}
	quickSort(arr[:head])
	quickSort(arr[head+1:])
}
