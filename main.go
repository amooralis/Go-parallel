// В данном коде мы определяем структуру Task, которая содержит информацию 
// о продолжительности выполнения задачи и о том, какие задачи должны быть выполнены до нее.
// Затем мы определяем функцию findOptimalOrder, которая перебирает все возможные порядки 
// выполнения задач и выбирает тот, который займет наименьшее количество времени. 
// Для генерации всех возможных порядков выполнения задач мы используем функцию generatePermutations, 
// которая рекурсивно генерирует все перестановки чисел от 0 до n-1. 
// Для проверки того, что порядок выполнения задач является допустимым, мы используем функцию isValidOrder,
// которая проверяет, что все задачи, от которых зависит текущая задача, уже были выполнены. 
// Для расчета времени выполнения задач мы используем функцию calculateTime, 
// которая для каждой задачи выбирает наименьшее время, когда все ее зависимости уже были выполнены, 
// и добавляет к этому времени продолжительность выполнения самой задачи.

// Пример использования данного кода показывает, как мы можем определить набор задач 
// и количество исполнителей, и как мы можем найти оптимальный порядок выполнения задач. 
// В данном примере мы определяем 5 задач и 2 исполнителя, (строки 33-42)
// и находим оптимальный порядок выполнения задач, который займет наименьшее количество времени.



package main

import (
	"fmt"
	"math"
)

type Task struct {
	duration int
	depends  []int
}

func main() {
	tasks := []Task{
		{duration: 2, depends: []int{}},
		{duration: 3, depends: []int{0}},
		{duration: 4, depends: []int{0}},
		{duration: 1, depends: []int{}},
		{duration: 5, depends: []int{1, 2}},
		{duration: 6, depends: []int{3}},
	}

	numWorkers := 3

	optimalOrder := findOptimalOrder(tasks, numWorkers)

	fmt.Println("Optimal order:", optimalOrder)
}

func findOptimalOrder(tasks []Task, numWorkers int) []int {
	numTasks := len(tasks)
	bestOrder := make([]int, numTasks)
	bestTime := math.MaxInt32

	permutations := generatePermutations(numTasks)

	for _, order := range permutations {
		if isValidOrder(order, tasks) {
			time := calculateTime(order, tasks, numWorkers)
			if time < bestTime {
				bestTime = time
				copy(bestOrder, order)
			}
		}
	}

	return bestOrder
}

func generatePermutations(n int) [][]int {
	if n == 1 {
		return [][]int{{0}}
	}

	perms := [][]int{}
	for _, perm := range generatePermutations(n - 1) {
		for i := 0; i < n; i++ {
			newPerm := make([]int, n)
			copy(newPerm, perm[:i])
			newPerm[i] = n - 1
			copy(newPerm[i+1:], perm[i:])
			perms = append(perms, newPerm)
		}
	}

	return perms
}

func isValidOrder(order []int, tasks []Task) bool {
	numTasks := len(tasks)
	completed := make([]bool, numTasks)

	for _, taskIndex := range order {
		task := tasks[taskIndex]
		for _, depIndex := range task.depends {
			if !completed[depIndex] {
				return false
			}
		}
		completed[taskIndex] = true
	}

	return true
}

func calculateTime(order []int, tasks []Task, numWorkers int) int {
	numTasks := len(tasks)
	times := make([]int, numTasks)

	for _, taskIndex := range order {
		task := tasks[taskIndex]
		minTime := math.MaxInt32
		for i := 0; i < numWorkers; i++ {
			prevTaskIndex := taskIndex - i - 1
			if prevTaskIndex < 0 {
				prevTaskIndex = -1
			}
			prevTime := 0
			if prevTaskIndex >= 0 {
				prevTime = times[prevTaskIndex]
			}
			minTime = min(minTime, prevTime)
		}
		times[taskIndex] = minTime + task.duration
	}

	return times[numTasks-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
