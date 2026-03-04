package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	maxretries = 4
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
	Firstname, _ := reader.ReadString('\n')
	Firstname = strings.TrimSpace(strings.Title(Firstname))
	if Firstname == "" {
		fmt.Println(style("ERROR❗❗ FIRSTNAME CANNOT BE EMPTY", "\033[1;31m", ""))
		fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
		fmt.Println("")
		fmt.Println("RE-ENTER YOUR FIRSTNAME")
		goto start
	}

	if strings.Contains(Firstname, " ") {
		fmt.Println(style("ERROR❗❗ FIRSTNAME CANNOT CONTAIN SPACES IN-BETWEEN", "\033[1;31m", ""))
		fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
		fmt.Println("")
		fmt.Println("RE-ENTER YOUR FIRSTNAME")
		goto start
	}
	for _, char := range Firstname {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
		} else {
			fmt.Println(" ")
			fmt.Println(style("ERROR❗❗ ONLY ALPHABETS ARE ALLOWED", "\033[1;31m", ""))
			fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
			fmt.Println(" ")
			fmt.Print("PLEASE RE-ENTER YOUR FIRSTNAME: ")
			continue
		}
	}

	// LASTNAME VERIFICATION

	fmt.Println(" ")

	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("INPUT YOUR LASTNAME: ")
start2:
	Lastname, _ := reader2.ReadString('\n')
	Lastname = strings.TrimSpace(strings.Title(Lastname))
	if Lastname == "" {
		fmt.Println(style("ERROR❗❗ LASTNAME CANNOT BE EMPTY", "\033[1;31m", ""))
		fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
		fmt.Println("")
		fmt.Print("PLEASE RE-ENTER YOUR LASTNAME: ")
		goto start2
	}

	if strings.Contains(Lastname, " ") {
		fmt.Println(style("ERROR❗❗ LASTNAME CANNOT CONTAIN SPACES IN-BETWEEN", "\033[1;31m", ""))
		fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
		fmt.Println("")
		fmt.Print("PLEASE RE-ENTER YOUR LASTNAME: ")
		goto start2

	}

	for _, char := range Lastname {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
		} else {
			fmt.Println(" ")
			fmt.Println(style("ERROR❗❗ ONLY ALPHABETS ARE ALLOWED", "\033[1;31m", ""))
			fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
			fmt.Println(" ")
			fmt.Print("PLEASE RE-ENTER YOUR LASTNAME: ")
			continue
		}
	}

	//PRINTING THE VERIFIED FIRSTNAME AND LASTNAME
	fmt.Println(" ")
	fmt.Printf(style("HELLO %s %s! YOU ARE VERIFIED\n", "\033[1;32m", ""), Firstname, Lastname)
	fmt.Println(" ")
	fmt.Println(style("=========================================", "\033[1;32m", ""))
	fmt.Println(" ")
	fmt.Println(style("TO ENTER VIP SECTION, YOU NEED TO INPUT AN ADMIN PASSWORD", "\033[1;33m", ""))

	// PASSWORD VERIFICATION

	reader3 := bufio.NewReader(os.Stdin)
	attempt := 1

	for attempt <= maxretries {
		fmt.Println(style("PLEASE ENTER YOUR ADMIN PASSWORD: ", "\033[1;33m", ""))

		Password, _ := reader3.ReadString('\n')
		Password = strings.TrimSpace(Password)

		if Password == "" {
			fmt.Println("PASSOWRD CANNOT BE LEFT EMPTY")
			fmt.Println("PLEASE TRY AGAIN")
			fmt.Println("RE-ENTER YOUR ADMIN PASSWORD")
			continue
		}

		if strings.Contains(Password, " ") {
			fmt.Println("PASSWORD CANNOT CONTAIN SPACES")
			fmt.Println("PLEASE TRY AGAIN")
			fmt.Println("RE-ENTER YOUR ADMIN PASSWORD")
			continue
		}

		fmt.Printf("INCORRECT PASSWORD. ATTEMPTS REMAINING: %d\n", maxretries-attempt)
		attempt++
	}

	if attempt == maxretries {
		fmt.Println(style("TO MANY FAILED ATTEMPTS. ACCESS DENIED🚫 YOU ARE NOT AUTHORIZED TO PROCEED. PROGRAM WILL TERMINATE NOW...", "\033[1;5;31m", ""))
	}

}
