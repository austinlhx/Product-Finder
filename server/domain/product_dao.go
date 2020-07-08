package domain

import (
	"fmt"
	"log"
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
	//c.Visit("https://www.bestbuy.com/site/searchpage.jsp?cp=2&st=xboxcontroller") //should i do this or naw
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

	// Find and visit all links
	c.OnHTML("div.a-section.a-spacing-medium", func(e *colly.HTMLElement) {
		temp := models.ProductFound{}
		//fmt.Println(e.Request.AbsoluteURL(e.Attr("href")))
		//e.Request.Visit(e.Attr("href"))
		//e.Request.Visit(e.ChildText("a-price"))
		temp.Name = e.ChildText("span.a-size-base-plus.a-color-base.a-text-normal")
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

		//e.Request.Visit(e.Text)
		if temp.Price < product.UpperBound && temp.Price > product.LowerBound {
			ch <- temp
		}
		

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.amazon.com/s?k=" + query + "&ref=nb_sb_noss_2")
	//log.Println(products)

}

/*func target(){
	c := colly.NewCollector()
	allProducts := make(map[string]string)
	// Find and visit all links
	c.OnHTML("div.Row-uds8za-0.cCApfd.styles__StyledProductCardRow-mkgs8k-1.gjLYKQ", func(e *colly.HTMLElement) {
		//e.Request.AbsoluteURL(e.Attr("href"))
		//e.Request.Visit(e.ChildText("a-price"))
		productName := e.ChildText("a.Link-sc-1khjl8b-0.styles__StyledTitleLink-mkgs8k-5.jhiHBx.h-display-block.h-text-bold.h-text-bs.flex-grow-one")
		price := e.ChildText("div.styles__StyledPricePromoWrapper-mkgs8k-9.koDuTx")
		//fmt.Println(price)
		allProducts[productName] = price
		//e.Request.Visit(e.Text)

	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.target.com/s?searchTerm=xbox+controller")
	fmt.Println(allProducts)
	//
}*/ //Target got some crazy protection???

/*
func walmart() { //Walmart got different tags for each item
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"),
	)
	allProducts := make(map[string]string)
	// Find and visit all links
	c.OnHTML("li.Grid-col.u-size-6-12.u-size-1-4-m.u-size-1-5-xl.search-gridview-first-grid-row-item", func(e *colly.HTMLElement) {
		//e.Request.AbsoluteURL(e.Attr("href"))
		//e.Request.Visit(e.ChildText("a-price"))
		productName := e.ChildText("span.price.display-inline-block.arrange-fit.price.price-main")
		price := e.ChildText("a.product-title-link.line-clamp.line-clamp-2.truncate-title")
		//fmt.Println(price)
		allProducts[productName] = price
		//e.Request.Visit(e.Text)
		//product-title-link line-clamp line-clamp-2 truncate-title
		//product-title-link line-clamp line-clamp-2 truncate-title
		//Grid-col u-size-6-12 u-size-1-4-m u-size-1-5-xl search-gridview-first-grid-row-item
		//Grid-col u-size-6-12 u-size-1-4-m u-size-1-5-xl search-gridview-last-col-item search-gridview-first-grid-row-item
		//Grid-col u-size-6-12 u-size-1-4-m u-size-1-5-xl search-gridview-first-grid-row-item
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.walmart.com/search/?query=xbox%20controller")
	fmt.Println(allProducts)
	//
}*/
