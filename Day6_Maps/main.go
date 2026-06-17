package main

import "fmt"

func main() {

	//-------------------
	employee := map[string]string{
		"name": "Shridhar",
		"city": "Bangalore",
	}

	fmt.Println(employee["name"])
	fmt.Println(employee["city"])
	fmt.Println("Employee Details:", employee)

	//-------------------

	employee["department"] = "IT"

	fmt.Println("Updated Employee Details:", employee)
	fmt.Println(employee["name"])
	fmt.Println(employee["city"])
	fmt.Println(employee["department"])

	//-------------------

	employee["city"] = "Mysore"

	fmt.Println("Updated Employee Details for city:", employee)
	fmt.Println(employee["name"])
	fmt.Println(employee["city"])
	fmt.Println(employee["department"])

	//-------------------

	delete(employee, "department")
	fmt.Println("Updated Employee Details for after delete department:", employee)

	//-------------------

	for key, value := range employee {
		fmt.Println("key: ", key, "value: ", value)
	}

	//-------------------

	employees := map[int]string{
		101: "Shridhar",
		102: "Ravi",
		103: "Ajay",
		104: "Vijay",
	}

	for _, value := range employees {
		fmt.Println("Employees name:", value)
	}

	//-------------------

	if _, exists := employees[102]; exists {
		fmt.Println("Employee Found")
	} else {
		fmt.Println("Employee Not Found")
	}

	//-------------------

	employees[105] = "Kiran"
	fmt.Println("Updated Employees Kiran Added:", employees)

	//-------------------

	delete(employees, 103)
	fmt.Println("Updated Employees after delete Ajay:", employees)

	fmt.Println("Employee Directory:")
	fmt.Println("-------------------")
	for key, value := range employees {
		fmt.Println("ID:", key, "Name:", value)
	}
}
