package book

import (
	"bufio"
	"fmt"
	// "go/scanner"
	"log"
	"os"
	"strconv"
	"strings"
)

func ChooseBook(){
	var nm int
	file, err:=os.Open("/home/abduazim/Projects/Golang/book/users_book.txt")
	if err!=nil{
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
	defer file.Close()

	fmt.Println("Siz tanlagan kitoblar")
	scanner2:=bufio.NewScanner(file)
	for scanner2.Scan(){
		line:=scanner2.Text()
		fmt.Println(line)
	}

	fmt.Println("Yana kitob band qilishni hohlaysizmi?")
	fmt.Println("[1]HA\t[2]Yoq")
	fmt.Scanln(&nm)
	if nm==1{
		List_of_Books()
	}else if nm==2{
		Book()
	}

}

func List_of_Books(){
	var num string
	cn:=0
	file3, err := os.Open("/home/abduazim/Projects/Golang/book/kutubxona.txt")
	if err != nil {
		log.Fatalf("Faylni ochishda xatolik: %v", err)
	}
	defer file3.Close()

	scanner2 := bufio.NewScanner(file3)
	var arr []string
	for scanner2.Scan() {
		line := scanner2.Text()
		arr = append(arr, line)
		fmt.Println(line)
	}

	fmt.Println("")
	fmt.Println("Yuqoridagilardan birini tanlang👆🏻👆🏻👆🏻")
	fmt.Scanln(&num)

	cnt := 0
	var slc []string
	var book string
	for _, work := range arr {
		natija := strings.Split(work, " ")
		if natija[0] == num {
			fmt.Println("Siz ", natija[1], " kitobni olishni xohlaysizmi?")
			fmt.Println("[1]HA\t[2]Yoq")
			fmt.Scanln(&cnt)
			if cnt == 1 {
				son, err := strconv.Atoi(natija[2])
				if err != nil {
					fmt.Println("Integerga olishga xatolik bor!!!!")
				}
				book=natija[1]
				sentence := natija[0] + " " + natija[1] + " " + strconv.Itoa(son-1)
				slc = append(slc, sentence)
				continue
			} else {
				List_of_Books()
			}
		}
		slc = append(slc, work)
	}

	file, err:=os.Create("users_book")
	if err!=nil{
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
	defer file.Close()

	_, err=file.WriteString(book+"\n")
	if err!=nil{
		fmt.Println("Error: ", err)
		os.Exit(0)
	}

	file4, err := os.Create("/home/abduazim/Projects/Golang/book/kutubxona.txt")
	if err != nil {
		fmt.Println("Faylda ochishda xatolik bor!!")
	}
	defer file4.Close()

	for _, char := range slc {
		yozuvchi := bufio.NewWriter(file4)
		fmt.Fprintln(yozuvchi, char)
		yozuvchi.Flush()
	}
	fmt.Println("Tanlab olgan kitoblarizni ro'yxatini kurishni xoxlaysizmi?")
	fmt.Println("[1]HA\t[2]YOQ")
	fmt.Scanln(&cn)
	if cn==1{
		ChooseBook()
	}else if cn==2{
		Book()
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
	}else if son==3{
		os.Exit(0)
	}
}
