package pivot_index

func pivotIndex(nums []int) int {
     for i:=0;i<len(nums);i++{
        res := sumthem(nums,i)
        if res == true {
            return i
        }
     }
     return -1
}
func sumthem(nums []int,index int)bool{
    lsum :=0
    rsum :=0
    for i:=0;i<index;i++{
        lsum += nums[i]
    }
    for i:=index+1;i<len(nums);i++{
        rsum += nums[i]
    }
    if rsum == lsum {
        return true
    }
    return false
}