package count_Pairs

// o(n^2)
func countPairs(nums []int, target int) int {
    c:=0
    for i:=0;i<len(nums);i++{
        for j:=i+1;j<len(nums);j++{
            if nums[i]+nums[j]< target{
                c++
            }
        }
    }
    return c
}