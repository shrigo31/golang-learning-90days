package main

import "fmt"

func main() {

	//Problem 1
	//---------------------------
	age := 25

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	//Problem 2
	//---------------------------

	salary := 55000

	if salary >= 30000 {
		fmt.Println("You are eligible for the loan.")
	} else {
		fmt.Println("You are not eligible for the loan.")
	}

	//Problem 3
	//---------------------------

	username := "admin"
	password := "1234"

	if username == "admin" && password == "1234" {
		fmt.Println("Login successful.")
	} else {
		fmt.Println("Login failed.")
	}

	//Problem 4
	//---------------------------

	number := 15

	if number%2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}

	//mini assignment
	//---------------------------

	employeeName := "Shridhar"

	var bonus int

	if salary >= 100000 {
		bonus = ((salary * 20) / 100)
	} else if salary >= 50000 {
		bonus = ((salary * 10) / 100)
	} else if salary >= 25000 {
		bonus = ((salary * 5) / 100)
	} else {
		bonus = 0
	}

	fmt.Println("Employee Bonus Details:")
	fmt.Println("--------------------------")
	fmt.Println("Name: ", employeeName)
	fmt.Println("Salary: ", salary)
	fmt.Println("Bonus: ", bonus)
}
