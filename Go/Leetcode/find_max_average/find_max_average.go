package find_max_average

// Correct but Time limit exceeded solution
// func findMaxAverage(nums []int, k int) float64 {
//     var max float64
//     sum := 0
//     av := 0.0
//     for i:=0;i<len(nums);i++{
//         sum = nums[i]
//         if len(nums)-i < k{
//             break
//         }
//         for j:=i+1;j<i+k;j++{
//             sum = nums[j] + sum
//         }
//         av = float64(sum)/float64(k)
//         if i == 0 {
//             max = av
//         }
//         if av > max {
//             max = av
//         }
//     }
//     return max
// }

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
    for i:=0;i<k;i++{
        sum += nums[i]
    }
    max_av := float64(sum)/float64(k)
    for i:=k;i<len(nums);i++{
        sum += nums[i]
        sum -= nums[i-k]
        max_av = max(max_av, float64(sum)/float64(k))
    }
    return max_av
}