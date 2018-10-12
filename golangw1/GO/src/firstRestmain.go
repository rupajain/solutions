package main
import {
	"fmt"
	"log"
	"net/http"
	"encoding/json"
}

type Article struct{
	Title String 'json:"Title"'
	Desc String 'json:"desc"'
	Content String 'json:"content"'
}
type Articles[] Article
func allArticles(w http.ResponseWriter, r * http.Request){
	articles:=Articles{Article{Title:"test",Desc:"test",content:"content1"}}
	fmt.println("Endpoint Hit:All content article endpoint")
	json.NewEncode(w).Encode(articles)
}
func homepage(w http.ResponseWriter, r *http.Request){
	fmt.fprintf(w,"homepage endpoint hit")
}
func handleRequests(){
	http.HandleFunc("/",homepage)
	http.HandleFunc("/articles",allArticles)
log.fatal(http.ListenAndServe(":8081",nil))
}


func main()
{
	handleRequest()
}