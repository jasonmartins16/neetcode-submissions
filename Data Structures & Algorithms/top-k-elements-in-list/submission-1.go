func topKFrequent(nums []int, k int) []int {
    n := len(nums)
    count := make(map[int]int)

    // count freq
    for _, num:= range nums {
        count[num]++
    }

    // buckets slice of slice when index = freq
    buckets := make([][]int, n+1)
    for num, freq := range count {
        buckets[freq] = append(buckets[freq], num)
    }

    // collect elements from bucket with highest freq to 1
    result := make([]int, 0, k)
    for freq := n; freq >= 1; freq-- {
        for _, num := range buckets[freq] {
            result = append(result, num)
            if len(result) == k {
                return result
            }
        }
    }
    return result
}   
