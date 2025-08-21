package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Вычисление выражений в скобках и заменяем их на результат вычислений
func Eval(expression string) (int, error) {
	expression = strings.ReplaceAll(expression, " ", "")

	for strings.Contains(expression, "(") {

		start := strings.LastIndex(expression, "(")           // Ищем самую правую открывающую скобку
		end := strings.Index(expression[start:], ")") + start // Ищем первую закрывающую скобку идущую после найденной открывающей
		if start+1 == end {
			return 0, fmt.Errorf("пустые скобки в выражении")
		}
		val, err := Eval(expression[start+1 : end])
		if err != nil {
			return 0, err
		}
		expression = expression[:start] + strconv.Itoa(val) + expression[end+1:]
	}

	return calcFlat(expression)

}

// Проверка на унарный минус
func Uns(i int, ch rune, expression string) bool {
	return ((ch == '-' || ch == '+') && (i == 0 ||
		expression[i-1] == '-' ||
		expression[i-1] == '+' ||
		expression[i-1] == '*' ||
		expression[i-1] == '/'))
}

// Вычисляем выражение без скобок
func calcFlat(expression string) (int, error) {
	nums := []int{}
	ops := []rune{}
	num := ""
	for i, ch := range expression {
		if ch >= '0' && ch <= '9' || Uns(i, ch, expression) {
			num += string(ch)
		} else {
			val, err := strconv.Atoi(num)
			if err != nil {
				return 0, fmt.Errorf("Ошибка преобразования строки в число: %s", num)
			}
			nums = append(nums, val)
			num = ""
			ops = append(ops, ch)
		}
	}

	if num != "" {
		val, err := strconv.Atoi(num)
		if err != nil {
			return 0, fmt.Errorf("Ошибка преобразования строки в число: %s", num)
		}
		nums = append(nums, val)
	}
	//Cначала считаем всем пары чисел которые перемножаются или делятся
	for i := 0; i < len(ops); {
		if ops[i] == '*' || ops[i] == '/' {
			a, b := nums[i], nums[i+1]
			var res int

			if ops[i] == '*' {
				res = a * b
			} else if ops[i] == '/' && b == 0 {
				return 0, fmt.Errorf("Деление на ноль")
			} else {
				res = a / b
			}

			nums[i] = res
			nums = append(nums[:i+1], nums[i+2:]...)
			ops = append(ops[:i], ops[i+1:]...)
		} else {
			i++
		}
	}

	//Считаем сложение и вычитание
	res := nums[0]
	for i, op := range ops {
		if op == '+' {
			res += nums[i+1]
		} else {
			res -= nums[i+1]
		}
	}

	return res, nil
}

func validEnter(i int, expression string) bool {
	return (((expression[i] == '+' ||
		expression[i] == '-') &&
		(expression[i+1] == '+' ||
			expression[i+1] == '-')) ||
		((expression[i] == '*' ||
			expression[i] == '/') &&
			(expression[i+1] == '*' ||
				expression[i+1] == '/')))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Введите команду(узнать команды /help):")

		if valid := scanner.Scan(); !valid {
			fmt.Println("Ошибка ввода")
			return
		}

		cmd := scanner.Text()
		if cmd == "/help" {
			fmt.Println("/help - список всех команд")
			fmt.Println("/solve - решение задачи")
			fmt.Println("/exit - выход из программы")
		} else if cmd == "/solve" {
			fmt.Println("Введите выражение:")

			if valid := scanner.Scan(); !valid {
				fmt.Println("Ошибка ввода")
				return
			}

			expression := scanner.Text()
			for i := 0; i < len(expression)-1; i++ {
				if validEnter(i, expression) {
					fmt.Println("Ошибка ввода")
					return
				}
			}
			result, err := Eval(expression)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			} else {
				fmt.Println("Результат", result)
			}
		} else if cmd == "/exit" {
			break
		} else {
			fmt.Println("Вы ввели несуществующую команду, чтобы посмотреть список команд введите /help")
		}
	}

}
