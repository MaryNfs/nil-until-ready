package max_area

// Time Limit Exceeded Solution
// func maxArea(height []int) int {
//     max := 0
//     for i:=0;i<len(height);i++{
//         for j:=i;j<len(height);j++{
//             min := min(height[i],height[j])
//             space := (j-i)*min
//             if space > max {
//                 max = space
//             }
//         }
//     }
//     return max
// }

// two pointers solution
func maxArea(height []int) int {
    max := 0
    j:=len(height)-1
    i:=0;
    for i<len(height) && j>=0 {
            min := min(height[i],height[j])
            space := (j-i)*min
            if space > max {
                max = space
            }
            if height[i]>height[j] {
                j--
            }else{
                i++
            }
    }
    return max
}