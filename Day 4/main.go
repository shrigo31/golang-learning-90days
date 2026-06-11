package main

import "fmt"

func printName() {
	fmt.Println("Shridhar")
}

func printCity(city string) {
	fmt.Println("City : ", city)
}

func multiply(a int, b int) int {
	return a * b
}

func isEven(number int) bool {
	return number%2 == 0
}

func calculateArea(length int, width int) int {
	return length * width
}

func calculateAnnualSalary(monthlySalary int) int {
	return monthlySalary * 12
}

func calculateBonus(monthlySalary int) int {
	if monthlySalary >= 50000 {
		return (monthlySalary * 10) / 100
	} else {
		return (monthlySalary * 5) / 100
	}
}

func displayEmployeeDetails(name string, monthlySalary int, annualSalary int, bonus int) {
	fmt.Println("Employee Details:")
	fmt.Println("-----------------")
	fmt.Println("Name: ", name)
	fmt.Println("Monthly Salary: ", monthlySalary)
	fmt.Println("Annual Salary: ", annualSalary)
	fmt.Println("Bonus: ", bonus)
}

func main() {
	printName()
	printCity("Bangalore")
	result := multiply(5, 10)
	fmt.Println("Multiplication Result:", result)
	fmt.Println("Is 5 Even?:", isEven(5))
	fmt.Println("Is 10 Even?:", isEven(10))
	fmt.Println("Area of Rectangle (5x10):", calculateArea(5, 10))
	fmt.Println("Annual Salary (5000):", calculateAnnualSalary(5000))
	fmt.Println("Bonus (49000):", calculateBonus(49000))
	fmt.Println("Bonus (51000):", calculateBonus(51000))
	displayEmployeeDetails("Shridhar", 55000, calculateAnnualSalary(55000), calculateBonus(55000))
}
