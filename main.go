package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	students := make(map[string]Student)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Добро пожаловать в журнал!")

	for {
		fmt.Print("\nВведите команду (help - для вывода всех команд): ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "add":
			addStudent(students, reader)
		case "list":
			if len(students) > 0 {
				printAllStudents(students)
			} else {
				fmt.Println("В базе данных пока нет студентов.")
			}
		case "filter":
			fmt.Print("Введите максимальный средний балл: ")
			thresholdStr, _ := reader.ReadString('\n')
			thresholdStr = strings.TrimSpace(thresholdStr)

			threshold, err := strconv.ParseFloat(thresholdStr, 64)
			if err != nil {
				fmt.Println("Некорректный ввод среднего балла. Пожалуйста, введите число.")
				continue
			}

			filteredStudents := filterStudentsByAvg(students, threshold)
			if len(filteredStudents) > 0 {
				fmt.Printf("Студенты со средним баллом ниже %.2f:\n", threshold)
				for _, student := range filteredStudents {
					printStudentInfo(student)
				}
			} else {
				fmt.Printf("Нет студентов со средним баллом ниже %.2f.\n", threshold)
			}
		case "help":
			fmt.Println("Доступные команды:")
			fmt.Println(" add - Добавить нового студента.")
			fmt.Println(" list - Вывести информацию о всех студентах.")
			fmt.Println(" filter - Отфильтровать студентов по среднему баллу (ниже заданного порога).")
			fmt.Println(" help - Показать список доступных команд.")
			fmt.Println(" exit - Выйти из программы.")
		case "exit":
			fmt.Println("Выход из программы. До свидания!")
			return
		default:
			fmt.Println("Неизвестная команда. Введите 'help' для просмотра списка команд.")
		}
	}
}
