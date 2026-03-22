package product_except_self

// First solution
func productExceptSelf(nums []int) []int {
    post := make([]int,len(nums))
    pre := make([]int,len(nums))
    prod :=1
    for i:=0;i<len(nums);i++{
        prod *= nums[i]
        pre[i]=prod
    }
    prod = 1
    for i:=len(nums)-1;i>=0;i--{
        prod *= nums[i]
        post[i]=prod
    }
    res := make([]int,0)
    xx := 0
    for i:=0;i<len(nums);i++{
        if i==0{
            xx = post[i+1]
        }else if i== len(nums)-1{
            xx = pre[i-1]
        }else{
            xx = pre[i-1]*post[i+1]
        }
        res = append(res,xx)
    }
    return res
}

// Second solution, less memory, less runTime
func productExceptSelf2(nums []int) []int {
    res := make([]int,len(nums))
    pre :=1
    res[0]=1
    for i:=0;i<len(nums)-1;i++{
        pre *= nums[i]
        res[i+1]=pre
    }
    post:=1
    for i:=len(nums)-1;i>=0;i--{
        res[i]=post*res[i]
        post *=nums[i]        
    }
    
    return res
}