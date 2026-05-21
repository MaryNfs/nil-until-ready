package contains_duplicate

func containsDuplicate(nums []int) bool {
    hash := make(map[int]int,0)
    for _,n:= range nums {
        _,ok := hash[n]
        if ok {
            return true
        }else{
            hash[n]++
        }
    }
    return false
}