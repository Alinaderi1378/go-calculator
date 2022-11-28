package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	run := true

	fmt.Printf("Hi there is a simple command Line calculator for IAUT \n this program will support this oprators \n addition \n subtract \n multiplication \n multiplication \n divide \n sin x \n cos x \n tan x \n log x \n pow x \n sqrt x \n---------------------------------------------------------------- \n")

	// reads user input
	reader := bufio.NewReader(os.Stdin)
	for run {
		fmt.Printf("please enter num1 :")
		first, _ := reader.ReadString('\n')

		var trimmed = trim(first, "\r\n")
		if shouldQuit(trimmed) {
			fmt.Printf("Quitting...")
			os.Exit(0)
		}
		num1, _ := strconv.ParseFloat(trimmed, 64)

		fmt.Printf("Enter your operand: ")
		op, _ := reader.ReadString('\n')

		var optrim = trim(op, "\r\n")
		if shouldQuit(optrim) {
			fmt.Printf("Quitting...")
			os.Exit(0)
		}
		var vo = validOp(optrim)
		for !vo {
			fmt.Printf("invalid op , please enter a valid operand :")
			op, _ := reader.ReadString('\n')
			optrim = trim(op, "\r\n")
			if shouldQuit(optrim) {
				fmt.Printf("Quitting...")
				os.Exit(0)
			}
			vo = validOp(optrim)
		}

		if optrim == "sin" || optrim == "cos" || optrim == "tan" || optrim == "log" || optrim == "sqrt" {
			fmt.Printf("answer: " + callCalculation_once(optrim, num1) + "\n--------------------------------------\n")

		} else {

			fmt.Printf("plese enter num2 : ")
			second, _ := reader.ReadString('\n')

			// حذف خطوط جدید با تریم
			trimmed = trim(second, "\r\n")
			if shouldQuit(trimmed) {
				fmt.Printf("exit")
				os.Exit(0)
			}
			num2, _ := strconv.ParseFloat(trimmed, 64)

			fmt.Printf("answer: " + callCalculation_twice(optrim, num1, num2) + "\n--------------------------------------\n")
		}

	}
}

// فراخوانی توابع ماشین حساب
func callCalculation_twice(op string, a, b float64) string {
	if op == "+" {
		return convertFloat64ToString(addition(a, b))
	} else if op == "-" {
		return convertFloat64ToString(subtract(a, b))
	} else if op == "*" || op == "x" || op == "X" {
		return convertFloat64ToString(multiply(a, b))
	} else if op == "/" {
		return convertFloat64ToString(divide(a, b))
	} else if op == "pow" {
		return convertFloat64ToString(float64(pow(a, b)))
	}
	return "Couldn't calculate correctly."
}
func callCalculation_once(op string, a float64) string {
	if op == "sin" {
		return convertFloat64ToString(float64(sin(a)))
	} else if op == "cos" {
		return convertFloat64ToString(float64(cos(a)))
	} else if op == "tan" {
		return convertFloat64ToString(float64(tan(a)))
	} else if op == "log" {
		return convertFloat64ToString(float64(logarithm(a)))
	} else if op == "sqrt" {
		return convertFloat64ToString(float64(sqrt(a)))
	}
	return "Couldn't calculate correctly."
}

// عملگر های شناخته شده
func validOp(op string) bool {

	validOps := []string{"+", "-", "*", "x", "X", "/", "sin", "cos", "tan", "log", "sqrt", "pow"}
	for _, validOp := range validOps {
		if validOp == op {
			return true
		}
	}
	return false
}

// تابع خروج
func shouldQuit(trimmed string) bool {
	if trimmed == "q" {
		return true
	}
	return false
}

func convertFloat64ToString(float float64) string {
	return strconv.FormatFloat(float, 'f', -1, 64)
}
func trim(value, cutset string) string {
	return strings.Trim(value, cutset)
}

// جمع
func addition(first, second float64) float64 {
	return first + second
}

// منها
func subtract(first, second float64) float64 {
	return first - second
}

// ضرب
func multiply(first, second float64) float64 {
	return first * second
}

// تقسیم
func divide(first, second float64) float64 {
	return first / second
}

// سینوس
func sin(first float64) float64 {
	return math.Sin(first)
}

// کوسینوس
func cos(first float64) float64 {
	return math.Cos(first)
}

// تانژانت
func tan(first float64) float64 {
	return math.Tan(first)
}

// لگاریتم
func logarithm(first float64) float64 {
	return math.Log10(first)
}

// جذر
func sqrt(first float64) float64 {
	return math.Sqrt(first)
}

// توان
func pow(first, second float64) float64 {
	return math.Pow(first, second)
}
