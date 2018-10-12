package main
import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
)
type Article struct{
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}
type Articles[] Article
func allArticles(w http.ResponseWriter,r * http.Request){
	articles := Articles{
		Article{ Title:"test Title",
	Desc:"Test description",
	Content:"Hello World" },
	}

fmt.Println("Endpoint Hit:All Articles Endpoint")


json.NewEncoder(w).Encode(articles)
}

func homepage(w http.ResponseWriter, r* http.Request){
	fmt.Fprintf(w,"Homepage endpoint Hit")
}

func handleRequests(){
	http.HandleFunc("/",homepage)
	http.HandleFunc("/articles",allArticles)
	log.Fatal(http.ListenAndServe(":8081",nil))
}
func main(){
	handleRequests()
	// res,err := http.Get("http://127.0.0.1:8200")
	// res.Header.Set("","")
	// if err !=nil{
	// 	log.Fatal(err)
	// }
	// res1, err :=ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// if err !=nil {
	// 	log.Fatal(err)
	// }
	//fmt.Printf("cccccc %s",res1)

	
}