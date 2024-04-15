package Signup

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Signup struct {
	First_name string
	Last_name  string
	Email      string
	Password   string
	// Price      int
}

func SIGNUP(s_up Signup) {
	var sentence string
	var nm []string
	fmt.Printf("First_Name:\t")
	fmt.Scanln(&s_up.First_name)
	fmt.Printf("Last_Name:\t")
	fmt.Scanln(&s_up.Last_name)
	fmt.Printf("Email:\t")
	fmt.Scanln(&s_up.Email)
	fmt.Printf("Password:\t")
	fmt.Scanln(&s_up.Password)

	if len(s_up.Password) < 8 {
		fmt.Println("Parol kamida 8 ta belgidan kam bo'lishi mumkin emas!!!")
		fmt.Println("Parolni qayta kiriting")
		fmt.Printf("Password:\t")
		fmt.Scanln(&s_up.Password)
	}
	// fmt.Printf("Price:\t")
	// fmt.Scanln(&s_up.Price)

	nm = append(nm, s_up.First_name, s_up.Last_name, s_up.Email, s_up.Password)

	fileName := "/home/abduazim/Projects/Golang/book/users.txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Faylni ochishda xatolik:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, s_up.Email) {
			fmt.Println("Bu email allaqachon ro'yxatdan o'tgan:", line)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Xatolik:", err)
		return
	}

	filee, err := os.OpenFile("/home/abduazim/Projects/Golang/book/users.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Fayl ochilmadi:", err)
		return
	}
	defer filee.Close()

	sentence = strings.Join(nm, " ")
	yozuvchi := bufio.NewWriter(filee)
	fmt.Fprintln(yozuvchi, sentence)
	yozuvchi.Flush()

	fmt.Println("Siz muvaffaqiyatli kirdingiz🥳🥳🥳")
	fmt.Println("")
	product(s_up)
}
func product(s_up Signup) {
	var num string
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
	for _, work := range arr {
		natija := strings.Split(work, " ")
		if natija[0] == num {
			fmt.Println("Siz ",natija[1]," kitobni olishni xohlaysizmi?")
			fmt.Println("[1]HA\t[2]Yoq")
			fmt.Scanln(&cnt)
			if cnt==1{
				son,err:=strconv.Atoi(natija[2])
				if err!=nil{
					fmt.Println("Integerga olishga xatolik bor!!!!")
				}
				sentence:=natija[0]+natija[1]+strconv.Itoa(son-1)
				slc=append(slc, sentence)
				continue
			}else{
				product(s_up)
			}
			slc=append(slc, work)
		}
	}


}

func SIgn_Up() {
	var s_up Signup
	SIGNUP(s_up)
}
