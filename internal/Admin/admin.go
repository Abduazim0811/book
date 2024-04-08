package admin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	// Id int
	Book string
	Price uint16
	cnt uint
}

func Prt(pr Product) {
	var sentence2 string
	var slc []string
	cnt := 0
	fmt.Println("Qanday kitob qo'shmoqchisiz!!!")
	fmt.Printf("Book: ")
	fmt.Scanln(&pr.Book)
	fmt.Printf("Price: ")
	fmt.Scanln(&pr.Price)
	fmt.Printf("Count: ")
	fmt.Scanln(&pr.cnt)

	lampochka := false

	file, err := os.OpenFile("/home/abduazim/Projects/Golang/registratsiya/warehouse.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Fayl ochilmadi: ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fmt.Println(scanner.Scan())
	for scanner.Scan() {
		cnt++
		fmt.Println(cnt)
		line := scanner.Text()
		if !strings.Contains(line, pr.Book) {
			nimadur := strconv.Itoa(int(pr.Price))
			if strings.Contains(line, nimadur) {
				lampochka = true
			}
		}
		
	}
	slc = append(slc, fmt.Sprintf("%d", cnt+1),pr.Book, fmt.Sprintf("%d", pr.Price))
	if !lampochka {
		sentence2 = strings.Join(slc, " ")
		yozuvchi := bufio.NewWriter(file)
		fmt.Fprintln(yozuvchi, sentence2)
		yozuvchi.Flush()
	}
}

func Admin() {
	var admin string
	var password string
	fmt.Printf("Email:\t")
	fmt.Scanln(&admin)
	fmt.Printf("Password:\t")
	fmt.Scanln(&password)

	file, err := os.Open("/home/abduazim/Projects/Golang/registratsiya/admin.txt")
	if err != nil {
		fmt.Println("Fayldi ochishda xatolik bor!!!")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, admin) {
			if strings.Contains(line, password) {
				fmt.Println("Admin panelga kirdiz!!!")
				var pr Product
				Prt(pr)
				return
			}else{
				fmt.Println("Parol notugri!!!")
			}
		}else{
			fmt.Println("Parol notugri!!!")
		}
	}

}
