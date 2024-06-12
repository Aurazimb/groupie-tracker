package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	Original, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 10000; i++ {
		go get(i, Original)
	}
	time.Sleep(10 * time.Second)
}

func get(i int, Original []byte) {
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)

	if string(res) != string(Original) {
		fmt.Println(string(Original))
		fmt.Println("Err")
	} else {
		fmt.Println("OK", i)
	}
}
