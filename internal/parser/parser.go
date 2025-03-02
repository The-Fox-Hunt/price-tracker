package parser

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jnc-dev/price-tracker/model"
	)
	

func Parse(url string) (*model.Product, error) {

	// 1. Отправляем HTTP-запрос
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 2. Проверяем статус код ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to load page: %d", resp.StatusCode)
	}

	// 3. Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to get body: %w", err)
	var product *model.Product

	// 4. Извлекаем данные о товаре (тут нужно написать парсер HTML)
	var product Product

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to get NewDocument goquery: %w", err)
	}

	product_title := doc.Find(".product__title").Text()
	priceText := doc.Find(".product__price-cur").Text()
	product_price := strings.TrimSpace(strings.Replace(priceText, "₽", "", -1))

	price, err := strconv.Atoi(strings.ReplaceAll(product_price, " ", ""))
	if err != nil {
		return nil, fmt.Errorf("failed to parse price: %w", err)
	}

	product = &Product{
		Name:  product_title,
		Price: price,
	}

	fmt.Println("Название:", product_title)
	fmt.Println("Цена:", product_price)

	return product, nil
}
