package main

import "fmt"

// 1
// 1 1
// 1 2 1
// 1 3 3 1
// 1 4 6 4 1
// f(m, n) = f(m - 1, n) + f(m - 1, n - 1)

func tri1(m int, n int) int {
	if m < 0 || n < 0 || n > m {
		return 0
	}
	if m <= 1 || n == 0 || m == n {
		return 1
	}
	return tri1(m - 1, n) + tri1(m - 1, n - 1)
}

func print_tri(m int, n int)  {
	if m < 0 || n < 0 || n > m {
		return
	}
	fmt.Println("[1]")
	if m > 1 {
		fmt.Println("[1 1]")
	}
	pre := make([]int, n + 1, n + 1)
	pre[0], pre[1] = 1, 1
	tmp := make([]int, n + 1, n + 1)
	var mc int
	for i := 2; i <= m; i++ {
		tmp[0] = 1
		mc = n
		if i < n {
			mc = i
		}
		for j := 1; j <= mc; j++ {
			if j == i {
				tmp[j] = 1
			} else {
				tmp[j] = pre[j] + pre[j - 1]
				//fmt.Printf("tmp[%d](%d) = pre[%d](%d) + pre[%d](%d)\n", j, tmp[j], j, pre[j], j - 1, pre[j - 1])
			}
		}
		
		//copy(pre, tmp)
		pre, tmp = tmp, pre
		fmt.Println(pre)
	}
}

func main()  {
	fmt.Printf("tri1(24, 5) = %d\n", tri1(24, 5))

	print_tri(24, 12)
}