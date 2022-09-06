// Напишите программу, которая считывает ввод с stdin,
// создаёт структуру student и записывает указатель на структуру
// в хранилище map[studentName] *Student.

package main

import (
	"28/pkg/storage"
	"28/pkg/student"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	var buffer = bufio.NewReader(os.Stdin)
	studentStorage := storage.NewStorage()
	fmt.Println("Введите имя возраст курс студента")

	for {
		fmt.Print(">>> ")
		line, err := buffer.ReadString('\n')

		if err == io.EOF {
			println("Студенты из хранилища")
			studentStorage.PrintStudents()
			break
		}

		lineFieldsArr := strings.Fields(line)

		if len(lineFieldsArr) != 3 {
			fmt.Println("Некорректный ввод")
			continue
		}

		studentName := lineFieldsArr[0]
		studentAge, errAge := strconv.Atoi(lineFieldsArr[1])
		studentGrade, errGrade := strconv.Atoi(lineFieldsArr[2])

		if errAge != nil || errGrade != nil {
			fmt.Println("Ошибка обработки числовых значений")
			continue
		}

		studentInfo := student.NewStudent()
		studentInfo.SetName(studentName)
		studentInfo.SetAge(studentAge)
		studentInfo.SetGrade(studentGrade)

		_, err = studentStorage.Get(studentInfo.GetName())
		if err != nil {
			studentStorage.Put(studentInfo)
		} else {
			fmt.Println("Студент с таким именем существует")
		}
	}
}
