package main

import "fmt"

func main() {
	var n,c,d,result int

	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&d)

	for i := 1; i<=n; i++ {
		if i%c == 0 && i%d != 0 {
			result = i
			fmt.Println(result)
			break
		}
	}

}
