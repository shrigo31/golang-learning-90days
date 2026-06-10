package main

import "fmt"

func main() {

	//Problem 1
	Name := "Shridhar"
	Age := 31
	City := "Bangalore"

	fmt.Println(Name)
	fmt.Println(Age)
	fmt.Println(City)
	//--------------------------

	//Problem 2
	var ProductName string = "Laptop"
	var Price float64 = 999.99
	var quantity int = 10

	fmt.Println(ProductName)
	fmt.Println(Price)
	fmt.Println(quantity)
	//--------------------------

	//Problem 3

	isMarried := true
	idEmployed := false
	fmt.Println(isMarried)
	fmt.Println(idEmployed)
	//--------------------------

	//Problem 4

	const CompanyName = "ABC Pvt Ltd"
	fmt.Println(CompanyName)

	//--------------------------

	employeeName := "Shridhar"
	employeeId := 1001
	employeeSalary := 55000
	employeeDepartment := "IT"
	isActive := true

	fmt.Println("Employee Details:")
	fmt.Println("--------------------")
	fmt.Println("Name : ", employeeName)
	fmt.Println("ID : ", employeeId)
	fmt.Println("Salary : ", employeeSalary)
	fmt.Println("Department : ", employeeDepartment)
	fmt.Println("Active : ", isActive)
}
