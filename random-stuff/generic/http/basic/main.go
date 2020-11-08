package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	target = `https://www.google.com/robots.txt`
)

func main() {
	resp, err := http.Get(target)
	if err != nil {
		log.Panicf("error: %s", err.Error())
	}

	fmt.Printf("GET http status %s (%v)\n", resp.Status, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("error: %s", err.Error())
	}
	fmt.Printf("response:\n%v\n", string(body))
	resp.Body.Close()

	resp, err = http.Head(target)
	if err != nil {
		log.Panicf("error: %s", err.Error())
	}
	resp.Body.Close()
	fmt.Printf("HEAD http status %s (%v)\n", resp.Status, resp.StatusCode)

	form := url.Values{}
	form.Add("foo", "bar")
	resp, err = http.Post(
		target,
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		log.Panicf("error: %s", err.Error())
	}
	resp.Body.Close()
	fmt.Printf("POST http status %s (%v)\n", resp.Status, resp.StatusCode)

	req, err := http.NewRequest("DELETE", target, nil)
	if err != nil {
		log.Panicf("error: %s", err.Error())
	}

	var client http.Client
	resp, err = client.Do(req)
	resp.Body.Close()
	fmt.Printf("DELETE http status %s (%v)\n", resp.Status, resp.StatusCode)

	req, err = http.NewRequest(
		"PUT",
		target,
		strings.NewReader(form.Encode()),
	)
	resp, err = client.Do(req)
	if err != nil {
		log.Panicf("error: %s", err.Error())
	}
	resp.Body.Close()
	fmt.Printf("PUT http status %s (%v)\n", resp.Status, resp.StatusCode)

}
