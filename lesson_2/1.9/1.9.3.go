package main
import "fmt"

func main() {
    var number int 

    fmt.Scan(&number)


     if number > 1000 && number < 10000 {
        number = number / 1000
     } else if number > 100 && number < 1000 {
        number = number / 100
     } else if number > 10 && number < 100 {
        number = number / 10
     } else if number == 10000 {
        number = number / 10000
     }


     fmt.Println(number)
}
