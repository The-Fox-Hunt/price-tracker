package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/The-Fox-Hunt/price-tracker/internal/parser"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите URL: ")
	url, _ := reader.ReadString('\n')
	url = strings.TrimSpace(url)

	product, err := parser.Parse(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Название: %s\nЦена: %d₽\n", product.Name, product.Price)
}
