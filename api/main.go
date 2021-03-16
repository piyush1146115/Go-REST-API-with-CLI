package api

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

var myRouter = mux.NewRouter().StrictSlash(true)
var wait time.Duration
var server *http.Server

type Article struct {
	Id      string `json: "Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
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

func CreateServer() {
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "The duration for which our server gracefully wait for existing connections to finish")
	flag.Parse()

	server = &http.Server{
		Addr:         "10000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      myRouter,
	}
}

func SetValue(P string) {
	server.Addr = ":" + P

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	GracefulShutDown()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	myRouter.HandleFunc("/", homePage)
}

func GracefulShutDown() {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	server.Shutdown(ctx)
	log.Println("Shutting down!!")
	os.Exit(0)
}

func init() {
	CreateDB()
	handleRequests()
}
