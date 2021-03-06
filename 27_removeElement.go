package main

func main() {

}

func removeElement(nums []int, val int) int {
	var fast, slow = 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}

		fast++
	}

	return slow
}
