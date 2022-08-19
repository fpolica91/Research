package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getProducts(name string, size string, page string, countryCode string, includeRange string, includeFixed string) *http.Response {
	uri, _ := url.Parse("https://giftcards.reloadly.com/products")

	values := uri.Query()
	values.Add("productName", name)
	values.Add("size", size)
	values.Add("page", page)
	values.Add("countryCode", countryCode)
	values.Add("includeRange", includeRange)
	values.Add("includeFixed", includeFixed)
	uri.RawQuery = values.Encode()

	var bearer = "Bearer " + os.Getenv("RELOADLY_API_KEY")
	response, _ := http.NewRequest("GET", uri.String(), nil)
	response.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(response)
	if err != nil {
		fmt.Println("Error on response.\n[ERRO] -", err)
		os.Exit(1)
	}

	return resp

}

func main() {

	if len(os.Args[1:]) == 0 {
		fmt.Println("Please provide a product name")
		os.Exit(1)
	}
	args := os.Args[1:]
	product := args[0]

	resp := getProducts(product, strconv.Itoa(10), strconv.Itoa(0), "US", strconv.FormatBool(true), strconv.FormatBool(true))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
