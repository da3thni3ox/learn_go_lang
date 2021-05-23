package main
import "fmt"

func main() {
	var n,numbers,i,result int

	fmt.Scan(&n)

	for  i<n{
		fmt.Scan(&numbers)

		if numbers%8 == 0 && numbers >= 10 && numbers < 100{
	             result += numbers
		}
	  i++
        }

	fmt.Println(result)
}
