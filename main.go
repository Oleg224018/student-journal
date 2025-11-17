package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("=== Журнал группы Тощев Олег ИС-323 💾 ===")
	students := make(map[string][]float64)

	for {
		fmt.Println("\nМеню:")
		fmt.Println("1. Добавить студента и оценки")
		fmt.Println("2. Посчитать средний балл студента")
		fmt.Println("3. Показать всех студентов с оценками")
		fmt.Println("4. Фильтр: студенты с средним баллом ниже 4")
		fmt.Println("5. Выход")

		var choice string
		fmt.Print("Ваш выбор: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			addStudent(students)
		case "2":
			calculateAverage(students)
		case "3":
			showAllStudents(students)
		case "4":
			filterByAverage(students)
		case "5":
			fmt.Println("До свидания!")
			os.Exit(0)
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}

func addStudent(students map[string][]float64) {
	var name string
	fmt.Print("Введите ФИО студента: ")
	fmt.Scanln(&name)

	fmt.Print("Введите оценки через пробел (например: 5 4 3): ")
	var input string
	fmt.Scanln(&input)

	grades, err := parseGrades(input)
	if err != nil {
		fmt.Printf("Ошибка при вводе оценок: %s\n", err)
		return
	}

	students[name] = grades
	fmt.Println("Студент добавлен!")
}

func calculateAverage(students map[string][]float64) {
	var name string
	fmt.Print("Введите ФИО студента: ")
	fmt.Scanln(&name)

	grades, exists := students[name]
	if !exists {
		fmt.Println("Студент не найден.")
		return
	}

	avg := average(grades)
	fmt.Printf("Средний балл студента %s: %.2f\n", name, avg)
}

func showAllStudents(students map[string][]float64) {
	if len(students) == 0 {
		fmt.Println("Нет студентов.")
		return
	}

	fmt.Println("\n--- Список студентов ---")
	for name, grades := range students {
		avg := average(grades)
		fmt.Printf("ФИО: %s | Оценки: %v | Средний балл: %.2f\n", name, grades, avg)
	}
}

func filterByAverage(students map[string][]float64) {
	fmt.Println("\n--- Студенты с средним баллом ниже 4 ---")
	found := false
	for name, grades := range students {
		avg := average(grades)
		if avg < 4.0 {
			fmt.Printf("ФИО: %s | Средний балл: %.2f\n", name, avg)
			found = true
		}
	}
	if !found {
		fmt.Println("Нет студентов с средним баллом ниже 4.")
	}
}

func parseGrades(input string) ([]float64, error) {
	fields := splitString(input)
	var grades []float64
	for _, s := range fields {
		grade, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}
	return grades, nil
}

func splitString(s string) []string {
	var parts []string
	current := ""
	for _, r := range s {
		if r == ' ' {
			if current != "" {
				parts = append(parts, current)
				current = ""
			}
		} else {
			current += string(r)
		}
	}
	if current != "" {
		parts = append(parts, current)
	}
	return parts
}
