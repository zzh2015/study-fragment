package main

import (
    "fmt"
)


func sum_1(n int) (int) {
    sum := 0
    for i := 0; i < n; i++ {
        if (i % 3 == 0) || (i % 5 == 0) {
            sum += i
        }
    }
    return sum
}

func sum_2(target, n int) (int) {
    sum := target / n

    return n*sum*(sum+1)/2
}

func main() {
    fmt.Println("sum=", sum_1(1000))
    fmt.Println("sum=", sum_2(999, 3) + sum_2(999, 5) - sum_2(999, 15))
}
