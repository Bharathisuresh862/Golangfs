package main
import "fmt"
type Student struct{
	Name string
	Rgno int
	Dept string 
}
func main(){
	//forPythonStyle()
	st:=Student{Name: "Student1",Rgno:12,Dept:"CS"}
	fmt.Println("Name:",st.Name,"\nRegistration no:",st.Rgno,"\nDepartment:",st.Dept)
}

func forCondiDemo(){
	var num int 
	var num2 int
	fmt.Scanln(&num)
	fmt.Scanln(&num2)


	if num>num2{
		fmt.Println("greater is",num)
	}else{
		fmt.Println("greater is",num2 )
	}

}

func forThreeVarDemo(){
    sum := 0
    for i:=0;i<5;i++{
	   sum+=i
    }
    fmt.Println(sum)
}

func forCondi2Demo(){
	n :=1
	for n < 5{
		n *= 2
	}
	fmt.Println(n)
}
func forPythonStyle(){
	strings :=[]string{"hello","world","golang","NIE"}
	for _, s:= range strings{
		fmt.Println(s)
	}
}
