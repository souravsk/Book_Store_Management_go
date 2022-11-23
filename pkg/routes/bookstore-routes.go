package routes

import (
	"github.com/gorilla/mux"                                       /* This is help to create routes so that we can routes to the methods */
	"github.com/souravsk/Book_Store_Management_go/pkg/controllers" /* Here we are importing the controller folder */
)

/* So here we are created routes means this will except the request from the user */

var RegisterBookStoreRoutes = func(router *mux.Router) { /* CREATED a variable of func type then created a router of mux.router type  */
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookid}", controllers.DeleteBook).Methods("DELETE")
}
