package pascal

import (
	"strconv"
)

// DrawTriangle creates the pascal triangle
func DrawTriangle(n int) string {
	temp := 1
	ans := ""

	for i := 0; i < n; i++ {

		for j := 1; j <= n-i; j++ {
			ans += " "
		}

		for k := 0; k <= i; k++ {

			if k == 0 || i == 0 {
				temp = 1
			} else {
				temp = temp * (i - k + 1) / k
			}
			
			ans = ans + " " + strconv.Itoa(temp)
		}

		ans = ans + "\n"
	}
	return ans
}
