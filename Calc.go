package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Input expression:")            // Show input text
	var scanner = bufio.NewScanner(os.Stdin)    // Initialization of new scanner for input line
	scanner.Scan()                              // Start scanning
	ex := scanner.Text()                        // Get scan result as Text
	ex = strings.ReplaceAll(ex, " ", "")        // Delete All Spaces from string
	ex = strings.ToTitle(ex)                    // Convert string to Upper Case
	reMath := regexp.MustCompile("[*]|/|-|[+]") // Regular expression for Math operands + - / *
	arg := reMath.Split(ex, -1)                 // Splitting a string into arguments by a Math operand
	// fmt.Println(arg)				// Unlock to view arguments

	if len(arg) < 2 { // Only 2 arguments must be
		fmt.Println("Error:\n1. Less than 2 arguments given\n2. Operand not given")
		os.Exit(1)
	}

	if len(arg) > 2 { // Only 2 arguments must be
		fmt.Println("Error:\n1. More than 2 arguments given\n2. More than 1 operand given")
		os.Exit(1)
	}

	var arabNums bool                                    // init helpers
	var argument1 int                                    // init argument 1 for calc function
	var argument2 int                                    // init argument 2 for calc function
	var operand string = string(ex[len(string(arg[0]))]) // Get an operand from input string for calc function

	//  Checking whether the arguments correspond to Arabic or Roman numerals
	matched1A, _ := regexp.MatchString(`[1-9]`, arg[0])
	matched1R, _ := regexp.MatchString(`^X$|^(IX|IV|V?I{0,3})$`, arg[0])
	matched2A, _ := regexp.MatchString(`[1-9]`, arg[1])
	matched2R, _ := regexp.MatchString(`^X$|^(IX|IV|V?I{0,3})$`, arg[1])

	// One or both arguments do not correspond conditions.
	if (!matched1A && !matched1R) || (!matched2R && !matched2A) {
		fmt.Println("Error:\nDude! You must use Arab from 1 to 10 or Roman digits from I to X only!.")
		os.Exit(1)
	}
	// Different number systems are used by arguments 1 and 2.
	if (matched1A && matched2R) || (matched1R && matched2A) {
		fmt.Println("Error:\nDifferent number systems are used at the same time.")
		os.Exit(1)
	}
	// Both arguments correspond Arabic system
	if matched1A && matched2A {
		arabNums = true
		argument1, _ = strconv.Atoi(arg[0])
		argument2, _ = strconv.Atoi(arg[1])
	}
	// Both arguments correspond Roman system
	if matched1R && matched2R {
		arabNums = false
		argument1 = RomanToInt(arg[0]) // convert argument 1 from Roman to Arabic
		argument2 = RomanToInt(arg[1]) // convert argument 2 from Roman to Arabic
	}
	// Check arguments for conditions corresponding
	if argument1 > 10 || argument1 < 1 || argument2 > 10 || argument2 < 1 {
		fmt.Println("Error:\nOne or both arguments are greater than 10 or lesser than 1.")
		os.Exit(1)
	}

	result := calc(argument1, argument2, operand) // Calculating... (in Arabic system)

	// Check result for negative. There are no negative numbers in the Roman system.
	if !arabNums && result <= 0 {
		fmt.Println("There are no negative numbers in the Roman system.")
		os.Exit(1)
	}
	// Show result
	if !arabNums {
		fmt.Println(intToRoman(result)) // Convert to Roman if arguments are Roman
	} else {
		fmt.Println(result) // Show as is. (Arabic)
	}
}

func calc(arg1 int, arg2 int, operand string) int { // Calc Function
	switch operand {
	case "+":
		return arg1 + arg2
	case "-":
		return arg1 - arg2
	case "*":
		return arg1 * arg2
	case "/":
		return arg1 / arg2
	default:
		fmt.Println("An impossible error has arisen. The world has gone mad!")
		os.Exit(1)
		return 0
	}
}

func RomanToInt(Str string) int { // Roman to Arabic integer Convertor for digits from 1 to 10
	var DigitList = []struct {
		IntValue   int
		RomanDigit string
	}{
		{10, "X"},
		{9, "IX"},
		{8, "VIII"},
		{7, "VII"},
		{6, "VI"},
		{5, "V"},
		{4, "IV"},
		{3, "III"},
		{2, "II"},
		{1, "I"},
	}
	for _, DigitList := range DigitList {
		if Str == DigitList.RomanDigit {
			return DigitList.IntValue
		}
	}
	return 0
}

func intToRoman(number int) string { // Arabic to Roman integer Convertor for digits from 1 to 100
	var DigitList = []struct {
		IntValue   int
		RomanDigit string
	}{
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
	for _, DigitList := range DigitList {
		for number >= DigitList.IntValue {
			roman.WriteString(DigitList.RomanDigit)
			number -= DigitList.IntValue
		}
	}
	return roman.String()
}
