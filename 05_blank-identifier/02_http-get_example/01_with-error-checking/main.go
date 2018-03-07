package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.geekwiseacademy.com/")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
	
	anotherWay()
}


func anotherWay() {
	var res *http.Response
	var page []byte
	var err error
	if res, err = http.Get("http://www.geekwiseacademy.com/"); err != nil {
		log.Fatal(err)
	} else {
		defer res.Body.Close()
	}

	if page, err = ioutil.ReadAll(res.Body); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf(string(page))
	}
}