package domain

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"../models"
	"github.com/gocolly/colly"
)

//SearchBloomingdales searches all of https://www.bloomingdales.com product info
func SearchBloomingdales(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	defer wg.Done()
	query := product.ProductName
	query = strings.ReplaceAll(query, " ", "%20")
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"),
	)

	c.OnHTML("li.small-6.medium-4.large-4.cell", func(e *colly.HTMLElement) {
		temp := models.ProductFound{}
		if name := e.ChildText("div.productDescription"); name != "" {
			temp.Name = e.ChildText("div.productDescription")
		} else {
			temp.Name = e.ChildText("div.productDescription.new-arrival")
		}
		temp.Link = "https://www.bloomingdales.com" + e.ChildAttr("a.productDescLink", "href")
		temp.Image = e.ChildAttr("img", "src")
		unfilteredPrice := e.ChildText("div.tnPrice span.regular")
		removeWhiteSpace := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
		price := removeWhiteSpace.ReplaceAllString(unfilteredPrice, "")
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

	c.Visit("https://www.bloomingdales.com/shop/search?keyword=" + query)

}

//SearchSaksFifth searches all of https://www.saksfifthavenue.com product info
func SearchSaksFifth(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	defer wg.Done()
	query := product.ProductName
	query = strings.ReplaceAll(query, " ", "%20")
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"),
	)

	c.OnHTML("div[data-url]", func(e *colly.HTMLElement) {
		temp := models.ProductFound{}
		temp.Name = e.ChildText("p.product-description")

		temp.Link = e.ChildAttr("div.product-text a.mainBlackText", "href")
		temp.Image = e.ChildAttr("img.pa-product-large", "src")

		price := e.ChildText("span.product-price")
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

	c.Visit("https://www.saksfifthavenue.com/search/EndecaSearch.jsp?bmArch=bmForm&bmForm=endeca_search_form_one&bmArch=bmIsForm&bmIsForm=true&bmHidden=submit-search&submit-search=&bmArch=bmSingle&bmSingle=N_Dim&bmHidden=N_Dim&N_Dim=0&bmArch=bmHidden&bmHidden=Ntk&bmHidden=Ntk&Ntk=Entire+Site&bmArch=bmHidden&bmHidden=Ntx&bmHidden=Ntx&Ntx=mode%2Bmatchpartialmax&bmHidden=PA&PA=TRUE&SearchString=" + query)

}

//SearchNeimanMarcus searches all of https://www.neimanmarcus.com product info
func SearchNeimanMarcus(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	defer wg.Done()
	query := product.ProductName
	query = strings.ReplaceAll(query, " ", "%20")
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"),
	)
	//div.product-thumbnail grid-33 tablet-grid-33 mobile-grid-50 grid-1600 enhancement
	c.OnHTML("div.product-thumbnail.grid-33.tablet-grid-33.mobile-grid-50.grid-1600.enhancement", func(e *colly.HTMLElement) {
		temp := models.ProductFound{}
		temp.Name = e.ChildText("h2 span.name")

		temp.Link = e.ChildAttr("a.product-thumbnail__link", "href")
		temp.Image = "https:" + e.ChildAttr("img", "src")

		price := e.ChildText("span.price-no-promo")
		price = strings.Replace(price, ",", "", -1)
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

	c.Visit("https://www.neimanmarcus.com/search.jsp?from=brSearch&l=" + query + "&q=" + query)

}
