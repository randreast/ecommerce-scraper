package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
)

func scraper(csvWriter *csv.Writer, limit int) {
	res, err := http.Get("https://www.tokopedia.com/p/handphone-tablet/handphone?ob=5&page=1")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".css-13l3l78 .e1nlzfl10").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		if i < limit {
			data := []string{
				s.Find(".css-1bjwylw").Text(),
				s.Find(".css-o5uqvq").Text(),
				s.Find(".css-1kr22w3").Text(),
			}

			csvWriter.Write(data)
		}
	})

	csvWriter.Flush()
}

func main() {
	godotenv.Load()

	// get limit product from .env
	limit, err := strconv.Atoi(os.Getenv("LIMIT"))
	if err != nil {
		log.Fatal(err)
	}

	// create csv files
	csvFile, err := os.Create("product.csv")
	if err != nil {
		log.Fatal(err)
	}

	// create csv writer
	csvWriter := csv.NewWriter(csvFile)

	// call function to read body and write to csv
	scraper(csvWriter, limit)

	csvFile.Close()
}
