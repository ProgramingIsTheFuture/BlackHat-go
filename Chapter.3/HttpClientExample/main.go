package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	r1, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer r1.Body.Close()

	log.Println(r1.Status)

	body, err := ioutil.ReadAll(r1.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	r2, err := http.Head("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer r2.Body.Close()

	form := url.Values{}
	form.Add("foo", "bar")
	// r3, err := http.Post("http://www.google.com/robots.txt", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	r3, err := http.PostForm("http://www.google.com/robots.txt", form)
	if err != nil {
		log.Fatalln(err)
	}
	defer r3.Body.Close()

	req, err := http.NewRequest("DELETE", "https://www.google.com/robots.txt", nil)
	if err != nil {
		log.Println(err)
	}
	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	req, err = http.NewRequest("PUT", "https://www.google.com/robots.txt", strings.NewReader(form.Encode()))
	if err != nil {
		log.Println(err)
	}

	resp, err = client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
}
