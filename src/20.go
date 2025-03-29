package main

import "fmt"

func main() {
    var input string = `2 + 3 * 4`
    result := evaluate(input)
    fmt.Println(result)
}

func evaluate(expression string) float64 {
    numbers := strings.Fields(expression)
    var result float64 = 0.0
    for _, number := range numbers {
        value, err := strconv.ParseFloat(number, 64)
        if err != nil {
            return 0.0 // handle parsing errors
        }
        switch number {
        case "+":
            result += value
        case "-":
            result -= value
        case "*":
            result *= value
        case "/":
            if value == 0.0 {
                return 0.0 // handle division by zero
            }
            result /= value
        default:
            fmt.Printf("Unknown operator: %s\n", number)
            return 0.0
        }
    }
    return result
}
