package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/souravsk/Book_Store_Management_go/pkg/utils"
	"github.com/souravsk/Book_Store_Management_go/pkg/models"
)

type NewBook models.Book /* here we have created struct NewBook of models.Book type so that we can access the struct that we have creatd in models pkg */


func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks) /* we are using marshal func to convert into json */
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Response){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookD
}