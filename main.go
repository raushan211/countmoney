package main

import "fmt"

func main() {
	n := 10
	fmt.Println(totalMoney(n))

}
func totalMoney(n int) int {
	sum := 0
	j := 1
	week := 1
	for i := 1; i <= n; i++ {

		if i%7 == 0 {
			sum = sum + j
			week++
			j = week
		} else {
			sum = sum + j
			j++
		}
	}
	return sum
}
