package main

import (
	"fmt"
	. "groupie/cmd/web/data"
	. "groupie/cmd/web/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	handler := http.HandlerFunc(Home)

	// Тут мы отправляем GET запрос на указанный URL
	Get, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Тут создаем пустой ответ для того чтобы на него записывать ответ с нашего хэндлера
	rr := httptest.NewRecorder()
	fmt.Println(rr, "тут мы принтим rr")
	// Тут мы отправляем GET запрос в наш Хэндлер и получаем ответ
	handler.ServeHTTP(rr, Get)
	// тут идет сравнение наших ответов и то что мы ожидаем
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GET handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Пример 2: POST запрос
	Post, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Используем тот же объект ResponseRecorder
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, Post)
	// отправляем POST запрос по нае
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("POST handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestArtist(t *testing.T) {
	go Parse()
	for len(Artists) == 0 {
	}
	handler := http.HandlerFunc(ArtistHandler)
	// Тут мы отправляем GET запрос на указанный URL
	Get, err := http.NewRequest("GET", "/artists/?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Тут создаем пустой ответ для того чтобы на него записывать ответ с нашего хэндлера
	rr := httptest.NewRecorder()
	fmt.Println(rr, "тут мы принтим rr")
	// Тут мы отправляем GET запрос в наш Хэндлер и получаем ответ
	handler.ServeHTTP(rr, Get)
	// тут идет сравнение наших ответов и то что мы ожидаем
	if status53 := rr.Code; status53 != http.StatusOK {
		t.Errorf("GET handler returned wrong status code: got %v want %v",
			status53, http.StatusOK)
	}
	// отправляем гет запрос только уже по артисту под номером 53
	Get, err = http.NewRequest("GET", "/artists/?id=53", nil)
	if err != nil {
		if err != nil {
			t.Fatal(err)
		}
	}
	// сощдание пустого ответа
	rr = httptest.NewRecorder()
	// туут уже тдет отправка нашего запроса по нашему хэндлеру
	handler.ServeHTTP(rr, Get)
	if rr.Code != http.StatusNotFound {
		t.Errorf("GET handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusNotFound)
	}
	// отправляем запрос на нечисловые значения
	Get, err = http.NewRequest("GET", "/artists/?id=asd", nil)
	if err != nil {
		if err != nil {
			t.Fatal(err)
		}
	}
	// сощдание пустого ответа
	rr = httptest.NewRecorder()
	// туут уже тдет отправка нашего запроса по нашему хэндлеру
	handler.ServeHTTP(rr, Get)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("GET handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusBadRequest)
	}
	// отправляем запрос на 0-го артиста
	Get, err = http.NewRequest("GET", "/artists/?id=0", nil)
	if err != nil {
		if err != nil {
			t.Fatal(err)
		}
	}
	// сощдание пустого ответа
	rr = httptest.NewRecorder()
	// туут уже тдет отправка нашего запроса по нашему хэндлеру
	handler.ServeHTTP(rr, Get)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("GET handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusBadRequest)
	}
	// Пример 2: POST запрос
	Post, err := http.NewRequest("POST", "/artists/", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Используем тот же объект ResponseRecorder
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, Post)
	// отправляем POST запрос по нае
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("POST handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusMethodNotAllowed)
	}
}
