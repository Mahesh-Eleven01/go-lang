// # server file of go lang
package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"I think server is listening at %s\n",r.URL.Path[1:])
}

func main(){
	http.HandleFunc("/",handler);
	log.Fatal(http.ListenAndServe(":8000",nil))
}
