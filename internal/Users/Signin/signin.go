package Signin

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type SignIn struct {
	Email    string
	Password string
	Price    int
}

func SIGNIN(s_in *SignIn) {
	// var num string
	file, err := os.Open("/home/abduazim/Projects/Golang/book/users.txt")
	if err != nil {
		fmt.Println("Faylni ochishda xatolik:", err)
		return
	}
	defer file.Close()

	fmt.Printf("Email:\t")
	fmt.Scanln(&s_in.Email)
	fmt.Printf("Password:\t")
	fmt.Scanln(&s_in.Password)
	fmt.Printf("Price:\t")
	fmt.Scanln(&s_in.Price)

	lampochka := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, s_in.Email) && strings.Contains(line, s_in.Password) {
			fmt.Println("Siz muvaffaqiyatli kirdingiz 🥳🥳🥳")
			lampochka = true
			break 
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scannerda xatolik:", err)
		return
	}

	if !lampochka {
		fmt.Println("Kirishda xatolik: email yoki parol noto'g'ri")
		os.Exit(0)
	}

// 	if lampochka {
// 		file3, err := os.Open("/home/abduazim/Projects/Golang/book/kutubxona.txt ")
// 		if err != nil {
// 			log.Fatalf("Faylni ochishda xatolik: %v", err)
// 		}
// 		defer file3.Close()

// 		scanner2 := bufio.NewScanner(file3)
// 		var arr []string
// 		for scanner2.Scan() {
// 			line := scanner2.Text()
// 			arr = append(arr, line)
// 			fmt.Println(line)
// 		}
// 		fmt.Println("")
// 		fmt.Println("Yuqoridagilardan birini tanlang👆🏻👆🏻👆🏻")
// 		fmt.Scanln(&num)

// 		cnt := 0
// 		for _, work := range arr {
// 			natija := strings.Split(work, " ")
// 			if natija[0] == num {
// 				if price, err := strconv.Atoi(natija[2]); err == nil && price <= s_in.Price {
// 					fmt.Println("Nechta olmoqchisiz ", natija[1])
// 					fmt.Scanln(&cnt)
// 					s_in.Price = s_in.Price - (price * cnt)
// 					fmt.Println("Sizning hisobizda ", s_in.Price, "mablag' bor")
// 				} else {
// 					fmt.Println("Hisobizda mag'lag' yetarli emas!!!")
// 				}
// 			}
// 		}
// 	}
}

func Sign_in() {
	var s_in SignIn
	SIGNIN(&s_in) 
}
// jjhlkjlkj