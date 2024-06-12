package main

import (
	"log"
	"net/http"
	"os"

	. "groupie/cmd/web/data"
	. "groupie/cmd/web/handlers"
	// . "groupie/cmd/web/struct"
)

func main() {
	err := fromMain(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

func fromMain(args []string) error {
	// Начинаем бесконечные попытки парсинга API на фоне
	go Parse()

	// Скрипт не пройдет дальше если массив артистов будет пуст(защита от отключения интернета с а данных на показ нет)

	for len(Artists) == 0 {
	}
	log.Println("\n\nSuccess parse. START SERVER\n")

	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/artists/", ArtistHandler)
	log.Println("http://localhost:80/")
	return http.ListenAndServe(":80", mux)
}
