package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

var dojo_url = os.Getenv("DEFECTDOJO_URL")
var dojo_username = os.Getenv("DEFECTDOJO_USERNAME")
var dojo_password = os.Getenv("DEFECTDOJO_PASSWORD")

var engagement = "test2"
var product_name = "Apple Accounting Software"

type Results struct {
	Results []struct {
		Id int `json:"id"`
	}

	// Engagement []Engagement
}

// type Product struct {
// 	Id []int `json:"id"`
// }

// type Engagement struct {
// 	Id int `json:"id"`
// }

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

func get_product(token string, product_name string) int {
	fmt.Println("Get product ID")
	product_encoded := url.QueryEscape(product_name)
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
}

func get_engagement(token string, product_id int, engagement string) string {
	fmt.Println("Get engagement")
	engagement_name := url.QueryEscape(engagement)
	var dojo_engangement_url = dojo_url + "/engagements/?name=" + engagement_name + "&product=" + fmt.Sprint(product_id)
	req, err := http.NewRequest("GET", dojo_engangement_url, nil)
	req.Header.Add("Authorization", "Token "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err, resp)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var data Results
	err = decoder.Decode(&data)
	if len(data.Results) == 0 {
		return ""
	}
	return fmt.Sprint(data.Results[0].Id)
}

func create_enganement(token string, product_id int, engagement string) string {
	fmt.Println("Create Engagement")
	var jsonData = []byte(`{
		"name": "test1",
		"description": "Test description",
		"product": "3",
		"version": "1.1.1",
		"reason": "reason",
		"tracker": "https://example.com",
		"threat_model": "true",
		"api_test": "true",
		"pen_test": "true",
		"check_list": "true",
		"status": "Not Started",
		"engagement_type": "CI/CD",
		"build_id": "1",
		"commit_hash": "xxx",
		"branch_tag": "xxx",
		"deduplication_on_engagement": "false",
		"source_code_management_uri": "https://example.com",
		"first_contacted": "2025-03-18",
		"target_start": "2025-03-18",
		"target_end": "2025-03-25",
		"tags": ["tag1,tag2"]
	}`)
	var dojo_engagement_url = dojo_url + "/engagements/"
	request, err := http.NewRequest("POST", dojo_engagement_url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Add("Authorization", "Token "+token)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	// fmt.Println(res)
	fmt.Println(res["id"])
	created_id := fmt.Sprint(res["id"])
	return created_id
}

func import_scan(token string, product_name string, engagement_id string) (*http.Request, error) {
	fmt.Println("Import scan")
	var dojo_import_url = dojo_url + "/import-scan/"
	path := "results.json"
	extraParams := map[string]string{
		"minimum_severity":   "Info",
		"active":             "true",
		"verified":           "true",
		"scan_type":          "Bandit Scan",
		"close_old_findings": "false",
		"push_to_jira":       "false",
		"product_name":       product_name,
		"scan_date":          "2025-03-18",
		"check_list":         "true",
		"status":             "Not Started",
		"engagement":         engagement_id,
		"engagement_name":    "test1",
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))

	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range extraParams {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	fmt.Println(path)
	req, err := http.NewRequest("POST", dojo_import_url, body)
	req.Header.Add("Authorization", "Token "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Println(resp.StatusCode)
	}
	// fmt.Println(req)
	return req, err
}

func main() {

	token := authenticate()
	fmt.Println(token)

	product_id := get_product(token, product_name)
	fmt.Println(product_id)

	engagement_id := get_engagement(token, product_id, engagement)
	if fmt.Sprint(engagement_id) == "" {
		engagement_id = create_enganement(token, product_id, engagement)
	}
	scan_import, _ := import_scan(token, product_name, engagement_id)
	fmt.Println(scan_import)

}
