package main

import(
	"fmt"
	"runtime"
	//"os"
	"os"
	"time"
)

//const pi float = math.Pi

func main() {
	fmt.Printf("hello, world\n")
	fmt.Print(addNumbers(3,2))
	fmt.Println()
	systemTest()
	timeTest()
	ifTest(3,2)
	switchTest("b")
	forTest(10)
	factorial(5, 1)
	continueTest(6)

}



// break default func interface select case defer go map struct chan else goto package switch const fallthrough if range type continue for import return var

func addNumbers(a int, b int)(sum int){
	sum = a+b
	return sum
}

func systemTest(){
	var goos string = runtime.GOOS
	fmt.Println("the operating system is ", goos)
	path := os.Getenv("PATH")
	fmt.Println("The path is ", path)
}

func timeTest(){
	fmt.Println(time.Now())
}

//start to learn chapter 5
// if-else structure

func ifTest(a int, b int){
	if a > b{
		fmt.Println("a is greater than b")
	}else if a < b{
		fmt.Println("b is greater than a")
	}else if a == b{
		fmt.Println("a is equal to b")
	}else{
		fmt.Println("error")
	}
}

func switchTest(text string){
	switch text{
	case "a":
		fmt.Print("a")
	case "b":
		fmt.Print("b")
	default:
		fmt.Print("error")
	}
	fmt.Println()

}

func forTest(times int){
	for i := 0; i < times; i++{
		fmt.Println(i)
	}

	for{
		fmt.Println("233333333")
		break
	}
}

func factorial(i int, s int){
	if i <= 1{
		fmt.Println(s)
	} else{
		s *= i
		i--
		factorial(i,s)
	}
}

func continueTest(j int){
	for i := 0; i < 10; i++{
		if i ==j{
			continue
		}
		fmt.Print(i, " ")
	}
}