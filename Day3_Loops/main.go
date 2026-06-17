package main

import "fmt"

func main() {

	//break statement is used to exit from the loop when a certain condition is met.
	fmt.Println("break statement")
	for i := 1; i <= 10; i++ {

		if i == 5 {
			break
		}

		fmt.Println(i)
	}

	//Continue statement is used to skip the current iteration of the loop when a certain condition is met.
	fmt.Println("continue statement")
	for i := 1; i <= 5; i++ {

		if i == 3 {
			continue
		}

		fmt.Println(i)
	}

	//nested loop is a loop inside another loop. The inner loop is executed for each iteration of the outer loop.
	fmt.Println("nested loop")
	for i := 1; i <= 3; i++ {

		for j := 1; j <= 3; j++ {

			fmt.Println(i, j)

		}
	}

	// //infinite loop is a loop that runs indefinitely until it is explicitly stopped. It can be created using the for loop without any condition.
	// fmt.Println("infinite loop")
	// for {
	// 	fmt.Println("This is an infinite loop")
	// }

	//Practice problem print 1 to 10 using for loop
	fmt.Println("Practice problem print 1 to 10 using for loop")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//Practice problem print 10 to 1 using for loop
	fmt.Println("Practice problem print 10 to 1 using for loop")
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	//Practice problem print even numbers using for loop
	fmt.Println("Practice problem print even numbers using for loop")
	for i := 1; i <= 10; i++ {

		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	//Practice problem print sum using for loop
	fmt.Println("Practice problem find sum")
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println("The sum is:", sum)

	//Practice problem tables
	fmt.Println("Practice problem print multiplication tables")
	num := 7
	for i := 1; i <= 10; i++ {
		fmt.Println(num, "x", i, "=", num*i)
	}

	//Practice problem employee attendance
	fmt.Println("Practice problem employee attendance")
	for employees := 1; employees <= 10; employees++ {
		fmt.Println("Employee", employees, ": Present")
	}

	fmt.Println("Practice problem employee attendance absent if even")
	for employees := 1; employees <= 10; employees++ {
		if employees%2 == 0 {
			fmt.Println("Employee", employees, ": Absent")
		} else {
			fmt.Println("Employee", employees, ": Present")
		}
	}

	//calculate the sum of 1 to 100 using loop

	//sum of natural numbers formula n(n+1)/2
	fmt.Println("Practice problem calculate the sum of 1 to 100 using loop")
	newSum := 0
	for i := 1; i <= 100; i++ {
		newSum += i
	}
	fmt.Println("The sum is:", newSum)

	// alternative carl gauss method
	fmt.Println("Practice problem calculate the sum of 1 to 100 using formula")
	n := 100
	carlSum := (n * (n + 1)) / 2
	fmt.Println("The sum is:", carlSum)

}
