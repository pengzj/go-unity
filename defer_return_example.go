package main
import (
        "fmt"
)

func a() (result int) {
        defer func() {
                result++
                fmt.Println("in a() ", result)
        }()
        return 0
}

func b() (r int) {
        t := 5
        defer func() {
                t = t + 5
                fmt.Println("in b() ", t)
        }()
        return t
}

func c()(r int) {
        defer func(r int) {
                r = r + 5
                fmt.Println("in c() ", r)
        }(r)
        return 1
}

func d() (result int) {

        return
}

func main() {
        val1 := a()
        val2 := b()
        val3 := c()
        val4 := d()
        fmt.Println(val1, val2, val3, val4)
}
