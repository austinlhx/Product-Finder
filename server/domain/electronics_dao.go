package domain

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"../models"
	"github.com/gocolly/colly"
)

//SearchBestBuy searches all of https://www.bestbuy.com product info
func SearchBestBuy(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	defer wg.Done()
	query := product.ProductName
	query = strings.ReplaceAll(query, " ", "%20")
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"),
	)

	c.OnHTML("li.sku-item", func(e *colly.HTMLElement) {
		temp := models.ProductFound{}

		temp.Image = e.ChildAttr("img.product-image", "src")
		temp.Link = "http://www.bestbuy.com" + e.ChildAttr("a.image-link", "href")
		temp.Name = e.ChildText("h4.sku-header")
		price := e.ChildText("div.priceView-hero-price.priceView-customer-price span[aria-hidden=true]")
		switch lengthPrice := len(price); {
		case lengthPrice > 5:
			temp.Price, _ = strconv.ParseFloat(price[1:6], 2)
		case lengthPrice == 5:
			temp.Price, _ = strconv.ParseFloat(price[1:5], 2)
		default:
			temp.Price = 0
		} 
		if temp.Price < product.UpperBound && temp.Price > product.LowerBound {
			ch <- temp
		}

	})

	c.OnRequest(func(r *colly.Request) {

		log.Println("Visiting", r.URL)
	})

	c.Visit("https://www.bestbuy.com/site/searchpage.jsp?st=" + query)
	//c.Visit("https://www.bestbuy.com/site/searchpage.jsp?cp=2&st=" + query)


}

//SearchAmazon searches all of https://www.amazon.com product info
func SearchAmazon(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	defer wg.Done()
	query := product.ProductName
	query = strings.ReplaceAll(query, " ", "%20")
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"),
	)
	c.OnHTML("div.a-section.a-spacing-medium", func(e *colly.HTMLElement) { 
		temp := models.ProductFound{}
		if e.ChildText("span.a-size-base-plus.a-color-base.a-text-normal") != "" {
			temp.Name = e.ChildText("span.a-size-base-plus.a-color-base.a-text-normal")
		} else {
			temp.Name = e.ChildText("span.a-size-medium.a-color-base.a-text-normal")
		}
		temp.Link = "http://www.amazon.com" + e.ChildAttr("a.a-link-normal.a-text-normal", "href")
		temp.Image = e.ChildAttr("img.s-image", "src")
		price := e.ChildText("span.a-offscreen")
		switch lengthPrice := len(price); {
		case lengthPrice > 5:
			temp.Price, _ = strconv.ParseFloat(price[1:6], 2)
		case lengthPrice == 5:
			temp.Price, _ = strconv.ParseFloat(price[1:5], 2)
		default:
			temp.Price = 0
		}
		if temp.Price < product.UpperBound && temp.Price > product.LowerBound {
			ch <- temp
		}

	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.Visit("https://www.amazon.com/s?k=" + query + "&ref=nb_sb_noss_2")

}

//SearchNewEgg searches all of https://www.newegg.com product info
func SearchNewEgg(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	defer wg.Done()
	query := product.ProductName
	query = strings.ReplaceAll(query, " ", "%20")
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"),
	)

	c.OnHTML("div.item-container", func(e *colly.HTMLElement) {
		temp := models.ProductFound{}
		temp.Name = e.ChildText("a.item-title")
		temp.Link = e.ChildAttr("a.item-title", "href")
		temp.Image = e.ChildAttr("img", "src")
		temp.Price, _ = strconv.ParseFloat((e.ChildText("li.price-current strong") + e.ChildText("li.price-current sup")), 2)
		if temp.Price < product.UpperBound && temp.Price > product.LowerBound {
			ch <- temp
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.Visit("https://www.newegg.com/p/pl?d=" + query)

}

//SearchBHPhotoVideo searches all of https://www.bhphotovideo.com product info
func SearchBHPhotoVideo(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	defer wg.Done()
	query := product.ProductName
	query = strings.ReplaceAll(query, " ", "%20")
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"),
	)
	c.OnHTML("div[data-selenium=miniProductPage]", func(e *colly.HTMLElement) {
		temp := models.ProductFound{}
		temp.Name = e.ChildText("h3[data-selenium=miniProductPageName]")
		temp.Link = "https://www.bhphotovideo.com" + e.ChildAttr("a[data-selenium=miniProductPageProductNameLink]", "href")
		temp.Image = e.ChildAttr("img[data-selenium=miniProductPageImg]", "src") //seems to not work
		price := e.ChildText("span[data-selenium=uppedDecimalPriceFirst]") + "." + e.ChildText("sup[data-selenium=uppedDecimalPriceSecond]")
		switch lengthPrice := len(price); {
		case lengthPrice > 5:
			temp.Price, _ = strconv.ParseFloat(price[1:6], 2)
		case lengthPrice == 5:
			temp.Price, _ = strconv.ParseFloat(price[1:5], 2)
		default:
			temp.Price = 0
		}
		if temp.Price < product.UpperBound && temp.Price > product.LowerBound {
			ch <- temp
		}

	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.Visit("https://www.bhphotovideo.com/c/search?Ntt=" + query + "&N=0&InitialSearch=yes&sts=ma")
}

//SearchAdorama searches all of https://www.adorama.com product info
func SearchAdorama(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	defer wg.Done()
	query := product.ProductName
	query = strings.ReplaceAll(query, " ", "%20")
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"),
	)


	c.OnHTML("div.item", func(e *colly.HTMLElement) {
		temp := models.ProductFound{}
		name := e.ChildText("h2 a.trackEvent")
		temp.Link = e.ChildAttr("h2 a.trackEvent", "href")
		temp.Image = "https://www.adorama.com" + e.ChildAttr("a.trackEvent img.productImage", "src")
		removeWhiteSpace := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
		temp.Name = removeWhiteSpace.ReplaceAllString(name, "")
		price := e.ChildText("strong.your-price")
		switch lengthPrice := len(price); {
		case lengthPrice > 5:
			temp.Price, _ = strconv.ParseFloat(price[1:6], 2)
		case lengthPrice == 5:
			temp.Price, _ = strconv.ParseFloat(price[1:5], 2)
		default:
			temp.Price = 0
		}
		if temp.Price < product.UpperBound && temp.Price > product.LowerBound {
			ch <- temp
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.Visit("https://www.adorama.com/l/?searchinfo="+ query + "&sel=Item-Condition_New-Items")
}

