package data

// Defining Author structure
type Author struct {
	AuthorId    int     `json:"authorId"`
	AuthorName  string  `json:"authorName"`
	AddressInfo Address `json:"addressInfo"`
	MyBooks     []Book  `json:"myBooks"`
}

// Defining Address structure
type Address struct {
	StreetNumber string `json:"streetNumber"`
	StreetName   string `json:"streetName"`
	City         string `json:"city"`
	Country      string `json:"country"`
}

// Defining Book structure
type Book struct {
	BookId       int    `json:"bookId"`
	Title        string `json:"title"`
	Language     string `json:"language"`
	NumberOfPage int    `json:"numberOfPage"`
	Price        int    `json:"price"`
	AuthorId     int    `json:"authorId"`
}

// Defining Credential structure
type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Defining Databases
type BookDB map[int]*Book
type AuthorDB map[int]*Author
type CredsDB map[string]string

// Declaring Databases
var BookList BookDB
var AuthorList AuthorDB
var CredentialList CredsDB

// Initializing Databases
func init() {
	BookList = make(BookDB)
	AuthorList = make(AuthorDB)
	CredentialList = make(CredsDB)
	GenerateAuthor()
	GenerateBook()
	GenerateCreds()
}

// Generating a demo Author
func GenerateAuthor() {
	AuthorList[1] = &Author{
		AuthorId:   1,
		AuthorName: "CALEB DOXSEY",
		AddressInfo: Address{
			StreetNumber: "10/A",
			StreetName:   "Sector 10",
			City:         "Dhaka",
			Country:      "Bangladesh",
		},
		MyBooks: []Book{},
	}
	//x, _ := json.Marshal(AuthorList[1])
	//fmt.Printf("%+v\n", string(x))
}

// Generating a demo Book
func GenerateBook() {
	BookList[1] = &Book{
		BookId:       1,
		Title:        "AN INTRODUCTION TO PROGRAMMING IN GO",
		Language:     "English",
		NumberOfPage: 161,
		Price:        200,
		AuthorId:     1,
	}
	AuthorList[1].MyBooks = append(AuthorList[1].MyBooks, *BookList[1])
}

// Generating a demo Credential
func GenerateCreds() {
	cred := Credential{
		"Mobarak",
		"123",
	}
	CredentialList[cred.Username] = cred.Password
}

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
