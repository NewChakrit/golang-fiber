package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("How are you?")

	// ----- Manual type declaration -----

	var name string = "New"
	var age int = 25
	var score float32 = 95.8
	var isPass bool = true
	fmt.Println("My name is ", name)
	fmt.Println("I'm", age, "years old")
	fmt.Println("My score are ", score, "point")
	fmt.Println("Pass exam : ", isPass)

	// ---- Type inference ----

	fmt.Println("\n")

	name2 := "New"
	age2 := 25
	score2 := 95.8
	isPass2 := true
	fmt.Println("My name is ", name2)
	fmt.Println("I'm", age2, "years old")
	fmt.Println("My score are ", score2, "point")
	fmt.Println("Pass exam : ", isPass2)

	fmt.Println("\n")

	name3 := "Chakrit"
	name3 = "Tnasinpaibool"
	fmt.Println(name3)

	fmt.Println("\n")
	// ---- Constant ----

	const name4 string = "Chanew"
	fmt.Println(name4)

	fmt.Println("\n")
	// ---- Check type ----
	name5 := "New"
	age5 := 25
	score5 := 95.8
	isPass5 := true
	fmt.Printf("My name is %v\n", name5)
	fmt.Printf("Age : %v\n", age5)

	fmt.Printf("Type name : %T\n", name5)
	fmt.Printf("Type age : %T\n", age5)
	fmt.Printf("Type score : %T\n", score5)
	fmt.Printf("Type isPass : %T\n", isPass5)

	fmt.Println("\n")

	// Math operation
	// var num1 int = 10
	// var num2 int = 8
	// or
	// var num1 = 10
	// var num2 = 8
	// or
	// var num1, num2 = 10, 8
	// or
	num1, num2 := 10, 8

	fmt.Println("Positive Result = ", num1+num2)
	fmt.Println("Minus Result = ", num1-num2)
	fmt.Println("Multiply = ", num1*num2)
	fmt.Println("Quotient = ", num1/num2)
	fmt.Println("Fraction = ", num1%num2)

	fmt.Println("\n")

	// ---- Compare operation ----
	fmt.Println("Equal or not : ", num1 == num2)
	fmt.Println("Not Equal or not : ", num1 != num2)
	fmt.Println(num1, ">", num2, "=", num1 > num2)
	fmt.Println(num1, "<", num2, "=", num1 < num2)
	fmt.Println(num1, ">=", num2, "=", num1 >= num2)
	fmt.Println(num1, "<=", num2, "=", num1 <= num2)

	// ---- Get values from keyboard with Scanf ----

	// Scanf scans text read from standard input, storing successive space-separated values into successive arguments as determined by the format. It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why. Newlines in the input must match newlines in the format. The one exception: the verb %c always scans the next rune in the input, even if it is a space (or tab etc.) or newline.

	// Syntax => fmt.Scanf(string_format, address_list)
	// func Scanf(format string, a ...any) (n int, err error)

	// string => %s
	// integer => %d
	// float => %f

	// var nameS string
	// fmt.Print("Please input student name. ")
	// fmt.Scanf("%s", &nameS)

	// fmt.Println("Hello ", nameS)

	// var scoreS int
	// fmt.Print("Please input student score. ")
	// fmt.Scanf("%d", &scoreS)

	// var scoreS float32
	// fmt.Print("Please input student score. ")
	// fmt.Scanf("%f", &scoreS)

	// fmt.Println(nameS, "has scores + mentality scores is ", scoreS+10, "point.")

	// ---- Condition ----

	// ---- if else ----

	// var scoreNew int

	// fmt.Print("Please input studen score ")
	// fmt.Scanf("%d", &scoreNew)
	// fmt.Println("New has score ", scoreNew, "point")

	// if scoreNew >= 50 {
	// 	fmt.Println("New is Pass exam.")
	// } else {
	// 	fmt.Println("New isn't Pass exam")
	// }

	// check number Odd or Even

	var number int
	fmt.Print("Please input number ")
	fmt.Scanf("%d", &number)

	// if number%2 == 0 {
	// 	fmt.Println(number, "is even number")
	// } else {
	// 	fmt.Println(number, "is odd number")
	// }

	if number == 1 {
		fmt.Println("Open an account")
	} else if number == 2 {
		fmt.Println("With draw money")
	} else {
		fmt.Println("Please input just 1 or 2.")
	}

}
