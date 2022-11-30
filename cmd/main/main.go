package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/souravsk/Book_Store_Management_go/pkg/routes"
)

func main(){
	r := mux.NewRouter() /* we created a r varibale of mux router */
	routes.RegisterBookStoreRoutes(r) /* then we call the this variable that is creaded by in routes dir and we are getting the values */
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r)) /* ListenAndServer is func that help to create server */
} 
