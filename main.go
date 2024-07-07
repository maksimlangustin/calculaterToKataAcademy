package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var num11 string
	var num12 string

	var number1 int
	var number2 int
	var oper string
	var plus int
	var res1 string
	var minu int
	var res2 string
	var umno int
	var res3 string
	var dely int
	var res4 string

	fmt.Scan(&num11)
	fmt.Scan(&oper)
	fmt.Scan(&num12)

	if (strings.ContainsAny(num11, "0123456789") && strings.ContainsAny(num12, "IVXCL")) ||
		(strings.ContainsAny(num12, "0123456789") && strings.ContainsAny(num11, "IVXCL")) {
		panic("разные типы цифр")
	}

	if strings.ContainsAny(num11, "0123456789") && strings.ContainsAny(num12, "0123456789") {
		num11 = strings.TrimSpace(num11)
		number1, _ := strconv.Atoi(num11)
		num12 = strings.TrimSpace(num12)
		number2, _ := strconv.Atoi(num12)

		if number1 > 10 || number2 > 10 {
			panic("калькулятор не может работать с числами больше 10")
		}

		switch oper {
		case "+":
			fmt.Print("=", number1+number2)
		case "-":
			fmt.Print("=", number1-number2)
		case "*":
			fmt.Print("=", number1*number2)
		case "/":
			if number2 == 0 {
				panic("на ноль делить нельзя")
			}
			fmt.Print("=", number1/number2)
		default:
			panic("неверный знак")
		}
	} else {

		number1 = romanToInt(num11)
		number2 = romanToInt(num12)

		if number1 > 10 || number2 > 10 {
			panic("калькулятор не может работать с числами больше 10")
		}

		plus = number1 + number2
		res1 = integerToRoman(plus)
		if number2 > number1 {
			panic("римские цифры не могут быть <= 0")
		}
		minu = number1 - number2
		res2 = integerToRoman(minu)
		umno = number1 * number2
		res3 = integerToRoman(umno)
		dely = number1 / number2
		res4 = integerToRoman(dely)
		switch oper {
		case "+":
			fmt.Print("=", res1)
		case "-":
			fmt.Print("=", res2)
		case "*":
			fmt.Print("=", res3)
		case "/":
			fmt.Print("=", res4)
		default:
			panic("неверный знак")
		}

	}
}

func integerToRoman(number int) string {
	maxRomanNumber := 1000
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}
func romanToInt(s string) int {
	rMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := 0
	for k := range s {
		if k < len(s)-1 && rMap[s[k:k+1]] < rMap[s[k+1:k+2]] {
			result -= rMap[s[k:k+1]]
		} else {
			result += rMap[s[k:k+1]]
		}
	}
	return result
}
