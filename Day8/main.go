package main

import "fmt"

func main() {

	//------add an object in struct
	employee := Employee{Name: "Shridhar", Age: 23}

	//----- call the created method
	employee.Display()

	//----- add an object in rectangle struct

	rectangle := Rectangle{Length: 10, Width: 10}

	//----- call the created method for area calculations
	fmt.Println(rectangle.Area())

	//---------------

	employee.UpdateAge()

	fmt.Println("Age must show updated age ie 35:------")
	employee.Display()

	//------------------------problem1

	student1 := Student{Name: "Shridhar", Age: 31}
	student1.Display()

	//----------------------Problem2

	product1 := Product{Name: "Laptop", Price: 999}
	product1.Display()

	//-------------------Problem 4

	circle1 := Circle{Radius: 10}
	fmt.Println(circle1.Diameter())

	//---------Problem 5

	employee.Greet()

	//Assignement
	account := BankAccount{
		AccountHolder: "Shridhar",
		Balance:       10000,
	}

	account.DisplayBalance()

	account.Deposit(5000)
	fmt.Println(account)

	account.Withdraw(3000)

	account.DisplayBalance()

}

//--------create struct
type Employee struct {
	Name string
	Age  int
}

//--------Create Method to display name and age
func (e Employee) Display() {
	fmt.Println("Name :", e.Name)
	fmt.Println("Age :", e.Age)
}

//--------------------------

type Rectangle struct {
	Length int
	Width  int
}

//--------Create Method for area calculation (problem 3)
func (r Rectangle) Area() int {
	return r.Length * r.Width
}

//------------------------------ without pointer you cant modify the object feilds, with pointer it is possible

func (e *Employee) UpdateAge() {
	e.Age = 35
}

//Problem 1---------

type Student struct {
	Name string
	Age  int
}

func (s Student) Display() {
	fmt.Println("Name :", s.Name)
	fmt.Println("Age :", s.Age)
}

//Problem 2-----------

type Product struct {
	Name  string
	Price int
}

func (p Product) Display() {
	fmt.Println("Name :", p.Name)
	fmt.Println("Price :", p.Price)
}

//-----------problem 4

type Circle struct {
	Radius int
}

func (c Circle) Diameter() int {
	return c.Radius * 2
}

//problem 5

func (e Employee) Greet() {
	fmt.Println("hello", e.Name)
}

// assignment

type BankAccount struct {
	AccountHolder string
	Balance       float64
}

//method1
func (b BankAccount) DisplayBalance() {
	fmt.Println("Account Holder", b.AccountHolder)
	fmt.Println("Balance :", b.Balance)
}

//method2

func (b *BankAccount) Deposit(amount float64) float64 {
	if amount < 0 {
		fmt.Println("Invalid amount")
	}
	if amount > 0 {
		b.Balance += amount
		fmt.Println("Amount is Credited")
	}
	return b.Balance
}

//method3

func (b *BankAccount) Withdraw(amount float64) float64 {
	if amount > b.Balance {
		fmt.Println("Insufficient Fund")
		return b.Balance
	} else {
		b.Balance -= amount
		fmt.Println("Amount is debited")
		return b.Balance
	}
}
