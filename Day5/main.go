package main

import "fmt"

func main() {

	numbers := [5]int{10, 20, 30, 40, 50}

	fmt.Println(numbers)

	//----------------

	fmt.Println("Length of array:", len(numbers))

	//----------------

	cities := []string{"Bangalore", "Mysore", "Hubli"}

	for index, value := range cities {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	//----------------

	marks := []int{80, 90, 70}
	marks = append(marks, 95, 100)

	fmt.Println("Marks:", marks)

	//----------------

	newMarks := []int{80, 90, 70, 95}
	var sum int
	for _, value := range newMarks {
		sum += value
	}

	fmt.Println("Total Marks:", sum)

	//----------------

	studentMarks := []int{85, 90, 78, 92, 88}

	total := totalMarks(studentMarks)
	average := averageMarks(studentMarks)
	highest := highestMarks(studentMarks)
	lowest := lowestMarks(studentMarks)

	fmt.Println("Total Marks:", total)
	fmt.Println("Average Marks:", average)
	fmt.Println("Highest Marks:", highest)
	fmt.Println("Lowest Marks:", lowest)

}

func totalMarks(marks []int) int {
	var sum int
	for _, value := range marks {
		sum += value
	}
	return sum
}

func averageMarks(marks []int) float64 {
	total := totalMarks(marks)
	count := len(marks)
	return float64(total) / float64(count)
}

func highestMarks(marks []int) int {
	highest := marks[0]
	for _, value := range marks {
		if value > highest {
			highest = value
		}
	}
	return highest
}

func lowestMarks(marks []int) int {
	lowest := marks[0]
	for _, value := range marks {
		if value < lowest {
			lowest = value
		}
	}
	return lowest
}
