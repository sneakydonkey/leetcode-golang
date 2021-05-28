func twoSum(nums []int, target int) []int {
    res := make([]int, 2)
    valueMap := make(map[int]int)
    for index, value := range nums {
        if preExistedIndex, ok := valueMap[target - value]; ok {
            res[0] = preExistedIndex
            res[1] = index
            return res
        } else {
            valueMap[value] = index
        }
    }
    return res
}
