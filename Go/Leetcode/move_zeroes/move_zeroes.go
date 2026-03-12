package move_zeroes

func moveZeroes(nums []int)  {
    j:=0
    for i:=0;i<len(nums);i++{
        if nums[i] == 0 && i<len(nums)-1{
            j=i+1
            for j<len(nums)-1 && nums[j]==0{
                j++
            }
            nums[i]=nums[j]
            nums[j]=0
        }
    }

}