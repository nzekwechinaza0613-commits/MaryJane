
package main

import(
    "fmt"
    "strconv"
)

func style(text, style, color string) string {
    return style + color + text + "\033[0m"
}

func squareRoot(a float64) float64 {

    if a == 0 {

        return 0
    }
    n := a
    for i := 0; i < 10; i++ {
        n = (n + a / n) / 2
    }
    return n
}

// func power(a float64) float64 {

// }

func main() {
    fmt.Println(" ")
    fmt.Println(style("********** WELCOME TO BLAISE PASCAL'S CALCULATOR🧮 ***********", "\033[1;33m", ""))
    fmt.Println(" ")

start :
    fmt.Print("FIRST VALUE: ")
    var first_value string
    fmt.Scanln(&first_value)

    if first_value == "off" {
        return
    }

    num1, err1 := strconv.ParseFloat(first_value, 64)

    if err1 != nil {
        fmt.Println(style("VALUE MUST BE AN INTEGER", "\033[1;31m", ""))
        fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
        fmt.Println(" ")
        goto start
    }


    fmt.Print("OPERATOR: ")
    var sign string
    fmt.Scanln(&sign)

    if sign == "off" {
        return
    }

    if sign == "sqrt" {
        fmt.Println(" ")
        fmt.Printf(style("√%.2f = %.2f\n", "\033[1;32m", ""), num1, squareRoot(num1))
        fmt.Println(" ")
        goto start
    }

start2 :
    fmt.Print("SECOND VALUE: ")
    var second_value string
    fmt.Scanln(&second_value) 

    if second_value == "off" {
        return
    }

    num2, err2 := strconv.ParseFloat(second_value, 64)

    if err2 != nil {
        fmt.Println(style("VALUE MUST BE AN INTEGER", "\033[1;31m", ""))
        fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
        fmt.Println(" ")
        goto start2
    }

    switch sign {
    case "+" :
        fmt.Println(" ")
        fmt.Printf(style("%.2f %s %.2f = %.2f\n", "\033[1;32m", ""), num1, sign, num2, num1 + num2)

    case "-" :
        fmt.Println(" ")
        fmt.Printf(style("%.2f %s %.2f = %.2f\n", "\033[1;32m", ""), num1, sign, num2, num1 - num2)

    case "*" :
        fmt.Println(" ")
        fmt.Printf(style("%.2f %s %.2f = %.2f\n", "\033[1;32m", ""), num1, sign, num2, num1 * num2)

    case "/" :
        fmt.Println(" ")
        fmt.Printf(style("%.2f %s %.2f = %.2f\n", "\033[1;32m", ""), num1, sign, num2, num1 / num2)

    case "%" :
        fmt.Println(" ")
        fmt.Printf(style("%v %s %v = %v\n", "\033[1;32m", ""), num1, sign, num2, int(num1) % int(num2))

    default :
        fmt.Println(style("INVALID ARITHMETIC SYNTAX", "\033[1;31m", ""))
        fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
        fmt.Println(" ")
        goto start
    }
    fmt.Println(" ")
    goto start
}
