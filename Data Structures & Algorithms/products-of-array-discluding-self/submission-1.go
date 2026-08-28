func productExceptSelf(nums []int) []int {
    n := len(nums)
    result := make([]int,n)

    //prfix pass
    result[0] = 1
    for i := 1; i < n; i++ {
        result[i] = result[i - 1] * nums[i - 1]
    }

    // suffix pass
    rightproduct := 1
    for i:= n - 1; i >= 0; i-- {
        result[i] *= rightproduct
        rightproduct *= nums[i]
    }
    return result
}
