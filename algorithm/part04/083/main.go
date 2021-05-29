package main

func main() {

}

func search(nums []int, target int) int {
	var l, mid, r = 0, 0, len(nums) - 1
	for l <= r {
		mid = l + (r-l)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] {
			if nums[l] <= target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && nums[r] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
