/*
    https://leetcode.com/problems/search-insert-position/
*/

func searchInsert(nums []int, target int) int {
    var left, right = -1, len(nums)
    for right - left > 1 {
        var mid = (left + right) / 2
        if nums[mid] > target {
            right = mid
        } else if nums[mid] < target {
            left = mid
        } else {
            return mid
        }
    }
    return right
}
