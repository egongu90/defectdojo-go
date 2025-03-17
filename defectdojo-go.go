package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

var dojo_url = os.Getenv("DEFECTDOJO_URL")
var dojo_username = os.Getenv("DEFECTDOJO_USERNAME")
var dojo_password = os.Getenv("DEFECTDOJO_PASSWORD")

type Results struct {
	Results []Product
}

type Product struct {
	Id int `json:"id"`
}

func authenticate() string {
	var dojo_auth_url = dojo_url + "/api-token-auth/"
	fmt.Println("Authenticate and return token")
	values := map[string]string{"username": dojo_username, "password": dojo_password}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(dojo_auth_url, "application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	token := fmt.Sprint(res["token"])
	return token
}

func get_product(token string) int {
	fmt.Println("Get product ID")
	product_encoded := url.QueryEscape("Apple Accounting Software")
	var dojo_product_url = dojo_url + "/products/?name_exact=" + product_encoded
	req, err := http.NewRequest("GET", dojo_product_url, nil)
	req.Header.Add("Authorization", "Token "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var data Results
	err = decoder.Decode(&data)
	if len(data.Results) >= 2 {
		log.Fatal("More than 1 product found")
	}
	return data.Results[0].Id

	// fmt.Println(data.Results[0].Id)
	// for _, product := range data.Results {
	// 	fmt.Println(product.Id)
	// }
}

func get_engagement() {
	fmt.Println("Get engagement")
}

func create_enganement() {
	fmt.Println("Create Engagement")
}

func import_scan() {
	fmt.Println("Import scan")
}
func main() {
	token := authenticate()
	fmt.Println(token)

	product_id := get_product(token)
	fmt.Println(product_id)

	get_engagement()
	create_enganement()
	import_scan()
}
