package main

import "fmt"

// variadic slices

type cost struct {
	day   int
	value float64
}

func getCostByDay(costs []cost) []float64 {
	costByDay := []float64{}
	for i := range costs {
		cost := costs[i]
		for cost.day >= len(costByDay) {
			costByDay = append(costByDay, 0.0)
		}
		costByDay[cost.day] += cost.value
	}
	return costByDay
}

func sums(nums ...float64) float64 {
	total := 0.0
	for i := range nums {
		total += nums[i]
	}
	return total
}

func test(nums ...float64) {
	total := sums(nums...)
	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month : %.2f \n", total)
	fmt.Println("......END REPORT.......")
}

func getMessagesCost(messages []string) []float64 {
	costs := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		cost := float64(len(message)) * 0.01
		costs[i] = cost
	}
	return costs
}

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}
func indexOfFirstBadWord(msg []string, badWrods []string) int {
	for i, word := range msg {
		for _, badWord := range badWrods {
			if word == badWord {
				return i
			}
		}
	}
	return -1
}
func test1(msg []string, badWords []string) {
	i := indexOfFirstBadWord(msg, badWords)
	for _, badWord := range badWords {
		fmt.Println("-", badWord)
	}
	fmt.Printf("Index : %v\n", i)

}

func main() {
	cost := getMessagesCost([]string{"Hello", "World", "!"})
	fmt.Println(cost)
	test(1.0, 2, 3.0, 4)
	test(5, 6.0, 7, 8.0)
	fmt.Println(createMatrix(3, 3))
	fmt.Println(createMatrix(2, 3))
	badWords := []string{"crap", "shoot", "dang", "frick"}
	msg := []string{"hey", "there", "pradeep"}
	test1(msg, badWords)
	msg = []string{"oh", "my", "frick"}
	test1(msg, badWords)
}
