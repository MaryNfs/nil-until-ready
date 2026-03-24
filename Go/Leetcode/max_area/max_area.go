package max_area

// Time Limit Exceeded Solution
func maxArea(height []int) int {
    max := 0
    for i:=0;i<len(height);i++{
        for j:=i;j<len(height);j++{
            min := min(height[i],height[j])
            space := (j-i)*min
            if space > max {
                max = space
            }
        }
    }
    return max
}