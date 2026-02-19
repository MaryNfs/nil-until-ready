package single_number

func singleNumber(nums []int) int {
    h := make(map[int]int,0)
    for _,n := range nums {
        h[n]++
    }
    for key,_ := range h{
        if h[key] ==1 {
            return key
        }
    }
    return 0
}
