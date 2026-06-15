package main

import "fmt"

func main() {

	student1 := Student{Name: "Shridhar", Age: 25}

	fmt.Println(student1.Name)
	fmt.Println(student1.Age)

	//-------------------

	product1 := Product{Name: "Laptop", Price: 999.99}

	fmt.Println(product1.Name)
	fmt.Println(product1.Price)

	//------------------- update the price

	product1.Price = 880
	fmt.Println(product1.Price)

	//-------------------

	displayStudent(student1)

	//-------------------

	Car1 := Car{Brand: "Hyundai", Model: "Grand i10 Nios Sportz", Year: 2026}
	fmt.Println("Ashuu's Car Name :", Car1.Model)
	fmt.Println("Ashuu's Car Brand :", Car1.Brand)
	fmt.Println("Ashuu's Car Manufactured year :", Car1.Year)

	//-------------------

	employee1 := Employee{ID: 101, Name: "Rahul", Department: "Physics", Salary: 40000}
	employee2 := Employee{ID: 102, Name: "Rajat", Department: "chemistry", Salary: 30000}
	employee3 := Employee{ID: 101, Name: "vaibhav", Department: "nano", Salary: 20000}

	displayEmployees(employee1)
	displayEmployees(employee2)
	displayEmployees(employee3)

	//----------------store employees in slice

	empSlice := []Employee{employee1, employee2, employee3}

	fmt.Println("Employee Details:-----------")

	for _, value := range empSlice {

		fmt.Println("ID:", value.ID, "Name :", value.Name, "Department :", value.Department, "Salary :", value.Salary)

	}

}

type Student struct {
	Name string
	Age  int
}

type Product struct {
	Name  string
	Price float64
}

func displayStudent(stu Student) {

	fmt.Println(stu.Name)
	fmt.Println(stu.Age)

}

type Car struct {
	Brand string
	Model string
	Year  int
}

//------------- assignmeant

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     int
}

func displayEmployees(emp Employee) {

	fmt.Println(emp.ID)
	fmt.Println(emp.Name)
	fmt.Println(emp.Department)
	fmt.Println(emp.Salary)

}
