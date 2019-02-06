package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	postTest()
	putTest()
	getTest()
	searchTest()
	deleteTest()

}

func postTest() {
	url := "http://localhost:1323/accounts/uber/contacts"

	payload := strings.NewReader("{\n\t\"number\" : \"+918867997430\",\n\t\"first_name\": \"raushan\",\n\t\"Last_name\": \"kumar\",\n\t\"company_name\":\"uber\",\n\t\"email\": \"raushan@google.com\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Basic dWJlcjpkamNmZGJ2aGVjempjd2RoY2Q=")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "73f4c615-424c-a22d-8879-bb2d1273f28a")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func putTest() {
	url := "http://localhost:1323/accounts/uber/contacts/1"

	payload := strings.NewReader("{\n\t\"number\" : \"+918867997430\",\n\t\"first_name\": \"raushan\",\n\t\"Last_name\": \"the kumar\",\n\t\"company_name\":\"uber\",\n\t\"email\": \"raushan@google.com\"\n}")

	req, _ := http.NewRequest("PUT", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Basic dWJlcjpkamNmZGJ2aGVjempjd2RoY2Q=")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "42f4980e-2c07-ed87-d573-924464db2842")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func getTest() {
	url := "http://localhost:1323/accounts/uber/contacts/1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Basic dWJlcjpkamNmZGJ2aGVjempjd2RoY2Q=")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "f19a6e63-8db9-1e93-4186-5a4a889c3a97")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func searchTest() {
	url := "http://localhost:1323/accounts/uber/contacts/search/raushan"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Basic dWJlcjpkamNmZGJ2aGVjempjd2RoY2Q=")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "e923d011-b4dd-3aba-297b-7bbe94e28214")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func deleteTest() {
	url := "http://localhost:1323/accounts/uber/contacts/1"

	req, _ := http.NewRequest("DELETE", url, nil)

	req.Header.Add("authorization", "Basic dWJlcjpkamNmZGJ2aGVjempjd2RoY2Q=")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "62db59ac-a077-bbdc-186c-e5a19f78dd80")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
