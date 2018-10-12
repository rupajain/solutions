package main
import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
	"bytes"
)



type Article struct{
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}
//type Articles[] Article
func allArticles(w http.ResponseWriter, r * http.Request){
	jsonData:=`{
		"options": {
			"version": "2"
		},
		"data": {
			"foo": "bar",
			"zip": "zap"
		  }
	  }`
	

    //articles:=Articles{Article{Title:"test",Desc:"test",Content:"content1"}}
    fmt.Println("Endpoint Hit:All content article endpoint")
	json.NewEncoder(w).Encode(jsonData)
	fmt.Println(w)
	request,err:=http.NewRequest("POST","http://127.0.0.1:8200/v1/secret/data/",bytes.NewBuffer(w))
}
func homepage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w,"homepage endpoint hit")
}
func handleRequests(){
    http.HandleFunc("/",homepage)
    http.HandleFunc("/articles",allArticles)
log.Fatal(http.ListenAndServe(":8081",nil))
}


func main(){

	handleRequests()
}
