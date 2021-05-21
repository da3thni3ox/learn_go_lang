package main
import "fmt"

func main() {
    var number,number_left_1,number_left_2,number_left_3 int
    var number_right_1,number_right_2,number_right_3 int
    var number_left_sum,number_right_sum int 

    fmt.Scan(&number)

    number_left_1 = number/100000
    number_left_2 = (number/10000)%10
    number_left_3 = (number/1000)%10
    number_right_1 = (number/100)%10
    number_right_2 = (number/10)%10%10
    number_right_3 = number%10%10%10

    number_left_sum = number_left_1+number_left_2+number_left_3
    number_right_sum = number_right_1+number_right_2+number_right_3


    if number_left_sum == number_right_sum {
        fmt.Println("YES")
    }else {
        fmt.Println("NO")
    }

}
