package maxOperations

func maxOperations(nums []int, k int) int {
    hmap := make(map[int]int,0)
    count := 0
    for i:=0;i<len(nums);i++{
        x := k-nums[i]
        _,ok:= hmap[x]
        if ok && hmap[x]>0 {
            hmap[x]--
            count++
        }else{
            hmap[nums[i]]++
        }
        
    }
      return count
}

