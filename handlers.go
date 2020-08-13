package main

import (
	"encoding/json"
	"net/http"
	"path"
)

func find(x string) int {
	for i, book := range books {
		if x == book.Id {
			return i
		}
	}
	return -1
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 {
		return getAll(w)
	}
	dataJson, err := json.Marshal(books[i])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
	return
}

func getAll(w http.ResponseWriter) (err error) {
	dataJson, err := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	book := Book{}
	json.Unmarshal(body, &book)
	books = append(books, book)
	w.WriteHeader(200)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 {
		return
	}
	bookOriginal := books[i]
	//----------
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	bookUpdated := Book{}
	json.Unmarshal(body, &bookUpdated)

	if bookUpdated.Id != "" {
		bookOriginal.Id = bookUpdated.Id
	}
	if bookUpdated.Title != "" {
		bookOriginal.Title = bookUpdated.Title
	}
	if bookUpdated.Edition != "" {
		bookOriginal.Edition = bookUpdated.Edition
	}
	if bookUpdated.Copyright != "" {
		bookOriginal.Copyright = bookUpdated.Copyright
	}
	if bookUpdated.Language != "" {
		bookOriginal.Language = bookUpdated.Language
	}
	if bookUpdated.Pages != "" {
		bookOriginal.Pages = bookUpdated.Pages
	}
	if bookUpdated.Author != "" {
		bookOriginal.Author = bookUpdated.Author
	}
	if bookUpdated.Publisher != "" {
		bookOriginal.Publisher = bookUpdated.Publisher
	}
	books[i] = bookOriginal
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 {
		return
	}
	books = append(books[:i], books[i+1:]...)
	w.WriteHeader(200)
	return
}
