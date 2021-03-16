package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var Router = mux.NewRouter()
var wait time.Duration
var server *http.Server

type Article struct{
	Id string `json: "Id"`
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func CreateDB() {
	Articles = []Article{
		Article{Id: "1", Title: "Test title", Desc: "Test Description", Content: "Hello World"},
		Article{Id: "2", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "3", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println(w, "Homepage Endpoint Hit")
}

func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
	GracefulShutDown()
}

func GracefulShutDown(){
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	server.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}

func init(){
	CreateDB()
	handleRequests()
}