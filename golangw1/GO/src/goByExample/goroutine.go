package main

import (
	"fmt"
"time"
)

func main(){

	done :=make(chan <-bool,2)
	//done :=make(<-chan bool,2)
	go func(){
	fmt.Println("series 1")
	time.Sleep(2* time.Second)

	fmt.Println("Done go routine")
	//done chan bool
	done <- true
	done<-true

	done<-true
}()
	fmt.Println("series 2")
	//result:=<-done
	//time.Sleep(2* time.Second)

//	<-done
	fmt.Println("done whole app")





}