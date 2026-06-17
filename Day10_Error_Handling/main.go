package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println(validateAge(17))
	fmt.Println(validateAge(18))
	fmt.Println(validateAge(19))

	fmt.Println(validateSalary(9000))
	fmt.Println(validateSalary(10000))
	fmt.Println(validateSalary(11000))

	fmt.Println(divide(10, 2))
	fmt.Println(divide(10, 0))
	fmt.Println(divide(5, 14))

	fmt.Println(validateName(""))
	fmt.Println(validateName("shridhar"))

	fmt.Println(checkPositive(10))
	fmt.Println(checkPositive(11))

	fmt.Println("TASK RESULT---------------")

	emp := Employee{Name: "Shridhar", Age: 18, Salary: 11000}

	fmt.Println(validateEmployee(emp))
}

func validateAge(age int) error {
	if age < 18 {
		return errors.New("age must be above 18")
	} else {
		return nil
	}
}

func validateSalary(salary int) error {
	if salary < 10000 {
		return errors.New("salaty is less then expected")
	} else {
		return nil
	}
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Not divisble by zero")
	} else {
		return a / b, nil
	}
}

func validateName(name string) error {
	if name == "" {
		return errors.New("Invalid Name/Empty name")
	} else {
		return nil
	}
}

func checkPositive(number int) error {
	if number/2 == 0 {
		return nil
	} else {
		return errors.New("Negative number")
	}
}

type Employee struct {
	Name   string
	Age    int
	Salary int
}

func validateAgeNew(e Employee) error {
	if e.Age >= 18 {
		return nil
	}
	return errors.New("Minor")
}

func validateSalaryNew(e Employee) error {
	if e.Salary >= 10000 {
		return nil
	}
	return errors.New("Salary is less then expected")
}

func validateNameNew(e Employee) error {
	if e.Name == "" {
		return errors.New("Name can not be empty")
	}
	return nil
}

func validateEmployee(emp Employee) error {
	if err := validateAgeNew(emp); err != nil {
		return err
	}
	if err := validateSalaryNew(emp); err != nil {
		return err
	}
	if err := validateNameNew(emp); err != nil {
		return err
	}
	return nil
}
