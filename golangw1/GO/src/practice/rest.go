package main
import(
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	
)
var(
	addr=flag.String("addr",":8200","http service address")
	data map[string] string
)

func main(){
	flag.Parse()
	data=map[string] string{}
	r:=httprouter.New()
	r.GET("/entry/:key",show)
	r.GET("/list",show)
	r.PUT("/entry/:key/:value",update)
	err:=http.ListenAndServe(*addr,r)

	if err!=nil{
		log.Fatal("ListenAndServe:",err)
	}
}

func show(w http.ResponseWriter,r *http.Request, p httprouter.Params){
	k:=p.ByName("key")
	if k == ""{
		fmt.Fprint(w,"Read list: %v ",data)
		return
	}

	fmt.Fprint(w,"ReadEntry: data[%s] = %s",k, data[k])
	
}
func update(w http.ResponseWriter, r *http.Request,p httprouter.Params){
k:=p.ByName("key")
v:=p.ByName("value")
data[k]=v 
fmt.Println("updated")
fmt.Fprint(w,"updates: data[%s] = %s",k,data[k])

}

