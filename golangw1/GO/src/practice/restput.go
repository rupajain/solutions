package main

import (
	"fmt"
	"net/http"
	
"encoding/json"
"log"
"math/rand"
"strconv"
"github.com/gorilla/mux"
)

// const{
// 	vault_token={""}
// }
//BOOK Struct(Model)
type Book struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author  `json:"author"`
}

type Author struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}
//init books var as a slice nook struct
var books []Book
//get all books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
json.NewEncoder(w).Encode(books)
}

//get a book
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params :=mux.Vars(r)//getparams
	//loop through books and find the ids

	for _, item :=range books{
		if item.ID==params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//get all books
func createBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var book Book
	_=json.NewDecoder(r.Body).Decode(&book)
	book.ID=strconv.Itoa(rand.Intn(10000000))
	books=append(books,book)
	json.NewEncoder(w).Encode(book)
}

//get all books
// func updateBooks(w http.ResponseWriter, r *http.Request){

// }
//get all books
// func deleteBooks(w http.ResponseWriter, r *http.Request){

// }


func main(){
	//init router
	fmt.Println("Hello World") 
	//http:ListenAndServe()
	//http.handle("/ListKeys")
	//http.handle("/create/{keys}")
	//http.handle("/read/{}")
	//err!=hi{exit}

books=append(books,Book{ID:"1",Isbn:"444",Title:"Book one",Author:&Author{Firstname:"rupa",Lastname:"jain"}})
books=append(books,Book{ID:"2",Isbn:"444",Title:"Book 2",Author:&Author{Firstname:"rupa",Lastname:"jain"}})	

r:=mux.NewRouter()

	//route handlers/endpoints
	r.HandleFunc("/api/books",getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/books",createBooks).Methods("POST")
//r.HandleFunc("/api/books/{id}",updatebook).Methods("PUT")
//r.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")
log.Fatal(http.ListenAndServe(":8080",r))
}