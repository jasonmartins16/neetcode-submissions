func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    set := make(map[int]bool)
    for _, num := range nums {
        set[num] = true
    }

    streak := 0 

    for num := range set {
        if !set[num - 1] {
            current_num := num
            current_streak := 1

            for set[current_num + 1] {
                current_num++
                current_streak++
            }

            if current_streak > streak {
                streak = current_streak
            }
        } 
    }
    return streak
}