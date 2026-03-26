package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(" ")
	fmt.Println("*************** WELCOME TO BLASIE PASCAL'S CALCULATOR ****************")
	fmt.Println(" ")

start:
	fmt.Print("FIRST VALUE: ")

	var first_value string
	fmt.Scanln(&first_value)

	 switch first_value {

	case "quit" :
		fmt.Println("GOODBYE")
		return

	case "help" :
		fmt.Println("-------- AVAILABLE COMMAND -----------")
		fmt.Println("<a> + <b>   → addition")
		fmt.Println("<a> - <b>   → subtraction")
		fmt.Println("<a> * <b>   → multiplication")
		fmt.Println("<a> / <b>   → division")
		fmt.Println(" ")
		goto start
	

	 }
	
	num1, err := strconv.ParseFloat(first_value, 64) 

	if err != nil {
		fmt.Println("VALUE MUST BE AN INTEGER")
		fmt.Println("TRY AGAIN")
		fmt.Println(" ")
		goto start

	}

operate:
	fmt.Print("OPERATOR: ")
	var sign string
	fmt.Scanln(&sign)
	
	if sign == "quit" {
		fmt.Println("GOODBYE")
		return
	}

second:
	fmt.Print("SECOND VALUE: ")
	var second_value string
	fmt.Scanln(&second_value)

	switch second_value {
	case "quit" :
		fmt.Println("GOODBYE")
		return

	case "help" :
		fmt.Println("-------- AVAILABLE COMMAND -----------")
		fmt.Println("<a> + <b>   → addition")
		fmt.Println("<a> - <b>   → subtraction")
		fmt.Println("<a> * <b>   → multiplication")
		fmt.Println("<a> / <b>   → division")
		fmt.Println(" ")
		goto second

	
	 }


	num2, err := strconv.ParseFloat(second_value, 64) 

	if err != nil {
		fmt.Println("VALUE MUST BE AN INTEGER")
		fmt.Println("TRY AGAIN")
		fmt.Println(" ")
		goto second
	}

	var result float64

	switch sign {

	case "+":
		fmt.Printf("RESULT: %.2f %s %.2f = %.2f\n", num1, sign, num2, result + num1 + num2)

	case "-" :
		fmt.Printf("RESULT: %.2f %s %.2f = %.2f\n", num1, sign, num2, result + num1 - num2)

	case "*" :
		fmt.Printf("RESULT: %.2f %s %.2f = %.2f\n", num1, sign, num2, result + num1 * num2)

	case "/" :
		if second_value == "0" {
			fmt.Println("THE DIVISOR CANNOT BE ZERO")
			fmt.Println("TRY AGAIN")
			goto second

		}
		fmt.Printf("RESULT: %.2f %s %.2f = %.2f\n", num1, sign, num2, result + num1 / num2)
		
	case "help":
		fmt.Println("-------- AVAILABLE COMMAND -----------")
		fmt.Println("<a> + <b>   → addition")
		fmt.Println("<a> - <b>   → subtraction")
		fmt.Println("<a> * <b>   → multiplication")
		fmt.Println("<a> / <b>   → division")
		fmt.Println(" ")
		goto operate

	default :
		fmt.Println("INVALID ARITHMETIC SYNTAX")
		fmt.Println("TRY AGAIN")
		goto start
	}
	fmt.Println(" ")
	goto start

}
