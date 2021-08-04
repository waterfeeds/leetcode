package main

func main() {

}

func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return len(nums)
	}
	var fast, slow = 1, 1
	for fast < len(nums) {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
