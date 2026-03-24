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

func main() {
	fmt.Println(" ")
	fmt.Println(style("********** WELCOME TO JOHN MCPHERSONS CALCULATOR ***********", "\033[1;33m", ""))
	fmt.Println(" ")

start :
	fmt.Print("FIRST VALUE : ")
	var first_value string
	fmt.Scanln(&first_value)

	num1, err1 := strconv.ParseFloat(first_value, 64)

	if err1 != nil {
		fmt.Println(style("VALUE MUST BE AN INTERGER", "\033[1;31m", ""))
		fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
		fmt.Println(" ")
		goto start
	}


	fmt.Print("OPERATOR : ")
	var sign string
	fmt.Scanln(&sign)

start2 :
	fmt.Print("SECOND VALUE : ")
	var second_value string
	fmt.Scanln(&second_value) 

	num2, err2 := strconv.ParseFloat(second_value, 64)

	if err2 != nil {
		fmt.Println(style("VALUE MUST BE AN INTERGER", "\033[1;31m", ""))
		fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
		fmt.Println(" ")
		goto start2
	}
	
	if sign == "+" {
		fmt.Printf("%.2f %s %.2f = %v\n", num1, sign, num2, num1 + num2)
	} else if sign == "-" {
		fmt.Printf("%.2f %s %.2f = %v\n", num1, sign, num2, num1 - num2)
	} else if sign == "*" {
		fmt.Printf("%.2f %s %.2f = %v\n", num1, sign, num2, num1 * num2)
	} else if sign == "/" {
		fmt.Printf("%.2f %s %.2f = %v\n", num1, sign, num2, num1 / num2)
	} else if sign == "%" {
		fmt.Printf("%v %s %v = %v\n", num1, sign, num2, int(num1) % int(num2))
	} else if sign == "sqrt" {
		fmt.Printf("√%.2f = %.2f\n", num1, squareRoot(num1))
	} else {
		fmt.Println(style("INVALID ARITHEMETIC SYNTAX", "\033[1;31m", ""))
		fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
		fmt.Println(" ")
		goto start
	}
}






// package main

// import "fmt"

// func main() {

// 	var result string

// 	result = "I am MJ and I am learning Go"

// 	fmt.Println(result)
// }
