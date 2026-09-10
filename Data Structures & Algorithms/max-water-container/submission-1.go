func maxArea(heights []int) int {
    left := 0
    right := len(heights) - 1
    max_area := 0
    for left < right {
        current_area := (right - left) * min(heights[left], heights[right])

        if heights[left] > heights[right] {
            right--
            max_area = max(max_area, current_area)
        } else if heights[left] < heights[right] {
            left++
            max_area = max(max_area, current_area)
        } else {
            max_area = max(max_area, current_area)
            left++
            right--
        }       
    }
    return max_area 
}
