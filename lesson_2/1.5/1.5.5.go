package main
import "fmt"

func main() {
   var number,result int 

   fmt.Scan(&number)

   number = number%100
   result = number/10
   fmt.Println(result)
}
