package main

import (
	"bufio"   //для работы с пользовательским вводом/выводом,
	"errors"  //для определения пользовательских ошибок
	"fmt"     //для форматированного вывода
	"os"      //используется для чтения пользовательского ввода из стандартного ввода (os.Stdin).
	"strconv" //строки в числа
	"strings"
)

func RomanToInteger(romanNumber string) (int, error) {
	romanNumbers := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	if romanNumbers[romanNumber] != 0 {
		return romanNumbers[romanNumber], nil
	} else {
		err := errors.New("Некорректное римское число от 1 (I) до 10 (X)")
		return 0, err
	}
}

func IntegerToRoman(number int) string {
	var result string
	romanNumbers := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	arabicNumbers := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	for i := 0; number > 0; i++ {
		for arabicNumbers[i] <= number {
			result += romanNumbers[i]
			number -= arabicNumbers[i]
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin) //считывает введенные данные
	inputData, _ := reader.ReadString('\n')
	inputData = strings.Trim(inputData, "\n") //разбивает строку на массив??
	var inputStrArray []string = strings.Split(inputData, " ")
	if len(inputStrArray) != 3 {
		// Проверка, что введены три значения через пробел
		fmt.Println("Некорректные входные данные")
		return
	}
	operators := map[string]int{"+": 1, "-": 1, "*": 1, "/": 1} //действия возможные
	_, check := operators[inputStrArray[1]]
	if !check {
		// Проверка, что "среднее" значение является оператором
		fmt.Println("Некорректные входные данные")
		return
	}
	numberOne, errNumberOne := strconv.Atoi(inputStrArray[0]) // преобразование в целые числа (numberOne и numberTwo) с помощью strconv.Atoi
	numberTwo, errNumberTwo := strconv.Atoi(inputStrArray[2])
	if errNumberOne == nil && errNumberTwo == nil && numberOne >= 1 && numberOne <= 10 && numberTwo >= 1 && numberTwo <= 10 {
		// Проверка, что введенные операнды являются арабскими числами от 1 до 10 включительно.

		arabicResult, _ := Calc(numberOne, numberTwo, inputStrArray[1])
		fmt.Println(arabicResult)
		return
	}
	romanNumberOne, errRomanNumberOne := RomanToInteger(inputStrArray[0])
	romanNumberTwo, errRomanNumberTwo := RomanToInteger(inputStrArray[2])
	if errRomanNumberOne == nil && errRomanNumberTwo == nil && (inputStrArray[1] != "-" || romanNumberOne > romanNumberTwo) {
		// Проверка, что введенные операнды являются римскими числами от 1 до 10 включительно,
		// и что для оператора "-" первое число больше второго.
		romanResult, _ := Calc(romanNumberOne, romanNumberTwo, inputStrArray[1])
		fmt.Println(IntegerToRoman(romanResult))
		return
	}
	// Все остальные вариации считаем некорректными
	fmt.Println("Некорректные входные данные")
}

func Calc(numOne, numTwo int, operator string) (int, error) {
	var result int
	switch operator {
	case "+":
		result = numOne + numTwo
	case "-":
		result = numOne - numTwo
	case "*":
		result = numOne * numTwo
	case "/":
		if numTwo != 0 {
			result = numOne / numTwo
		} else {
			err := errors.New("Делить на 0 нельзя")
			return 0, err
		}
	}
	return result, nil
}
