package main

import (
	"fmt"
)

func main() {
	fmt.Println("PLEASE ENTER YOUR DETAILS TO PROCEED")
	fmt.Println("INPUT YOUR FIRSTNAME")
	start1:
	var Firstname string
	fmt.Scanln(&Firstname)
	for _, char := range Firstname {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') == true {
		} else {
			fmt.Println("ERROR; ONLY ALPHABETS ARE ALLOWED❗❗❗")
			fmt.Println("TRY AGAIN")
			fmt.Println("PLEASE RE-ENTER YOUR FIRSTNAME")
			goto start1
		}
	}
	fmt.Println("PLEASE ENTER YOUR LASTNAME")
	start2:
	var Lastname string
	fmt.Scanln(&Lastname)
	for _, char := range Lastname {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') == true {
		} else {
			fmt.Println("ERROR; ONLY ALPHABETS ARE ALLOWED❗❗❗.")
			fmt.Println("TRY AGAIN")
			fmt.Println("PLEASE RE-ENTER YOUR LASTNAME")
			goto start2
		}
	}
	fmt.Println("ADMIN PASSWORD REQUIRED TO PROCEED")
	var Password int
	fmt.Scanln(&Password)
	if len(fmt.Sprintf("%d", Password)) != 6 {
		fmt.Println("OH SNAP❌❗ WRONG PASSWORD FORMAT. ACCESS DENIED🚫")
		return
	}
	fmt.Printf("SPECIAL GUEST DETECTED! Hello %s %s!,WELCOME TO 𝗛 𝗘 𝗔 𝗩 𝗘 𝗡 👑🕊️ \n", Firstname, Lastname)
	fmt.Println("YOU HAVE BEEN GRANTED ACCESS TO THE GATES OF 𝗛 𝗘 𝗔 𝗩 𝗘 𝗡 DO ENJOY YOUR STAY HERE💫")


}


// func main() {
// 	fmt.Print("Please, input your name: ")
// 	var name string
// 	fmt.Scan(&name)
// 	fmt.Println("Welcome, ", name)
// }
