package main

import (
    "fmt"
)

func main() {
    var n int = 10
    result := ""
    if n > 0 {
        for i := 1; i <= n; i++ {
            result += fmt.Sprintf("%d * x %d = ", i, i)
            for j := 1; j <= n; j++ {
                result += fmt.Sprintf("%d", (i * j))
            }
            result += "\n"
        }
    }
    fmt.Println(result)
}
