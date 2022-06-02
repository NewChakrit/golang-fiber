package main

import "fmt"

func main() {
	fmt.Println("Hello P'Kong")
	fmt.Println("How are you?")

	//manual type declaration

	var name string = "New"
	var age int = 25
	var score float32 = 95.8
	var isPass bool = true
	fmt.Println("My name is ", name)
	fmt.Println("I'm", age, "years old")
	fmt.Println("My score are ", score, "point")
	fmt.Println("Pass exam : ", isPass)

	//type inference

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
	//constant

	const name4 string = "Chanew"
	fmt.Println(name4)

	fmt.Println("\n")
	//check type
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
}

// golang is static type . Variable must has value.
// has 2 types
// - manual type declaration
// - type inference

// ---- Manual type declaration ----

// var <name> <type>

// ---- Type inference ----

// name:= 'New'
