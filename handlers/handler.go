package handlers

import (
	"api-server/data"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

// This http handler function handle the request of adding new Book
func HandleAddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var newBook data.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), 400)
		return
	}
	if _, ok := data.AuthorList[newBook.AuthorId]; ok != true {
		w.WriteHeader(http.StatusForbidden)
		http.Error(w, "This author is not available first add the author!", 400)
		return
	}
	newBook.AddBook()

	err = json.NewEncoder(w).Encode(data.BookList)
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
	if _, ok := data.BookList[id]; ok == false {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Invalid Book id", 404)
		return
	}
	oldBook := data.BookList[id]
	for i := 0; i < len(data.AuthorList[oldBook.AuthorId].MyBooks); i++ {
		if data.AuthorList[oldBook.AuthorId].MyBooks[i].BookId == oldBook.BookId {
			data.AuthorList[oldBook.AuthorId].MyBooks = append(data.AuthorList[oldBook.AuthorId].MyBooks[:i], data.AuthorList[oldBook.AuthorId].MyBooks[i+1:]...)
			break
		}
	}
	delete(data.BookList, id)

	err = json.NewEncoder(w).Encode(data.BookList)
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
	if _, ok := data.BookList[id]; ok == false {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Invalid Book id", 404)
		return
	}

	oldBook := data.BookList[id]
	var newBook data.Book
	err = json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, err.Error(), 400)
		return
	}

	if oldBook.AuthorId != newBook.AuthorId {
		if _, ok := data.AuthorList[newBook.AuthorId]; ok == false {
			w.WriteHeader(http.StatusNotFound)
			http.Error(w, "Invalid author id", 404)
			return
		}
		for i := 0; i < len(data.AuthorList[oldBook.AuthorId].MyBooks); i++ {
			if data.AuthorList[oldBook.AuthorId].MyBooks[i].BookId == oldBook.BookId {
				data.AuthorList[oldBook.AuthorId].MyBooks = append(data.AuthorList[oldBook.AuthorId].MyBooks[:i], data.AuthorList[oldBook.AuthorId].MyBooks[i+1:]...)
				break
			}
		}
		newBook.BookId = id
		data.BookList[id] = &newBook
		data.AuthorList[newBook.AuthorId].MyBooks = append(data.AuthorList[newBook.AuthorId].MyBooks, newBook)

		err = json.NewEncoder(w).Encode(data.BookList)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			http.Error(w, err.Error(), 400)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	newBook.BookId = id
	data.BookList[id] = &newBook
	for i := range data.AuthorList[oldBook.AuthorId].MyBooks {
		if data.AuthorList[oldBook.AuthorId].MyBooks[i].BookId == oldBook.BookId {
			data.AuthorList[oldBook.AuthorId].MyBooks[i] = newBook
			break
		}
	}
	err = json.NewEncoder(w).Encode(data.BookList)
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
	err := json.NewEncoder(w).Encode(data.BookList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// This http handler function handle the request of adding new Author
func HandleAddAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var newAuthor data.Author
	err := json.NewDecoder(r.Body).Decode(&newAuthor)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), 400)
		return
	}
	newAuthor.AddAuthor()
	err = json.NewEncoder(w).Encode(data.AuthorList)
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
	if _, ok := data.AuthorList[id]; ok == false {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Invalid Author id", 404)
		return
	}
	oldAuthor := data.AuthorList[id]
	for _, i := range oldAuthor.MyBooks {
		delete(data.BookList, i.BookId)
	}
	delete(data.AuthorList, id)

	err = json.NewEncoder(w).Encode(data.AuthorList)
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
	err := json.NewEncoder(w).Encode(data.AuthorList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}
