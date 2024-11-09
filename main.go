package main

import (
	// "bufio"
	"fmt"
	// "os"
	"time"
	// "strconv"
	// "strings"
)

// jswTOken:="123456789" --->not allowed
var token string = "123456789"

const LoginToken string = "fagjiegnov"

func main() {
	// var uname string = "sai"
	// fmt.Println(uname)
	// fmt.Printf("variable is of type %T \n",uname)

	// var defaultVar string;
	// fmt.Println(defaultVar)
	// fmt.Printf("variable is of type %T \n",defaultVar)
	//implicit type
	// var web="www.google.com"
	// fmt.Println(web)

	//no var style
	// noVar := 10
	// fmt.Println(noVar)

	// fmt.Println(LoginToken)

	// ** take input from user
	// var userName string

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("enter your name")

	// //comma ok syntax or error ok syntax
	// userName, _ = reader.ReadString('\n')
	// fmt.Println("welcome", userName)

	//converting datatype
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter rating here:")
	// rating, _ := reader.ReadString('\n')
	// fmt.Println("rating",rating);

	// numRating,err:=strconv.ParseFloat(strings.TrimSpace(rating),64)
	// if(err!=nil){
	// 	fmt.Println(err)
	// }else{
	// 	fmt.Println(numRating)
	// }

	// *** Handling time
	fmt.Println("Time handling in GO")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println("Formated presentTIme is",presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate:= time.Date(2020, time.August, 12,23,23,1,1,time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
