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

	// var number int
	// fmt.Print("Please input number ")
	// fmt.Scanf("%d", &number)

	// if number%2 == 0 {
	// 	fmt.Println(number, "is even number")
	// } else {
	// 	fmt.Println(number, "is odd number")
	// }

	// if number == 1 {
	// 	fmt.Println("Open an account")
	// } else if number == 2 {
	// 	fmt.Println("With draw money")
	// } else {
	// 	fmt.Println("Please input just 1 or 2.")
	// }

	// ---- Switch ... Case ----

	// var number int
	// fmt.Print("Please input number . ")
	// fmt.Scanf("%d", &number)

	// switch number {
	// case 1:
	// 	fmt.Println("Open an account")

	// case 2:
	// 	fmt.Println("With draw money")
	// default:
	// 	fmt.Println("Please input just 1 or 2")
	// }

	fmt.Println("\n")

	// ---- Array ----
	// has 2 type
	// 1. set number of members
	var numberA [4]int = [4]int{10, 20, 30, 40}
	fmt.Println(numberA, "\n")
	fmt.Println(numberA[0])
	fmt.Println(numberA[1])
	fmt.Println(numberA[2])
	fmt.Println(numberA[3])
	fmt.Println(len(numberA), "\n")
	// short hand
	nameA := [3]string{"Somchay", "Sompong", "Sommai"}
	fmt.Println(nameA, "\n")
	fmt.Println(nameA[0])
	fmt.Println(nameA[1])
	fmt.Println(nameA[2])
	fmt.Println(len(nameA), "\n")

	// nameB := [3]int{} //result [ 0, 0, 0]
	nameB := [3]int{100} // result [100, 0, 0]
	fmt.Println("Member of nameB are", nameB)

	var numbers [4]int
	// fmt.Println("Members in numbers array is ", numbers, "\n") //result = [0 ,0 ,0 ,0]
	// create members
	numbers[0] = 23
	numbers[1] = 99
	numbers[2] = 69
	numbers[3] = 55
	fmt.Println("Members in numbers array are ", numbers, "\n")

	// count length array
	numberA2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Length of numberA2 is ", len(numberA2))

	petsName := [2]string{"Chokdee", "Meaw"}
	fmt.Println("Length of petsName are ", len(petsName))

	// 2. Not set number of members
	nameC := [...]string{"Hi", "I'", "m", "New."}
	sizeOfNameC := len(nameC)
	fmt.Println(nameC)
	fmt.Println("Size of nameC are ", sizeOfNameC, '\n')

	// Slices
	// Slices is like array but can chang size of member in array (dynamic size)

	nameS := []string{"new", "chakrit"}
	// append member in Slice
	nameS = append(nameS, "Sompong")
	nameS = append(nameS, "Nong")

	fmt.Println(nameS)

	// query member in Slice
	fmt.Println(nameS[0])
	fmt.Println(nameS[1])
	fmt.Println(nameS[2])
	fmt.Println(nameS[3], "\n")

	// query by set length
	// syntax : slice[low:high]
	// slice[first Index, ก่อน last index*]
	fmt.Println(nameS[0:3]) //result [new , chakrit, sompong]
	fmt.Println(nameS[1:2]) //result [chakrit]

	fmt.Println(nameS[1:]) // index 1 to last member //[chakrit, sompng ,nong]
	fmt.Println(nameS[:1]) // index 0 to 1 //[new]
	fmt.Println("\n")

	// change value in slice
	nameS[0] = "Oum"
	nameS[1] = "Praeploy"

	fmt.Println(nameS)    //same value
	fmt.Println(nameS[:]) //same value

	// ---- Map ----
	//Maps is like array. It save data is even type(key, value) but use key is index. If we know key property .We can know that key value
	// syntax : var <name> map [key_type]value_type

	// var country map[string]string

	country := map[string]string{}
	country["TH"] = "Thailand"
	country["EN"] = "England"

	fmt.Println("\n")

	fmt.Println(country)       // map[EN:England TH:Thailand]
	fmt.Println(country["TH"]) //Thailand
	fmt.Println(country["EN"]) //England

	fmt.Println("\n")

	coin := map[string]string{}
	coin["BTC"] = "Bitcoin"
	coin["ETC"] = "Ether"
	fmt.Println(coin["BTC"])
	fmt.Println(coin["ETC"])

	fmt.Println("\n")

	population := map[string]int{}
	population["Thailand"] = 72000000
	population["China"] = 2000000000

	fmt.Println("Thaland has population = ", population["Thailand"])
	fmt.Println("China has population = ", population["China"])

	// one line
	coin2 := map[string]string{"ETC": "Ether", "BTC": "Bitcoin"}
	fmt.Println(coin2["ETC"])
	fmt.Println(coin2["BTC"])

	fmt.Println("\n")

	// ----- Create Map -----

	// countrys := [3]string{"Thailand, China", "Japan"}

	// countrys := map[string]string{"TH": "Thailand", "JP": "Japan", "CH": "China"} //or
	countrys := map[string]string{}
	countrys["TH"] = "Thailand"
	countrys["JP"] = "Japan"
	countrys["CH"] = "China"

	fmt.Println(countrys["TH"])
	fmt.Println(countrys["JP"])
	fmt.Println(countrys["CH"])
	// fmt.Println(countrys[""]) //

	fmt.Println(countrys) //map[CH:China JP:Japan TH:Thailand]

	// check map
	//--value* find value in key -- Result = value
	//--check* check key in map -- Result = bool

	value, check := countrys["JP"]
	value2, isKey := countrys["TH"]

	if check {
		fmt.Println("Does coruntrys variable in key JP has value :", check, "\n , and value is :", value)

	} else if isKey {
		fmt.Println("Does coruntrys variable in key JP has value :", isKey, "\n , and value is :", value2)

	} else {
		fmt.Println("Does coruntrys variable in key JP has value :", check)
	}

	// ----- For ... Loop -----
	// syntax : for <create default value>; <condition>; <change value in variable>{order if condition is True}
	// 1. Set default value.
	// 2. if condition is true ,do repeat order in statement.
	// 3. if condition is false ,get out for loop.

	for count := 1; count <= 3; count++ {
		fmt.Println("Hello Go programming.", count)
	}

	fmt.Println("\n")

	for count := 10; count >= 2; count-- {
		fmt.Println("Hello Go programming.", count)
	}

	// ----- break and continue -----

	for count := 1; count <= 10; count++ {
		// fmt.Println(count) //1-5
		// if count == 5 {
		// 	break
		// }    OR

		// if count == 5 {
		// 	break
		// }
		// fmt.Println(count) //1-4

		if count == 4 {
			continue
		}
		fmt.Println(count)

	}

	fmt.Println("End Loop")

	fmt.Println("\n")

	// ----- Range -----

	numberR := []int{10, 20, 30, 40, 50}

	for index := 0; index < len(numberR); index++ {
		// fmt.Println(numberR[index])
		fmt.Println("Index : ", index, ", Value : ", numberR[index])
	}

	fmt.Println("\n")

	for index, value := range numberR {
		fmt.Println("Index : ", index, ", Value : ", value)
	}
	for _, value := range numberR { // if don't want index = "_" => it mean ignore.
		fmt.Println("Value : ", value)
	}

	fmt.Println("\n")

	for index, _ := range numberR {
		fmt.Println("Index : ", index)
	}

	language := map[string]string{"TH": "Thai", "EN": "Englist"}

	for key, value := range language {
		fmt.Println("Key: ", key, ", Value: ", value)
	}

	showMessage("New")
	showMessage("Oum")

	total(50, 40)

	delivery := getDelivery()
	fmt.Println("Delivery : ", delivery)

	totalPrice := getTotalCart(9, 69)

	fmt.Println("Totalprice : ", totalPrice)
}

// ----- Function -----

func showMessage(name string) {
	fmt.Println("Hello ", name)
}

func total(num1, num2 int) {
	fmt.Println("Total : ", num1+num2)
}

func getDelivery() int {
	return 50
}

func getTotalCart(num1, num2 int) int {
	total := num1 + num2
	return total
}
