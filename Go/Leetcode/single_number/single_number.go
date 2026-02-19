package single_number

// Map solution
// func singleNumber(nums []int) int {
//     h := make(map[int]int,0)
//     for _,n := range nums {
//         h[n]++
//     }
//     for key,_ := range h{
//         if h[key] ==1 {
//             return key
//         }
//     }
//     return 0
// }

// XOR Solution
func singleNumber(nums []int) int {
    res :=0
    for _,n := range nums {
        res = n^res
    }
    return res
}