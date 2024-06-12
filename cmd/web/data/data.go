package data

import (
	"encoding/json"
	"fmt"
	. "groupie/cmd/web/struct"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

var mutex sync.Mutex

const (
	artistsURL  = "https://groupietrackers.herokuapp.com/api/artists"
	relationURL = "https://groupietrackers.herokuapp.com/api/relation"
)

var Artists []Artist

func Parse() {
	for true {
		tempArtists, err := artistsParce()
		if err != nil {
			log.Println("Internet Error, try again...")
			return
		}

		relations, err := realtionsParce()
		if err != nil {
			fmt.Println("парсинг relations")
			log.Println(err)
			continue

		}
		for i := 0; i < len(tempArtists); i++ {
			tempArtists[i].Relations = relations[i]
		}

		if fmt.Sprint(Artists) != fmt.Sprint(tempArtists) {
			mutex.Lock()

			Artists = tempArtists

			mutex.Unlock()
		}

		time.Sleep(5 * time.Minute)
	}
}

func artistsParce() ([]Artist, error) {
	var res []Artist

	resp, err := http.Get(artistsURL)
	if err != nil {
		// log.Printf("Error: GETRequest error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(data, &res)

	if err != nil {
		log.Println(err)
		return res, err
	}

	return res, nil
}

func realtionsParce() ([]Relation, error) {
	var withIndex RelationsIndex
	var outIndex []Relation
	resp, err := http.Get(relationURL)
	if err != nil {
		return outIndex, fmt.Errorf("GETRequest error: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return outIndex, fmt.Errorf("error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &withIndex)
	if err != nil {
		return outIndex, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return withIndex.Index, nil
}
