package main

import (
	"fmt"
	signup "Book/internal/Users/Signup"
	signin "Book/internal/Users/Signin"
	admin "Book/internal/Admin"
	book "Book/internal/Users/Book"
)

func main(){
	num:=0
	fmt.Println("[1]Users  [2]Admin")
	fmt.Scanln(&num)
	switch num {
	case 1:
		var num2 int
		fmt.Println("[1]Signup  [2]Signin")
		fmt.Scanln(&num2)
		if num2==1{
			signup.SIgn_Up()
		}else if num2==2{
			signin.Sign_in()
		}
		book.Book()
	case 2:
		admin.Admin()
	default:
		fmt.Println("Notugri son kiritdiz!!!")		
	}
}