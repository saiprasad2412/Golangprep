package main

import (
	// "bufio"
	"fmt"
	// "sort"
	// "os"
	// "time"
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
	// fmt.Println("Time handling in GO")
	// presentTime := time.Now()
	// fmt.Println(presentTime)
	// fmt.Println("Formated presentTIme is",presentTime.Format("01-02-2006 15:04:05 Monday"))

	// createdDate:= time.Date(2020, time.August, 12,23,23,1,1,time.UTC)
	// fmt.Println(createdDate.Format("01-02-2006 Monday"))

	// ***Pointers in GO
	// fmt.Println("Pointers in GO")

	// var ptr *int
	// fmt.Println(ptr)

	// myNum:=24;
	// var ptr = &myNum;
	// fmt.Println(*ptr);

	// Array in GO --->we dont use to much of array in go language
	// fmt.Println("Array in GO")

	// var fruits[4]string
	// fruits[0]="A"
	// fruits[1]="B"
	// fruits[2]="C"

	// fmt.Println("Array -->", fruits)

	// var vegList=[3]string{"P","B","M"}
	// fmt.Println("Veglist & Length -->",vegList, len(vegList))

	// slices in GO --> uses more in go
	// fmt.Println("Slices in Go")
	// var fruitList= []string{"A","B"}
	// fruitList= append(fruitList, "Mango","Banana")
	// fmt.Println("ffff-->",fruitList)

	// highscores:= make([]int , 4)
	// highscores[0]=1
	// highscores[1]=1
	// highscores[2]=1
	// highscores[3]=1
	// // highscores[4]=1 --->wrong out of range
	// highscores=append(highscores, 2,3)

	// fmt.Println(highscores)

	// sort.Ints(highscores)
	// fmt.Println(highscores)

	// *** How to remove values from slice based on index
	// var course= []string{"react","JS","C#","C++"}
	// fmt.Println(course)
	// var deleteINdex int =2;
	// course=append(course[:deleteINdex],course[deleteINdex+1:]...)
	// fmt.Println(course)

	// *** Maps in go
	// fmt.Println("maps in go")

	// languages:=make(map[string]string)

	// languages["JS"]="Javascript"
	// languages["RB"]="Ruby"
	// languages["PY"]="Python"

	// fmt.Println(languages["JS"])
	// delete(languages, "RB")

	// fmt.Println(languages)

	// for key,value:=range languages{
	// 	fmt.Printf("for key %v -> value is %v \n",key, value)
	// }

	// *** Structs in go
	// fmt.Println("Structs in GO")

	// bot1 := User{"Bot1", "Bot1@g.com", true, 1}
	// fmt.Println(bot1)
	// fmt.Printf("Details of bot1 are : %+v",bot1)

	// 	result:=addfn(10,5)
	// 	fmt.Println(result)

	// 	proResult:=proaddfn(1,2,3,4,5,6,7)
	// 	fmt.Println(proResult)

	// }

	// func addfn(val1 int , val2 int)int{
	// 	return val1+val2
	// }
	// func proaddfn(values ...int)int{
	// 	total:=0
	// 	for _,val:=range values{
	// 		total+=val
	// 	}
	// 	return total
	// }
	// bot1 := User{"bot1", "bot1@g.com", true, 1}
	// bot1.GetUserInfo()

	// *** Defer in go 
	defer fmt.Println("four")
	defer fmt.Println("three")
	fmt.Println("one")
	fmt.Println("Two")

}

// *** methods
// func (u User) GetUserInfo() {
// 	fmt.Println("user info", u.Name)
// }

// type User struct {
// 	Name   string
// 	Email  string
// 	Status bool
// 	Age    int
// }
