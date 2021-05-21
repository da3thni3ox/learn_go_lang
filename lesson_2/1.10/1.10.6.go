package main
import "fmt"

func main() {
  var numbers,result int

    for true {
     fmt.Scan(&numbers)

     if numbers < 10 {
            continue
     }else if numbers > 100 {
            break
     }else{
	    result = numbers
        }
	    fmt.Println(result)
     }

}
