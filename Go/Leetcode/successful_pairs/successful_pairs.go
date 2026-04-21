package successful_pairs

func successfulPairs(spells []int, potions []int, success int64) []int {
    res := make([]int,0)
    for _,s := range spells {
        r:=0
        for _,p := range potions {
            num := s*p
            if int64(num) >= success {
                r ++
            }
        }
        res =append(res, r)
    }
    return res
}