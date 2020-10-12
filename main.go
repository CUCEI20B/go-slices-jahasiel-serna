package main

import "fmt"

func main()  {
	var size, n, sum int
	fmt.Scan(&size)
	s := make([]int, 0, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&n)
		s = append(s, n)
	}
	for _, j := range s {
		sum += j
	}
	fmt.Println(sum)
}