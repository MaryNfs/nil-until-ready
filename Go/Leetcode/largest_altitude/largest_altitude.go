package largest_altitude

func largestAltitude(gain []int) int {
    alt := make([]int, 0)
    alt = append(alt,0)
    pre := 0
    for i:=0;i<len(gain);i++{
        pre += gain[i]
        alt = append(alt,pre)
    }
    max := alt[0]
    for i:=0;i<len(alt);i++{
        if alt[i]> max{
            max = alt[i]
        }
    }
    return max
}