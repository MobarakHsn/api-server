package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// This function get the first unused id for Book
func GetAvailableBookId() int {
	for i := 1; ; i++ {
		_, ok := BookList[i]
		if !ok {
			return i
		}
	}
}

// This method adds a new Book
func (book *Book) AddBook() {
	id := GetAvailableBookId()
	book.BookId = id
	AuthorList[book.AuthorId].MyBooks = append(AuthorList[book.AuthorId].MyBooks, *book)
	BookList[book.BookId] = book
}

// This http handler function handle the request of adding new Book
func HandleAddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), 400)
		return
	}
	if _, ok := AuthorList[newBook.AuthorId]; ok != true {
		w.WriteHeader(http.StatusForbidden)
		http.Error(w, "This author is not available first add the author!", 400)
		return
	}
	newBook.AddBook()

	err = json.NewEncoder(w).Encode(BookList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// This http handler function handle the request of deleting an existing Book
func HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	id, err := strconv.Atoi(chi.URLParam(r, "bookId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), 400)
		return
	}
	if _, ok := BookList[id]; ok == false {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Invalid Book id", 404)
		return
	}
	oldBook := BookList[id]
	for i := 0; i < len(AuthorList[oldBook.AuthorId].MyBooks); i++ {
		if AuthorList[oldBook.AuthorId].MyBooks[i].BookId == oldBook.BookId {
			AuthorList[oldBook.AuthorId].MyBooks = append(AuthorList[oldBook.AuthorId].MyBooks[:i], AuthorList[oldBook.AuthorId].MyBooks[i+1:]...)
			break
		}
	}
	delete(BookList, id)

	err = json.NewEncoder(w).Encode(BookList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// This http handler function handle the request of updating an existing Book
func HandleUpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	id, err := strconv.Atoi(chi.URLParam(r, "bookId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), 400)
		return
	}
	if _, ok := BookList[id]; ok == false {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Invalid Book id", 404)
		return
	}

	oldBook := BookList[id]
	var newBook Book
	err = json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, err.Error(), 400)
		return
	}

	if oldBook.AuthorId != newBook.AuthorId {
		if _, ok := AuthorList[newBook.AuthorId]; ok == false {
			w.WriteHeader(http.StatusNotFound)
			http.Error(w, "Invalid author id", 404)
			return
		}
		for i := 0; i < len(AuthorList[oldBook.AuthorId].MyBooks); i++ {
			if AuthorList[oldBook.AuthorId].MyBooks[i].BookId == oldBook.BookId {
				AuthorList[oldBook.AuthorId].MyBooks = append(AuthorList[oldBook.AuthorId].MyBooks[:i], AuthorList[oldBook.AuthorId].MyBooks[i+1:]...)
				break
			}
		}
		newBook.BookId = id
		BookList[id] = &newBook
		AuthorList[newBook.AuthorId].MyBooks = append(AuthorList[newBook.AuthorId].MyBooks, newBook)

		err = json.NewEncoder(w).Encode(BookList)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			http.Error(w, err.Error(), 400)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	newBook.BookId = id
	BookList[id] = &newBook
	for i := range AuthorList[oldBook.AuthorId].MyBooks {
		if AuthorList[oldBook.AuthorId].MyBooks[i].BookId == oldBook.BookId {
			AuthorList[oldBook.AuthorId].MyBooks[i] = newBook
			break
		}
	}
	err = json.NewEncoder(w).Encode(BookList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// This http handler function handle the request of getting all the Books
func HandleGetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	err := json.NewEncoder(w).Encode(BookList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// This function get the first unused id for Author
func GetAvailableAuthorId() int {
	for i := 1; ; i++ {
		_, ok := AuthorList[i]
		if !ok {
			return i
		}
	}
}

// This method adds a new Author
func (author *Author) AddAuthor() {
	id := GetAvailableAuthorId()
	author.AuthorId = id
	AuthorList[id] = author
}

// This http handler function handle the request of adding new Author
func HandleAddAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var newAuthor Author
	err := json.NewDecoder(r.Body).Decode(&newAuthor)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), 400)
		return
	}
	newAuthor.AddAuthor()
	err = json.NewEncoder(w).Encode(AuthorList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// This http handler function handle the request of deleting an existing Author
func HandleDeleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	id, err := strconv.Atoi(chi.URLParam(r, "authorId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), 400)
		return
	}
	if _, ok := AuthorList[id]; ok == false {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Invalid Author id", 404)
		return
	}
	oldAuthor := AuthorList[id]
	for _, i := range oldAuthor.MyBooks {
		delete(BookList, i.BookId)
	}
	delete(AuthorList, id)

	err = json.NewEncoder(w).Encode(AuthorList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// This http handler function handle the request of getting all the Author
func HandleGetAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	err := json.NewEncoder(w).Encode(AuthorList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}
