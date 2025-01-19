package errors

import (
	"fmt"
	"log"
)

func Check(e error){
	if e != nil{
		panic(e)
	}
}
func CheckLog(e error){
	if e != nil{
		log.Fatal(e)
	}
}
func CheckCustom(e error){
	if e != nil{
		fmt.Println(e)
	}
}