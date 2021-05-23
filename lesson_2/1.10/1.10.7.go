package main
import "fmt"

func main() {
    var x,p,y,i int

    fmt.Scan(&x)
    fmt.Scan(&p)
    fmt.Scan(&y)

    for x<y {
         x+=int((x*p)/100)
         i++
    }

        fmt.Println(i)

}
