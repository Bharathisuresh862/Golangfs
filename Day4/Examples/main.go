package main
import(
	"fmt"
)

func f1(){
	fmt.Println("this is beginning of f1 func")

	fmt.Println("This is the end of f1 func")
}

func f2(){
	fmt.Println("this is beginning of f2 func")

	fmt.Println("This is the end of f2 func")
}

func f3(){
	fmt.Println("this is beginning of f3 func")

	fmt.Println("This is the end of f3 func")
}

func main(){
	fmt.Println("Start of main")
	f1()
	defer f2() //after finishing all otherfunc execution only defer will be executed
	f3()
	fmt.Println("end of main")
}