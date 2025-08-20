package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func eval(expression string) int {
	expression = strings.ReplaceAll(expression, " ", "")

	for strings.Contains(expression, "(") {

		start := strings.LastIndex(expression, "(")
		end := strings.Index(expression[start:], ")") + start
		val := eval(expression[start+1 : end])
		expression = expression[:start] + strconv.Itoa(val) + expression[end+1:]
	}

	return calcFlat(expression)

}

func calcFlat(expression string) int {
	nums := []int{}
	ops := []rune{}
	num := ""

	for i, ch := range expression {
		if ch >= '0' && ch <= '9' || (ch == '-' && (i == 0 || expression[i-1] == '-' || expression[i-1] == '+' || expression[i-1] == '*' || expression[i-1] == '/')) {
			num += string(ch)
		} else {
			val, _ := strconv.Atoi(num)
			nums = append(nums, val)
			num = ""
			ops = append(ops, ch)
		}
	}

	if num != "" {
		val, _ := strconv.Atoi(num)
		nums = append(nums, val)
	}

	for i := 0; i < len(ops); {
		if ops[i] == '*' || ops[i] == '/' {
			a, b := nums[i], nums[i+1]
			var res int
			if ops[i] == '*' {
				res = a * b
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

	res := nums[0]
	for i, op := range ops {
		if op == '+' {
			res += nums[i+1]
		} else {
			res -= nums[i+1]
		}
	}

	return res
}

func main() {

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Введите команду(узнать команды /help):")

		if valid := scanner.Scan(); !valid {
			fmt.Println("Ошибка ввода")
			return
		}

		expression := scanner.Text()

		result := eval(expression)

		fmt.Println("Результат", result)
	}

}
