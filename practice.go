package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"

)

const (
	max = 4
)

func style(text, style, color string) string {
	return style + color + text + "\033[0m"
}

func containsAtSymbol(email string) bool {
	return strings.Contains(email, "@")
}

func main() {

	//FIRSNAME VERIFICATION

	fmt.Println(" ")
	fmt.Println(style("********** IDENTITY VERIFICATION ***********", "\033[1;33m", ""))
	fmt.Println(" ")
	fmt.Println(style("PLEASE ENTER YOUR DETAILS TO PROCEED ", "\033[1;33m", ""))
	fmt.Println(" ")
	reader := bufio.NewReader(os.Stdin)	
	fmt.Print("INPUT YOUR FIRSTNAME: ")
	start:
	firstname, _ := reader.ReadString('\n')
	firstname = strings.TrimSpace(firstname)
	if strings.Contains(firstname, " ") || firstname == "" {
		fmt.Println(style("ERROR❗❗ FIRSTNAME CANNOT BE EMPTY OR CONTAIN SPACES", "\033[1;31m", ""))
		fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
		fmt.Println("")
		fmt.Println("RE-ENTER YOUR FIRSTNAME")
		goto start
	}
	for _, char := range firstname {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
		} else {
			fmt.Println(" ")
			fmt.Println(style("ERROR❗❗ ONLY ALPHABETS ARE ALLOWED", "\033[1;31m", ""))
			fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
			fmt.Println(" ")
			fmt.Print("PLEASE RE-ENTER YOUR FIRSTNAME: ")
			goto start	
		}
	}

	// LASTNAME VERIFICATION

	fmt.Println(" ")

	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("INPUT YOUR LASTNAME: ")
	start2:
	lastname, _ := reader2.ReadString('\n')
	lastname = strings.TrimSpace(lastname)
	if strings.Contains(lastname, " ") || lastname == "" {
		fmt.Println(style("ERROR❗❗ LASTNAME CANNOT BE EMPTY OR CONTAIN SPACES", "\033[1;31m", ""))
		fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
		fmt.Println("")
		fmt.Print("PLEASE RE-ENTER YOUR LASTNAME: ")
		goto start2
	}
	for _, char := range lastname {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
		} else {
			fmt.Println(" ")
			fmt.Println(style("ERROR❗❗ ONLY ALPHABETS ARE ALLOWED", "\033[1;31m", ""))
			fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
			fmt.Println(" ")
			fmt.Print("PLEASE RE-ENTER YOUR LASTNAME: ")
			goto start2	
		}
	}	

	//PRINTING THE VERIFIED FIRSTNAME AND LASTNAME
	fmt.Println(" ")
	fmt.Printf(style("HELLO %s %s! YOU ARE VERIFIED\n", "\033[1;32m", ""), firstname, lastname)
	fmt.Println(" ")
	fmt.Println(style("=========================================", "\033[1;32m", ""))

	// PASSWORD VERIFICATION
	reader3 := bufio.NewReader(os.Stdin)
	start3:
	for attempts := 1; attempts <= max; attempts++ {
	fmt.Println(style("PLEASE ENTER YOUR ADMIN PASSWORD: ", "\033[1;33m", ""))
	fmt.Println("YOU HAVE", max - attempts+1, "ATTEMPT(S) LEFT")
	return
	}
	password, _ := reader3.ReadString('\n')
	password = strings.TrimSpace(password)
	if len(password) < 8 {
			fmt.Println(style("ERROR❗❗ PASSWORD MUST BE AT LEAST 8 CHARACTERS LONG", "\033[1;31m", ""))
			fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
			fmt.Println(" ")
			goto start3
		}
	if attempts == max {
		fmt.Println(style("YOU HAVE EXCEEDED THE MAXIMUM NUMBER OF ATTEMPTS. ACCESS DENIED.", "\033[1;31m", ""))
			fmt.Println(style("PLEASE CONTACT THE ADMINISTRATOR FOR ASSISTANCE", "\033[1;31m", ""))
			fmt.Println(" ")
			return
		}
		continue
	}
	Password, _ := reader3.ReadString('\n')
	Password = strings.TrimSpace(Password)
	if strings.Contains(Password, " ") || Password == "" {
		fmt.Println(style("ERROR❗❗ PASSWORD CANNOT BE EMPTY OR CONTAIN SPACES", "\033[1;31m", ""))
		fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
		fmt.Println(" ")
		goto start3
	}
var hasUpper bool
var hasLower bool
var hasNumber bool
var hasSymbol bool

	for _, char := range Password {

		if char >= 'A' && char <= 'Z' {
			hasUpper = true
		}

		if char >= 'a' && char <= 'z' {
			hasLower = true
		}

		if char >= '0' && char <= '9' {
			hasNumber = true
		}

		if char == '!' || char == '@' || char == '#' {
			hasSymbol = true
		}
		if !hasUpper || !hasLower || !hasNumber || !hasSymbol {
		fmt.Println(style("PASSWORD IS TOO WEAK", "\033[1;31m", ""))
		fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
		fmt.Println(" ")
		goto start3
		}

	reader4 := bufio.NewReader(os.Stdin)
	fmt.Print(style("CONFIRM PASSWORD: ", "\033[1;33m", ""))
	confirm, _ := reader4.ReadString('\n')
	confirm = strings.TrimSpace(confirm)
		if Password != confirm {
			fmt.Println(style("PASSWORDS DO NOT MATCH", "\033[1;31m", ""))
			fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
			goto start3
		}
	}
}


