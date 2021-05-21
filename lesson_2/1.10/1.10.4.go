package main
import "fmt"

func main() {
	var n,result,max int
		n = 1

		for n !=0 {
			fmt.Scan(&n)

		if n > max {
			max=n
			result = 0
		}

		if n == max {
			result++
		}
	}

	fmt.Println(result)
}
