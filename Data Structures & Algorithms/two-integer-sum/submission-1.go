func twoSum(nums []int, target int) []int {
    count := make(map[int]int)
    for i, num := range nums {
        compliment := target - num
        if prevIndex, found := count[compliment]; found {
            return []int{prevIndex, i}
        }
        count[num] = i
    }
    return nil
}
