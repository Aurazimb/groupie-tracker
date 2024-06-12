package handlers

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"

	. "groupie/cmd/web/data"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	file := "ui/pages/index.html"
	// чтобы не было магических значений

	tmpl, err := template.ParseFiles(file)
	if err != nil {
		log.Print(err.Error())
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, Artists)
	if err != nil {
		log.Print(err.Error())
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artists/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	file := "ui/pages/artist_details.html"
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		log.Print(err.Error())
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	id := r.URL.Query().Get("id")
	match, err := regexp.MatchString("^[1-9][0-9]*$", id)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	if !match {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}

	art_id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	if art_id <= 0 || art_id > len(Artists) {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	err = tmpl.Execute(w, Artists[art_id-1])
	if err != nil {
		log.Print(err.Error())
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func ErrorHandler(w http.ResponseWriter, err_num int) {
	w.WriteHeader(err_num)
	tmpl, err := template.ParseFiles("ui/pages/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", err_num)
		return
	}
	text := http.StatusText(err_num)
	num := strconv.Itoa(err_num)
	text = num + "\n" + text

	err = tmpl.Execute(w, text)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
