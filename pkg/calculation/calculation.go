package calculation

import (
	"slices"
	"strconv"
)

var (
	numbers  = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	symbols  = []rune{'+', '-', '*', '/'}
	brackets = []rune{'(', ')'}
)

func checkString(expression string) error {
	//наличие цифр вообще
	var numbers_marker bool = false
	for _, el := range expression {
		if slices.Contains(numbers, el) {
			numbers_marker = true
			break
		}
	}
	if !numbers_marker {
		return ErrNoNumbers
	}

	//проверка "цифра или символ или скобка"
	for _, el := range expression {
		if !slices.Contains(numbers, el) && !slices.Contains(symbols, el) && !slices.Contains(brackets, el) {
			return ErrUnexpectedSymbol
		}
	}

	//проверка "два знака подряд и знак стоит последний"
	for i, el := range expression {
		if slices.Contains(symbols, el) && i != len(expression)-1 {
			if slices.Contains(symbols, rune(expression[i+1])) {
				return ErrTwoOperationsInARow
			}
		} else if slices.Contains(symbols, el) && i == len(expression)-1 {
			return ErrOperationAtTheEnd
		}
	}

	//проверка "количество открытых и закрытых скобок"
	var bracket_counter int
	for _, el := range expression {
		switch el {
		case '(':
			bracket_counter += 1
		case ')':
			bracket_counter -= 1
		}
	}
	if bracket_counter != 0 {
		return ErrBracketsProblem
	}

	return nil
}

func makeSlice(expression string) []string { //разбиваем выражение на слайс
	var exp_slice []string
	var str string
	var start_marker bool = true
	var multiply_division_marker bool = false
	var minus_counter int
	var brackets_counter int

	for _, el := range expression {
		if el == '(' {
			if start_marker {
				str += "+"
				start_marker = false
			}
			brackets_counter += 1
			str += string(el)
			continue
		} else if el == ')' {
			brackets_counter -= 1
			str += string(el)
			continue
		}
		if el == '-' && !start_marker {
			if multiply_division_marker && minus_counter == 0 {
				str += string(el)
				continue
			}
			if brackets_counter == 0 {
				exp_slice = append(exp_slice, str)
				str = string(el)
				minus_counter = 0
				multiply_division_marker = false
				continue
			}
		} else if el == '+' && !start_marker {
			if brackets_counter == 0 {
				multiply_division_marker = false
				exp_slice = append(exp_slice, str)
				str = string(el)
				continue
			}
		} else if el == '*' {
			if brackets_counter == 0 {
				exp_slice = append(exp_slice, str)
				str = string(el)
				if multiply_division_marker {
					continue
				}
				multiply_division_marker = true
				continue
			}
		} else if el == '/' {
			if brackets_counter == 0 {
				exp_slice = append(exp_slice, str)
				str = string(el)
				if multiply_division_marker {
					continue
				}
				multiply_division_marker = true
				continue
			}
		}
		if start_marker && el != '-' {
			str += "+"
		}
		str += string(el)
		start_marker = false
		multiply_division_marker = false
	}
	exp_slice = append(exp_slice, str)
	start_marker = true

	//fmt.Println(exp_slice)

	return exp_slice
}

func bracketsOperations(exp_slice []string) ([]string, error) { //работа со скобками
	for i, el := range exp_slice {
		if el[1] == '(' {
			el_without_brackets := string(el[2 : len(el)-1])
			//fmt.Println("Brackets off:", el_without_brackets)
			result, err := Calc(el_without_brackets)
			if err != nil {
				return nil, err
			}
			if el[0] == '+' {
				exp_slice[i] = strconv.FormatFloat(result, 'f', -1, 64)
			} else {
				exp_slice[i] = string(el[0]) + strconv.FormatFloat(result, 'f', -1, 64)
			}
		}
	}
	//fmt.Println("Final:", exp_slice)
	return exp_slice, nil
}

func multiplyAndDivision(exp_slice []string) ([]string, error) { //умножение и деление
	var tres float64
	i := 0

	for i < len(exp_slice) {
		el := exp_slice[i]
		if el[0] == '*' {
			el1_float, _ := strconv.ParseFloat(el[1:], 64)
			el2_float, _ := strconv.ParseFloat(exp_slice[i-1], 64)
			tres = el1_float * el2_float
			exp_slice[i-1] = strconv.FormatFloat(tres, 'f', -1, 64)
			if i == len(exp_slice) {
				exp_slice = exp_slice[:i]
			} else {
				exp_slice = append(exp_slice[:i], exp_slice[i+1:]...)
			}
		} else if el[0] == '/' {
			el1_float, _ := strconv.ParseFloat(el[1:], 64)
			if el1_float == 0 {
				return nil, ErrDivisionByZero
			}
			el2_float, _ := strconv.ParseFloat(exp_slice[i-1], 64)
			tres = el2_float / el1_float
			exp_slice[i-1] = strconv.FormatFloat(tres, 'f', -1, 64)
			if i == len(exp_slice) {
				exp_slice = exp_slice[:i]
			} else {
				exp_slice = append(exp_slice[:i], exp_slice[i+1:]...)
			}
		} else {
			i++
		}
	}
	return exp_slice, nil
}

func additionAndSubtraction(exp_slice []string) float64 { //складываем все числа между собой
	var result float64

	for _, el := range exp_slice {
		t, _ := strconv.ParseFloat(el, 64)
		result += t
	}

	return result
}

func Calc(expression string) (float64, error) {

	check_result := checkString(expression)
	if check_result != nil {
		return 0, check_result
	}

	exp_slice := makeSlice(expression)

	exp_slice, err := bracketsOperations(exp_slice)
	if err != nil {
		return 0, err
	}

	exp_slice, err = multiplyAndDivision(exp_slice)
	if err != nil {
		return 0, err
	}

	result := additionAndSubtraction(exp_slice)

	return result, nil
}
