package main

import (
	"fmt"
)

const (
	maxretry = 3
)

func main() {
	fmt.Println("PLEASE ENTER YOUR DETAILS TO PROCEED")
	fmt.Print("INPUT YOUR FIRSTNAME: ")
	start1:
	var Firstname string
	fmt.Scanln(&Firstname)
	if Firstname == "" {
		fmt.Println("ERROR; FIRSTNAME CANNOT BE EMPTY❗❗❗")
		fmt.Println("TRY AGAIN")
		fmt.Print("PLEASE RE-ENTER YOUR FIRSTNAME: ")
		goto start1
	}
	for _, char := range Firstname {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') == true {
		} else {
			fmt.Println("ERROR; ONLY ALPHABETS ARE ALLOWED❗❗❗")
			fmt.Println("TRY AGAIN")
			fmt.Print("PLEASE RE-ENTER YOUR FIRSTNAME: ")
			goto start1
		}
	}
	fmt.Print("INPUT YOUR LASTNAME: ")
	start2:
	var Lastname string
	fmt.Scanln(&Lastname)
	if Lastname == "" {
		fmt.Println("ERROR; LASTNAME CANNOT BE EMPTY❗❗❗")
		fmt.Println("TRY AGAIN")
		fmt.Print("PLEASE RE-ENTER YOUR LASTNAME: ")
		goto start2
	}
	for _, char := range Lastname {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') == true {
		} else {
			fmt.Println("ERROR; ONLY ALPHABETS ARE ALLOWED❗❗❗.")
			fmt.Println("TRY AGAIN")
			fmt.Print("PLEASE RE-ENTER YOUR LASTNAME: ")
			goto start2
		}
	}
	fmt.Println("ADMIN PASSWORD REQUIRED TO PROCEED")
	var Password string
	for attempt := 1; attempt <= maxretry; attempt++ {
	fmt.Print("PLEASE ENTER YOUR 6-DIGIT PASSWORD: ")
	fmt.Println("YOU HAVE ", maxretry-attempt+1, "ATTEMPTS LEFT")
	fmt.Scanln(&Password)
	if len(Password) != 6 {
		fmt.Println("OH SNAP❌❗ WRONG PASSWORD FORMAT. PASSWORD IS INCORRECT")
		fmt.Println("TRY AGAIN")
		if attempt == maxretry {
			fmt.Println("TOO MANY ATTEMPTS. ACCESS DENIED 🚫 INTRUDER ALERT.")
			fmt.Println("TERMINATING PROGRAM...")
			fmt.Println("PROGRAM TERMINATED SUCCESSFULLY.")
			return
		}
		continue
	}
	for _, char := range Password {
		if char >= '0' && char <= '9' {
		} else {
			fmt.Println("PASSWORD MUST CONTAIN ONLY DIGITS❗❗❗")
			return
		}
	fmt.Printf("SPECIAL GUEST DETECTED! Hello %s %s!,WELCOME TO 𝗛 𝗘 𝗔 𝗩 𝗘 𝗡 👑🕊️ \n", Firstname, Lastname)
	fmt.Println("YOU HAVE BEEN GRANTED ACCESS TO THE GATES OF 𝗛 𝗘 𝗔 𝗩 𝗘 𝗡 DO ENJOY YOUR STAY HERE💫")
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
