package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func addStudent(students map[string]Student, reader *bufio.Reader) {
	fmt.Print("Введите фамилию и имя студента: ")
	fio, _ := reader.ReadString('\n')
	fio = strings.TrimSpace(fio)

	if _, exists := students[fio]; exists {
		fmt.Println("Студент с таким ФИО уже существует.")
		return
	}

	var grades []int
	fmt.Println("Введите оценки студента через пробел. Нажмите Enter еще раз для завершения.")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			if len(grades) == 0 {
				fmt.Println("Оценки не были введены. Введите оценки или нажмите Enter еще раз для пропуска.")
				continue
			}
			break
		}

		gradesStrSlice := strings.Fields(input)
		for _, gradeStr := range gradesStrSlice {
			grade, err := strconv.Atoi(gradeStr)
			if err != nil {
				fmt.Printf("Некорректный ввод оценки '%s'. Пожалуйста, вводите только числа.\n", gradeStr)
				continue
			}
			if grade < 1 || grade > 5 {
				fmt.Printf("Оценка %d вне допустимого диапазона (1-5)\n", grade)
				continue
			}
			grades = append(grades, grade)
		}
	}

	student := Student{
		FIO:    fio,
		Grades: grades,
	}
	student.calculateAverage()
	students[fio] = student
	fmt.Println("Студент", fio, "успешно добавлен!")
}

func filterStudentsByAvg(students map[string]Student, threshold float64) []Student {
	var filteredStudents []Student
	for _, student := range students {
		if student.Avg < threshold {
			filteredStudents = append(filteredStudents, student)
		}
	}
	return filteredStudents
}

func printStudentInfo(student Student) {
	fmt.Printf(" ФИ: %s, Оценки: %v, Средний балл: %.2f\n", student.FIO, student.Grades, student.Avg)
}

func printAllStudents(students map[string]Student) {
	fmt.Println("Список всех студентов:")
	for _, student := range students {
		printStudentInfo(student)
	}
}
