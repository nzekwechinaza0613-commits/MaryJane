package main

import (
	"fmt"
)
func style(text, style, color string) string {
	return style + color + text + "\033[0m"
}

const (
	maxretry = 3
)

func main() {
	fmt.Println(" ")
	fmt.Println(style("********** IDENTITY VERIFICATION ***********", "\033[1;33m", ""))
	fmt.Println(" ")
	fmt.Println(style("PLEASE ENTER YOUR DETAILS TO PROCEED ", "\033[1;33m",""))
	fmt.Println(" ")
	fmt.Print("INPUT YOUR FIRSTNAME: ")
	start1:
	var Firstname string
	fmt.Scanln(&Firstname)
	if Firstname == "" || Firstname == " " {
		fmt.Println(" ")
		fmt.Println(style("ERROR❗❗ FIRSTNAME CANNOT BE EMPTY OR CONTAIN SPACES", "\033[1;31m",""))
		fmt.Println(style("TRY AGAIN", "\033[1;31m",""))
		fmt.Println(" ")
		fmt.Print(("PLEASE RE-ENTER YOUR FIRSTNAME: "))
		goto start1
	}
	for _, char := range Firstname {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') == true {
		}else {
			fmt.Println(" ")
			fmt.Println(style("ERROR❗❗ ONLY ALPHABETS ARE ALLOWED", "\033[1;31m",""))
			fmt.Println(style("TRY AGAIN", "\033[1;31m",""))
			fmt.Println(" ")
			fmt.Print("PLEASE RE-ENTER YOUR FIRSTNAME: ")
			goto start1
		}
	}
	fmt.Print("INPUT YOUR LASTNAME: ")
	start2:
	var Lastname string
	fmt.Scanln(&Lastname)
	if Lastname == "" || Lastname == " " {
		fmt.Println(" ")
		fmt.Println(style("ERROR❗❗ LASTNAME CANNOT BE EMPTY OR CONTAIN SPACES", "\033[1;31m",""))
		fmt.Println(style("TRY AGAIN", "\033[1;31m",""))
		fmt.Println(" ")
		fmt.Print("PLEASE RE-ENTER YOUR LASTNAME: ")
		goto start2
	}
	for _, char := range Lastname {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') == true {
		}else {
			fmt.Println(" ")
			fmt.Println(style("ERROR❗❗ ONLY ALPHABETS ARE ALLOWED", "\033[1;31m",""))
			fmt.Println(style("TRY AGAIN", "\033[1;31m",""))
			fmt.Println(" ")
			fmt.Print("PLEASE RE-ENTER YOUR LASTNAME: ")
			goto start2
		}
	}
	fmt.Println(" ")
	fmt.Printf(style("HELLO %s %s! YOU ARE VERIFIED\n", "\033[1;32m",""), Firstname, Lastname)
	fmt.Println(" ")
	fmt.Println(style("=========================================", "\033[1;32m",""))
	fmt.Println(" ")
	fmt.Println(style("TO ENTER VIP SECTION, YOU NEED TO INPUT AN ADMIN PASSWORD", "\033[1;33m",""))
	fmt.Println(" ")
	var Password string
	for attempt := 1; attempt <= maxretry; attempt++ {
	fmt.Println(style("PLEASE ENTER YOUR 6-DIGIT ADMIN PASSWORD: ", "\033[1;33m",""))
	fmt.Println("YOU HAVE ", maxretry-attempt+1, "ATTEMPTS LEFT")
	fmt.Scanln(&Password)
	if len(Password) != 6 {
		fmt.Println(style("OH SNAP❌❗ WRONG PASSWORD FORMAT. PASSWORD IS INCORRECT", "\033[1;31m",""))
		fmt.Println(style("TRY AGAIN", "\033[1;31m",""))
		fmt.Println(" ")
		if attempt == maxretry {
			fmt.Println(style("ATTEMPTS EXHAUSTED. ACCESS DENIED🚫 YOU ARE NOT AUTHORIZED TO PROCEED. PROGRAM WILL TERMINATE NOW...", "\033[1;5;31m",""))
			fmt.Println(" ")
			fmt.Println(style("LOADING......................", "\033[1;31m",""))
			fmt.Println(style("TERMINATING PROGRAM...", "\033[1;31m",""))
			fmt.Println(style("PROGRAM TERMINATED SUCCESSFULLY.", "\033[1;32m",""))
			return
		}
		continue
	}
	for _, char := range Password {
		if char >= '0' && char <= '9' {
		} else {
			fmt.Println("PASSWORD MUST CONTAIN ONLY DIGITS ❗❗❗", style("", "\033[1;31m",""))
			return
		} 
	fmt.Println(style("PASSWORD ACCEPTED ✅", "\033[1;32m",""))
	fmt.Println(style("VIP STATUS CONFIRMED ✅", "\033[1;32m",""))
	fmt.Println(" ")
	fmt.Printf(style("MY CHEIF %s %s!\n", "\033[1;32m",""), Firstname, Lastname)
	fmt.Println(style("WELCOME TO 𝗛 𝗘 𝗔 𝗩 𝗘 𝗡 👑🕊️ \n", "\033[1;32m",""))
	fmt.Println(style("THE SYSTEM BOWS TO YOU RESPECTFULLY OF COURSE", "\033[1;32m",""))
	fmt.Println(style("|></></></></></></></></></></></></></></></></></></></></><|", "\033[1;30m",""))
	return
		}
	}
}




// func main() {
// 	fmt.Print("Please, input your name: ")
// 	var name string
// 	fmt.Scan(&name)
// 	fmt.Println("Welcome, ", name)
// }
