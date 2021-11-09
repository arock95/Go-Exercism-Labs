package wordy

import (
	"strconv"
	"strings"
)

var (
	operators = [4]string{"plus", "minus", "mb", "db"}
)

func Answer(question string) (int, bool) {
	// parse string and test validity
	words, valid := validity(question)

	if !valid {
		return 0, false
	}

	// take action
	result := findValue(words)

	// report result
	return result, true

}

func validity(question string) ([]string, bool) {
	question = strings.ReplaceAll(question, "multiplied by", "mb")
	question = strings.ReplaceAll(question, "divided by", "db")
	words := strings.Split(question[:len(question)-1], " ")

	if words[0] != "What" && words[1] != "is" {
		return nil, false
	}
	if len(words) == 2 || len(words) == 4 {
		return nil, false
	}

	// test number operator number [operator number...]
	equation := words[2:]

	for k, v := range equation {
		if k%2 == 0 {
			if _, err := strconv.Atoi(v); err != nil {
				return nil, false
			}
		} else if !validOperator(v) {
			return nil, false
		}
	}

	return equation, true

}

func validOperator(op string) bool {
	for _, v := range operators {
		if v == op {
			return true
		}
	}
	return false
}

func findValue(equation []string) int {
	if len(equation) == 1 {
		ret, _ := strconv.Atoi(equation[0])
		return ret
	}

	total := 0
	for k, v := range equation {
		if k%2 == 1 { //hit an operator
			op1, _ := strconv.Atoi(equation[k-1])
			op2, _ := strconv.Atoi(equation[k+1])

			switch v {
			case "plus":
				if k == 1 {
					total = op1 + op2
				} else {
					total = total + op2
				}
			case "minus":
				if k == 1 {
					total = op1 - op2
				} else {
					total = total - op2
				}
			case "mb":
				if k == 1 {
					total = op1 * op2
				} else {
					total = total * op2
				}
			case "db":
				if k == 1 {
					total = op1 / op2
				} else {
					total = total / op2
				}
			}
		}
	}
	return total
}
