package main

import "fmt"

const milesTokm=1.60934

func main(){
	var miles float64
	fmt.Print("Enter miles:")
	fmt.Scanf("%f",&miles)

	kilometers:=miles*milesTokm


	fmt.Println(kilometers)


	var name="rupa"

	fmt.Println("Hello my name is:",name)

	var n string
	fmt.Printf("enter n ....")
	fmt.Scanf("%s",&n)
	fmt.Println("Hello my name is",n)






}