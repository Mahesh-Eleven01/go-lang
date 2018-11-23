package router

/**
 * Routing for go-lang , angular routing
 * Created: M Mahesh
 * Date of Creation : 11 November 2018
*/

import (
 	"fmt"
 	"encoding/json"
 	"github.com/go-lang/logger"
 	"github.com/gorilla/mux"
 	"net/http"
)

type Route struct{
	Name string
	Method string
	Pattern string
	Description string
	HandlerFunc http.HandlerFunc
}

type Url struct {
	Endpoint string `json:"endpoint"`
	Parameters string `json:"parameters"`
	Description string `json:"description"`
	Name string `json:"name"`
}

type Routes []Route

var routes = Routes{

	Route{
		Name:"GetAll",
		Method:"GET",
		Pattern:"/getall",
		Description:"Getting all the existing endpoints",
		HandlerFunc:GetAllEndPoints,
	},

	Route{
		Name:"GetRegister",
		Method:"GET",
		Pattern:"/register",
		Description:"Register Get Route",
		HandlerFunc: RegisterGetHandler,
	},

	Route{
		Name:"PostRegister",
		Method:"POST",
		Pattern: "/register",
		Description: "Register Post Route",
		HandlerFunc: RegisterPostHandler,
	},

	Route{
		"GetLogin",
		"GET",
		"/login",
		"Login get route",
		 LoginGetHandler,
	},
	
	Route{
		"PostLogin",
		"POST",
		"/login",
		"Login Post Route",
		 LoginPostHandler,
	},

	Route{
		"GetTransactions",
		"GET",
		"/transactions",
		"Getting all the transactions",
		 TransactionGetHandler,
	},

	Route {
		"GetBlocks",
		"GET",
		"/blocks",
		"Getting all the blocks from our testnet",
		 BlockGetHandler,
	},

}


func RegisterGetHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello from register but getting doubt")
}

func RegisterPostHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Post Route from Register Page")
}

func LoginGetHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello from Login but getting doubt")
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Post Route from Login Page")
}


func TransactionGetHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello from transactions but getting doubt")
}

func BlockGetHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Get Route from Block Page")
}


func GetAllEndPoints(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Access-Control-Allow-Origin","*")
	json.NewEncoder(w).Encode(urllist)
}

var urllist []Url

func NewRouter() *mux.Router {
	MakeUrlList();
	r := mux.NewRouter()

	for _,route := range routes {

		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		r.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// r.HandleFunc("/",handler).Methods("GET")
	// r.HandleFunc("/admin",router.RegisterGetHandler).Methods("GET")
	// r.HandleFunc("/admin",router.RegisterPostHandler).Methods("POST")

	return r
}

func MakeUrlList(){
	for _,route := range routes {
		urllist = append(urllist,Url{
				Endpoint : route.Pattern,
				Parameters: "",
				Description: route.Description,
				Name: route.Name,

			})
	}
}


