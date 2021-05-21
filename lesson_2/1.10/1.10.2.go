package main
import "fmt"

func main() {
  var a,b,sum,i int

    fmt.Scan(&a)
    fmt.Scan(&b)

     i=a

     for i<=b {
          sum+=i
          i++
     }
     fmt.Println(sum)
}
