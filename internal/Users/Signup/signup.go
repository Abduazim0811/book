package Signup

import (
	"bufio"
	"fmt"
	"os"
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

	nm = append(nm, s_up.First_name, s_up.Last_name, s_up.Email, s_up.Password)

	fileName := "/home/abduazim/Projects/Golang/book/users.txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Faylni ochishda xatolik:", err)
		os.Exit(0)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, s_up.Email) {
			fmt.Println("Bu email allaqachon ro'yxatdan o'tgan:", line)
			os.Exit(0)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Xatolik:", err)
		os.Exit(0)
	}

	filee, err := os.OpenFile("/home/abduazim/Projects/Golang/book/users.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Fayl ochilmadi:", err)
		os.Exit(0)
	}
	defer filee.Close()

	sentence = strings.Join(nm, " ")
	yozuvchi := bufio.NewWriter(filee)
	fmt.Fprintln(yozuvchi, sentence)
	yozuvchi.Flush()

	fmt.Println("Siz muvaffaqiyatli kirdingiz🥳🥳🥳")
	fmt.Println("")
}


func SIgn_Up() {
	var s_up Signup
	SIGNUP(s_up)
}
