package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func errCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	url := "https://techcrunch.com/"

	response, err := http.Get(url)

	defer response.Body.Close()

	errCheck(err)

	if response.StatusCode > 400 {
		fmt.Println("Status code: ", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)

	errCheck(err)

	river, err := doc.Find("div.river").Html()

	errCheck(err)

	fmt.Println(river)
}
