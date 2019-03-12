package sort

func InsertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; j >= 1; j-- {
			if nums[j] < nums[j-1] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			} else {
				break
			}
		}
	}
}
