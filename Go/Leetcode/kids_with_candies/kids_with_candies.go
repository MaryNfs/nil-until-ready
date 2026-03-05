package kids_with_candies

func kidsWithCandies(candies []int, extraCandies int) []bool {
    res := make([]bool,len(candies))
    for k,v := range candies {
        t := v + extraCandies
        res[k]=true
        for i:=0;i<len(candies);i++{
            if candies[i] > t{
               res[k]=false
            }
        }
    }
    return res
}
