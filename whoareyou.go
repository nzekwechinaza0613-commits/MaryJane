package main

import (
	"fmt"
)

func main() {
	fmt.Println("Please input your firstname")
	var Firstname string
	fmt.Scanln(&Firstname)
	for _, char := range Firstname {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') == true {
		} else {
			fmt.Println("ERROR; ONLY ALPHABETS ARE ALLOWED❗❗❗")
			return
		}
	}
	fmt.Println("Please input your lastname")
	var Lastname string
	fmt.Scanln(&Lastname)
	for _, char := range Lastname {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') == true {
		} else {
			fmt.Println("ERROR; ONLY ALPHABETS ARE ALLOWED❗❗❗.")
			return
		}
	}
	fmt.Println("Admin Password Required")
	var Password string
	fmt.Scanln(&Password)
	if Password != "12345" {
		fmt.Println("OH SNAP❌ Password is incorrect")
		return
	}

	fmt.Printf("SPECIAL GUEST DETECTED! Hello %s %s!,WELCOME TO 𝗛 𝗘 𝗔 𝗩 𝗘 𝗡 👑🕊️ \n", Firstname, Lastname)

}

// func main() {
// 	fmt.Print("Please, input your name: ")
// 	var name string
// 	fmt.Scan(&name)
// 	fmt.Println("Welcome, ", name)
// }
