package main

import (
	"net/http"
	  "fmt"
	  "encoding/json"
	  _ "github.com/mattn/go-sqlite3"
	  "strconv"
	  "time"
)

type BookData struct {
	Id         int
Title      string
Authur     string
Publisher  string
Popub      string
Catogory   string
Language   string
ISBN       string
Dop        string
Pages      string
Price      string
Genre      string
Format     string
Remarks    string

};

var count int64 = 10 

func addBook(w http.ResponseWriter, r *http.Request){
	fmt.Print("book_controller:addBook")

	if err := db.Ping(); err != nil {
		fmt.Print("book_controller:addBook ERROR::DB closed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
  	}

	var book BookData
	err := json.NewDecoder(r.Body).Decode(&book)
	fmt.Print("Testing %+v\n",book)

	//fmt.Print("Testing %+v\n", getDBTableCreationString())	
	
	if err != nil {
		fmt.Println("book_controller:addBook ERROR::HTTP JSON decode failed")
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Println("book_controller:addBook DB is open")
	nextIndex := getMaxBookIndex() + 1
	bookId := book.Title + "_" + strconv.Itoa(nextIndex)

	t := time.Now()
	formattedTime := t.Format(time.RFC3339)
	
	_, err = db.Exec("insert into book (bookindex, id, dateofentry, author, title, language, level, publisher, placeofpublication, dateofpublication, pages, price, source, isbn, genre, booktype, format, remarks, numofcopies, ischeckedout) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
	nextIndex, bookId, formattedTime , book.Authur, book.Title, 1, 1, book.Publisher, book.Popub, book.Dop, 100, book.Price, "source", book.ISBN, 1, "booktype", 1, "remark", count, 0)

	if err!= nil {
		fmt.Println("book_controller:addBook ERROR::DB Error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("book_controller:addBook DB insetion successful")

}

func getMaxBookIndex() int{
	if err := db.Ping(); err != nil {
		fmt.Println("book_controller:getMaxBookIndex ERROR::DB closed")
		return -1
	}

	var maxIndex int = -1
	err := db.QueryRow("select max(bookindex) from book;").Scan(&maxIndex)

	if err!= nil {
		fmt.Println("book_controller:addgetMaxBookIndexBook ERROR::DB closed")
		return -1
	}

	fmt.Println("book_controller:getMaxBookIndex maxIndex=%d", maxIndex)
	return maxIndex  

}

func isDBConnected() bool{
	if err := db.Ping(); err != nil {
		fmt.Println("book_controller:addBook ERROR::DB closed")
		return false
	  }
	  return true
}
func getDBTableCreationString() string{
	return  "   CREATE TABLE  checkouts (                                                  " +
			"	userindex            int    ,                                               " +
			"	bookindex            int  NOT NULL  ,                                       " +
			"	userid                 varchar(10)    ,                                       " +
			"	username                 varchar(50)    ,                                       " +
			"	bookid                 varchar(10)    ,                                       " +
				"	booktitle                 varchar(50)    ,                                       " +
			"	checkedout           varchar(20)    ,                                              " +
			"	duedate              varchar(20)    ,                                              "  +
			"	CONSTRAINT pk_checkouts PRIMARY KEY ( bookindex )                           " +
			" ) ;                                                               " +
			"                                                                                " +
			"   CREATE TABLE  checkouthistory (                                                  " +
			"	userindex            int    ,                                               " +
			"	bookindex            int  NOT NULL  ,                                       " +
			"	userid                 varchar(10)    ,                                       " +
			"	username                 varchar(50)    ,                                       " +
			"	bookid                 varchar(10)    ,                                       " +
				"	booktitle                 varchar(50)    ,                                       " +
			"	checkedout           varchar(20)    ,                                              " +
			"	duedate              varchar(20)                                                 "  +
			" ) ;                                                               " +
			"              " +
			"   CREATE TABLE book (                                                            " +
			"	bookindex            int  NOT NULL  ,                                       " +
			"	id                   varchar(10) UNIQUE ,                                         " +
			"	dateofentry      varchar(20)    ,                                          " +
			"	author               varchar(50)    ,                                       " +
			"	title                varchar(50)    ,                                       " +
			"	language             int    ,                                               " +
			"	level                int    ,                                               " +
			"	publisher            varchar(50)    ,                                       " +
			"	placeofpublication   varchar(50)    ,                                       " +
			"	dateofpublication    varchar(20)    ,                                              "  +
			"	pages                int    ,                                               " +
			"	price                varchar(10)    ,                                       " +
			"	source               varchar(50)    ,                                       " +
			"	isbn                 varchar(50)    ,                                       " +
			"	genre                int    ,                                               " +
			"	booktype             int    ,                                               " +
			"	format               int    ,                                               " +
			"	remarks              varchar(50)    ,                                       " +
			"	numofcopies          int    ,                                       " +
			"	ischeckedout              int    ,                                       " +
			"	CONSTRAINT pk_book PRIMARY KEY ( bookindex )                                " +
			" );                                                                                    "  +
			"CREATE TABLE  bookcategory (                                               " +
			"	bookcategoryindex    int  NOT NULL  ,                                       " +
			"	id                   varchar(10)  UNIQUE  ,                                       " +
			"	name                 varchar(20)    ,                                       " +
			"	remarks              varchar(20)    ,                                       " +
			"	CONSTRAINT pk_bookcategory PRIMARY KEY ( bookcategoryindex )                " +
			" );                                                               " +
			"                                                                                " +
			"CREATE TABLE  bookformat (                                                 " +
			"	formatindex          int  NOT NULL  ,                                       " +
			"	id                   varchar(10)  UNIQUE  ,                                       " +
			"	name                 varchar(20)    ,                                       " +
			"	remarks              varchar(20)    ,                                       " +
			"	CONSTRAINT pk_bookformat PRIMARY KEY ( formatindex )                        " +
			" );                                                               " +
			"                                                                                " +
			"CREATE TABLE  booktypes (                                                  " +
			"	booktypeindex        int  NOT NULL  ,                                       " +
			"	id                   varchar(10)   UNIQUE ,                                       " +
			"	name                 varchar(20)    ,                                       " +
			"	remarks              varchar(20)    ,                                       " +
			"	CONSTRAINT pk_booktypes PRIMARY KEY ( booktypeindex )                       " +
			" );                                                               " +
			"                                                                                " +
			"CREATE TABLE  fines (                                                      " +
			"	rate                 float    ,                                             " +
			"	remarks              varchar(20)                                            " +
			" ) ;                                                               " +
			"                                                                                " +
			"CREATE TABLE  genre (                                                      " +
			"	genreindex           int  NOT NULL  ,                                       " +
			"	id                   varchar(10)   UNIQUE ,                                       " +
			"	name                 varchar(20)    ,                                       " +
			"	remarks              varchar(20)    ,                                       " +
			"	CONSTRAINT pk_genre PRIMARY KEY ( genreindex )                              " +
			" ) ;                                                               " +
			"                                                                                " +
			"CREATE TABLE  level (                                                      " +
			"	levelindex           int  NOT NULL  ,                                       " +
			"	id                   varchar(10)   UNIQUE ,                                       " +
			"	name                 varchar(20)    ,                                       " +
			"	remarks              varchar(20)    ,                                       " +
			"	CONSTRAINT pk_level PRIMARY KEY ( levelindex )                              " +
			" );                                                               " +
			"                                                                                " +
			"CREATE TABLE  libuser (                                                     " +
			"	userindex            int  NOT NULL  ,                                       " +
			"	id                   varchar(10)   UNIQUE ,                                        " +
			"	name                 varchar(50)    ,                                       " +
			"	address              varchar(100)    ,                                      " +
			"	membercategory       int    ,                                               " +
			"	grade                varchar(10)    ,                                       " +
			"	year                 int    ,                                               " +
			"	level                int    ,                                               " +
			"	gender               int    ,                                               " +
			"	dob                  varchar(20)    ,                                              " +
			"	contactnum           varchar(15)    ,                                       " +
			"	other                varchar(50)    ,                                       " +
			"	CONSTRAINT pk_user PRIMARY KEY ( userindex )                                " +
			" ) ;                                                               " +
			"                                                                                " +
			"                                                                                " +
			"CREATE TABLE  usercategory (                                               " +
			"	usercategoryindex    int  NOT NULL  ,                                       " +
			"	id                  varchar(10)   UNIQUE ,                                     " +
			"	name                 varchar(20)    ,                                       " +
			"	remarks              varchar(20)    ,                                       " +
			"	CONSTRAINT pk_usercategory PRIMARY KEY ( usercategoryindex )                " +
			" ) ;   ";

}