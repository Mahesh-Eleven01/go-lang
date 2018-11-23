// # server file of go lang
package main

import (

	"log"
	"net/http"
	"github.com/go-lang/routes"
	"io"
)

func handler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/plain")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, "Hello this is generated from admin previlized")
}

func main(){

	r:= router.NewRouter();

	// http.HandleFunc("/",handler);
	log.Fatal(http.ListenAndServe(":8000",r))
}
