package book

import (
	"bufio"
	"fmt"
	"os"
)

func ChooseBook(){

}

func List_of_Books(){
	file, err:=os.Open("/home/abduazim/Projects/Golang/book/kutubxona.txt")
	if err!=nil{
		fmt.Println("Fayldi ochishda xatolik bor: ", err)
		os.Exit(0)
	}
	defer file.Close()

	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		line:=scanner.Text()
		fmt.Println(line)
	}


}

func Book() {
	son := 0
	fmt.Println("[1]Tanlab olingan kitoblar\n[2]Kitoblar ro'yxati\n[3]Chiqish")
	fmt.Scanln(&son)
	if son==1{
		ChooseBook()
	}else if son==2{
		List_of_Books()
	}
}
