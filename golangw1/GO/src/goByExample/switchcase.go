package main
import "fmt"
import "time"


func main(){
i :=2

fmt.Print("write ",i,"as")

switch i{
case 1:
	fmt.Println("one")
case 2:
	fmt.Println("two")

case 3:
	fmt.Println("three")
}

switch time.Now().Weekday(){
case time.Saturday,time.Sunday:
	fmt.Println("it is the weekend")
default:
	fmt.Println("ti is a weekday")
}


t:=time.Now()


switch{
case t.Hour()<12:
	fmt.Println("it is before noon")
default:
fmt.Println("it is after noon")

}

whatAmI := func(i interface{}){
	switch t :=i.(type){
	case bool:
		fmt.Println("i am a bool")
	case int:
		fmt.Println("i am an int")
	default:
		fmt.Printf("dont know type..",t)
	}
}

whatAmI(true)
whatAmI(1)
whatAmI("hey")
}