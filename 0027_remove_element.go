/*
  https://leetcode.com/problems/remove-element/
*/
func removeElement(nums []int, val int) int {
    index := 0
    for i := 0; i < len(nums); i++ {
        if val != nums[i] {
            nums[index] = nums[i]
            index++
        }
    }
    return index
}
