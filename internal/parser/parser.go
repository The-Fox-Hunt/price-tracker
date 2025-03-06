package parser

//https://zielinskiandrozen.ru/product/duhi-kontsentrirovannye-dubovyy-moh-ambra-50ml

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/The-Fox-Hunt/price-tracker/internal/model"
)

func Parse(url string) (*model.Product, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to load page: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to get body: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to get NewDocument goquery: %w", err)
	}

	err = os.WriteFile("debug.html", body, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	}
	fmt.Println("DEBUG: HTML сохранён в debug.html")

	product_title := strings.TrimSpace(doc.Find(".product__title").Text())
	//priceText := doc.Find("span.product__price-cur").Text()
	priceText, exists := doc.Find("span.product__price-cur").Attr("data-product-card-price-from-cart")
	if !exists {
		return nil, fmt.Errorf("price not found in attribute")
	}

	product_price := strings.TrimSpace(strings.Replace(priceText, "₽", "", -1))

	fmt.Println("DEBUG: priceText =", priceText)

	price, err := strconv.Atoi(strings.ReplaceAll(product_price, " ", ""))
	if err != nil {
		return nil, fmt.Errorf("failed to parse price: %w", err)
	}

	var product = &model.Product{
		Name:  product_title,
		Price: price,
	}

	fmt.Println("Название:", product_title)
	fmt.Println("Цена:", price)

	return product, nil
}
