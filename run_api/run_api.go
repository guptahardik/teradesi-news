package run_api

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/robtec/newsapi/api"
)

func Give_news(query string, country string, category string) {
	godotenv.Load() //load .env file for api key
	//The API Has TWO Endpoints: Top Headlines and Everything
	//Top Headlines: https://newsapi.org/docs/endpoints/top-headlines
	//Everything: https://newsapi.org/docs/endpoints/everything
	//--------------------------------------------------------------------------------------------------------//
	//sortBy
	// The order to sort the articles in. Possible options: relevancy, popularity, publishedAt.
	// relevancy = articles more closely related to q come first.
	// popularity = articles from popular sources and publishers come first.
	// publishedAt = newest articles come first.
	// Default: publishedAt

	//--------------------------------------------------------------------------------------------------------//
	// language - en, fr, de, es, it, nl, no, pt, ru, se, ud, zh

	httpClient := http.Client{}
	key := os.Getenv("key")
	url := "https://newsapi.org"

	client, err := api.New(&httpClient, key, url)
	if err != nil {
		panic(err)
	}
	opts := api.Options{Country: country, Category: category, Q: query}
	topHeadlines, err := client.TopHeadlines(opts)

	// Different options
	//moreOpts := api.Options{Language: "en", Q: query, SortBy: "popularity"}

	// Get Everything with options from above
	//everything, err := client.Everything(moreOpts)
	if err != nil {
		panic(err)
	}
	if len(topHeadlines.Articles) > 0 {
		for _, article := range topHeadlines.Articles {
			println(article.Title)
		}
	}

}
