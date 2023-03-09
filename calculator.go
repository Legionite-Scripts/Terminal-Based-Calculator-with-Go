package main
import ("fmt")

func main(){
	modeSelect := []string{"Golang Calculator\nSelect the mode you want : ","1.Addition","2.Subtraction","3.Multiplication","4.Division","5.Find Percentage\n"}
	for _, val := range modeSelect{
		fmt.Printf("%v\n", val)
	 }
	fmt.Print("Enter Choice : ")
	var enterChoice string
	fmt.Scanln(&enterChoice)

	switch enterChoice{
	case "1" : addFunction()
	case "2" : subtractFunction()
	case "3" : multiplyFunction()
	case "4" : divisionFunction()
	case "5" : percentFunction()
	default : fmt.Println("Invalid Input")
	}
}

func addFunction(){
	fmt.Print("Enter First number : ")
	var firstNumber int
	fmt.Scanln(&firstNumber)

	fmt.Print("Enter Second number : ")
	var secondNumber int
	fmt.Scanln(&secondNumber)

	fmt.Print(firstNumber + secondNumber)
}

func subtractFunction(){
	fmt.Print("Enter First number : ")
	var firstNumber int
	fmt.Scanln(&firstNumber)

	fmt.Print("Enter Second number : ")
	var secondNumber int
	fmt.Scanln(&secondNumber)

	fmt.Print(firstNumber - secondNumber)
}

func multiplyFunction(){
	fmt.Print("Enter First number : ")
	var firstNumber int
	fmt.Scanln(&firstNumber)

	fmt.Print("Enter Second number : ")
	var secondNumber int
	fmt.Scanln(&secondNumber)

	fmt.Print(firstNumber * secondNumber)
}

func divisionFunction(){
	fmt.Print("Enter First number : ")
	var firstNumber float64
	fmt.Scanln(&firstNumber)

	fmt.Print("Enter Second number : ")
	var secondNumber float64
	fmt.Scanln(&secondNumber)

	fmt.Print(firstNumber / secondNumber)
}

func percentFunction(){
	fmt.Print("Enter First number : ")
	var firstNumber int
	fmt.Scanln(&firstNumber)

	fmt.Print("Enter Second number : ")
	var secondNumber int
	fmt.Scanln(&secondNumber)

	fmt.Print(firstNumber % secondNumber)
}
