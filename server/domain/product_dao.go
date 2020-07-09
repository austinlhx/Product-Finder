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

func SearchProduct(product models.Product) *models.Products {
	//TODO: Verify if info is correct
	tempUpper, _ := strconv.ParseFloat(product.UpperBound, 2)
	tempLower, _ := strconv.ParseFloat(product.LowerBound, 2)
	p := &models.Products{
		ProductName: product.ProductName,
		ProductType: product.ProductType,
		UpperBound:  tempUpper,
		LowerBound:  tempLower,
	}
	log.Println(p)
	return p
}

//type temp should be different
func SearchBestBuy(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	defer wg.Done()
	query := product.ProductName
	query = strings.ReplaceAll(query, " ", "%20")
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"),
	)
	//allProducts := make(map[string]string)
	// Find and visit all links
	c.OnHTML("li.sku-item", func(e *colly.HTMLElement) {
		temp := models.ProductFound{}
		//fmt.Println(e.Request.AbsoluteURL(e.Attr("href")))
		//e.Request.Visit(e.Attr("href"))
		//e.Request.Visit(e.ChildText("a-price"))\
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
		} //calling <span aria-hidden= true>Price</span>
		//fmt.Println(price)
		if temp.Price < product.UpperBound && temp.Price > product.LowerBound {
			ch <- temp
		}

	})

	c.OnRequest(func(r *colly.Request) {
		//r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.bestbuy.com/site/searchpage.jsp?st=" + query)
	//c.Visit("https://www.bestbuy.com/site/searchpage.jsp?cp=2&st=" + query) //should i do this or naw
	//fmt.Println(allProducts)
	//log.Println(products)

}

func SearchAmazon(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	defer wg.Done()
	query := product.ProductName
	query = strings.ReplaceAll(query, " ", "%20")
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"),
	)
	c.OnHTML("div.a-section.a-spacing-medium", func(e *colly.HTMLElement) { //a-size-medium a-color-base a-text-normal
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
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.amazon.com/s?k=" + query + "&ref=nb_sb_noss_2")

}

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
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.newegg.com/p/pl?d=" + query)

}

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
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.bhphotovideo.com/c/search?Ntt=" + query + "&N=0&InitialSearch=yes&sts=ma")
}

func SearchAdorama(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	defer wg.Done()
	query := product.ProductName
	query = strings.ReplaceAll(query, " ", "%20")
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"),
	)

	// Find and visit all links
	c.OnHTML("div.item", func(e *colly.HTMLElement) {
		temp := models.ProductFound{}
		name := e.ChildText("h2 a.trackEvent")
		temp.Link = e.ChildAttr("h2 a.trackEvent", "href")
		temp.Image = "https://www.adorama.com" + e.ChildAttr("a.trackEvent img.productImage", "src")
		removeWhiteSpace := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
		temp.Name = removeWhiteSpace.ReplaceAllString(name, "")
		//price := e.ChildText("span[data-selenium=uppedDecimalPriceFirst]") + "." + e.ChildText("sup[data-selenium=uppedDecimalPriceSecond]")
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
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.adorama.com/l/?searchinfo="+ query + "&sel=Item-Condition_New-Items")
	//fmt.Println(allProducts)
}
