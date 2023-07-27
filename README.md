## RESTful HTTP API server using [Go](https://github.com/golang), [Cobra CLI](https://github.com/spf13/cobra), [Go-chi](https://github.com/go-chi/chi)

### Description

This is a basic RESTful API server to handle book and book author, build with Golang. This API server is implemented with Cobra CLI for running the API from the CLI and also used go-chi instead of Go net/http.


------------ 

### Installation

- `git clone https://github.com/MobarakHsn/api-server.git`
- `cd api-server`
- `go install api-server`

---------------

### Run by CLI Commands

- start the API in default port : 8081 by `api-server serve`
- start the API in your given port by `api-server serve -p=8088`, give your port number in the place of 8088

--------------

### The Endpoints of this REST API

| Endpoint                | Function        | Method | StatusCode                                    | Authentication |
|-------------------------|-----------------|--------|-----------------------------------------------|----------------|
| `/Login`                | Login           | POST   | StatusOK, StatusUnauthorized                  | Basic          |
| `/Book`                 | ShowAllBooks    | GET    | StatusOK, StatusUnauthorized                  | JWT            |
| `/Book`                 | AddBook         | POST   | StatusOK, StatusUnauthorized                  | JWT            |
| `/Book/{id}`            | DeleteBook      | DELETE | StatusOK, StatusNoContent, StatusUnauthorized | JWT            |
| `/Book/{id}`            | UpdateBook      | PUT    | StatusOK, StatusNoContent, StatusUnauthorized | JWT            |
| `/Author`               | ShowAllAuthors  | GET    | StatusOK, StatusUnauthorized                  | JWT            |
| `/Author`               | AddAuthor       | POST   | StatusOK, StatusUnauthorized                  | JWT            |
| `/Author/{id}`          | DeleteBook      | DELETE | StatusOK, StatusNoContent, StatusUnauthorized | JWT            |

----------------

### Data Model

* Author Model
```
    type Author struct {
		AuthorId    int     `json:"authorId"`
		AuthorName  string  `json:"authorName"`
		AddressInfo Address `json:"addressInfo"`
		MyBooks     []Book  `json:"myBooks"`
	}

```

* Address Model
```
    type Address struct {
		StreetNumber string `json:"streetNumber"`
		StreetName   string `json:"streetName"`
		City         string `json:"city"`
		Country      string `json:"country"`
	}

```

* Book Model
```
    type Book struct {
		BookId       int    `json:"bookId"`
		Title        string `json:"title"`
		Language     string `json:"language"`
		NumberOfPage int    `json:"numberOfPage"`
		Price        int    `json:"price"`
		AuthorId     int    `json:"authorId"`
	}
```

* Credentials Model
```
   type Credentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

```

----------------

### JWT Authentication

- Implemented JWT authentication
- First of all user need to hit `/Login` endpoint with basic authentication then a token will be given and with that token for specific time user can do all other requests
----------------

#### Run the API server

- `curl -X POST -d '{"username":"Mobarak","password":"123"}' http://localhost:8081/Login`

#### Get all Books

- `curl -X GET -H "Authorization: ${TOKEN}" http://localhost:8081/Book`


#### Get all Authors

- `curl -X GET -H "Authorization: ${TOKEN}" http://localhost:8081/Author`


#### add new Book

```
curl -X POST -d '{"title":"Basic GO","language":"English","numberOfPage":161,"price":200,"authorId":1}' -H "Authorization: ${TOKEN}" http://localhost:8081/Book
```

#### add new Author

```
curl -X POST -d '{"authorName":"Sabbir","addressInfo":{"streetNumber":"10/A","streetName":"Sector 10","city":"Dhaka","country":"Bangladesh"},"myBooks":[]}' -H "Authorization: ${TOKEN}" http://localhost:8081/Author
```

#### Update any Book

```
curl -X PUT  -d '{"title":"Introduction to GO","language":"English","numberOfPage":161,"price":200,"authorId":2}' -H "Authorization: ${TOKEN}" http://localhost:8081/Book/1
```

#### Delete a Book
```
curl -X DELETE -H "Authorization: ${TOKEN}" http://localhost:8081/Book/1`
```

#### Delete a Author
```
curl -X DELETE -H "Authorization: ${TOKEN}" http://localhost:8081/Author/1`
```
-----------------

#### Docker image

- `docker pull mobarak9239/api-server:latest`
- `docker run -p <local port number>:8081 <image id>`

----------------

### API Endpoints Testing

- Primarily tested the API endpoints by [Postman](https://github.com/postmanlabs)
